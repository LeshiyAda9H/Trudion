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

func LoginHandler(c *gin.Context) {
	// get the email and password from the request
	var body struct {
		Email    string
		Password string
	}
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
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(body.Password))
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

func RegisterHandler(c *gin.Context) {
	var err error
	var body struct {
		Name     string
		Email    string
		Password string
		Gender   string
	}

	// check incoming data
	if err = c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	var user models.User
	result := initializers.DB.Where("email = ?", body.Email).First(&user)
	if result.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email is already in use",
		})
		return
	}

	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// create a new user
	user = models.User{
		Username:      body.Name,
		Email:         body.Email,
		PasswordHash:  string(hashedPassword),
		Gender:        body.Gender,
		AccountStatus: "active",
		OnlineStatus:  "online",
	}
	result = initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created"})
}
