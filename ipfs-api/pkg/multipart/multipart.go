package multipart

import (
	"io"
	"log/slog"
	"mime/multipart"

	"github.com/Gabi-Limberea/educhain/ipfs-api/pkg/zip"
)

// AddFileToForm adds a file to the multipart form data
func AddFileToForm(
	multipartWriter *multipart.Writer, file zip.FileWithName, fieldName string,
) error {
	fw, err := multipartWriter.CreateFormFile(fieldName, file.Name)
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

func AddFilesInZipToForm(
	multipartWriter *multipart.Writer, archive io.ReaderAt, archiveSize int64, fieldName string,
) error {
	files, err := zip.UnzipArchive(archive, archiveSize)
	if err != nil {
		slog.Error("error unzipping archive", "msg", err)
		return err
	}

	for _, file := range files {
		if err = AddFileToForm(multipartWriter, file, fieldName); err != nil {
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
