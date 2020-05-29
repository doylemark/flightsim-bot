package discord

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/doylemark/flightsim-bot/commands/botcommands/ping"
	"github.com/doylemark/flightsim-bot/commands/botcommands/submit"
	"github.com/doylemark/flightsim-bot/commands/botcommands/tally"
	"github.com/doylemark/flightsim-bot/commands/database"
)

// MessageCreate Handles all incoming messages to the bot and routes them to commands
func MessageCreate(Session *discordgo.Session, message *discordgo.MessageCreate) {

	if message.Author.ID == Session.State.User.ID {
		return
	}

	if !strings.HasPrefix(message.Content, "!") {
		return
	}

	if hasCmd(message.Content, "!ping") {
		ping.Ping(Session, message)
	}

	if hasCmd(message.Content, "!submit") {
		submit.Submit(Session, message)
	}

	if hasCmd(message.Content, "!tally") {
		if !checkPerms(Session, message) {
			return
		}
		tally.Tally(Session, message)
	}

	if hasCmd(message.Content, "!start") {
		if !checkPerms(Session, message) {
			return
		}
		database.RemoveCompetition()
	}

}

func hasCmd(message string, cmd string) bool {
	if strings.HasPrefix(message, cmd) {
		return true
	}
	return false
}

func checkPerms(Session *discordgo.Session, message *discordgo.MessageCreate) bool {
	hasPerms, _ := MemberHasPermission(Session, message.Author.ID, 8)

	if hasPerms {
		fmt.Println("User has perms")
		return true
	}

	Session.ChannelMessageSend(message.ChannelID, "You do not have permission to perform that operation")
	return false
}
