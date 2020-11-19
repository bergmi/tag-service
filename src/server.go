package main

import (
	"os"

	"github.com/bergmi/tag-service/controllers"
	"github.com/bergmi/tag-service/models"

	"github.com/gin-gonic/gin"
)

// init gets called before the main function
func init() {
	gin.SetMode(gin.ReleaseMode)
	os.Setenv("PORT", "3000")
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
