package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"src/handlers"
	"src/initializers"
)

func init() {
	initializers.LoadConfig()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	router.GET("/", handlers.RootHandler)
	router.GET("/users", handlers.GetUsers)
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)
	router.POST("/verify/email", handlers.VerifyEmail)
	//router.POST("/verify/token", handlers.VerifyToken)

	// run server
	port := os.Getenv("PORT")
	log.Printf("Server is running at %s port.", port)
	err := router.Run(":" + port)
	if err != nil {
		log.Fatalln("Can't run server.")
		return
	}
}
