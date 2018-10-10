package utils

import (
	"strings"

	"github.com/target/flottbot/model"
)

// GetChannelIDs helps find a channel by name, if we have 'cached' it
func GetChannelIDs(wantChannels []string, bot *model.Bot) []string {
	channels := []string{}

	for _, channel := range wantChannels {
		channelMatch := bot.Channels[strings.ToLower(channel)]
		if len(channelMatch) > 0 {
			channels = append(channels, channelMatch)
		} else {
			bot.Log.Debugf("Channel '%s' does not exist", channel)
		}
	}

	return channels
}
