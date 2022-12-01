// Golang
// MongoDB
// CRUD: Examples

package main

import (
	"context"
	//"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	firstname string `bson:"firstname"`
	lastname  string `bson:"lastname"`
}

type Podcast struct {
	ID     primitive.ObjectID `bson:"_id, omitempty"`
	Title  string             `bson:"title, omitempty"`
	Author string             `bson:"author, omitempty"`
	Tags   []string           `bson:"tags, omitempty"`
}

type Episode struct {
	ID          primitive.ObjectID `bson:"_id, omitempty"`
	Podcast     primitive.ObjectID `bson:"podcast, omitempty"`
	Title       string             `bson:"title, omitempty"`
	Description string             `bson:"desc, omitempty"`
	Duration    int32              `bson:"duration, omitempty"`
}

const uri = "mongodb://localhost:27017"

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	// CRUD examples separated for easier reading
	//
	// Examples of Creating, returns pointers of 2 collections
	pC, eC := create(client, ctx)

	// Examples of Retrieving
	//retrieve(client, ctx, pC, eC)

	// Examples of Updating
	//update(client, ctx, pC, eC)

	// Examples of Deleting
	//del(client, ctx, pC, eC)

	// ==== Examples with native structs ====

	mongoPodcast := Podcast{
		Title:  "The MongoDB Podcast",
		Author: "Nick Namey",
		Tags:   []string{"mongodb", "nosql"},
	}

	insertResult, err := pC.InsertOne(ctx, mongoPodcast)
	if err != nil {
		panic(err)
	}

	log.Printf("inserted %v!\n", insertResult.InsertedID)

	// Examples with native structs
	var podcasts []Podcast
	podcastCursor, err := pC.Find(ctx, bson.M{})
	if err != nil {
		panic(err)
	}
	if err = podcastCursor.All(ctx, &podcasts); err != nil {
		panic(err)
	}

	log.Println(podcasts)

	var episodes []Episode
	episodeCursor, err := eC.Find(ctx, bson.M{})
	if err != nil {
		panic(err)
	}
	if err = episodeCursor.All(ctx, &episodes); err != nil {
		panic(err)
	}
	log.Println(episodes)

}
