package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"src/initializers"
	"src/models"
)

// @Summary Ping
// @Description ping server
// @Produce  json
// @Success 200 {object} string
// @Router /api/v1 [get]
func RootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Server is running.",
	})
}

// @Summary UserProfile
// @Description get user profile information
// @Security ApiKeyAuth
// @Produce  json
// @Success 200 {object} string
// @Failure 500 {object} string
// @Router /api/v1/profile [get]
func UserProfile(c *gin.Context) {
	UserIdentity(c)

	id, _ := c.Get("userId")
	c.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary GetUsers
// @Description get all users
// @Produce  json
// @Success 200 {object} string
// @Failure 500 {object} string
// @Router /api/v1/users [get]
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
