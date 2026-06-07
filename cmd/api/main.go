package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/Palak178-creator/hennge-microblog-api/config"
	"github.com/Palak178-creator/hennge-microblog-api/internal/handlers"
	"github.com/Palak178-creator/hennge-microblog-api/internal/middleware"
	"github.com/Palak178-creator/hennge-microblog-api/internal/models"
)

func main() {
	config.InitDB()

	err := config.DB.AutoMigrate(
		&models.User{},
		&models.Post{},
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Database migrated successfully!")

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "HENNGE Microblog API Running",
		})
	})

	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)

	router.GET("/profile", middleware.AuthMiddleware(), handlers.Profile)

	router.POST("/posts", middleware.AuthMiddleware(), handlers.CreatePost)
	router.GET("/posts", handlers.GetPosts)

	log.Println("Server running on port 8080")
	router.Run(":8080")
}