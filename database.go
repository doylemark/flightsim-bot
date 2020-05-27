package main

import (
	"context"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var winnersCollection *mongo.Collection
var ctx = context.TODO()

func connectDb() {
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

	collection = client.Database("flightsim").Collection("screenshots")
	winnersCollection = client.Database("flightsim").Collection("winners")
}

func saveSubmission(id string, url string, author string) {
	_, err := collection.InsertOne(context.TODO(), bson.D{
		{Key: "url", Value: url},
		{Key: "id", Value: id},
		{Key: "uid", Value: author},
	})

	if err != nil {
		fmt.Println(err)
		return
	}
}

func removeCompetition() {
	_, err := collection.DeleteMany(context.TODO(), bson.D{{}})
	_, err = winnersCollection.DeleteMany(context.TODO(), bson.D{{}})

	if err != nil {
		fmt.Println(err)
		return
	}
}

type post struct {
	URL string `json:"url"`
	ID  string `json:"id"`
	UID string `json:"uid"`
}

func getEntries(session *discordgo.Session) ([]*post, error) {
	result, err := collection.Find(context.TODO(), bson.D{{}}, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var posts []*post

	// get all posts in database
	for result.Next(context.TODO()) {
		var post post
		err = result.Decode(&post)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		posts = append(posts, &post)
	}

	return posts, nil

}

func storeWinners(winners []*reaction) {

	for i := 0; i < 5; i++ {
		_, err := winnersCollection.InsertOne(context.TODO(), bson.D{
			{Key: "uid", Value: winners[i].UID},
		})

		if err != nil {
			fmt.Println(err)
			return
		}
	}

}

func getPrevWinners(session *discordgo.Session) ([]*reaction, error) {
	result, err := winnersCollection.Find(context.TODO(), bson.D{{}}, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var prevWinners []*reaction

	// get all posts in database
	for result.Next(context.TODO()) {
		var winner reaction
		err = result.Decode(&winner)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		prevWinners = append(prevWinners, &winner)
	}

	return prevWinners, nil
}
