package model

import "github.com/sirupsen/logrus"

// Bot is a struct representation of bot.yml
type Bot struct {
	// Bot fields
	User            *BotUser
	Rules           map[string]*Rule
	IncomingMessage chan<- *Message
	OutgoingMessage <-chan *Message
	Remotes         *map[string]map[string]interface{} `mapstructure:"remotes" binding:"required"`
	Debug           bool                               `mapstructure:"debug"`
	LogJSON         bool                               `mapstructure:"log_json"`
	Metrics         bool                               `mapstructure:"metrics"`
	HelpIntroText   string                             `mapstructure:"help_intro_text,omitempty"`
	AllowUnsafeExec bool                               `mapstructure:"allow_unsafe_exec"`
	// System
	Log logrus.Logger
}

// Configure ...
func (b *Bot) Configure() {
	return
}
