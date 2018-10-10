package model

// Message is the struct of the main data structure being passed around for each message generated
type Message struct {
	Rule            *Rule
	Bot             *Bot
	Type            MessageType
	Remote          string
	ChannelID       string
	ChannelName     string
	Input           string
	Output          string
	ErrorOutput     string
	Timestamp       string
	ThreadTimestamp string
	StartTime       int64
	EndTime         int64
	IsError         bool
	BotMentioned    bool
}

// MessageType is used to differentiate between different message types
type MessageType int

// Supported MessageTypes
const (
	MsgTypeUnknown MessageType = iota
	MsgTypeDirect
	MsgTypeChannel
	MsgTypePrivateChannel
)
