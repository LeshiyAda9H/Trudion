package main

import (
	"github.com/gin-gonic/gin"
	"log"
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
	router.POST("/register", handlers.RegisterHandler)
	router.POST("/login", handlers.LoginHandler)

	// run server
	log.Println("Server is running at 8080 port.")
	err := router.Run(":8080")
	if err != nil {
		log.Fatalln("Can't run server.")
		return
	}
}
