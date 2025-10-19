package services

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
)

func NewImageService(url string) *ImageService {
	imageBlob, err := urlToBytes(url)
	if err != nil {
		return &ImageService{}
	}

	return &ImageService{Blob: imageBlob}
}

func (is ImageService) GetBase64() (string, error) {
	if is.Blob == nil {
		return "", fmt.Errorf("ImageService.Blob is nil")
	}

	encodedString := base64.StdEncoding.EncodeToString(is.Blob)
	mimeType := http.DetectContentType(is.Blob)
	dataUri := fmt.Sprintf("data:%s;base64,%s", mimeType, encodedString)

	return dataUri, nil
}

func urlToBytes(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Failed to load image url: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Received unexpected status: %w", err)
	}

	imageData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read image: %w", err)
	}

	return imageData, nil
}
