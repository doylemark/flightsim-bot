package database

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
)

// Post defines an entry
type Post struct {
	URL string `json:"url"`
	ID  string `json:"id"`
	UID string `json:"uid"`
}

// GetEntries gets all submissions to the competition
func GetEntries(session *discordgo.Session) ([]*Post, error) {
	result, err := Collection.Find(context.TODO(), bson.D{{}}, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var posts []*Post

	// get all posts in database
	for result.Next(context.TODO()) {
		var post Post
		err = result.Decode(&post)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		posts = append(posts, &post)
	}

	return posts, nil

}

// Reaction is the props returned when checking reaction counts on discord api
type Reaction struct {
	URL   string
	UID   string
	Count int
}

// StoreWinners adds the top 5 winners to the database for future reference
func StoreWinners(winners []*Reaction) {

	var winnersCount int

	if len(winners) > 5 {
		winnersCount = 5
	} else {
		winnersCount = len(winners)
	}

	for i := 0; i < winnersCount; i++ {
		_, err := WinnersCollection.InsertOne(context.TODO(), bson.D{
			{Key: "uid", Value: winners[i].UID},
		})

		if err != nil {
			fmt.Println(err)
			return
		}
	}

}

// GetPrevWinners fetches a list of last competitions winners from the db
func GetPrevWinners(session *discordgo.Session) ([]*Reaction, error) {
	result, err := WinnersCollection.Find(context.TODO(), bson.D{{}}, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var prevWinners []*Reaction

	for result.Next(context.TODO()) {
		var winner Reaction
		err = result.Decode(&winner)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		prevWinners = append(prevWinners, &winner)
	}

	return prevWinners, nil
}
