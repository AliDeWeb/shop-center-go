package db

import (
	"context"
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

	return result, data, nil
}
