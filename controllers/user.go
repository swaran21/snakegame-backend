package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaran21/snakegame-backend/db"
	"github.com/swaran21/snakegame-backend/models"
)

// Login or register a user with a given username.
func Login(c *gin.Context) {
	var payload struct {
		Username string `json:"username"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil || payload.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
		return
	}

	var user models.User
	result := db.DB.Where("username = ?", payload.Username).First(&user)
	if result.Error != nil {
		// If not found, create a new user
		user = models.User{Username: payload.Username}
		db.DB.Create(&user)
	}

	c.JSON(http.StatusOK, gin.H{
		"username":  user.Username,
		"highScore": user.HighScore,
	})
}

// Update the user's high score if the current score is greater.
func UpdateScore(c *gin.Context) {
	var payload struct {
		Username string `json:"username"`
		Score    int    `json:"score"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil || payload.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username and score are required"})
		return
	}

	var user models.User
	result := db.DB.Where("username = ?", payload.Username).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if payload.Score > user.HighScore {
		user.HighScore = payload.Score
		db.DB.Save(&user)
	}

	c.JSON(http.StatusOK, gin.H{
		"username":  user.Username,
		"highScore": user.HighScore,
	})
}

// GetTopScores returns the top 10 users sorted by high score descending.
func GetTopScores(c *gin.Context) {
	var users []models.User
	result := db.DB.Order("high_score desc").Limit(10).Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch top scores"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"topScores": users,
	})
}
