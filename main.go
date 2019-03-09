package main

import (
	"net/http"

	migrationController "go-arch/app/controllers"
	. "go-arch/config"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	port := ":" + ENV["APP_PORT"]

	route.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Welcome to the beginning of nothingness",
		})
	})

	route.GET("migrate", migrationController.Migrate)

	route.Run(port)
}
