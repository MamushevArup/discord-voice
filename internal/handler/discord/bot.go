package discord

import (
	"github.com/bwmarrin/discordgo"
	"sync"
	"time"
)

var (
	rtpVersion  uint8 = 2
	payloadType uint8 = 0x78
)

const (
	sampleRate          = 48000
	channelCount        = 2
	recordDuration      = 1 * time.Minute
	checkChannelSize    = 2 * time.Second
	disclaimerRecording = "I am recording this voice channel"
)

type Bot struct {
	session        *discordgo.Session
	joinedChannels map[string]bool
	mu             sync.Mutex
}

func NewBot(session *discordgo.Session) *Bot {
	return &Bot{
		session:        session,
		joinedChannels: make(map[string]bool),
	}
}
