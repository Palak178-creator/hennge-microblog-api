package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Palak178-creator/hennge-microblog-api/config"
	"github.com/Palak178-creator/hennge-microblog-api/internal/models"
)

func Profile(c *gin.Context) {

	userID := c.GetFloat64("user_id")

	var user models.User

	if err := config.DB.First(&user, uint(userID)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
	})
}