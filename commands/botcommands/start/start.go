package start

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/doylemark/flightsim-bot/commands/env"
)

// Start ends the current competition and begins a new one by cleaning the db
func Start(session *discordgo.Session, message *discordgo.MessageCreate) {
	session.ChannelMessageDelete(message.ChannelID, message.ID)
	startCompetitionMessage(session)
}

func startCompetitionMessage(session *discordgo.Session) {
	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{},
		Title:  "New Screenshot Competition",
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name: "Announcement",
				Value: `You may submit by going to <#230761348831641600> and typing !submit, you must attach an
								image/include an image link with this message 
								`,
			},
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	_, err := session.ChannelMessageSendEmbed(env.ScreenshotChannelID, embed)

	if err != nil {
		fmt.Println(err)
		return
	}
}
