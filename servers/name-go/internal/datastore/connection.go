package datastore

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
)

func GetMongoClient() (*mongo.Client, error) {
	uri := os.Getenv("auth_mongo_uri")

	//connect
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	//error check
	if err != nil {
		return nil, err
	}

	//ping
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
		return nil, err
	}

	//return
	return client, nil
}
