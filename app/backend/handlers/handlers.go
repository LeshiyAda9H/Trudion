package handlers

import (
	"fmt"
	"net/http"
	"src/initializers"
	"src/models"
	"src/repository"
	"src/types"
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
// @Tags profile
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
// @Security ApiKeyAuth
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

	UserIdentity(c)

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

	// userSession := sessions.Default(c)
	userId, _ := c.Get("userId")

	fmt.Println(userId)

	result := initializers.DB.Table("users").Where("user_id <> ?", userId).Order("RANDOM()").Limit(countInt).Find(&users)
	// result := initializers.DB.Table("users").Order("RANDOM()").Limit(countInt).Find(&users)
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

	// setStartStep(c)
	userSession := sessions.Default(c)
	if userSession.Get("start") == nil {
		setStartStep(c)
	}
	start, _ := userSession.Get("start").(int)
	step := userSession.Get("step").(int)
	dbSize := userSession.Get("dbSize").(int)
	userId := userSession.Get("userId").(uint)

	pageStr, correct := c.GetQuery("page")
	if !correct {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empty page query",
		})
		return
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 || (dbSize+pageSize-2)/pageSize < page {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid page value",
		})
		return
	}

	// var rowIndices [pageSize]int
	var userRow int
	if initializers.DB.Raw(getUserRow, userId).Scan(&userRow).Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed find user ID",
		})
		return
	}

	var rowIndices = make([]int, pageSize)
	var dbIterator = (start + page*pageSize*step) % dbSize
	for i := range pageSize {
		dbIterator = (dbIterator + step) % dbSize
		if dbIterator == start {
			rowIndices[i] = start + 1
			break
		}
		rowIndices[i] = dbIterator + 1
	}

	// var result [pageSize]models.UserPage
	var result = make([]models.UserPage, pageSize)
	if err := initializers.DB.Raw(getPageQuery, userId, rowIndices).Scan(&result).Error; err != nil {
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

// @Summary Notifications
// @Description get user notifications information
// @Security ApiKeyAuth
// @Produce  json
// @Success 200 {object} string
// @Failure 400,401,500 {object} string
// @Router /api/v1/notifications [get]
func GetNotifications(c *gin.Context) {
	UserIdentity(c)

	if c.IsAborted() {
		return
	}

	userId, _ := c.Get("userId")
	var result []models.Notification

	initializers.DB.Where("user_id = ?", userId).Find(&result)

	c.JSON(http.StatusOK, gin.H{
		"notifications": result,
	})

}

// @Summary Handshake
// @Description send match request
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param input body types.HandshakePayload true "recipient id"
// @Success 200 {object} string
// @Failure 400,500 {object} string
// @Router /api/v1/handshake [post]
func Handshake(c *gin.Context) {
	UserIdentity(c)

	if c.IsAborted() {
		return
	}

	var body types.HandshakePayload
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to read body"})
		return
	}

	senderId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	if senderId == body.RecipientId {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't handshake with yourself"})
		return
	}

	// get sender model
	var sender models.User
	if initializers.DB.Where("user_id = ?", senderId).First(&sender).Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to get sender from database"})
		return
	}

	// check if the recipient exists
	var recipient models.User
	if initializers.DB.Where("user_id = ?", body.RecipientId).First(&recipient).Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "recipient not found"})
		return
	}

	// check if the recipient sends a handshake to the sender
	if initializers.DB.Where("sender_id = ? AND recipient_id = ?", body.RecipientId, sender.UserId).First(&models.Like{}).Error == nil {
		c.JSON(http.StatusOK, gin.H{"message": "mutually"})

		if initializers.DB.Where("sender_id = ? AND recipient_id = ?", sender.UserId, body.RecipientId).First(&models.Like{}).Error == nil {
			return
		}

		// create a handshake
		initializers.DB.Create(&models.Like{
			SenderID:    sender.UserId,
			RecipientID: body.RecipientId,
		})

		// create a notification for the recipient
		initializers.DB.Create(&models.Notification{
			UserID:  body.RecipientId,
			Message: "У вас взаимная симпатия с пользователем " + sender.Username,
		})
		initializers.DB.Create(&models.MatchList{
			FirstPersonID:  body.RecipientId,
			SecondPersonID: sender.UserId,
		})
		// SendNotification(recipient.UserId, "У вас взаимная симпатия с пользователем "+sender.Username)

	} else {
		c.JSON(http.StatusOK, gin.H{"message": "handshake successful"})

		if initializers.DB.Where("sender_id = ? AND recipient_id = ?", sender.UserId, body.RecipientId).First(&models.Like{}).Error == nil {
			return
		}

		// create a handshake
		initializers.DB.Create(&models.Like{
			SenderID:    sender.UserId,
			RecipientID: body.RecipientId,
		})

		// create a notification for the recipient
		initializers.DB.Create(&models.Notification{
			UserID:  body.RecipientId,
			Message: "Пользователь " + sender.Username + " хочет познакомиться с вами",
		})
		// SendNotification(recipient.UserId, "Пользователь "+sender.Username+" хочет познакомиться с вами")

	}
}

