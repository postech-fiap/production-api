package main

import (
	"github.com/gin-gonic/gin"
	"github.com/postech-fiap/production-api/cmd/config"
	repositoryAdapter "github.com/postech-fiap/production-api/cmd/repository"
	"github.com/postech-fiap/production-api/internal/adapter/handler/http"
	"github.com/postech-fiap/production-api/internal/adapter/handler/http/middlewares"
	"github.com/postech-fiap/production-api/internal/adapter/repository"
	"github.com/postech-fiap/production-api/internal/core/service"
)

func main() {
	// config
	configuration, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	// repositories
	mongoClient, err := repositoryAdapter.OpenConnection(configuration)
	if err != nil {
		panic(err)
	}
	defer repositoryAdapter.CloseConnection()

	mongoRepository := repository.NewMongoRepository(mongoClient)

	// service
	orderService := service.NewOrderService(mongoRepository)

	// handler
	orderHandler := http.NewOrderHandler(orderService)

	router := gin.New()
	router.Use(middlewares.ErrorHandler)
	router.GET("/order", orderHandler.List)
	router.POST("/order", orderHandler.Insert)
	router.PUT("/order/:id/status", orderHandler.SetStatus)
	router.Run()
}
