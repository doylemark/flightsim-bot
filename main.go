package main

import (
	"github.com/doylemark/flightsim-bot/commands/discord"
	"github.com/doylemark/flightsim-bot/commands/env"
)

func main() {
	env.LoadEnvVars()
	discord.Connect(env.DiscordToken)
}
