package model

// BotUser contains information about the Flottbot User
type BotUser struct {
	ID       string `mapstructure:"id,omitempty"`
	Name     string `mapstructure:"name,omitempty"`
	Email    string `mapstructure:"email,omitempty"`
	Version  string `mapstructure:"version,omitempty"`
	Type     string `mapstructure:"type,omitempty"`
	Channels Channels
}
