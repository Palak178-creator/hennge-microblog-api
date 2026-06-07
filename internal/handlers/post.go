package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Palak178-creator/hennge-microblog-api/config"
	"github.com/Palak178-creator/hennge-microblog-api/internal/models"
)

// CREATE POST
func CreatePost(c *gin.Context) {

	var input struct {
		Content string `json:"content"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userID := c.GetFloat64("user_id")

	post := models.Post{
		Content: input.Content,
		UserID:  uint(userID),
	}

	config.DB.Create(&post)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Post created successfully",
		"post":    post,
	})
}

// GET ALL POSTS (with user info)
func GetPosts(c *gin.Context) {

	var posts []models.Post

	config.DB.Preload("User").Find(&posts)

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}