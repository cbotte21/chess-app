package datastore

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
)

func GetMongoClient() (*mongo.Client, error) {
	//Verify env variables exist
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Could not load enviroment variables!")
	}
	uri, uriPresent := os.LookupEnv("username_mongo_uri")

	if !uriPresent {
		fmt.Println("could not find {username_mongo_uri} environment variable")
		os.Exit(1)
	}

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
