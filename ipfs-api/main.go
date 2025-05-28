package main

import (
	"bytes"
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
	multipart2 "github.com/Gabi-Limberea/educhain/ipfs-api/pkg/multipart"
	"github.com/Gabi-Limberea/educhain/ipfs-api/pkg/zip"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

const templateIPFSNodeHost = "http://%s:%s/api/v0"

var (
	nodeIP   string
	nodePort string
	nodeHost string
)

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
		if err = multipart2.AddFilesInZipToForm(
			multipartWriter, file, fileHeader.Size, "",
		); err != nil {
			slog.Error("error adding files to form", "msg", err)
			http.Error(w, "something went wrong", http.StatusBadRequest)
			return
		}
	} else {
		if err = multipart2.AddFileToForm(
			multipartWriter, zip.FileWithName{
				Name: fileHeader.Filename, ReadCloser: file,
			}, "",
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

	resp, err := http.Post(
		nodeHost+"/add", "multipart/form-data; boundary="+multipartWriter.Boundary(), toSend,
	)
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
		timeoutCtx, http.MethodPost,
		fmt.Sprintf("%s/get?arg=%s&archive=true&compress=true", nodeHost, cid), nil,
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
	zipFile, err := zip.ConvertTarGzToZip(resp.Body)
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
