package main

import (
	"github.com/bergmi/tag-service/controllers"
	"github.com/bergmi/tag-service/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/tags", controllers.FindTags)
	r.GET("/tags/:postId", controllers.FindTagsForPost)
	r.POST("/tags", controllers.CreateTag)

	// Run the server
	r.Run()
}
