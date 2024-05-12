package aws

import (
	"os"
)

type CloudService interface {
	UploadOne(file *os.File) (uploadID string, err error)
	GetOneUrl(uploadID string) (url string, err error)
}
