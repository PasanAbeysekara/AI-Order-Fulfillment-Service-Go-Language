package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var orders = []Order{}

func GetOrders(c *gin.Context) {
	c.JSON(http.StatusOK, orders)
}

func CreateOrder(c *gin.Context) {
	var order Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order.ID = len(orders) + 1
	order.Status = "Pending"
	orders = append(orders, order)
	c.JSON(http.StatusCreated, order)
}

func RegisterOrderRoutes(router *gin.Engine) {
	orderGroup := router.Group("/orders")
	{
		orderGroup.GET("/", GetOrders)
		orderGroup.POST("/", CreateOrder)
	}
}
