package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

func stopRecord(voiceChannelZero, stop chan struct{}, timer *time.Timer) {

	for {
		select {
		// wait for signal when members reach 0
		case <-voiceChannelZero:
			close(stop)
			fmt.Println(stopRecordMessage)
			return
		case <-timer.C:
			close(stop)
			fmt.Println(stopTimeMessage)
			return
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
				member, err := session.GuildMember(guild.ID, vs.UserID)
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
			fmt.Println(channelOnlyBots)
			ticker.Stop()
			voiceChannelZero <- struct{}{}
			return
		}
	}
}
