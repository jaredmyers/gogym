// Golang
// MongoDB
// CRUD: Examples of Updating

package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

func update(client *mongo.Client, ctx context.Context, pC *mongo.Collection, eC *mongo.Collection) {

	// ==== update operations ====

	/*
		id, _ := primitive.ObjectIDFromHex("")

		// updating one
		result, err := pC.UpdateOne(
			ctx,
			bson.M{"_id": id},
			bson.D{
				{"$set", bson.D{{"author", "Henry Cavall"}}},
			},
		)

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Updated %v Documents\n", result.ModifiedCount)
	*/

	// updating many
	result, err := pC.UpdateMany(
		ctx,
		bson.M{"title": "Polyglot Dev Podcast"},
		bson.D{
			{"$set", bson.D{{"author", "Charles Cavall"}}},
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Updated %v Documents\n", result.ModifiedCount)

	/*
		// full doc replace based on the search criteria (author)
		result, err = pC.ReplaceOne(
			ctx,
			bson.M{"author": "Charles Cavall"},
			bson.M{
				"title":  "Tommy T Show",
				"author": "Big C Cavall",
			},
		)

		log.Printf("Updated %v Documents\n", result.ModifiedCount)
	*/
}
