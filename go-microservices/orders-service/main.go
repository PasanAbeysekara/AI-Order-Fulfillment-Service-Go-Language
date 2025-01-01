package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	RegisterOrderRoutes(r)

	r.Run(":8082")
}
