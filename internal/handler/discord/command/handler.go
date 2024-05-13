package command

import (
	"github.com/MamushevArup/ds-voice/internal/usecase"
	"github.com/bwmarrin/discordgo"
	"os"
	"path/filepath"
	"strings"
)

func (b *Execute) HandleCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages sent by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Check if the message starts with "!"
	if len(m.Content) > 0 && m.Content[0] == '!' {
		words := strings.Fields(m.Content[1:])
		command := words[0]

		switch command {
		case "help":
			b.helpCom(s, m)
		case "download-ogg":
			b.oggAudio(words, s, m)
		case "download-wav":
			id := b.getID(words, s, m)
			b.wavAudio(id, s, m)
		case "download-mp3":
			id := b.getID(words, s, m)
			b.mp3Audio(id, s, m)
		}
	}
}

func (b *Execute) mp3Audio(id string, s *discordgo.Session, m *discordgo.MessageCreate) {
	filePath, err := b.storage.DownloadAudio(id, usecase.FormatMP3)
	if err != nil {
		b.log.Errorf("Failed to download audio (uuid) %v: %v", id, err)
		return
	}

	// Open the local file
	cleanedPath := filepath.Clean(filePath)
	f, err := os.Open(cleanedPath)
	if err != nil {
		b.log.Errorf("Failed to open file: %v", err)
		return
	}

	defer os.Remove(filePath) //nolint:errcheck

	// Create a new message with the audio file as an attachment
	msg := &discordgo.MessageSend{
		Files: []*discordgo.File{
			{
				Name:   id + ".mp3",
				Reader: f,
			},
		},
	}

	// Send the message to a channel
	_, err = s.ChannelMessageSendComplex(m.ChannelID, msg)
	if err != nil {
		b.log.Errorf("Failed to send message (channelID) %v: %v", m.ChannelID, err)
		return
	}

}

func (b *Execute) getID(words []string, s *discordgo.Session, m *discordgo.MessageCreate) string {
	if len(words) < 2 {
		_, err := s.ChannelMessageSend(m.ChannelID, "Please specify the audio ID")
		if err != nil {
			b.log.Errorf("Failed to send message (channelID) %v: %v", m.ChannelID, err)
			return ""
		}
		b.log.Errorf("Invalid command (channelID) %v : %v", m.ChannelID, m.Content)
		return ""
	}
	return words[1]
}

func (b *Execute) wavAudio(id string, s *discordgo.Session, m *discordgo.MessageCreate) {
	filePath, err := b.storage.DownloadAudio(id, usecase.FormatWAV)
	if err != nil {
		b.log.Errorf("Failed to download audio (uuid) %v: %v", id, err)
		return
	}

	// Open the local file
	f, err := os.Open(filePath)
	if err != nil {
		b.log.Errorf("Failed to open file: %v", err)
		return
	}

	defer os.Remove(filePath) //nolint:errcheck

	// Create a new message with the audio file as an attachment
	msg := &discordgo.MessageSend{
		Files: []*discordgo.File{
			{
				Name:   id + ".wav",
				Reader: f,
			},
		},
	}

	// Send the message to a channel
	_, err = s.ChannelMessageSendComplex(m.ChannelID, msg)
	if err != nil {
		b.log.Errorf("Failed to send message (channelID) %v: %v", m.ChannelID, err)
		return
	}

}

func (b *Execute) oggAudio(words []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	uuid := b.getID(words, s, m)

	filePath, err := b.storage.DownloadAudio(uuid, usecase.FormatOGG)
	if err != nil {
		b.log.Errorf("Failed to download audio (uuid) %v: %v", uuid, err)
		return
	}

	// Open the local file
	f, err := os.Open(filePath)
	if err != nil {
		b.log.Errorf("Failed to open file: %v", err)
		return
	}

	defer os.Remove(filePath) //nolint:errcheck

	// Create a new message with the audio file as an attachment
	msg := &discordgo.MessageSend{
		Files: []*discordgo.File{
			{
				Name:   uuid + ".ogg",
				Reader: f,
			},
		},
	}

	// Send the message to a channel
	_, err = s.ChannelMessageSendComplex(m.ChannelID, msg)
	if err != nil {
		b.log.Errorf("Failed to send message (channelID) %v: %v", m.ChannelID, err)
		return
	}
}

func (b *Execute) helpCom(s *discordgo.Session, m *discordgo.MessageCreate) {
	txt, err := os.ReadFile(helpDir)
	if err != nil {
		b.log.Errorf("Failed to read help file: %v", err)
		return
	}
	msg, err := s.ChannelMessageSend(m.ChannelID, string(txt))
	if err != nil {
		// Handle error
		b.log.Errorf("Failed to send message: %v", err)
		return
	}
	b.log.Infof("Message sent: %v, author %v", msg.ID, msg.Author.ID)
}
