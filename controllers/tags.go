package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/bergmi/tag-service/models"
)

type CreateTagInput struct {
	Name    string `json:"name" binding:"required"`
	Post 	  uint   `json:"post" binding:"required"`
}

type UpdateTagInput struct {
	Name    string `json:"name"`
	Post 	  uint   `json:"post"`
}

// GET /tags
// Find all tags
func FindTags(c *gin.Context) {
	var tags []models.Tag
	models.DB.Find(&tags)

	c.JSON(http.StatusOK, gin.H{"data": tags})
}

// GET /tags/:postId
// Find a tag
func FindTagsForPost(c *gin.Context) {
	// Get model if exist
	var tags []models.Tag
	if err := models.DB.Where("post = ?", c.Param("postId")).Find(&tags).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tags})
}

// POST /tags
// Create new tags
func CreateTag(c *gin.Context) {
	// Validate input
	var input CreateTagInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create 5qt
	tag := models.Tag{Name: input.Name, Post: input.Post}
	models.DB.Create(&tag)

	c.JSON(http.StatusOK, gin.H{"data": tag})
}
