package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(ctx context.Context) *mongo.Database {
	connection := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, connection)
	if err != nil {
		panic(err)
	}

	return client.Database("learn_mongo")
}
