package main

import (
	"context"

	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectDB() *mongo.Client {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var mongoUrl string = os.Getenv("MONGO_DB_URI")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))

	if err != nil {
		log.Fatal(err)
	}
	MongoClient = client

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	return client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("Cluster0").Collection(collectionName)
	return collection
}
