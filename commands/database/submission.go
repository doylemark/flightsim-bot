package database

import (
	"fmt"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

// SaveSubmission Saves a competition entry to the database
func SaveSubmission(id string, url string, author string) {
	_, err := Collection.InsertOne(context.TODO(), bson.D{
		{Key: "url", Value: url},
		{Key: "id", Value: id},
		{Key: "uid", Value: author},
	})

	if err != nil {
		fmt.Println(err)
		return
	}
}