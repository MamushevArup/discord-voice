package usecase

import (
	"github.com/MamushevArup/ds-voice/adapters/avatar/aws"
	"sync"
)

type Format string

const (
	FormatMP3 Format = "mp3"
	FormatWAV Format = "wav"
	FormatOGG Format = "ogg"
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
