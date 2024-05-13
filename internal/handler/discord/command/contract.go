package command

import (
	"github.com/MamushevArup/ds-voice/internal/usecase"
	"github.com/MamushevArup/ds-voice/pkg/logger"
	"github.com/bwmarrin/discordgo"
)

type storage interface {
	DownloadAudio(uploadID string, format usecase.Format) (string, error)
}

const helpDir = "help.txt"

type Execute struct {
	session *discordgo.Session
	storage storage
	log     *logger.Logger
}

func NewBot(session *discordgo.Session, storage storage, logger *logger.Logger) *Execute {
	return &Execute{
		session: session,
		storage: storage,
		log:     logger,
	}
}
