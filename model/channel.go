package model

// Channel is a channel
type Channel struct {
	ID        string `mapstructure:"id,omitempty"`
	Name      string `mapstructure:"name,omitempty"`
	IsPrivate bool   `mapstructure:"is_private,omitempty"`
	Type      string `mapstructure:"type,omitempty"`
}

// Channels is a collection of channels
type Channels []Channel
