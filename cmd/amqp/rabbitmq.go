package amqp

import (
	"fmt"
	"github.com/postech-fiap/production-api/cmd/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

var connection *amqp.Connection = nil
var channel *amqp.Channel = nil

func OpenConnection(config *config.Config) (*amqp.Channel, error) {
	var err error = nil

	url := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		config.RabbitMQ.Username,
		config.RabbitMQ.Password,
		config.RabbitMQ.Host,
		config.RabbitMQ.Port)

	connection, err = amqp.Dial(url)
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
