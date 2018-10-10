package core

import (
	"strings"

	"github.com/target/flottbot/model"
	"github.com/target/flottbot/remote/cli"
	"github.com/target/flottbot/remote/discord"
	"github.com/target/flottbot/remote/slack"
)

// Outputs determines where messages are output based on fields set in the bot.yml
// TODO: Refactor to keep remote specifics in remote/
func Outputs(outputMsgs <-chan model.Message, hitRule <-chan model.Rule, bot *model.Bot) {
	remoteCLI := &cli.Client{}
	remoteDiscord := &discord.Client{}
	remoteSlack := &slack.Client{}
	for {
		message := <-outputMsgs
		rule := <-hitRule
		service := message.Service
		switch service {
		case model.MsgServiceChat, model.MsgServiceScheduler:
			chatApp := strings.ToLower(bot.ChatApplication)
			switch chatApp {
			case "discord":
				if service == model.MsgServiceScheduler {
					bot.Log.Warn("Scheduler does not currently support Discord")
					break
				}
				remoteDiscord = &discord.Client{Token: bot.DiscordToken}
				remoteDiscord.Send(message, bot)
			case "slack":
				// Create Slack client
				remoteSlack = &slack.Client{
					Token:             bot.SlackToken,
					VerificationToken: bot.SlackVerificationToken,
					WorkspaceToken:    bot.SlackWorkspaceToken,
				}
				if service == model.MsgServiceChat {
					if bot.InteractiveComponents {
						remoteSlack.InteractiveComponents(nil, &message, rule, bot)
					}
					remoteSlack.Reaction(message, rule, bot)
				}
				remoteSlack.Send(message, bot)
			default:
				bot.Log.Debugf("Chat application %s is not supported", chatApp)
			}
		case model.MsgServiceCLI:
			remoteCLI.Send(message, bot)
		case model.MsgServiceUnknown:
			bot.Log.Error("Found unknown service")
		default:
			bot.Log.Errorf("No service found")
		}
	}
}
