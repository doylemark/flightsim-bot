package start

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/doylemark/flightsim-bot/commands/database"
	"github.com/doylemark/flightsim-bot/commands/env"
)

// Start ends the current competition and begins a new one by cleaning the db
func Start(session *discordgo.Session, message *discordgo.MessageCreate) {
	session.ChannelMessageDelete(message.ChannelID, message.ID)
	database.RemoveCompetition()
	startCompetitionMessage(session)
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

	_, err := session.ChannelMessageSendEmbed(env.ScreenshotChannelID, embed)

	if err != nil {
		fmt.Println(err)
		return
	}
}
