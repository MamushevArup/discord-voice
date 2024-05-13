package discord

import (
	"github.com/MamushevArup/ds-voice/internal/handler/discord/command"
	"github.com/MamushevArup/ds-voice/internal/handler/discord/voice"
	"github.com/MamushevArup/ds-voice/internal/usecase"
	"github.com/MamushevArup/ds-voice/pkg/logger"
	"github.com/bwmarrin/discordgo"
)

// storage defines method for use cases
type storage interface {
	UploadAudio(filePath string) (string, error)
	DownloadAudio(uploadID string, format usecase.Format) (string, error)
}

type commandH interface {
	HandleCommand(s *discordgo.Session, m *discordgo.MessageCreate)
}

type voiceI interface {
	HandleChannelCreate(s *discordgo.Session, c *discordgo.ChannelCreate)
}

type Bot struct {
	command commandH
	voiceI  voiceI
	session *discordgo.Session
	log     *logger.Logger
}

func NewBot(session *discordgo.Session, storage2 storage, logger *logger.Logger) *Bot {
	return &Bot{
		command: command.NewBot(session, storage2, logger),
		voiceI:  voice.New(session, storage2, logger),
		log:     logger,
		session: session,
	}
}
