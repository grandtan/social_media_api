package controllers

import (
	"net/http"
	"social_media_app/database"
	"social_media_app/models"
	"social_media_app/utils"

	"github.com/gin-gonic/gin"
)

// CreatePost handles the creation of a new post
func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid input")
		return
	}

	if post.UserID == 0 || post.Content == "" {
		utils.RespondError(c, http.StatusBadRequest, "UserID and Content are required")
		return
	}

	if err := database.DB.Create(&post).Error; err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusOK, post)
}

// GetPost handles fetching a post by ID
func GetPost(c *gin.Context) {
	var post models.Post
	if err := database.DB.First(&post, c.Param("id")).Error; err != nil {
		utils.RespondError(c, http.StatusNotFound, "Post not found")
		return
	}
	utils.RespondJSON(c, http.StatusOK, post)
}

// GetPosts handles fetching all posts
func GetPosts(c *gin.Context) {
	var posts []models.Post
	database.DB.Find(&posts)
	utils.RespondJSON(c, http.StatusOK, posts)
}

// UpdatePost handles updating a post by ID
func UpdatePost(c *gin.Context) {
	var post models.Post
	if err := database.DB.First(&post, c.Param("id")).Error; err != nil {
		utils.RespondError(c, http.StatusNotFound, "Post not found")
		return
	}

	if err := c.ShouldBindJSON(&post); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid input")
		return
	}

	database.DB.Save(&post)
	utils.RespondJSON(c, http.StatusOK, post)
}

// DeletePost handles deleting a post by ID
func DeletePost(c *gin.Context) {
	var post models.Post
	if err := database.DB.Delete(&post, c.Param("id")).Error; err != nil {
		utils.RespondError(c, http.StatusNotFound, "Post not found")
		return
	}
	utils.RespondJSON(c, http.StatusOK, "Post deleted")
}
