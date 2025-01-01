package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	RegisterDeliveryRoutes(r)

	r.Run(":8083")
}
