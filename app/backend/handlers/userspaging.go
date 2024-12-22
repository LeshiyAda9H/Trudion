package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"src/initializers"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	getPageQuery = `
	SELECT * 
    FROM (
        SELECT *, ROW_NUMBER() OVER (ORDER BY user_id) AS row_num
        FROM users
		WHERE user_id <> ?
    ) AS numbered_rows
    WHERE row_num IN ?;
	`
	getUserRow = `
	SELECT row_num
	FROM (
		SELECT user_id, ROW_NUMBER() OVER (ORDER BY user_id) AS row_num
		FROM users
	) AS numbered_rows
	WHERE user_id = ?;
	`
)

var (
	// start      = -1
	// step       = -1
	primeSteps = []int{1, 2, 3, 5, 7, 11, 13, 17, 19}
	stepsCount = 9
	// dbSize     = -1
	pageSize = 3
)

func setStartStep(c *gin.Context) {
	var count int64
	userSession := sessions.Default(c)
	userId, correct := userSession.Get("userId").(uint)
	if !correct {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User ID not found",
		})
	}

	initializers.DB.Table("users").Where("is_banned = ?", false).Where("user_id <> ?", userId).Count(&count)
	if count <= 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "No unbanned users",
		})
		return
	}

	ind := rand.Intn(stepsCount - 1)
	if count == int64(primeSteps[ind]) {
		ind = (ind + 1) % stepsCount
	}

	start := rand.Intn(int(count))
	step := primeSteps[ind]
	dbSize := int(count)
	pageSize, _ = strconv.Atoi(os.Getenv("PAGESIZE"))

	fmt.Println("userid", userSession.Get("userId").(uint))
	userSession.Set("start", start)
	userSession.Set("step", step)
	userSession.Set("dbSize", dbSize)
	err := userSession.Save()
	if err != nil {
		panic(err)
	}

}
