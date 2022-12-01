// Golang
// MongoDB
// CRUD: Examples of Deleting

package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

func del(client *mongo.Client, ctx context.Context, pC *mongo.Collection, eC *mongo.Collection) {

	// ==== delete one ====
	result, err := pC.DeleteOne(ctx, bson.M{"title": "Polyglot Dev Podcast"})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("deleteone remove %v document(s)\n", result.DeletedCount)

	// ==== delete many ====
	result, err = eC.DeleteMany(ctx, bson.M{"duration": 25})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("DeleteMany removed %v document(s)\n", result.DeletedCount)

	// ==== drop collection + all docs inside ====
	if err = pC.Drop(ctx); err != nil {
		log.Fatal(err)
	}

	if err = eC.Drop(ctx); err != nil {
		log.Fatal(err)
	}

}
