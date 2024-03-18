package consumer

import (
	"encoding/json"
	"fmt"
	"github.com/postech-fiap/production-api/internal/adapter/queue/consumer/dto"
	"github.com/postech-fiap/production-api/internal/adapter/queue/consumer/mapper"
	"github.com/postech-fiap/production-api/internal/core/port"
	amqp "github.com/rabbitmq/amqp091-go"
	"gopkg.in/go-playground/validator.v9"
)

const orderReceivedQueueName = "pedido-recebido"

type orderQueueConsumer struct {
	channel      *amqp.Channel
	orderUseCase port.OrderUseCaseInterface
}

func NewOrderQueueConsumer(channel *amqp.Channel, orderUseCase port.OrderUseCaseInterface) *orderQueueConsumer {
	o := orderQueueConsumer{
		channel:      channel,
		orderUseCase: orderUseCase,
	}

	o.createQueue()

	return &o
}

func (o *orderQueueConsumer) Listen() {
	go func() {
		msgs, err := o.channel.Consume(orderReceivedQueueName, "", false, false, false, false, nil)
		if err != nil {
			panic(err)
		}

		for message := range msgs {
			var orderDTO dto.NewOrderMessage
			err := json.Unmarshal(message.Body, &orderDTO)
			if err != nil {
				fmt.Println(err)
				continue
			}

			validate := validator.New()
			err = validate.Struct(orderDTO)
			if err != nil {
				fmt.Println(err, message.MessageId, string(message.Body))
				continue
			}

			newOrder := mapper.MapNewOrderMessageToDomain(&orderDTO)

			err = o.orderUseCase.Insert(newOrder)
			if err != nil {
				fmt.Println(err)
				continue
			}

			err = message.Ack(false)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}()
}

func (o *orderQueueConsumer) createQueue() {
	_, err := o.channel.QueueDeclare(orderReceivedQueueName, false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
}
