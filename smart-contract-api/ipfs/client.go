package ipfs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"mime/multipart"
	"net/http"
	"os"

	multipart2 "github.com/Gabi-Limberea/educhain/ipfs-api/pkg/multipart"
	"github.com/Gabi-Limberea/educhain/ipfs-api/pkg/zip"
)

type Client struct {
	host string
	port string
}

type File struct {
	Name string `json:"Name"`
	Hash string `json:"Hash"`
}

func NewClient() *Client {
	host := os.Getenv("IPFS_API_HOST")
	if host == "" {
		log.Fatal("IPFS_API_HOST environment variable not set")
	}

	port := os.Getenv("IPFS_API_PORT")
	if port == "" {
		log.Fatal("IPFS_API_PORT environment variable not set")
	}

	return &Client{host: host, port: port}
}

func (c *Client) url() string {
	basePath := "/api/file"
	return fmt.Sprintf("http://%s:%s", c.host, c.port) + basePath
}

func (c *Client) UploadFile(file io.ReadCloser, fileName string) (*File, error) {
	toSend := new(bytes.Buffer)
	multipartWriter := multipart.NewWriter(toSend)

	fileWithName := zip.FileWithName{
		Name:       fileName,
		ReadCloser: file,
	}

	if err := multipart2.AddFileToForm(multipartWriter, fileWithName, "file"); err != nil {
		slog.Error("failed to add file to form", "error", err)
		return nil, err
	}

	if err := multipartWriter.Close(); err != nil {
		slog.Error("error closing form writer", "msg", err)
		return nil, err
	}

	resp, err := http.Post(
		c.url(), "multipart/form-data; boundary="+multipartWriter.Boundary(), toSend,
	)
	if err != nil {
		slog.Error("failed to upload file", "error", err)
		return nil, err
	}

	defer resp.Body.Close()

	var files []File
	if err = json.NewDecoder(resp.Body).Decode(&files); err != nil {
		return nil, fmt.Errorf("failed to decode file upload response: %w", err)
	}
	if len(files) == 0 {
		return nil, fmt.Errorf("no file uploaded")
	}

	return &files[0], nil
}
