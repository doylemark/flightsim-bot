package main

// import (
// 	"fmt"
// 	"sort"
// 	"strconv"
// 	"time"

// 	"github.com/bwmarrin/discordgo"
// )

// type reaction struct {
// 	URL   string
// 	UID   string
// 	Count int
// }

// func getReactionCounts(session *discordgo.Session, posts []*post) ([]*reaction, error) {

// 	var reactions []*reaction

// 	for _, post := range posts {
// 		reactedUsers, err := session.MessageReactions(ScreenshotChannelID, post.ID, "✅", 100)

// 		if err != nil {
// 			fmt.Println("Submission appears to have been deleted:", err)
// 		}

// 		messageResult := reaction{post.URL, post.UID, len(reactedUsers)}

// 		reactions = append(reactions, &messageResult)
// 	}
// 	return reactions, nil
// }

// func getWinners(reactions []*reaction) []*reaction {
// 	sort.Slice(reactions[:], func(i, j int) bool {
// 		return reactions[j].Count < reactions[i].Count
// 	})

// 	return reactions
// }

// func postWinners(winners []*reaction, session *discordgo.Session) {

// 	var embedFields []*discordgo.MessageEmbedField

// 	for i := 0; i < 5; i++ {
// 		// add roles to winners
// 		winnersRoleID := "549990190794407937"
// 		session.GuildMemberRoleAdd(GuildID, winners[i].UID, winnersRoleID)

// 		embed := discordgo.MessageEmbedField{
// 			Value: "<@" + winners[i].UID + "> with **" + strconv.Itoa(winners[i].Count) + "** votes ⭐️",
// 			Name:  "✈️ **" + strconv.Itoa(i+1) + "**",
// 		}

// 		embedFields = append(embedFields, &embed)
// 	}

// 	embed := &discordgo.MessageEmbed{
// 		Author:    &discordgo.MessageEmbedAuthor{},
// 		Title:     "Screenshot Competition Results",
// 		Fields:    embedFields,
// 		Timestamp: time.Now().Format(time.RFC3339),
// 	}

// 	_, err := session.ChannelMessageSendEmbed(ScreenshotChannelID, embed)

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

// func removeRoles(prevWinners []*reaction, session *discordgo.Session) {
// 	for _, prevWinner := range prevWinners {
// 		session.GuildMemberRoleRemove(GuildID, prevWinner.UID, "549990190794407937")
// 	}
// }
