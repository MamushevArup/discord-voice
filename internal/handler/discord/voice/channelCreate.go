package voice

import (
	"github.com/bwmarrin/discordgo"
	"os"
	"time"
)

func (v *Record) HandleChannelCreate(session *discordgo.Session, channel *discordgo.ChannelCreate) {

	v.session = session
	v.log.Info("Joining voice channel")

	vc, err := session.ChannelVoiceJoin(channel.GuildID, channel.ID, true, false)
	if err != nil {
		v.log.Errorf("Failed to join voice channel (channelID) %v: %v", channel.ID, err)
		return
	}

	// Wait for 10 seconds before starting the recording
	// Send a message to the voice channel to notify that the recording is starting
	// timer record is specified timer after it reach the time it stop record
	// Continuously check if the voice channel is empty
	time.Sleep(waitFor)

	// Create a new stop channel
	v.stop = make(chan struct{})

	v.log.Info("Recording started")

	_, err = session.ChannelMessageSend(channel.ID, disclaimerRecording)
	if err != nil {
		v.log.Errorf("Failed to send message to voice channel (channelID) %v: %v", channel.ID, err)
		return
	}

	voiceChannelZero := make(chan struct{})

	timer := time.NewTimer(recordDuration)

	go v.stopRecord(voiceChannelZero, v.stop, timer)

	go v.checkMemberSize(session, channel, voiceChannelZero)

	filePaths, err := v.recordVoice(vc.OpusRecv)
	if err != nil {
		v.log.Errorf("Failed to record voice (channelID) %v: %v", channel.ID, err)
		return
	}

	for _, filePath := range filePaths {
		audioID, err := v.uploader.UploadAudio(filePath)
		if err != nil {
			v.log.Errorf("Failed to save data from channelID %v: %v", channel.ID, err)
			continue
		}
		os.Remove(filePath) //nolint:errcheck

		_, err = session.ChannelMessageSend(channel.ID, downloadMessage+audioID)
		if err != nil {
			v.log.Errorf("Failed to send message to voice channel (channelID) %v: %v", channel.ID, err)
		}

		v.log.Infof("Uploaded successfully ID: %v", audioID)
	}
}
