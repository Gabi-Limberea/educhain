package zip

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log/slog"
	"strings"
)

// ConvertTarGzToZip converts a tar.gz file to a zip archive
func ConvertTarGzToZip(tarGz io.ReadCloser) ([]byte, error) {
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

func UnzipArchive(archive io.ReaderAt, archiveSize int64) ([]FileWithName, error) {
	zipReader, err := zip.NewReader(archive, archiveSize)
	if err != nil {
		slog.Error("error creating zip reader", "msg", err)
		return nil, err
	}

	var files []FileWithName
	for _, file := range zipReader.File {
		fileName := file.Name
		if toks := strings.Split(file.Name, "."); len(toks) == 2 {
			fileName = toks[0]
		}

		slog.Info(
			"Opening file from zip", "file name", fileName, "size",
			file.UncompressedSize64,
		)
		fileReader, err := file.Open()
		if err != nil {
			slog.Error("error opening file in zip", "msg", err)
			return nil, err
		}

		files = append(
			files, FileWithName{
				Name: fileName, ReadCloser: fileReader,
			},
		)
	}

	return files, nil
}

type FileWithName struct {
	Name string
	io.ReadCloser
}
