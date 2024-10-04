package main

import (
	"log"

	"crud_go/config"
	database "crud_go/db"
	"crud_go/handlers"
	"crud_go/models"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db, err := database.Connect(cfg.DBUser, cfg.DBPassword, cfg.DBName)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	if err := db.AutoMigrate(&models.Post{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	r := gin.Default()
	postHandler := handlers.PostHandler{DB: db}

	r.POST("/posts", postHandler.CreatePost)
	r.GET("/posts", postHandler.GetPosts)
	r.GET("/posts/:id", postHandler.GetPost)
	r.PUT("/posts/:id", postHandler.UpdatePost)
	r.DELETE("/posts/:id", postHandler.DeletePost)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	log.Println("Server running on http://localhost:8080")
}
