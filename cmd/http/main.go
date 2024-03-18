package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/postech-fiap/production-api/cmd/amqp"
	"github.com/postech-fiap/production-api/cmd/config"
	repositoryAdapter "github.com/postech-fiap/production-api/cmd/repository"
	"github.com/postech-fiap/production-api/internal/adapter/handler/http"
	"github.com/postech-fiap/production-api/internal/adapter/handler/http/middlewares"
	"github.com/postech-fiap/production-api/internal/adapter/queue/consumer"
	"github.com/postech-fiap/production-api/internal/adapter/queue/publisher"
	"github.com/postech-fiap/production-api/internal/adapter/repository"
	"github.com/postech-fiap/production-api/internal/core/usecase"
)

func main() {
	// config
	configuration, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	// repository
	mongoClient, err := repositoryAdapter.OpenConnection(configuration)
	if err != nil {
		panic(err)
	}
	defer repositoryAdapter.CloseConnection()

	mongoRepository := repository.NewMongoRepository(mongoClient)

	// amqp
	AMQPChannel, err := amqp.OpenConnection(configuration)
	if err != nil {
		panic(err)
	}
	defer amqp.CloneConnection()

	// queue publisher
	orderQueuePublisher := publisher.NewOrderQueuePublisher(AMQPChannel)

	// usecase
	orderUseCase := usecase.NewOrderUserCase(mongoRepository, orderQueuePublisher)

	// service
	pingService := http.NewPingService()
	orderService := http.NewOrderService(orderUseCase)

	// queue consumer
	orderQueueConsumer := consumer.NewOrderQueueConsumer(AMQPChannel, orderUseCase)
	orderQueueConsumer.Listen()

	router := gin.New()
	router.Use(middlewares.ErrorService)
	router.GET("/ping", pingService.Ping)
	router.GET("/order", orderService.List)
	router.PUT("/order/:id/status", orderService.UpdateStatus)

	address := fmt.Sprintf("%s:%s", configuration.Server.Host, configuration.Server.Port)
	router.Run(address)
}
