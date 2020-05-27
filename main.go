package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

// ScreenshotChannelID The channel the bot will send messages to
const ScreenshotChannelID = "550043127696588829"

// GuildID The guild id of the server the screenshot competitions will be taking place in
const GuildID = "216217565272211456"

func main() {
	connectDb()
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	token := os.Getenv("DISCORD_BOT_TOKEN")

	Session, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println(err)
		return
	}

	Session.AddHandler(messageCreate)

	err = Session.Open()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Bot is now online, press CTRL+C to terminate the process!")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	Session.Close()
}

func messageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}

	if message.Content == "!ping" {
		msgTime, err := time.Parse(time.RFC3339, string(message.Timestamp))
		responseTime := time.Now().Sub(msgTime)

		if err != nil {
			fmt.Println(err)
			return
		}

		session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("🏓 Pong! response time was %s 😍", responseTime.String()))
	}

	if strings.HasPrefix(message.Content, "!submit") {
		session.ChannelMessageDelete(message.ChannelID, message.ID)
		if len(message.Attachments) > 0 {

			sendSubmit(message.Attachments[0].URL, session, message)

		} else if len(strings.Fields(message.Content)) > 1 {

			if len(strings.Fields(message.Content)) > 2 {
				session.ChannelMessageSend(message.ChannelID, "You may only attach one link!")
				return
			}

			sendSubmit(strings.Fields(message.Content)[1], session, message)

		} else {
			session.ChannelMessageSend(message.ChannelID, "You must provide an image URL or attach an image to your submission")
		}
	}

	if strings.HasPrefix(message.Content, "!tally") {
		session.ChannelMessageDelete(message.ChannelID, message.ID)
		entries, _ := getEntries(session)
		reactions, _ := getReactionCounts(session, entries)
		sorted := getWinners(reactions)
		if len(sorted) > 0 {
			postWinners(sorted, session)
		} else {
			session.ChannelMessageSend(ScreenshotChannelID, "Failed to compute competition results as no enteries could be located in the database")
		}

		storeWinners(sorted)
	}

	if strings.HasPrefix(message.Content, "!start") {
		session.ChannelMessageDelete(message.ChannelID, message.ID)
		// removeCompetition()
		prevWinners, _ := getPrevWinners(session)
		removeRoles(prevWinners, session)
		startCompetitionMessage(session)
	}
}
