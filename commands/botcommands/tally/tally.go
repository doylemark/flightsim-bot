package tally

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/doylemark/flightsim-bot/commands/database"
	"github.com/doylemark/flightsim-bot/commands/env"
)

// Tally gets all submissions from database, calculates winners and posts to the GuildChannelID
func Tally(session *discordgo.Session, message *discordgo.MessageCreate) {
	session.ChannelMessageDelete(message.ChannelID, message.ID)
	prevWinners, _ := database.GetPrevWinners(session)
	if len(prevWinners) > 0 {
		removeRoles(prevWinners, session)
	}

	entries, _ := database.GetEntries(session)
	reactions, _ := getReactionCounts(session, entries)
	sorted := getWinners(reactions)
	if len(sorted) > 0 {
		notifyWinners(sorted, session)
	} else {
		session.ChannelMessageSend(env.ScreenshotChannelID, "Failed to compute competition results as no enteries could be located in the database")
	}

	database.StoreWinners(sorted)
}

func getReactionCounts(session *discordgo.Session, posts []*database.Post) ([]*database.Reaction, error) {
	var reactions []*database.Reaction

	for _, post := range posts {
		reactedUsers, err := session.MessageReactions(env.ScreenshotChannelID, post.ID, "✅", 100)

		if err != nil {
			fmt.Println("Submission appears to have been deleted:", err)
		}

		messageResult := database.Reaction{URL: post.URL, UID: post.UID, Count: len(reactedUsers)}

		reactions = append(reactions, &messageResult)
	}
	return reactions, nil
}

func getWinners(reactions []*database.Reaction) []*database.Reaction {
	sort.Slice(reactions[:], func(i, j int) bool {
		return reactions[j].Count < reactions[i].Count
	})

	return reactions
}

// notifyWinners adds role to winner members and notifies them in Screenshot Channel
func notifyWinners(winners []*database.Reaction, session *discordgo.Session) {
	var embedFields []*discordgo.MessageEmbedField

	var winnersCount int

	if len(winners) > 5 {
		winnersCount = 5
	} else {
		winnersCount = len(winners)
	}

	for i := 0; i < winnersCount; i++ {
		if i > len(winners)-1 {
			break
		}
		// add roles to winners
		session.GuildMemberRoleAdd(env.GuildID, winners[i].UID, env.WinnerRoleID)

		embed := discordgo.MessageEmbedField{
			Value: "<@" + winners[i].UID + "> with **" + strconv.Itoa(winners[i].Count) + "** votes ⭐️",
			Name:  "✈️ **" + strconv.Itoa(i+1) + "**",
		}

		embedFields = append(embedFields, &embed)
	}

	embed := &discordgo.MessageEmbed{
		Author:    &discordgo.MessageEmbedAuthor{},
		Title:     "Screenshot Competition Results",
		Fields:    embedFields,
		Timestamp: time.Now().Format(time.RFC3339),
	}

	_, err := session.ChannelMessageSendEmbed(env.ScreenshotChannelID, embed)

	if err != nil {
		fmt.Println(err)
	}
}

func removeRoles(prevWinners []*database.Reaction, session *discordgo.Session) {
	for _, prevWinner := range prevWinners {
		session.GuildMemberRoleRemove(env.GuildID, prevWinner.UID, env.WinnerRoleID)
	}
}
