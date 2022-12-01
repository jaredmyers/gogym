// Golang
// MongoDB
// CRUD: Examples of Creating

package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// examples of creating db, collections, and  documents in in MongoDB
// returning ptrs to the collections for the other examples: (cRUD)
func create(client *mongo.Client, ctx context.Context) (*mongo.Collection, *mongo.Collection) {

	// ==== Creating documents ===

	db := client.Database("media")
	podcastsCollection := db.Collection("podcasts")
	episodesCollection := db.Collection("episodes")

	podcastResult, err := podcastsCollection.InsertOne(ctx, bson.D{
		{Key: "title", Value: "Polyglot Dev Podcast"},
		{Key: "author", Value: "Mc Fats"},
		{Key: "tags", Value: bson.A{"development", "programming", "coding"}},
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(podcastResult.InsertedID)

	episodeResult, err := episodesCollection.InsertMany(ctx, []interface{}{
		bson.D{
			{"podcast", podcastResult.InsertedID},
			{"title", "Episode 1"},
			{"desc", "This is the first episode"},
			{"duration", 25},
		},
		bson.D{
			{"podcast", podcastResult.InsertedID},
			{"title", "Episode 2"},
			{"desc", "This is the second episode"},
			{"duration", 32},
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(episodeResult.InsertedIDs)

	return podcastsCollection, episodesCollection
}
