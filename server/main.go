package main

import (
	"log"

	"github.com/ayushsarode/SnapNotes/server/config"
	"github.com/ayushsarode/SnapNotes/server/routes"

	"github.com/ayushsarode/SnapNotes/server/models"

	"github.com/gofiber/fiber/v2"
)

func main() {
    // Load environment variables
    config.LoadEnv()

    // Initialize database connection
    config.InitDB()

    // Auto migrate models (create tables if they don't exist)
    config.GetDB().AutoMigrate(&models.User{}, &models.Note{})

    app := fiber.New()

    // Setup routes
    routes.Setup(app)

    // Start server
    log.Fatal(app.Listen(":3000"))
}
