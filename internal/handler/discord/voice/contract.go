package voice

import (
	"github.com/MamushevArup/ds-voice/pkg/logger"
	"github.com/bwmarrin/discordgo"
)

type uploader interface {
	UploadAudio(filePath string) (string, error)
}

type Record struct {
	session  *discordgo.Session
	uploader uploader
	log      *logger.Logger
	stop     chan struct{}
}

func New(session *discordgo.Session, upl uploader, logger *logger.Logger) *Record {
	return &Record{
		session:  session,
		uploader: upl,
		log:      logger,
		stop:     make(chan struct{}),
	}
}
