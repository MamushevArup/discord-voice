package usecase

import (
	"fmt"
	"os"
	"path/filepath"
)

func (u *UseCase) UploadAudio(filePath string) (string, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	clean := filepath.Clean(filePath)
	file, err := os.Open(clean)
	if err != nil {
		return "", fmt.Errorf("failed to open file %v", err)
	}
	defer file.Close()

	uploadID, err := u.s3.UploadOne(file)
	if err != nil {
		return "", fmt.Errorf("failed to upload file to aws %v", err)
	}

	return uploadID, nil
}
