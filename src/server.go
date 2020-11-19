package main

import (
	"log"

	"github.com/bergmi/tag-service/controllers"
	"github.com/bergmi/tag-service/models"

	"github.com/gin-gonic/gin"
	// Import godotenv for .env variables
	"github.com/joho/godotenv"
)

// init gets called before the main function
func init() {
    // Log error if .env file does not exist
    if err := godotenv.Load(); err != nil {
        log.Printf("No .env file found")
    }

		gin.SetMode(gin.ReleaseMode)
}

func main() {
	r := gin.New()
	r.Use(gin.Recovery())

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/tags", controllers.FindTags)
	r.GET("/tags/:postId", controllers.FindTagsForPost)
	r.POST("/tags", controllers.CreateTag)

	// Run the server
	r.Run()
}
