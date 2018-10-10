package remote

import (
	"context"

	"github.com/target/flottbot/model"
)

// Remote - this interface allows us to keep the bot "remote agnostic" meaning
// that the bot should not care about what specific remote (e.g. Slack or Discord)
// it is reading/sending messages from. It is up to the developer to implement
// the details of how messages should be read/sent in the specific package for
// the remote (e.g. see '/remote/slack/remote.go')
// Each remote will share generic functions as seen below that will be evoked
// by the bot (e.g. see '/core/remotes.go' or '/core/outputs.go'), and each function
// will be implemented in detail within its specific remote package.
type Remote interface {
	Login() (*model.BotUser, error)

	Channels() (*model.Channels, error)

	Reaction(message model.Message, rule model.Rule, bot *model.Bot)

	Read(inputMsgs chan<- model.Message, rules map[string]model.Rule, bot *model.Bot)

	Send(message model.Message, bot *model.Bot)

	// InteractiveComponents(inputMsgs chan<- model.Message, message *model.Message, rule model.Rule, bot *model.Bot)
}

// Reaction enables the bot to add emoji reactions to messages
func Reaction(c context.Context, message model.Message, rule model.Rule, bot *model.Bot) {
	FromContext(c).Reaction(message, rule, bot)
}

func Channels(c context.Context) (*model.Channels, error) {
	return FromContext(c).Channels()
}

func Login(c context.Context) (*model.BotUser, error) {
	return FromContext(c).Login()
}

// Read enables the bot to read messages from a remote
func Read(c context.Context, inputMsgs chan<- model.Message, rules map[string]model.Rule, bot *model.Bot) {
	FromContext(c).Read(inputMsgs, rules, bot)
}

// Send enables the bot to send messages to a remote
func Send(c context.Context, message model.Message, bot *model.Bot) {
	FromContext(c).Send(message, bot)
}

// InteractiveComponents enables the bot to listen to Interactive Components coming from a remote
// func InteractiveComponents(c context.Context, inputMsgs chan<- model.Message, message *model.Message, rule model.Rule, bot *model.Bot) {
// 	FromContext(c).InteractiveComponents(inputMsgs, message, rule, bot)
// }
