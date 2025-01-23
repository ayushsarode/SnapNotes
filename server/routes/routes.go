package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/ayushsarode/SnapNotes/server/controllers"
    "github.com/ayushsarode/SnapNotes/server/middleware"
)

func Setup(app *fiber.App) {
    api := app.Group("/api")

    api.Post("/register", controllers.Register)  // Register user
    api.Post("/login", controllers.Login)        // User login
    api.Get("/notes", middleware.AuthRequired, controllers.GetNotes)
    api.Post("/notes", middleware.AuthRequired, controllers.CreateNote)
    api.Put("/notes/:id", middleware.AuthRequired, controllers.UpdateNote)
    api.Delete("/notes/:id", middleware.AuthRequired, controllers.DeleteNote)
}
