package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// ScreenshotChannelID is the ID of the channel which the bot should send submissions to
var ScreenshotChannelID string

// GuildID is the ID of the guild the bot should send submissions to
var GuildID string

// DiscordToken is the Discord Bot Token used for authorizing API requests
var DiscordToken string

// WxToken is the token for the AVWX API, only required if !metar command will be used
var WxToken string

// WinnerRoleID is the ID of the Discord role that will be assigned to the winners of the screenshot competitions
var WinnerRoleID string

// LoadEnvVars loads env variables from .env and exposes them to other packages
func LoadEnvVars() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ScreenshotChannelID = os.Getenv("SCREENSHOT_CHANNEL_ID")
	GuildID = os.Getenv("GUILD_ID")
	DiscordToken = os.Getenv("DISCORD_BOT_TOKEN")
	WxToken = os.Getenv("WX_API_TOKEN")
	WinnerRoleID = os.Getenv("WINNER_ROLE_ID")
}
