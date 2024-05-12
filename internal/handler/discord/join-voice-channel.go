package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"time"
)

func (b *Bot) handleChannelCreate(session *discordgo.Session, channel *discordgo.ChannelCreate) {

	b.session = session

	fmt.Println("Joining voice channel")

	vc, err := session.ChannelVoiceJoin(channel.GuildID, channel.ID, true, false)
	if err != nil {
		log.Printf("Failed to join voice channel: %v", err)
		return
	}

	// Wait for 10 seconds before starting the recording
	// Send a message to the voice channel to notify that the recording is starting
	// timer record is specified timer after it reach the time it stop record
	// Continuously check if the voice channel is empty
	time.Sleep(waitFor)

	// Create a new stop channel
	b.stop = make(chan struct{})

	_, err = session.ChannelMessageSend(channel.ID, disclaimerRecording)
	if err != nil {
		log.Printf("Failed to send message to voice channel: %v", err)
		return
	}

	voiceChannelZero := make(chan struct{})

	timer := time.NewTimer(recordDuration)

	go stopRecord(voiceChannelZero, b.stop, timer)

	go checkMemberSize(session, channel, voiceChannelZero)

	filePaths, err := b.recordVoice(vc.OpusRecv)
	if err != nil {
		log.Printf("Failed to record voice: %v", err)
	}

	for _, filePath := range filePaths {
		audioID, err := b.storage.UploadAudio(filePath)
		if err != nil {
			log.Printf("unable to save data %v", err)
		}
		os.Remove(filePath) //nolint:errcheck

		_, err = session.ChannelMessageSend(channel.ID, downloadMessage+audioID)
		if err != nil {
			log.Printf("failed to sent message %v", err)
		}

		fmt.Printf("Uploaded successfully ID: %v", audioID)
	}
}
