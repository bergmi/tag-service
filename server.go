package main

import (
	"github.zhaw.com/stoppdan/SCAD-HS20-Team1/P06/tag-service/controllers"
	"github.zhaw.com/stoppdan/SCAD-HS20-Team1/P06/tag-service/models"

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
	r.DELETE("/tags/:id", controllers.DeleteT)

	// Run the server
	r.Run()
}
