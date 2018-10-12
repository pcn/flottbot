package model

// Action defines the structure for Actions used within Rules
type Action struct {
	Name             string            `mapstructure:"name"`
	Type             ActionType        `mapstructure:"type" binding:"required"`
	Timeout          int               `mapstructure:"timeout"`
	ExposeJSONFields map[string]string `mapstructure:"expose_json_fields"`
	Response         string            `mapstructure:"response"`
	LimitToChannels  []string          `mapstructure:"limit_to_channels"`
	Message          string            `mapstructure:"message"`
	Reaction         string            `mapstructure:"update_reaction" binding:"omitempty"` // TODO: make its own struct
	ActionHTTP
	ActionExec
}

// ActionType is used to differentiate between different message types
type ActionType int

// Supported MessageTypes
// TODO: Consider adding Reaction
const (
	ActionTypeHTTP MessageType = iota
	ActionTypeExec
	ActionTypeMsg
	ActionTypeLog
)

// ActionHTTP are fields needed for a HTTP actions
type ActionHTTP struct {
	URL           string                 `mapstructure:"url"`
	Auth          Auth                   `mapstructure:"auth"`
	Method        string                 `mapstructure:"method"`
	QueryData     map[string]interface{} `mapstructure:"query_data"`
	CustomHeaders map[string]string      `mapstructure:"custom_headers"`
}

// ActionExec are fields needed for Exec actions
type ActionExec struct {
	Cmd string `mapstructure:"cmd"`
}

// Auth is a basic Auth data structure
type Auth struct {
	Type  string `mapstructure:"type"`
	User  string `mapstructure:"user"`
	Pass  string `mapstructure:"pass"`
	Token string `mapstructure:"token"`
}