// @Summary UpdateProfile
// @Tags profile
// @Description update user profile information
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param input body types.UpdateProfilePayload true "profile information"
// @Success 200 {object} string
// @Failure 400,500 {object} string
// @Router /api/v1/profile [patch]
func UpdateProfile(c *gin.Context) {
	UserIdentity(c)

	if c.IsAborted() {
		return
	}

	var body map[string]interface{}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to read body"})
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

	allowedFields := map[string]bool{
		"username":  true,
		"gender":    true,
		"biography": true,
	}

	updateData := make(map[string]interface{})
	for key, val := range body {
		if allowedFields[key] {
			updateData[key] = val
		}
	}

	if len(updateData) > 0 {
		if err := initializers.DB.Model(user).Updates(updateData).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
			return
		}
	}

	// update labels
	if labels, ok := body["label"].([]interface{}); ok {
		initializers.DB.Where("user_id = ?", user.UserId).Delete(&models.UserLabel{})
		for _, label := range labels {
			labelStr, ok := label.(string)
			if !ok {
				continue
			}
			var labelInfo models.LabelInfo
			initializers.DB.Where("label_name = ?", labelStr).First(&labelInfo)
			initializers.DB.Create(&models.UserLabel{
				UserID:  user.UserId,
				LabelID: labelInfo.LabelID,
			})
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "profile updated successfully"})
}

// @Summary DeleteUser
// @Tags profile
// @Description delete user
// @Security ApiKeyAuth
// @Produce  json
// @Success 200 {object} string
// @Failure 400,500 {object} string
// @Router /api/v1/deleteuser [delete]
func DeleteUser(c *gin.Context) {
	UserIdentity(c)

	if c.IsAborted() {
		return
	}

	id, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user ID not found"})
		return
	}

	user, err := repository.GetUserByID(id.(uint))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't find user profile"})
		return
	}

	if initializers.DB.
		Delete(&user).
		Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
		return
	}

	if initializers.DB.
		Where("user_id = ?", id).
		Delete(&models.UserLabel{}).
		Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user labels"})
		return
	}

	if initializers.DB.
		Where("user_id = ?", id).
		Delete(&models.Notification{}).
		Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user notifications"})
		return
	}

	if initializers.DB.
		Where("sender_id = ? OR recipient_id = ?", id, id).
		Delete(&models.Like{}).
		Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user likes"})
		return
	}

	if initializers.DB.Model(&models.Message{}).
		Where("sender_id = ?", id).
		Update("sender_id = ?", nil).
		Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user messages"})
		return
	}

	if initializers.DB.Model(&models.Message{}).
		Where("recipient_id = ?", id).
		Update("recipient_id", nil).
		Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user messages"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}

// @Summary GetMatches
// @Description get user matches
// @Security ApiKeyAuth
// @Produce  json
// @Success 200 {object} string
// @Failure 400,500 {object} string
// @Router /api/v1/matches [get]
func GetMatches(c *gin.Context) {
	UserIdentity(c)

	if c.IsAborted() {
		return
	}

	id, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user ID not found"})
		return
	}

	var matches []models.MatchList
	if err := initializers.DB.Where("first_person_id = ? OR second_person_id = ?", id, id).Order("match_date desc").Find(&matches).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get matches"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"matches": matches})
}
