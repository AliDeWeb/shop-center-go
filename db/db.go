package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongo(uri string) *mongo.Client {
	clientOption := options.Client().ApplyURI(uri)

	client, err := mongo.NewClient(clientOption)
	if err != nil {
		log.Fatal("Error creating Mongo client", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Error connecting to MongoDB", err)
	}

	log.Println("Connected to MongoDB")

	return client
}
