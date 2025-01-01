package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var inventory = []Product{
	{ID: 1, Name: "Laptop", Quantity: 50},
	{ID: 2, Name: "Mouse", Quantity: 150},
}

func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, inventory)
}

func AddProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	inventory = append(inventory, product)
	c.JSON(http.StatusCreated, product)
}

func UpdateProductQuantity(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedProduct Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, p := range inventory {
		if p.ID == id {
			inventory[i].Quantity = updatedProduct.Quantity
			c.JSON(http.StatusOK, inventory[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}

func RegisterInventoryRoutes(router *gin.Engine) {
	inventoryGroup := router.Group("/inventory")
	{
		inventoryGroup.GET("/products", GetProducts)
		inventoryGroup.POST("/products", AddProduct)
		inventoryGroup.PUT("/products/:id", UpdateProductQuantity)
	}
}
