package datastore

import (
	"context"
	"github.com/cbotte21/name-go/internal/schema"
)

func Find[T schema.Schema[any]](schema T) (T, error) {
	client, err := GetMongoClient()
	if err != nil {
		return schema, err
	}
	var result T
	collection := client.Database(schema.Database()).Collection(schema.Collection())
	err = collection.FindOne(context.TODO(), schema).Decode(&result)

	if err != nil {
		return schema, err
	}
	return result, nil
}

func Create[T schema.Schema[any]](schema T) error {
	client, err := GetMongoClient()
	if err != nil {
		return err
	}

	collection := client.Database(schema.Database()).Collection(schema.Collection())
	_, err = collection.InsertOne(context.TODO(), schema)

	return err
}

func Update[X, Y schema.Schema[any]](filter X, updated Y) error {
	client, err := GetMongoClient()
	if err != nil {
		return err
	}

	collection := client.Database(filter.Database()).Collection(filter.Collection())
	_, err = collection.UpdateOne(context.TODO(), filter, updated)

	return err
}

func Delete[T schema.Schema[any]](schema T) error {
	client, err := GetMongoClient()
	if err != nil {
		return err
	}

	collection := client.Database(schema.Database()).Collection(schema.Collection())
	_, err = collection.DeleteOne(context.TODO(), schema)

	return err
}
