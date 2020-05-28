package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/doylemark/flightsim-discord-bot/commands/database"
	"github.com/doylemark/flightsim-discord-bot/commands/discord"
)

// ScreenshotChannelID The channel the bot will send messages to
var ScreenshotChannelID string

// GuildID The guild id of the server the screenshot competitions will be taking place in
var GuildID string

func main() {
	database.ConnectDb()
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ScreenshotChannelID = os.Getenv("SCREENSHOT_CHANNEL_ID")
	GuildID = os.Getenv("GUILD_ID")
	Token := os.Getenv("DISCORD_BOT_TOKEN")

	discord.Connect(Token)


}

// func messageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
// 	if message.Author.ID == session.State.User.ID {
// 		return
// 	}

// 	if message.Content == "!ping" {
// 		msgTime, err := time.Parse(time.RFC3339, string(message.Timestamp))
// 		responseTime := time.Now().Sub(msgTime)

// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
	
// 		session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("🏓 Pong! response time was %s 😍", responseTime.String()))
// 	}

// 	if strings.HasPrefix(message.Content, "!submit") {
// 		if len(message.Attachments) > 0 {

// 			sendSubmit(message.Attachments[0].URL, session, message)

// 		} else if len(strings.Fields(message.Content)) > 1 {

// 			if len(strings.Fields(message.Content)) > 2 {
// 				session.ChannelMessageSend(message.ChannelID, "You may only attach one link!")
// 				return
// 			}

// 			sendSubmit(strings.Fields(message.Content)[1], session, message)

// 		} else {
// 			session.ChannelMessageSend(message.ChannelID, "You must provide an image URL or attach an image to your submission")
// 		}
// 	}

// 	if strings.HasPrefix(message.Content, "!tally") {
		
// 		hasPermissions, _ := memberHasPermission(session, GuildID, message.Author.ID, 8)

// 		if !hasPermissions {
// 			session.ChannelMessageSend(message.ChannelID, "You do not have the required permissions to perform this command")
// 			return
// 		}

// 		session.ChannelMessageDelete(message.ChannelID, message.ID)
// 		entries, _ := getEntries(session)
// 		reactions, _ := getReactionCounts(session, entries)
// 		sorted := getWinners(reactions)
// 		if len(sorted) > 0 {
// 			postWinners(sorted, session)
// 		} else {
// 			session.ChannelMessageSend(ScreenshotChannelID, "Failed to compute competition results as no enteries could be located in the database")
// 		}

// 		storeWinners(sorted)
// 	}

// 	if strings.HasPrefix(message.Content, "!start") {

// 		hasPermissions, _ := memberHasPermission(session, GuildID, message.Author.ID, 8)

// 		if !hasPermissions {
// 			session.ChannelMessageSend(message.ChannelID, "You do not have the required permissions to perform this command")
// 			return
// 		}


// 		session.ChannelMessageDelete(message.ChannelID, message.ID)
// 		// removeCompetition()
// 		prevWinners, _ := getPrevWinners(session)
// 		if len(prevWinners) > 0 {
// 			removeRoles(prevWinners, session)
// 		}
// 		startCompetitionMessage(session)
// 	}
// }
