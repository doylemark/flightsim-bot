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

func startCompetitionMessage(session *discordgo.Session) {
	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{},
		Title:  "New Screenshot Competition",
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name: "Announcement",
				Value: `A new screenshot competition has just begun, previous winner roles have been removed. 
								 You may submit by going to <#230761348831641600> and typing !submit, you must attach an
								 image/include an image link with this message 
								`,
			},
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	_, err := session.ChannelMessageSendEmbed(ScreenshotChannelID, embed)

	if err != nil {
		fmt.Println(err)
		return
	}
}
