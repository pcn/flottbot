package model

import "github.com/sirupsen/logrus"

// Bot is a struct representation of bot.yml
type Bot struct {
	// Bot fields
	User           *BotUser
	Remotes        map[string]map[string]interface{} `mapstructure:"remotes" binding:"required"`
	Debug          bool                              `mapstructure:"debug,omitempty"`
	LogJSON        bool                              `mapstructure:"log_json,omitempty"`
	Metrics        bool                              `mapstructure:"metrics,omitempty"`
	CustomHelpText string                            `mapstructure:"custom_help_text,omitempty"`
	// System
	Log logrus.Logger
}
