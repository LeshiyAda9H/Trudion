package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"src/initializers"
	"src/models"
	"src/repository"
	"src/types"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	salt      = "hK3B6z4a2pwPUmQdcrCxE9"
	secretKey = "b0nN_4;>-dj!PCtIYdqw0`C1&Uw;gS"
	tokenTTL  = 12 * time.Hour
)

// var (
// 	websocketStorage  sync.Map
// 	websocketUpgrader = websocket.Upgrader{
// 		CheckOrigin: func(r *http.Request) bool {
// 			return true
// 		},
// 	}
// )

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

	// create user session
	userSession := sessions.Default(c)
	userSession.Clear()
	userSession.Set("userId", user.UserId)
	err = userSession.Save()
	if err != nil {
		panic(err)
	}

	//save websocket connection
	// connection, err := websocketUpgrader.Upgrade(c.Writer, c.Request, nil)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "Failed to connect web socket",
	// 	})
	// } else {
	// 	websocketStorage.Store(user.UserId, connection)
	// }
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
		Biography:     body.Biography,
		AccountStatus: "active",
		OnlineStatus:  "online",
	}

	// check if email is already in use
	_, err = repository.GetUserByEmail(body.Email)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is already in use"})
		return
	}

	// Save image
	file, err := c.FormFile("image")
	if err == nil {
		user.Image = filepath.Join("uploads", file.Filename)
		if err := c.SaveUploadedFile(file, user.Image); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
			return
		}
	}

	// Save user to database
	if initializers.DB.Create(&user).Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Process labels if provided
	if len(body.Label) > 0 {
		for _, labelName := range body.Label {
			var label models.LabelInfo

			// Check if label already exists
			if err := initializers.DB.Where("label_name = ?", labelName).First(&label).Error; err != nil {
				// If not found, create a new label
				if errors.Is(err, gorm.ErrRecordNotFound) {
					label = models.LabelInfo{LabelName: labelName}
					if initializers.DB.Create(&label).Error != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create label"})
						return
					}
				} else {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check label existence"})
					return
				}
			}

			// Create UserLabel relationship
			userLabel := models.UserLabel{
				UserID:  user.UserId,
				LabelID: label.LabelID,
			}

			if initializers.DB.Create(&userLabel).Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to associate user with label"})
				return
			}
		}
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
