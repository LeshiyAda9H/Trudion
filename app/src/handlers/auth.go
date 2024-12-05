package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"src/initializers"
	"src/models"
	"src/repository"
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

// RegisterHandler is a handler for POST /register
func RegisterHandler(c *gin.Context) {
	var newUser models.User

	// check incoming data
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error() + "\n"})
		return
	}

	existingUser, err := repository.FindUserByEmail(newUser.Email)
	if err == nil && existingUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email is already in use\n",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error hashing password\n",
		})
		return
	}
	newUser.PasswordHash = string(hashedPassword)
	newUser.RegistrationDate = time.Now()
	if newUser.AccountStatus == "" {
		newUser.AccountStatus = "active"
	}

	if err := repository.CreateUser(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "User registration failed: " + err.Error() + "\n",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Пользователь успешно зарегистрирован",
		"user": gin.H{
			"id":       newUser.UserId,
			"username": newUser.Username,
			"email":    newUser.Email,
		},
	})
}
