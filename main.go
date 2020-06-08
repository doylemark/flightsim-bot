package main

import (
	"github.com/doylemark/flightsim-bot/commands/database"
	"github.com/doylemark/flightsim-bot/commands/discord"
	"github.com/doylemark/flightsim-bot/commands/env"
)

func main() {
	env.LoadEnvVars()
	database.Connect()
	discord.Connect(env.DiscordToken)
}
