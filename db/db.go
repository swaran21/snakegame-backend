package db

import (
	"log"

	"github.com/swaran21/snakegame-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=aizen dbname=snakegame_db port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	DB = database

	// Auto-migrate the User model
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to auto-migrate:", err)
	}

	log.Println("Database connection established.")
}
