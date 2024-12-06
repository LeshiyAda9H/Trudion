package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"src/initializers"
	"src/models"
)

func RootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Server is running.",
	})
}

func GetUsers(c *gin.Context) {
	var users []models.User
	result := initializers.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get users",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
