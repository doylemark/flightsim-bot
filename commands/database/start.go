package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

// RemoveCompetition removes all entries from the screenshots collection
func RemoveCompetition() {
	_, err := Collection.DeleteMany(context.TODO(), bson.D{{}})
	_, err = WinnersCollection.DeleteMany(context.TODO(), bson.D{{}})

	if err != nil {
		fmt.Println(err)
		return
	}
}
