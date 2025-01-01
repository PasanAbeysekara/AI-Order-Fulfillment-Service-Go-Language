package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PredictionRequest struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
}

type PredictionResponse struct {
	Route string `json:"route"`
	ETA   string `json:"eta"`
}

func PredictRoute(c *gin.Context) {
	var req PredictionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, PredictionResponse{
		Route: "Optimal Route",
		ETA:   "2 hours",
	})
}

func RegisterAIRoutes(router *gin.Engine) {
	aiGroup := router.Group("/ai")
	{
		aiGroup.POST("/predict", PredictRoute)
	}
}
