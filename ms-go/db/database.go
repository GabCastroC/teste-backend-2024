package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Collection
var client *mongo.Client

func Connection() *mongo.Collection {
	var err error

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://admin:password@mongodb:27017")

	// Connect to MongoDB
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("Failed to connect to MongoDB: ", err)
		return nil
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		fmt.Println("Failed to ping MongoDB: ", err)
		return nil
	}

	// Set the database and collection variables
	db = client.Database("teste_backend").Collection("products")

	fmt.Println("Connected to MongoDB!")

	return db
}

func Disconnect() {
    err := client.Disconnect(context.Background())
    if err != nil {
        fmt.Println("Failed to disconnect from MongoDB:", err)
    }
}

