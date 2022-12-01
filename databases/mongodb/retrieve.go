// Golang
// MongoDB
// CRUD: Examples of Retrieving

package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func retrieve(client *mongo.Client, ctx context.Context, pC *mongo.Collection, eC *mongo.Collection) {

	// ========== Retrieving

	// ==== list databases ====
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(databases)

	cursor, err := eC.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	// ==== iterate in batches ====
	for cursor.Next(ctx) {
		var episode bson.M
		if err = cursor.Decode(&episode); err != nil {
			log.Fatal(err)
		}

		fmt.Println(episode)
	}

	// ================== find one
	var podcast bson.M
	if err = pC.FindOne(ctx, bson.M{}).Decode(&podcast); err != nil {
		log.Fatal(err)
	}

	fmt.Println(podcast)
	// ===================

	// =================== find all documents with duation = 25
	filterCursor, err := eC.Find(ctx, bson.M{"duration": 25})
	if err != nil {
		log.Fatal(err)
	}

	// load all filtered results into new bson.M slice
	var episodesFiltered []bson.M
	if err = filterCursor.All(ctx, &episodesFiltered); err != nil {
		log.Fatal(err)
	}

	fmt.Println(episodesFiltered)
	// =================

	// ================= sorting on find
	opts := options.Find()
	opts.SetSort(bson.D{{"duration", 1}})
	sortCursor, err := eC.Find(ctx, bson.D{
		{"duration", bson.D{
			{"$gt", 24},
		}},
	}, opts)

	// load all sorted results into new bson.M slice
	var episodesSorted []bson.M
	if err = sortCursor.All(ctx, &episodesSorted); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodesSorted)

	/*
		// loading all documents into cursor
		var episodes []bson.M
		if err = cursor.All(ctx, &episodes); err != nil {
			log.Fatal(err)
		}

		for _, episode := range episodes {
			fmt.Println(episode["title"])

		}
	*/
}
