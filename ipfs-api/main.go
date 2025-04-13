package main

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"log"
	"log/slog"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Gabi-Limberea/educhain/ipfs-api/constants"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

const templateIPFSNodeHost = "http://%s:%s/api/v0"

var (
	nodeIP   string
	nodePort string
	nodeHost string
)

type fileWithName struct {
	Name string
	io.ReadCloser
}

// convertTarGzToZip converts a tar.gz file to a zip archive
func convertTarGzToZip(tarGz io.ReadCloser) ([]byte, error) {
	uncompressedTar, err := gzip.NewReader(tarGz)
	if err != nil {
		return nil, fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer uncompressedTar.Close()

	tarReader := tar.NewReader(uncompressedTar)
	buf := new(bytes.Buffer)
	zipArchive := zip.NewWriter(buf)

	// I expect this loop to only run once, as the tar file is expected to contain a single file, but it is written to
	// support multiple files as well
	for {
		var rawFileBytes []byte

		// read the next file header from the tar archive
		header, err := tarReader.Next()
		if err == io.EOF {
			slog.Info("Reached end of tar file")
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to go to next file in tar: %w", err)
		}

		// initialize file buffer
		rawFileBytes = make([]byte, header.Size)
		slog.Info("Reading file from tar", "file name", header.Name, "size", header.Size)
		if _, err = tarReader.Read(rawFileBytes); err != nil && err != io.EOF {
			return nil, fmt.Errorf("failed to read file from tar: %w", err)
		}

		// create file in zip archive with the name defined in the tar header
		slog.Info("Writing file to zip", "file name", header.Name, "size", header.Size)
		file, err := zipArchive.Create(header.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to create zip file: %w", err)
		}

		// write the file bytes to the file in the zip archive
		_, err = file.Write(rawFileBytes)
		if err != nil {
			return nil, fmt.Errorf("failed to write file to zip: %w", err)
		}
	}

	// write and close the zip archive
	if err = zipArchive.Close(); err != nil {
		return nil, fmt.Errorf("failed to write zip: %w", err)
	}
	slog.Info("File written to zip", "size", buf.Len())

	return buf.Bytes(), nil
}

// addFileToForm adds a file to the multipart form data
func addFileToForm(multipartWriter *multipart.Writer, file fileWithName) error {
	fw, err := multipartWriter.CreateFormFile("", file.Name)
	if err != nil {
		slog.Error("error creating form field", "msg", err)
		return err
	}

	if _, err = io.Copy(fw, file); err != nil {
		slog.Error("error copying form field data", "msg", err)
		return err
	}

	return nil
}

func addFilesInZipToForm(multipartWriter *multipart.Writer, archive io.ReaderAt, archiveSize int64) error {
	files, err := unzipArchive(archive, archiveSize)
	if err != nil {
		slog.Error("error unzipping archive", "msg", err)
		return err
	}

	for _, file := range files {
		if err = addFileToForm(multipartWriter, file); err != nil {
			slog.Error("error adding file to form", "msg", err)
			return err
		}
		if err = file.Close(); err != nil {
			slog.Error("error closing file", "msg", err)
			return err
		}
	}

	return nil
}

func unzipArchive(archive io.ReaderAt, archiveSize int64) ([]fileWithName, error) {
	zipReader, err := zip.NewReader(archive, archiveSize)
	if err != nil {
		slog.Error("error creating zip reader", "msg", err)
		return nil, err
	}

	var files []fileWithName
	for _, file := range zipReader.File {
		slog.Info("Opening file from zip", "file name", file.Name, "size", file.UncompressedSize64)
		fileReader, err := file.Open()
		if err != nil {
			slog.Error("error opening file in zip", "msg", err)
			return nil, err
		}

		files = append(
			files, fileWithName{
				Name: file.Name, ReadCloser: fileReader,
			},
		)
	}

	return files, nil
}

func handleFileUpload(w http.ResponseWriter, r *http.Request) {
	slog.Info("Received request to upload file")
	slog.Info("Request headers", "content-type", r.Header.Get("Content-Type"))

	if err := r.ParseMultipartForm(constants.MB512); err != nil {
		slog.Error("error parsing form", "msg", err)
		http.Error(w, "failed to parse form", http.StatusBadRequest)
		return
	}

	toSend := new(bytes.Buffer)
	multipartWriter := multipart.NewWriter(toSend)

	// I don't care for the file header, since I want to upload the file based on its content. This is to avoid
	// users from leaking sensitive information mentioned in the file name
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		slog.Error("error getting form file", "msg", err)
		http.Error(w, "failed to get form file", http.StatusBadRequest)
		return
	}

	slog.Info(
		"Received file", "file name", fileHeader.Filename, "size", fileHeader.Size, "content type",
		fileHeader.Header.Get("Content-Type"),
	)
	if fileHeader.Header.Get("Content-Type") == "application/zip" {
		slog.Info("File is a zip archive")
		if err = addFilesInZipToForm(multipartWriter, file, fileHeader.Size); err != nil {
			slog.Error("error adding files to form", "msg", err)
			http.Error(w, "something went wrong", http.StatusBadRequest)
			return
		}
	} else {
		if err = addFileToForm(
			multipartWriter, fileWithName{
				Name: fileHeader.Filename, ReadCloser: file,
			},
		); err != nil {
			slog.Error("error adding file to form", "msg", err)
			http.Error(w, "something went wrong", http.StatusBadRequest)
			return
		}
	}

	if err := multipartWriter.Close(); err != nil {
		slog.Error("error closing form writer", "msg", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	resp, err := http.Post(nodeHost+"/add", "multipart/form-data; boundary="+multipartWriter.Boundary(), toSend)
	if err != nil {
		slog.Error("error executing POST request", "msg", err)
		http.Error(w, "failed to upload file", http.StatusInternalServerError)
		return
	}

	buf := new(bytes.Buffer)
	if _, err = buf.ReadFrom(resp.Body); err != nil {
		slog.Error("error reading response body", "msg", err)
		http.Error(w, "failed to read IPFS response", http.StatusInternalServerError)
		return
	}

	splitResults := strings.Split(buf.String(), "\n")
	strResults := fmt.Sprintf("[%s]", strings.Join(splitResults[:len(splitResults)-1], ","))

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(strResults))
}

func handleFileDownload(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cid := vars["cid"]

	slog.Info("Received request to download file", "cid", cid)

	timeoutCtx, cancel := context.WithTimeout(r.Context(), 1*time.Minute)
	defer cancel()

	// create request to the IPFS node with timeout context to avoid blocking when cid cannot be found
	req, err := http.NewRequestWithContext(
		timeoutCtx, http.MethodPost, fmt.Sprintf("%s/get?arg=%s&archive=true&compress=true", nodeHost, cid), nil,
	)
	if err != nil {
		slog.Error("error creating request", "msg", err)
		http.Error(w, "failed to create request", http.StatusInternalServerError)
		return
	}

	// execute request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// if context deadline exceeded, cid (most likely) not found
		if errors.Is(err, context.DeadlineExceeded) {
			slog.Error("file not found in IPFS", "cid", cid)
			http.Error(w, "file not found", http.StatusNotFound)
			return
		}

		slog.Error("error executing POST request", "msg", err)
		http.Error(w, "failed to download file", http.StatusInternalServerError)
		return
	}

	slog.Info("Received response from IPFS", "status", resp.StatusCode)

	defer resp.Body.Close()

	// response is expected to be a tar.gz file, convert it to zip for better compatibility
	zipFile, err := convertTarGzToZip(resp.Body)
	if err != nil {
		slog.Error("error converting tar.gz to zip", "msg", err)
		http.Error(w, "failed to convert file", http.StatusInternalServerError)
		return
	}

	// set headers to trigger attachment download in browsers
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename="+cid+".zip")
	w.WriteHeader(http.StatusOK)
	if _, err := io.Copy(w, bytes.NewReader(zipFile)); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}

	slog.Info("File download completed", "cid", cid)
}

func main() {
	nodeIP = os.Getenv("IPFS_NODE_IP")
	if nodeIP == "" {
		log.Fatal("IPFS_NODE_IP environment variable not set")
	}

	nodePort = os.Getenv("IPFS_NODE_PORT")
	if nodePort == "" {
		log.Fatal("IPFS_NODE_PORT environment variable not set")
	}

	nodeHost = fmt.Sprintf(templateIPFSNodeHost, nodeIP, nodePort)

	slog.Info("IPFS API client is starting...", "IPFS_NODE_IP", nodeIP, "IPFS_NODE_PORT", nodePort)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/file/{cid}", handleFileDownload).Methods(http.MethodGet)
	router.HandleFunc("/api/file", handleFileUpload).Methods(http.MethodPost)
	router.HandleFunc(
		"/healthz", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		},
	).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
}
