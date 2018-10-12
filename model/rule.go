package model

// Rule is a struct representation of the .yml rules
type Rule struct {
	Actions            []Action `mapstructure:"actions" binding:"required"`
	ExposeArgs         []string `mapstructure:"expose_args" binding:"required"`
	OutputToChannels   []string `mapstructure:"output_to_channels" binding:"omitempty"`
	OutputToUsers      []string `mapstructure:"output_to_users" binding:"omitempty"`
	AllowUsers         []string `mapstructure:"allow_users" binding:"omitempty"`
	AllowUserGroups    []string `mapstructure:"allow_usergroups" binding:"omitempty"`
	IgnoreUsers        []string `mapstructure:"ignore_users" binding:"omitempty"`
	IgnoreUserGroups   []string `mapstructure:"ignore_usergroups" binding:"omitempty"`
	Name               string   `mapstructure:"name" binding:"required"`
	Trigger            string   `mapstructure:"trigger" binding:"omitempty"` // TODO: denote regex usage in some way
	Schedule           string   `mapstructure:"schedule"`
	FormatOutput       string   `mapstructure:"format_output"`
	Reaction           string   `mapstructure:"reaction" binding:"omitempty"` // TODO: move to action
	HelpText           string   `mapstructure:"help_text"`
	DirectMessageOnly  bool     `mapstructure:"direct_message_only" binding:"required"`
	StartMessageThread bool     `mapstructure:"start_message_thread" binding:"omitempty"`
	IncludeInHelp      bool     `mapstructure:"include_in_help" binding:"required"`
	Active             bool     `mapstructure:"active" binding:"required"`
	Debug              bool     `mapstructure:"debug" binding:"required"`
	Vars               map[string]string
	RemoveReaction     string // TODO: move to action?
}
