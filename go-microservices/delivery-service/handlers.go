package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var deliveries = []Delivery{}

func GetDeliveries(c *gin.Context) {
	c.JSON(http.StatusOK, deliveries)
}

func ScheduleDelivery(c *gin.Context) {
	var delivery Delivery
	if err := c.ShouldBindJSON(&delivery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	delivery.ID = len(deliveries) + 1
	delivery.Status = "Scheduled"
	deliveries = append(deliveries, delivery)
	c.JSON(http.StatusCreated, delivery)
}

func RegisterDeliveryRoutes(router *gin.Engine) {
	deliveryGroup := router.Group("/deliveries")
	{
		deliveryGroup.GET("/", GetDeliveries)
		deliveryGroup.POST("/", ScheduleDelivery)
	}
}
