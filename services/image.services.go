package services

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ImageServices struct {} 

func NewImageService() *ImageServices {
	return &ImageServices{}
}

func (is ImageServices) UrlToBase64(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("Failed to load image url: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Received unexpected status: %w", err)
	}

	imageData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("Failed to read image: %w", err)
	}

	encodedString := base64.StdEncoding.EncodeToString(imageData)
	mimeType := http.DetectContentType(imageData)
	dataUri := fmt.Sprintf("data:%s;base64,%s", mimeType, encodedString)

	return dataUri, nil
}
