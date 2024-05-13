package voice

import "time"

var (
	rtpVersion  uint8 = 2
	payloadType uint8 = 0x78
)

const (
	sampleRate   = 48000
	channelCount = 2

	recordDuration   = 10 * time.Second
	waitFor          = 5 * time.Second
	checkChannelSize = 2 * time.Second

	disclaimerRecording = "I am recording this voice channel"
	downloadMessage     = "Audio file is saved you can download it with this ID: "

	stopRecordMessage = "Record stopped no users or all of them are bots"
	stopTimeMessage   = "Record stopped time limit reached"
	channelOnlyBots   = "Record stopped all users are bots"
)

