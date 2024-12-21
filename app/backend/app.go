package main

import (
	"log"
	"net/http"
	"os"
	"src/handlers"
	"src/initializers"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "src/docs"
)

func init() {
	initializers.LoadConfig()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

// @title	Trudion API
// @version 1.0
// @description API Service for Trudion Application
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	storage := cookie.NewStore([]byte("secret"))
	storage.Options(sessions.Options{
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
	})
	router.Use(sessions.Sessions("user_session", storage))

	// Add swagger documentation
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/api/v1", handlers.RootHandler)
	router.GET("/api/v1/users", handlers.GetUsers)
	router.GET("/api/v1/profile", handlers.GetUserProfile)
	router.GET("/api/v1/usersnumber", handlers.GetUsersNumber)
	router.GET("/api/v1/userspage", handlers.GetUsersPage)

	router.POST("/api/v1/handshake", handlers.Handshake)
	router.POST("/api/v1/register", handlers.SignUp)
	router.POST("/api/v1/login", handlers.SignIn)
	router.POST("/api/v1/verify/email", handlers.VerifyEmail)

	router.PATCH("/api/v1/profile", handlers.UpdateProfile)

	router.DELETE("/api/v1/deleteuser", handlers.DeleteUser)

	// run server
	port := os.Getenv("PORT")
	log.Printf("Server is running at %s port.", port)
	err := router.Run(":" + port)
	if err != nil {
		log.Fatalln("Can't run server.")
		return
	}
}
