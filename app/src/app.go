package main

import (
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

	router.GET("/", handlers.RootHandler)
	router.GET("/users", handlers.GetUsers)
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)

	// run server
	port := os.Getenv("PORT")
	log.Printf("Server is running at %s port.", port)
	err := router.Run(":" + port)
	if err != nil {
		log.Fatalln("Can't run server.")
		return
	}
}
