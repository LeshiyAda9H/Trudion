package handlers

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"src/models"
	"src/repository"
	"time"
)

func LoginHandler(c *gin.Context) {
	var loginReq models.LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload\n"})
		return
	}

	token, err := AuthenticateUser(loginReq.Email, loginReq.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token\n"})
		log.Fatalln("Error: ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func AuthenticateUser(email, password string) (string, error) {
	user, err := repository.FindUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password\n")
	}

	// Compare passwords hashes
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", errors.New("invalid email or password\n")
	}

	// generate JWT token
	token, err := generateJWT(user.UserId)
	if err != nil {
		return "", err
	}

	return token, nil
}

func generateJWT(userId uint) (string, error) {
	// TODO: Заменить секретный ключ на безопасный и сохранить его в надежном месте
	secretKey := []byte("qwerty12345")

	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
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
