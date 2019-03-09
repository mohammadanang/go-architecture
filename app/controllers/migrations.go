package controllers

import (
	userModel "go-arch/app/models"
	. "go-arch/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Migrate(c *gin.Context) {
	user := userModel.User{}
	result := DB.AutoMigrate(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": result.Error,
		"data":    result.Value,
	})
}
