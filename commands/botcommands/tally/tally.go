package tally

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/doylemark/flightsim-bot/commands/env"
	"github.com/doylemark/flightsim-bot/commands/store"
)

// Tally gets all submissions from database, calculates winners and posts to the GuildChannelID
func Tally(session *discordgo.Session, message *discordgo.MessageCreate) {
	session.ChannelMessageDelete(message.ChannelID, message.ID)

	entries := store.Submissions
	reactions, _ := getReactionCounts(session, entries)
	sorted := getWinners(reactions)

	if len(sorted) > 0 {
		notifyWinners(sorted, session)
	} else {
		session.ChannelMessageSend(env.ScreenshotChannelID, "Failed to compute competition results as no enteries could be located in the database")
	}
}

type reaction struct {
	URL   string
	UID   string
	Count int
}

func getReactionCounts(session *discordgo.Session, posts []*store.Submission) ([]*reaction, error) {
	var reactions []*reaction

	for _, post := range posts {
		reactedUsers, err := session.MessageReactions(env.ScreenshotChannelID, post.ID, "✅", 100, "", "")

		if err != nil {
			fmt.Println("Submission appears to have been deleted:", err)
		}

		messageResult := reaction{URL: post.URL, UID: post.UID, Count: len(reactedUsers)}

		reactions = append(reactions, &messageResult)
	}
	return reactions, nil
}

func getWinners(reactions []*reaction) []*reaction {
	sort.Slice(reactions[:], func(i, j int) bool {
		return reactions[j].Count < reactions[i].Count
	})

	return reactions
}

// notifyWinners adds role to winner members and notifies them in Screenshot Channel
func notifyWinners(winners []*reaction, session *discordgo.Session) {
	message := "**Screenshot Competition Results** \n"

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

		description := "<@" + winners[i].UID + "> with **" + strconv.Itoa(winners[i].Count) + "** votes ⭐️ \n"

		message = message + description
	}

	_, err := session.ChannelMessageSend(env.ScreenshotChannelID, message)

	if err != nil {
		fmt.Println(err)
	}
}
