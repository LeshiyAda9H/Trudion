package handlers

import (
	"fmt"
	"net/http"
	"src/initializers"
	"src/models"
	"src/repository"
	"strconv"

	"github.com/gin-contrib/sessions"
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
	Username     string   `json:"username"`
	Gender       string   `json:"gender"`
	Biography    string   `json:"biography"`
	Labels       []string `json:"label"`
	OnlineStatus string   `json:"online_status"`
}

// @Summary UserProfile
// @Description get user profile information
// @Security ApiKeyAuth
// @Produce  json
// @Success 200 {object} string
// @Failure 400,401,500 {object} string
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

	// Get labels
	var userLabels []models.UserLabel
	err = initializers.DB.Preload("Label").Where("user_id = ?", user.UserId).Find(&userLabels).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user labels"})
		return
	}

	var labels []string
	for _, userLabel := range userLabels {
		labels = append(labels, userLabel.Label.LabelName)
	}

	response := userProfileResponse{
		Username:     user.Username,
		Gender:       user.Gender,
		Biography:    user.Biography,
		Labels:       labels,
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
// @Description get number of users
// @Produce  json
// @Param usersnumber query int true "Number of users to fetch"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/usersnumber [get]
func GetUsersNumber(c *gin.Context) {
	//ALTERNATIVE:
	// result := initializers.DB.Exec("SELECT * FROM users ORDER BY RANDOM() LIMIT 5;")
	// tmp := []int{5, 6, 8}
	// result := initializers.DB.Table("users").Where("user_id in ?", tmp).Find(&users)

	var users []models.UserPage
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

	userSession := sessions.Default(c)
	userId := userSession.Get("userId")

	fmt.Println(userId)

	result := initializers.DB.Table("users").Where("user_id <> ?", userId).Order("RANDOM()").Limit(countInt).Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get users",
		})
		return
	}
	for idx, user := range users {
		var userLabels []models.UserLabel
		err := initializers.DB.Preload("Label").Where("user_id = ?", user.UserId).Find(&userLabels).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user labels"})
			return
		}

		var labels []string
		for _, userLabel := range userLabels {
			labels = append(labels, userLabel.Label.LabelName)
		}

		users[idx].Labels = labels
	}
	c.JSON(http.StatusOK, gin.H{
		"result": users,
	})
}

// @Summary GetUsersPage
// @Description get page of users
// @Produce  json
// @Param page query int true "Page number of users"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/userspage [get]
func GetUsersPage(c *gin.Context) {
	if start == -1 || step == -1 || dbSize == -1 {
		setStartStep(c)
	}

	pageStr, correct := c.GetQuery("page")
	if !correct {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empty page query",
		})
		return
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 || (dbSize+pageSize)/pageSize < page {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid page value",
		})
		return
	}

	// var rowIndices [pageSize]int
	var rowIndices []int = make([]int, pageSize)
	var dbIterator = (start + (page-1)*pageSize*step) % dbSize
	for i := range pageSize {
		dbIterator = (dbIterator + step) % dbSize
		if dbIterator == start {
			rowIndices[i] = start + 1
			break
		}
		rowIndices[i] = dbIterator + 1
	}

	// var result [pageSize]models.UserPage
	var result []models.UserPage = make([]models.UserPage, pageSize)
	if err := initializers.DB.Raw(getPageQuery, rowIndices).Scan(&result).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed get users page",
		})
		return
	}

	usersOnPage := pageSize
	if page > dbSize/pageSize {
		usersOnPage = dbSize % pageSize
	}

	c.JSON(http.StatusOK, gin.H{
		"page":        page,
		"users_count": usersOnPage,
		"result":      result,
	})
}
