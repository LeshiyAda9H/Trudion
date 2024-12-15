package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	usersNumber         = "usersNumber"
)

func UserIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "empty authorization header"})
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header"})
		return
	}

	// parse token
	userId, err := parseToken(headerParts[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.Set(userCtx, userId)
}

// func UsersCount(c *gin.Context) {
// 	header := c.GetHeader("GetUsersNumber")
// 	if header == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "empty number users header"})
// 	}

// 	headerParts := strings.Split(header, " ")
// 	if len(headerParts) != 2 {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid number users header"})
// 		return
// 	}

// 	usersInputNumber, err := strconv.Atoi(headerParts[1])
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
// 	}

// 	c.Set(usersNumber, usersInputNumber)

// }
