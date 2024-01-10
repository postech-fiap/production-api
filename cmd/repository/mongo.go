package repository

import (
	"context"
	"fmt"
	"github.com/postech-fiap/producao/cmd/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var client *mongo.Client = nil

func OpenConnection(config *config.Config) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := fmt.Sprintf("mongodb://%s:%s", config.Database.Host, config.Database.Port)
	credential := options.Credential{
		Username: config.Database.Username,
		Password: config.Database.Password,
	}

	cli, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetAuth(credential))
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
