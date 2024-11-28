package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"src/handlers"
	"src/models"
	"src/repository"
)

func DBMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		dbConn, err := repository.InitializeDatabase()
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to connect to the database\n"})
			c.Abort()
			return
		}

		err = models.MigrateAll(dbConn)
		if err != nil {
			log.Fatal("Migration error: ", err, "\n")
		}

		c.Set("dbConn", dbConn)
		c.Next()
	}
}

func main() {
	router := gin.Default()

	router.Use(DBMiddleware())

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
