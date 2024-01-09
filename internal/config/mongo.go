package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var client *mongo.Client = nil

func OpenConnection() (*mongo.Client, error) {
	credential := options.Credential{
		Username: "root",
		Password: "example",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cli, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credential))
	if err != nil {
		return nil, err
	}

	client = cli

	err = testConnection()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func CloseConnection() {
	if err := client.Disconnect(context.Background()); err != nil {
		panic(err)
	}
	fmt.Println("Mongo disconnected")
}

func testConnection() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	return client.Ping(ctx, readpref.Primary())
}
