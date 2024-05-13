package voice

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/pion/rtp"
	"github.com/pion/webrtc/v3/pkg/media"
	"github.com/pion/webrtc/v3/pkg/media/oggwriter"
)

func (v *Record) recordVoice(c chan *discordgo.Packet) (map[uint32]string, error) {

	v.log.Info("Recording voice")

	files := make(map[uint32]media.Writer)
	filePaths := make(map[uint32]string)
	for {
		select {
		case <-v.stop:
			// Close all file writers
			for _, f := range files {
				f.Close() // nolint: errcheck
			}
			v.log.Info("Close file writers, recording stopped")
			return filePaths, nil
		case p, ok := <-c:
			if !ok {
				v.log.Error("Failed to read from channel, giving up on recording")
				return filePaths, nil
			}
			file, ok := files[p.SSRC]
			if !ok {
				var err error
				filePath := fmt.Sprintf("%d.ogg", p.SSRC)
				file, err = oggwriter.New(filePath, sampleRate, channelCount)
				if err != nil {
					return nil, fmt.Errorf("failed to create file %s, giving up on recording: %v", filePath, err)
				}
				files[p.SSRC] = file
				filePaths[p.SSRC] = filePath
			}
			rtp := createPionRTPPacket(p)
			err := file.WriteRTP(rtp)
			if err != nil {
				return nil, fmt.Errorf("failed to write to file %s, giving up on recording: %v", file, err)
			}
		}
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
