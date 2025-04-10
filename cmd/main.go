package main

import (
	"github.com/gin-gonic/gin"
	"order-service/db"
	"order-service/handler"
)

func main() {
	router := gin.Default()
	orderHandler := handler.NewHandler()

	defer db.CloseConnection()

	// Order routes
	router.GET("/orders", orderHandler.GetAllOrders)
	router.GET("/orders/:id", orderHandler.GetOrderByID)
	router.POST("/orders", orderHandler.CreateOrder)
	router.PATCH("/orders/:id", orderHandler.UpdateOrder)
	router.DELETE("/orders/:id", orderHandler.DeleteOrder)

	router.Run("0.0.0.0:8083")
}
