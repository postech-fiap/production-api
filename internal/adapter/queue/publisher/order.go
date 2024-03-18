package publisher

import (
	"context"
	"encoding/json"
	"github.com/postech-fiap/production-api/internal/adapter/queue/publisher/mapper"
	"github.com/postech-fiap/production-api/internal/core/domain"
	amqp "github.com/rabbitmq/amqp091-go"
)

const orderNewStatusQueueName = "status-pedido-alterado"

type orderQueuePublisher struct {
	channel *amqp.Channel
}

func NewOrderQueuePublisher(channel *amqp.Channel) *orderQueuePublisher {
	o := orderQueuePublisher{
		channel: channel,
	}

	o.createQueue()

	return &o
}

func (o *orderQueuePublisher) PublishNewStatus(order *domain.Order) error {
	dto := mapper.DomainToOrderNewStatusMessage(order)

	body, err := json.Marshal(dto)
	if err != nil {
		return err
	}

	return o.channel.PublishWithContext(context.Background(),
		"",
		orderNewStatusQueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
}

func (o *orderQueuePublisher) createQueue() {
	_, err := o.channel.QueueDeclare(orderNewStatusQueueName, false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
}
