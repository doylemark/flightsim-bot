package discord

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/doylemark/flightsim-discord-bot/commands/botcommands/ping"
)

// MessageCreate Handles all incoming messages to the bot and routes them to commands
func MessageCreate(Session *discordgo.Session, message *discordgo.MessageCreate) {
	fmt.Println(message.Content)

	if message.Author.ID == Session.State.User.ID {
		return
	}

	if !strings.HasPrefix(message.Content, "!") {
		return
	} 

	if hasCmd(message.Content, "!ping") {
		ping.Ping(Session, message)
	}

	
}

func hasCmd(message string, cmd string) bool {
	if strings.HasPrefix(message, cmd) {
		return true
	}
	return false
}