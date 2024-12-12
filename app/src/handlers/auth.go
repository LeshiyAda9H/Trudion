package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"src/initializers"
	"src/models"
	"time"
)

type verifyEmailInput struct {
	Email string `json:"email" binding:"required"`
}

// @Summary VerifyEmail
// @Tags auth
// @Description check if email is available
// @Accept json
// @Produce json
// @Param input body verifyEmailInput true "Email"
// @Success 200 {object} string
// @Failure 400,500 {object} string
// @Failure default {object} string
// @Router /api/v1/verify/email [post]
func VerifyEmail(c *gin.Context) {
	var body verifyEmailInput
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read json body"})
		return
	}

	var count int64
	initializers.DB.Model(&models.User{}).Select("email").Where("email = ?", body.Email).Count(&count)
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is already in use"})
		return
	}

	// TODO: validate email via email service (sending email with a code)
	// ...

	c.JSON(http.StatusOK, gin.H{"available": true})
}

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary SignIn
// @Tags auth
// @Description login
// @Accept json
// @Produce json
// @Param input body signInInput true "Email and password"
// @Success 200 {object} string
// @Failure 400,500 {object} string
// @Failure default {object} string
// @Router /api/v1/login [post]
func SignIn(c *gin.Context) {
	// get the email and password from the request
	var err error
	var body signInInput
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	// look up requested user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.UserId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	// compare the password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
		return
	}

	// generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.UserId,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// return the token
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Token generated"})
}

type signUpInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
}

// @Summary SignUp
// @Tags auth
// @Description create account
// @Accept json
// @Produce json
// @Param input body signUpInput true "account info"
// @Success 200 {object} string
// @Failure 400,500 {object} string
// @Failure default {object} string
// @Router /api/v1/register [post]
func SignUp(c *gin.Context) {
	var err error
	var body signUpInput

	// check incoming data
	if err = c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// create a new user
	var user = models.User{
		Username:      body.Username,
		Email:         body.Email,
		PasswordHash:  string(hashedPassword),
		Gender:        body.Gender,
		AccountStatus: "active",
		OnlineStatus:  "online",
	}

	// Save user to database
	if initializers.DB.Create(&user).Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created"})
}
