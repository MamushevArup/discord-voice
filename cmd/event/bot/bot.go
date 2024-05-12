package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"time"
)

var (
	rtpVersion  uint8 = 2
	payloadType uint8 = 0x78
)

const (
	sampleRate     = 48000
	channelCount   = 2
	recordDuration = 1 * time.Minute
)

type Bot struct {
	session        *discordgo.Session
	joinedChannels map[string]bool
}

func NewBot(session *discordgo.Session) *Bot {
	return &Bot{
		session:        session,
		joinedChannels: make(map[string]bool),
	}
}

// StartBot main entry point for `frontend of the discord`
func (b *Bot) StartBot() error {

	err := b.session.Open()
	if err != nil {
		fmt.Println("Error opening connection:", err)
		return err
	}

	b.session.AddHandler(b.handleChannelCreate)

	return nil
}
