package usecase

import (
	"fmt"
	"github.com/MamushevArup/ds-voice/internal/usecase/mp3"
	"github.com/MamushevArup/ds-voice/internal/usecase/ogg"
	"github.com/MamushevArup/ds-voice/internal/usecase/wav"
)

func (u *UseCase) DownloadAudio(uploadID string, format Format) (string, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	url, err := u.s3.GetOneUrl(uploadID)
	if err != nil {
		return "", fmt.Errorf("failed to download file from aws %v", err)
	}

	switch format {
	case FormatWAV:
		return wav.Download(url, uploadID)
	case FormatOGG:
		return ogg.Download(url, uploadID)
	case FormatMP3:
		return mp3.Download(url, uploadID)
	}
	return "", fmt.Errorf("unsupported format: %v", format)
}
