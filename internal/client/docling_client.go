package client

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

type DoclingClient struct {
	Endpoint string
}

func NewDoclingClient(url string) *DoclingClient {
	return &DoclingClient{Endpoint: url}
}

func (c *DoclingClient) ExtractText(fileName string, fileData []byte) (string, error) {
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)

	fileWriter, err := writer.CreateFormFile("files", fileName)
	if err != nil {
		return "", err
	}

	if _, err := fileWriter.Write(fileData); err != nil {
		return "", err
	}

	if err := writer.Close(); err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", c.Endpoint+"/v1/convert/file", buf)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("docling error: %s", string(body))
	}

	return string(body), nil
}
