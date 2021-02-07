package submit

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/doylemark/flightsim-bot/commands/env"
	"github.com/doylemark/flightsim-bot/commands/store"
)

// Submit submits a screenshot to the competition
func Submit(session *discordgo.Session, message *discordgo.MessageCreate) {
	if len(message.Attachments) > 0 {
		addSubmission(message.Attachments[0].URL, session, message)
	} else if len(strings.Fields(message.Content)) > 1 {
		if len(strings.Fields(message.Content)) > 2 {
			session.ChannelMessageSend(message.ChannelID, "You may only attach one link!")
			return
		}

		addSubmission(strings.Fields(message.Content)[1], session, message)
	} else {
		session.ChannelMessageSend(message.ChannelID, "You must provide an image URL or attach an image to your submission")
	}
}

func addSubmission(url string, session *discordgo.Session, message *discordgo.MessageCreate) {
	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{},
		Title:  "Screenshot Competition Entry",
		Image: &discordgo.MessageEmbedImage{
			URL: url,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text:    message.Author.Username,
			IconURL: "https://cdn.discordapp.com/avatars/" + message.Author.ID + "/" + message.Author.Avatar,
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}
	sentMsg, err := session.ChannelMessageSendEmbed(env.ScreenshotChannelID, embed)

	session.ChannelMessageSend(message.ChannelID, "Your submission has been recorded")

	if err != nil {
		fmt.Println(err)
		return
	}

	err = session.MessageReactionAdd(env.ScreenshotChannelID, sentMsg.ID, "✅")

	if err != nil {
		fmt.Println(err)
	}

	store.AddSubmission(&store.Submission{URL: url, ID: sentMsg.ID, UID: message.Author.ID})
}
