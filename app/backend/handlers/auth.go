package handlers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"src/initializers"
	"src/models"
	"src/repository"
	"src/types"
	"time"
)

const (
	salt      = "hK3B6z4a2pwPUmQdcrCxE9"
	secretKey = "b0nN_4;>-dj!PCtIYdqw0`C1&Uw;gS"
	tokenTTL  = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId uint `json:"user_id"`
}

// @Summary VerifyEmail
// @Tags auth
// @Description check if email is available
// @Accept json
// @Produce json
// @Param input body types.VerifyEmailPayload true "Email"
// @Success 200 {object} string
// @Failure 400,500 {object} string
// @Failure default {object} string
// @Router /api/v1/verify/email [post]
func VerifyEmail(c *gin.Context) {
	var body types.VerifyEmailPayload
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

// @Summary SignIn
// @Tags auth
// @Description login
// @Accept json
// @Produce json
// @Param input body types.SignInPayload true "Email and password"
// @Success 200 {object} string
// @Failure 400,500 {object} string
// @Failure default {object} string
// @Router /api/v1/login [post]
func SignIn(c *gin.Context) {
	// get the email and password from the request
	var err error
	var body types.SignInPayload
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	// look up requested user
	user, err := repository.GetUserByEmail(body.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
		return
	}

	// compare the password
	passWithSalt := append([]byte(body.Password), []byte(salt)...)
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), passWithSalt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
		return
	}

	// generate JWT token
	token, err := generateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// return the token
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// @Summary SignUp
// @Tags auth
// @Description create account
// @Accept json
// @Produce json
// @Param input body types.SignUpPayload true "account info"
// @Success 200 {object} string
// @Failure 400,500 {object} string
// @Failure default {object} string
// @Router /api/v1/register [post]
func SignUp(c *gin.Context) {
	var err error
	var body types.SignUpPayload

	// check incoming data
	if err = c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	// hash the password
	passWithSalt := append([]byte(body.Password), []byte(salt)...)
	hashedPassword, err := bcrypt.GenerateFromPassword(passWithSalt, bcrypt.DefaultCost)
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

	// check if email is already in use
	_, err = repository.GetUserByEmail(body.Email)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is already in use"})
		return
	}

	// Save user to database
	if initializers.DB.Create(&user).Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// generate JWT token
	token, err := generateToken(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func generateToken(user *models.User) (string, error) {
	// generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.UserId,
	})

	return token.SignedString([]byte(secretKey))
}

func parseToken(accessToken string) (uint, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}

		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, fmt.Errorf("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}
