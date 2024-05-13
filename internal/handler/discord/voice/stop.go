package voice

import (
	"github.com/bwmarrin/discordgo"
	"time"
)

func (v *Record) stopRecord(voiceChannelZero, stop chan struct{}, timer *time.Timer) {
	for {
		select {
		// wait for signal when members reach 0
		case <-voiceChannelZero:
			close(stop)
			v.log.Info(stopRecordMessage)
			return
		case <-timer.C:
			close(stop)
			v.log.Info(stopTimeMessage)
			return
		}
	}
}

func (v *Record) checkMemberSize(session *discordgo.Session, channel *discordgo.ChannelCreate, voiceChannelZero chan struct{}) {
	// ticker go every timer and check for channel size
	ticker := time.NewTicker(checkChannelSize)

	for range ticker.C {
		guild, err := session.State.Guild(channel.GuildID)
		if err != nil {
			v.log.Errorf("Failed to get guild info: %v", err)
			return
		}

		// Check if all members in the voice channel are bots
		onlyBotsInChannel := true
		for _, vs := range guild.VoiceStates {
			if vs.ChannelID == channel.ID {
				member, err := session.GuildMember(guild.ID, vs.UserID)
				if err != nil {
					v.log.Errorf("Failed to get member (memberID) %v, info: %v", member.User.ID, err)
					return
				}
				if !member.User.Bot {
					onlyBotsInChannel = false
					break
				}
			}
		}

		if onlyBotsInChannel {
			v.log.Info(channelOnlyBots)
			ticker.Stop()
			voiceChannelZero <- struct{}{}
			return
		}
	}
}
