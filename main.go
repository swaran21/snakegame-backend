package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaran21/snakegame-backend/controllers"
	"github.com/swaran21/snakegame-backend/db"
)

func main() {
	// Connect to the database
	db.Connect()

	router := gin.Default()

	// Configure CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://snake-game-frontend-flgk.onrender.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// API endpoints
	api := router.Group("/api")
	{
		api.POST("/login", controllers.Login)
		api.POST("/updateScore", controllers.UpdateScore)
		api.GET("/topScores", controllers.GetTopScores) // New endpoint for top scores
	}

	// Run server on port 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
