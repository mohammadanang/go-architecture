package main

import (
	"net/http"

	env "go-arch/config"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	port := ":" + env.Env["PORT"]

	route.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Welcome to the beginning of nothingness",
		})
	})

	route.Run(port)
}
