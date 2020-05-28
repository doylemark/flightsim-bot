package database

import (
	"context"
	"fmt"
	"time"

	// "github.com/bwmarrin/discordgo"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collection is the standard database collection for storing of all competition entries
var Collection *mongo.Collection

// WinnersCollection is the collection for storing the previous competitions winners. This should only ever have 5 documents
var WinnersCollection *mongo.Collection

// ConnectDb Connects app to the database
func ConnectDb() {
	fmt.Println("Attempting to connect to DB!")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/vatsim-backend"))

	if err != nil {
		fmt.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Connected to DB!")

	// defer client.Disconnect(ctx)

	Collection = client.Database("flightsim").Collection("screenshots")
	WinnersCollection = client.Database("flightsim").Collection("winners")
}

// var ctx = context.TODO()

// func removeCompetition() {
// 	_, err := Collection.DeleteMany(context.TODO(), bson.D{{}})
// 	_, err = WinnersCollection.DeleteMany(context.TODO(), bson.D{{}})

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// }

// type post struct {
// 	URL string `json:"url"`
// 	ID  string `json:"id"`
// 	UID string `json:"uid"`
// }

// func getEntries(session *discordgo.Session) ([]*post, error) {
// 	result, err := Collection.Find(context.TODO(), bson.D{{}}, nil)

// 	if err != nil {
// 		fmt.Println(err)
// 		return nil, err
// 	}

// 	var posts []*post

// 	// get all posts in database
// 	for result.Next(context.TODO()) {
// 		var post post
// 		err = result.Decode(&post)

// 		if err != nil {
// 			fmt.Println(err)
// 			return nil, err
// 		}
// 		posts = append(posts, &post)
// 	}

// 	return posts, nil

// }

// func storeWinners(winners []*reaction) {
// 	for i := 0; i < 5; i++ {
// 		if i > len(winners) {
// 			break
// 		}

// 		_, err := WinnersCollection.InsertOne(context.TODO(), bson.D{
// 			{Key: "uid", Value: winners[i].UID},
// 		})

// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 	}

// }

// func getPrevWinners(session *discordgo.Session) ([]*reaction, error) {
// 	result, err := WinnersCollection.Find(context.TODO(), bson.D{{}}, nil)

// 	if err != nil {
// 		fmt.Println(err)
// 		return nil, err
// 	}

// 	var prevWinners []*reaction

// 	for result.Next(context.TODO()) {
// 		var winner reaction
// 		err = result.Decode(&winner)

// 		if err != nil {
// 			fmt.Println(err)
// 			return nil, err
// 		}
// 		prevWinners = append(prevWinners, &winner)
// 	}

// 	return prevWinners, nil
// }
