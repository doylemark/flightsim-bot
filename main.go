package main

import (
	"github.com/doylemark/flightsim-bot/commands/database"
	"github.com/doylemark/flightsim-bot/commands/discord"
	"github.com/doylemark/flightsim-bot/commands/env"
)

// ScreenshotChannelID The channel the bot will send messages to
var ScreenshotChannelID string

// GuildID The guild id of the server the screenshot competitions will be taking place in
var GuildID string

func main() {
	env.LoadEnvVars()
	database.Connect()
	discord.Connect(env.DiscordToken)
}
