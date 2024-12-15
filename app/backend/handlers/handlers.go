package handlers

import (
	"net/http"
	"src/initializers"
	"src/models"
	"src/repository"
	"strconv"

	"github.com/gin-gonic/gin"
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

type userProfileResponse struct {
	Username     string `json:"username"`
	Gender       string `json:"gender"`
	Biography    string `json:"biography"`
	Label        string `json:"label"`
	OnlineStatus string `json:"online_status"`
}

// @Summary UserProfile
// @Description get user profile information
// @Security ApiKeyAuth
// @Produce  json
// @Success 200 {object} string
// @Failure 400,401 {object} string
// @Router /api/v1/profile [get]
func GetUserProfile(c *gin.Context) {
	UserIdentity(c)

	if c.IsAborted() {
		return
	}

	id, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	user, err := repository.GetUserByID(id.(uint))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't find user profile"})
		return
	}

	response := userProfileResponse{
		Username:     user.Username,
		Gender:       user.Gender,
		Biography:    user.Biography,
		Label:        "ExampleLabel",
		OnlineStatus: user.OnlineStatus,
	}

	c.JSON(http.StatusOK, response)
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

// @Summary GetUsersNumber
// @Description get numver of users
// @Security ApiKeyAuth
// @Accept query
// @Produce  json
// @Param GetUsersNumber header int true "Get number of users"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/userscount [get]
func GetUsersNumber(c *gin.Context) {
	//ALTERNATIVE:
	// result := initializers.DB.Exec("SELECT * FROM users ORDER BY RANDOM() LIMIT 5;")
	// tmp := []int{5, 6, 8}
	// result := initializers.DB.Table("users").Where("user_id in ?", tmp).Find(&users)

	// UserIdentity(c)
	type UserPage struct {
		Username     string   `gorm:"size:20;not null" json:"username"`
		Gender       string   `gorm:"size:255;not null;check:gender IN ('male', 'female', 'prefer_not_to_say');default:prefer_not_to_say" json:"gender"`
		Biography    string   `gorm:"type:text;not null;default:' '" json:"biography"`
		Labels       []string `gorm:"type:text[]" json:"labels"'`
		OnlineStatus string   `gorm:"size:255;not null;check:online_status IN ('online', 'offline', 'away');default:'offline'" json:"online_status"`
	}

	var users []UserPage
	count, correct := c.GetQuery("usersnumber")
	if !correct {
		c.JSON(http.StatusBadRequest, gin.H{"error": "empty users number"})
		return
	}

	countInt, err := strconv.Atoi(count)
	if err != nil || countInt < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong users number"})
		return
	}

	result := initializers.DB.Table("users").Order("RANDOM()").Limit(countInt).Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get users",
		})
		return
	}
	for idx, _ := range users {
		users[idx].Labels = []string{"sport", "lenguage", "prog"}
	}
	c.JSON(http.StatusOK, gin.H{
		"result": users,
	})
}
