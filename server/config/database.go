package config

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
)

var DB *gorm.DB

func SetupDatabase() {
	// Load .env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get DB connection details from .env
	dsn := "user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " host=" + os.Getenv("DB_HOST") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable"
	var err2 error
	DB, err2 = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err2 != nil {
		log.Fatal("Failed to connect to the database")
	}

	// Migrate the database
	DB.AutoMigrate(&Note{}, &User{})
}
