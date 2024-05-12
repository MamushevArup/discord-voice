package discord

import (
	"github.com/bwmarrin/discordgo"
)

// storage defines method for use cases
type storage interface {
	UploadAudio(filePath string) (string, error)
}

type Bot struct {
	session        *discordgo.Session
	storage        storage
	joinedChannels map[string]bool
	stop           chan struct{}
}

func NewBot(session *discordgo.Session, storage2 storage) *Bot {
	return &Bot{
		storage:        storage2,
		session:        session,
		joinedChannels: make(map[string]bool),
	}
}
