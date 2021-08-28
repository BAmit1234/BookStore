package mongodb

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Collection {
	clientOptions := options.Client().ApplyURI("mongodb://db:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("ERRROR", err, os.Getenv("MONGODB_URI"))
		os.Exit(1)

	}
	fmt.Println("mongo connected")
	collection := client.Database("mongo").Collection("book")

	return collection
}
