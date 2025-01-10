package db

import (
	"context"
	"log"

	"github.com/alideweb/shop-center-go/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOne[T any](collection string, data T) (*mongo.InsertOneResult, T, error) {
	dbName := config.ServerEnvsConfig.MongoDbName
	coll := MongoClient.Database(dbName).Collection(collection)

	result, err := coll.InsertOne(context.Background(), data)
	if err != nil {
		return nil, data, err
	}

	log.Printf("Inserted a document:\nresult: %v\ndata:%v\nerr:%v", result, data, err)

	return result, data, nil
}
