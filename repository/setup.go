package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//declare app collection in mongoDb
var (
	UserCollection *mongo.Collection
	TodoCollection *mongo.Collection
	Client         *mongo.Client
	Error          error
)

//declare server and credential to connect mongodb
var (
	HOST    = os.Getenv("HOST")
	UID     = os.Getenv("UID")
	PWD     = os.Getenv("PWD")
	DB_NAME = os.Getenv("Database")
)

func Setup() {
	URI := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", UID, PWD, HOST)

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)

	clientOptions := options.Client().
		ApplyURI(URI).
		SetServerAPIOptions(serverAPIOptions)

	Client, Error = mongo.Connect(context, clientOptions)

	UserCollection = Client.Database(DB_NAME).Collection("users")
	TodoCollection = Client.Database(DB_NAME).Collection("todos")

	if Error != nil {
		log.Fatal(Error)
	}
}
