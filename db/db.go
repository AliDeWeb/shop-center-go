package db

import (
	"context"
	"github.com/alideweb/shop-center-go/config"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoClient *mongo.Client
	once        sync.Once
)

func ConnectToMongo(uri string) {
	once.Do(func() {
		wg := sync.WaitGroup{}

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

		MongoClient = client

		// Create Indexes
		wg.Add(1)
		go func() {
			defer wg.Done()

			err := createUniqueIndexes("user", []string{"email"})
			if err != nil {
				log.Printf("Error creating indexes for collection %s:", err)
			}
		}()
		wg.Wait()
	})

}

func createUniqueIndexes(collection string, fields []string) error {
	dbName := config.ServerEnvsConfig.MongoDbName
	coll := MongoClient.Database(dbName).Collection(collection)

	var indexes []mongo.IndexModel

	for _, field := range fields {
		index := mongo.IndexModel{
			Keys:    bson.D{{Key: field, Value: 1}},
			Options: options.Index().SetUnique(true),
		}
		indexes = append(indexes, index)
	}

	_, err := coll.Indexes().CreateMany(context.Background(), indexes)
	if err != nil {
		return err
	}

	return nil
}
