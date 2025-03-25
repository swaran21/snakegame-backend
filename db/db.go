package db

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/swaran21/snakegame-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect function loads environment variables and establishes a database connection
func Connect() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Retrieve environment variables
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	// Create the DSN (Data Source Name) string
	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=" + sslmode

	// Open a connection to the database
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
