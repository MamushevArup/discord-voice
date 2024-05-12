package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/pion/rtp"
	"github.com/pion/webrtc/v3/pkg/media"
	"github.com/pion/webrtc/v3/pkg/media/oggwriter"
	"log"
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
	time.Sleep(10 * time.Second)

	// Send a message to the voice channel to notify that the recording is starting
	_, err = session.ChannelMessageSend(channel.ID, disclaimerRecording)
	if err != nil {
		log.Printf("Failed to send message to voice channel: %v", err)
		return
	}

	voiceChannelZero := make(chan struct{})

	// timer record is specified timer after it reach the time it stop record
	timer := time.NewTimer(recordDuration)

	go stopRecord(voiceChannelZero, vc, timer)

	// Continuously check if the voice channel is empty
	go checkMemberSize(session, channel, voiceChannelZero)

	err = recordVoice(vc.OpusRecv)
	if err != nil {
		log.Printf("Failed to record voice: %v", err)
	}
}

func stopRecord(voiceChannelZero chan struct{}, vc *discordgo.VoiceConnection, timer *time.Timer) {

	for {
		select {
		// wait for signal when members reach 0
		case <-voiceChannelZero:
			vc.Close()
			fmt.Println("Record stopped no users or all of them are bots")
		case <-timer.C:
			vc.Close()
			fmt.Println("Record stopped time limit reached")
		}
	}
}

func checkMemberSize(session *discordgo.Session, channel *discordgo.ChannelCreate, voiceChannelZero chan struct{}) {
	// ticker go every timer and check for channel size
	ticker := time.NewTicker(checkChannelSize)

	for range ticker.C {
		guild, err := session.State.Guild(channel.GuildID)
		if err != nil {
			log.Printf("Failed to get guild info: %v", err)
			return
		}

		// Check if all members in the voice channel are bots
		onlyBotsInChannel := true
		for _, vs := range guild.VoiceStates {
			if vs.ChannelID == channel.ID {
				member, err := session.State.Member(channel.GuildID, vs.UserID)
				if err != nil {
					log.Printf("Failed to get member info: %v", err)
					return
				}
				if !member.User.Bot {
					onlyBotsInChannel = false
					break
				}
			}
		}

		if onlyBotsInChannel {
			fmt.Println("Voice channel only has bots. Stopping recording.")
			ticker.Stop()
			voiceChannelZero <- struct{}{}
			return
		}
	}
}

func recordVoice(c chan *discordgo.Packet) error {
	fmt.Println("Recording voice", len(c))

	files := make(map[uint32]media.Writer)
	for p := range c {
		file, ok := files[p.SSRC]
		if !ok {
			var err error
			file, err = oggwriter.New(fmt.Sprintf("%d.ogg", p.SSRC), sampleRate, channelCount)
			if err != nil {
				return fmt.Errorf("failed to create file %d.ogg, giving up on recording: %v", p.SSRC, err)
			}
			files[p.SSRC] = file
		}
		rtp := createPionRTPPacket(p)
		err := file.WriteRTP(rtp)
		if err != nil {
			return fmt.Errorf("failed to write to file %d.ogg, giving up on recording: %v", p.SSRC, err)
		}
	}

	for _, f := range files {
		f.Close() // nolint: errcheck
	}

	return nil
}

func createPionRTPPacket(p *discordgo.Packet) *rtp.Packet {
	return &rtp.Packet{
		Header: rtp.Header{
			Version:        rtpVersion,
			PayloadType:    payloadType,
			SequenceNumber: p.Sequence,
			Timestamp:      p.Timestamp,
			SSRC:           p.SSRC,
		},
		Payload: p.Opus,
	}
}
