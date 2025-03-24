package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/swaran21/snakegame-backend/controllers"
	"github.com/swaran21/snakegame-backend/db"
)

func main() {
	// Connect to the database
	db.Connect()

	router := gin.Default()

	// API endpoints
	api := router.Group("/api")
	{
		api.POST("/login", controllers.Login)
		api.POST("/updateScore", controllers.UpdateScore)
	}

	// Run server on port 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
