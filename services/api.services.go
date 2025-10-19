package services

import (
	"fmt"
	"io"
	"net/http"
)

func FetchApi(req *http.Request) ([]byte, error) {
	// Handle response
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Request error: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API call failed with status %d: %s", res.StatusCode, req.URL.String())
	}

	// Retrieve body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read body: %w", err)
	}

	return body, nil
}
