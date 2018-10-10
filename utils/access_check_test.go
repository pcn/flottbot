package utils

import (
	"testing"

	"github.com/target/flottbot/model"
)

func TestCanTrigger(t *testing.T) {
	type args struct {
		currentUserName string
		currentUserID   string
		rule            model.Rule
		bot             *model.Bot
	}

	testBot := new(model.Bot)
	testBot.ChatApplication = "slack"

	discordBot := new(model.Bot)
	discordBot.ChatApplication = "discord"

	strangeBot := new(model.Bot)
	strangeBot.ChatApplication = "strange"

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"No restrictions", args{"jane.doe", "F123456", model.Rule{}, testBot}, true},
		{"User is allowed", args{"jane.doe", "F123456", model.Rule{AllowUsers: []string{"john.doe", "jane.doe"}}, testBot}, true},
		{"User not allowed", args{"jane.doe", "F123456", model.Rule{AllowUsers: []string{"john.doe", "jack.jill"}}, testBot}, false},
		{"User is ignored", args{"jane.doe", "F123456", model.Rule{IgnoreUsers: []string{"jane.doe", "jack.jill"}}, testBot}, false},
		{"User not in ignore list", args{"jane.doe", "F123456", model.Rule{IgnoreUsers: []string{"john.doe", "jack.jill"}}, testBot}, true},
		{"User is allowed but ignored", args{"jane.doe", "F123456", model.Rule{AllowUsers: []string{"jane.doe"}, IgnoreUsers: []string{"jane.doe", "jack.jill"}}, testBot}, false},
		{"User is not allowed and ignored", args{"john.doe", "F123456", model.Rule{AllowUsers: []string{"jane.doe"}, IgnoreUsers: []string{"john.doe", "jack.jill"}}, testBot}, false},
		{"Group - Workspace Token not supplied", args{"jane.doe", "F123456", model.Rule{AllowUserGroups: []string{"admins"}}, testBot}, false},
		{"Group - Discord - Not supported", args{"jane.doe", "F123456", model.Rule{AllowUserGroups: []string{"admins"}}, discordBot}, false},
		{"Group - Chat network not supported", args{"jane.doe", "F123456", model.Rule{AllowUserGroups: []string{"admins"}}, strangeBot}, false},
		// TODO: figure out how to test this below:
		// {"User in allow group but ignored", args{"jane.doe", "F123456", model.Rule{}, testBot}, false},
		// {"User in ignore group but allowed", args{"jane.doe", "F123456", model.Rule{}, testBot}, false},
		// {"User in ignore group and allow group", args{"jane.doe", "F123456", model.Rule{}, testBot}, false},
		// {"User in allow group and not ignored", args{"jane.doe", "F123456", model.Rule{}, testBot}, true},
		// {"User is not in allow group and not ignored", args{"jane.doe", "F123456", model.Rule{}, testBot}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanTrigger(tt.args.currentUserName, tt.args.currentUserID, tt.args.rule, tt.args.bot); got != tt.want {
				t.Errorf("CanTrigger() = %v, want %v", got, tt.want)
			}
		})
	}
}
