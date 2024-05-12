package usecase

import (
	"fmt"
	"github.com/MamushevArup/ds-voice/adapters/avatar/aws"
	"os"
	"sync"
)

type UseCase struct {
	s3 aws.CloudService
	mu sync.Mutex
}

func New(s3 aws.CloudService) *UseCase {
	return &UseCase{
		s3: s3,
	}
}

func (u *UseCase) UploadAudio(filePath string) (string, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	file, err := os.Open(filePath)
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
