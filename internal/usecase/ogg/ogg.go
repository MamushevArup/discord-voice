package ogg

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func Download(url, uploadID string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to download file from URL %v: %v", url, err)
	}
	defer resp.Body.Close() // nolint:errcheck

	// Create a temporary file
	tempDir := os.TempDir()

	filePath := filepath.Join(tempDir, uploadID+".ogg")

	out, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create temporary file: %v", err)
	}

	defer out.Close() // nolint:errcheck

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to write to temporary file: %v", err)
	}
	return filePath, nil
}
