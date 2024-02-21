package amqp

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

var connection *amqp.Connection = nil
var channel *amqp.Channel = nil

func OpenConnection() (*amqp.Channel, error) {
	var err error = nil

	connection, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}

	channel, err = connection.Channel()

	return channel, err
}

func CloneConnection() {
	channel.Close()
	connection.Close()
}
