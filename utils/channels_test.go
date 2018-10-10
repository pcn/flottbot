package utils

import (
	"reflect"
	"testing"

	"github.com/target/flottbot/model"
)

func TestGetChannelIDs(t *testing.T) {
	type args struct {
		wantChannels   []string
		activeChannels map[string]string
		bot            *model.Bot
	}

	// For Channel Exists
	ChannelExistsIn := []string{"testing", "testing-channel"}

	ChannelExistsActive := make(map[string]string)
	ChannelExistsActive["testing"] = "123"
	ChannelExistsActive["testing-channel"] = "456"

	ChannelExistsWant := []string{"123", "456"}

	// For Channel Doesn't Exist
	ChannelDoesNotExistIn := []string{"not"}

	ChannelDoesNotExistActive := make(map[string]string)
	ChannelDoesNotExistActive["testing"] = "123"
	ChannelDoesNotExistActive["testing-channel"] = "456"

	ChannelDoesNotExistWant := []string{}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Basic", args{}, []string{}},
		{"Channel exists", args{wantChannels: ChannelExistsIn, bot: &model.Bot{Channels: ChannelExistsActive}}, ChannelExistsWant},
		{"Channel does not exist", args{wantChannels: ChannelDoesNotExistIn, bot: &model.Bot{Channels: ChannelDoesNotExistActive}}, ChannelDoesNotExistWant},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetChannelIDs(tt.args.wantChannels, tt.args.bot); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetChannelIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}
