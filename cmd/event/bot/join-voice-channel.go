package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/pion/rtp"
	"github.com/pion/webrtc/v3/pkg/media"
	"github.com/pion/webrtc/v3/pkg/media/oggwriter"
	"log"
	"time"
)

func (b *Bot) joinVoiceChannel(guildID string, channelID string) {
	vc, err := b.session.ChannelVoiceJoin(guildID, channelID, true, false)
	if err != nil {
		log.Printf("Failed to join voice channel: %v", err)
		return
	}

	go func() {
		time.Sleep(10 * time.Second)
		fmt.Println("Leaving voice channel")
		vc.Close()
	}()
	handleVoice(vc.OpusRecv)
}

func (b *Bot) handleChannelCreate(session *discordgo.Session, channel *discordgo.ChannelCreate) {

	b.session = session

	fmt.Println("Joining voice channel")

	vc, err := session.ChannelVoiceJoin(channel.GuildID, channel.ID, true, false)
	if err != nil {
		log.Printf("Failed to join voice channel: %v", err)
		return
	}

	go func() {
		time.Sleep(recordDuration)
		fmt.Println("Leaving voice channel")
		vc.Close()
	}()
	handleVoice(vc.OpusRecv)
}

func handleVoice(c chan *discordgo.Packet) {
	fmt.Println("Recording voice", len(c))

	files := make(map[uint32]media.Writer)
	for p := range c {
		file, ok := files[p.SSRC]
		if !ok {
			var err error
			file, err = oggwriter.New(fmt.Sprintf("%d.ogg", p.SSRC), sampleRate, channelCount)
			if err != nil {
				fmt.Printf("failed to create file %d.ogg, giving up on recording: %v\n", p.SSRC, err)
				return
			}
			files[p.SSRC] = file
		}
		rtp := createPionRTPPacket(p)
		err := file.WriteRTP(rtp)
		if err != nil {
			fmt.Printf("failed to write to file %d.ogg, giving up on recording: %v\n", p.SSRC, err)
		}

	}

	for _, f := range files {
		f.Close() // nolint: errcheck
	}
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
