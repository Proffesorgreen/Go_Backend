package data

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client
var taskCollection *mongo.Collection
var userCollection *mongo.Collection

func ConnectDB() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://kirubellegese:RnLhJWoMdVE9Trd4@go-backend-cluster.0dunmke.mongodb.net/?retryWrites=true&w=majority&appName=Go-Backend-Cluster").SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	// Send a ping to confirm a successful connection
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	Client = client
	taskCollection = client.Database("Task_Database").Collection("Task_Five")
	userCollection = client.Database("Task_Database").Collection("user_collection")
}

func DisconnectDB() {
	if Client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := Client.Disconnect(ctx); err != nil {
			log.Println("Error disconnecting MongoDB:", err)
		} else {
			fmt.Println("Disconnected from MongoDB.")
		}
	}
}
