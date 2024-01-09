package main

import (
	"github.com/gin-gonic/gin"
	"github.com/postech-fiap/producao/internal/adapter/handler/http"
	"github.com/postech-fiap/producao/internal/adapter/handler/http/middlewares"
	"github.com/postech-fiap/producao/internal/adapter/repository"
	"github.com/postech-fiap/producao/internal/config"
	"github.com/postech-fiap/producao/internal/core/service"
)

func main() {
	// repositories
	mongoClient, err := config.OpenConnection()
	if err != nil {
		panic(err)
	}
	defer config.CloseConnection()

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
