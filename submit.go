package main

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func sendSubmit(url string, session *discordgo.Session, message *discordgo.MessageCreate) {
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
	sentMsg, err := session.ChannelMessageSendEmbed(ScreenshotChannelID, embed)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = session.MessageReactionAdd(ScreenshotChannelID, sentMsg.ID, "✅")

	if err != nil {
		fmt.Println(err)
	}

	saveSubmission(sentMsg.ID, url, message.Author.ID)
}
