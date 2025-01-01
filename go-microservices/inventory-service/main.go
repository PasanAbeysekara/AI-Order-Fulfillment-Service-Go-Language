package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	RegisterInventoryRoutes(r)

	r.Run(":8081")
}
