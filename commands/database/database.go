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

// Connect Connects app to the database
func Connect() {
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
