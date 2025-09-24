package database

import (
	"fmt"
	"log"
	"os"
	"todo-backend/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}

	// Build DSN
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	// Connect
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate
	err = DB.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Fatal("Failed to migrate:", err)
	}

	log.Println("✅ Database connected!")
}
