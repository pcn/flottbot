package model

// BotUser contains information about the Flottbot User
type BotUser struct {
	ID       string
	Name     string `mapstructure:"name,omitempty"`
	Email    string
	Version  string
	Type     string
	Channels []Channel
}
