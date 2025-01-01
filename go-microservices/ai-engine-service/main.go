package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	RegisterAIRoutes(r)

	r.Run(":8084")
}
