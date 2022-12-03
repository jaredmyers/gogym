package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jaredmyers/gogym/api/gin_mongo/controllers"
	"github.com/jaredmyers/gogym/api/gin_mongo/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server         *gin.Engine
	userService    *services.UserService
	userController controllers.UserController
	ctx            context.Context
	userCollection *mongo.Collection
	client         *mongo.Client
	err            error
)

func init() {
	ctx = context.TODO()

	conn := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err = mongo.Connect(ctx, conn)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("mongo connection established...")

	userCollection = client.Database("userdb").Collection("user")
	userService = services.NewUserService(userCollection, ctx)
	userController = controllers.NewUserController(userService)
	server = gin.Default()
}

func main() {

	defer client.Disconnect(ctx)

	basepath := server.Group("/v1")
	userController.RegisterUserRoutes(basepath)

	log.Panic(server.Run(":8000"))

}
