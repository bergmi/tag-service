package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.zhaw.com/stoppdan/SCAD-HS20-Team1/P06/tag-service/models"
)

type CreateTagInput struct {
	Name    string `json:"name" binding:"required"`
	PostId 	int    `json:"postId" binding:"required"`
}

type UpdateTagInput struct {
	Name    string `json:"name"`
	PostId 	uint   `json:"postId"`
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
	var tag models.Tag
	if err := models.DB.Where("postId = ?", c.Param("postId")).First(&tag).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tag})
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
	tag := models.Tag{Name: input.Name, PostId: input.PostId}
	models.DB.Create(&tag)

	c.JSON(http.StatusOK, gin.H{"data": tag})
}

// DELETE /tags/:id
// Delete a tag
func DeleteTag(c *gin.Context) {
	// Get model if exist
	var tag models.Tag
	if err := models.DB.Where("id = ?", c.Param("id")).First(&tag).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&tag)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// DELETE /tags/:postId
// Delete a tag for postId
func DeleteTagForPostId(c *gin.Context) {
	// Get model if exist
	var tag models.Tag
	if err := models.DB.Where("postId = ?", c.Param("postId")).First(&tag).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&tag)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
