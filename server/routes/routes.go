package routes

import (
	"github.com/gofiber/fiber/v2"
	"note-taking-app/controllers"
	"note-taking-app/middleware"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/notes", middleware.Protected(), controllers.CreateNote)
	app.Get("/notes", middleware.Protected(), controllers.GetNotes)
	app.Put("/notes/:id", middleware.Protected(), controllers.UpdateNote)
	app.Delete("/notes/:id", middleware.Protected(), controllers.DeleteNote)
}
