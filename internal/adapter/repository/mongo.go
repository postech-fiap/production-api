package repository

import (
	"context"
	"errors"
	"github.com/postech-fiap/producao/internal/core/domain"
	"github.com/postech-fiap/producao/internal/core/port"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type mongoRepository struct {
	client          *mongo.Client
	orderCollection *mongo.Collection
	transactionTime time.Duration
}

func NewMongoRepository(client *mongo.Client) port.OrderRepositoryInterface {
	database := client.Database("local")
	return &mongoRepository{
		client:          client,
		orderCollection: database.Collection("orders"),
		transactionTime: time.Second * 1,
	}
}

func (m *mongoRepository) List() ([]domain.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), m.transactionTime)
	defer cancel()

	filter := bson.D{{"status", bson.D{{"$ne", domain.FINISHED}}}}
	opts := options.Find().SetSort(bson.D{{"status", 1}})
	cursor, err := m.orderCollection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	var orders []domain.Order
	err = cursor.All(context.Background(), &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (m *mongoRepository) Get(id int64) (*domain.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), m.transactionTime)
	defer cancel()

	filter := bson.D{{"id", id}}
	var order domain.Order
	err := m.orderCollection.FindOne(ctx, filter).Decode(&order)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return &order, nil
}

func (m *mongoRepository) Insert(order *domain.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), m.transactionTime)
	defer cancel()

	_, err := m.orderCollection.InsertOne(ctx, *order)
	if err != nil {
		return err
	}

	return nil
}

func (m *mongoRepository) UpdateStatus(order *domain.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), m.transactionTime)
	defer cancel()

	filter := bson.D{{"id", order.ID}}
	update := bson.D{{"$set", bson.D{{"status", order.Status}}}}
	_, err := m.orderCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
