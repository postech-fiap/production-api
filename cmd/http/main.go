package main

import (
	"github.com/gin-gonic/gin"
	"github.com/postech-fiap/production-api/cmd/config"
	repositoryAdapter "github.com/postech-fiap/production-api/cmd/repository"
	"github.com/postech-fiap/production-api/internal/adapter/handler/http"
	"github.com/postech-fiap/production-api/internal/adapter/handler/http/middlewares"
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

	// usecase
	orderUseCase := usecase.NewOrderUserCase(mongoRepository)

	// service
	pingService := http.NewPingService()
	orderService := http.NewOrderService(orderUseCase)

	router := gin.New()
	router.Use(middlewares.ErrorService)
	router.GET("/ping", pingService.Ping)
	router.GET("/order", orderService.List)
	router.POST("/order", orderService.Insert)
	router.PUT("/order/:id/status", orderService.SetStatus)
	router.Run()
}
