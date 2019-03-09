package controllers

import (
	userModel "go-arch/app/models"
	. "go-arch/config"
    "net/http"
    "errors"

	"github.com/gin-gonic/gin"
)

func Migrate(c *gin.Context) {
    user := userModel.User{}
    errorMsg := errors.New("")

    result := DB.AutoMigrate(&user)
    if result.Error == nil {
        errorMsg = errors.New("Migration error occured")
    }

	c.JSON(http.StatusOK, gin.H{
		"message": errorMsg,
		"data":    result.Value,
	})
}
