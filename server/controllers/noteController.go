package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"note-taking-app/models"
	"note-taking-app/config"
	"strconv"
	"github.com/gofiber/fiber/v2/middleware/jwt"
)

// Create a new note
func CreateNote(c *fiber.Ctx) error {
	var note models.Note
	if err := c.BodyParser(&note); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user := c.Locals("user")
	note.UserID = user.(uint)

	if err := config.DB.Create(&note).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(note)
}

// Get all notes for the current user
func GetNotes(c *fiber.Ctx) error {
	user := c.Locals("user")

	var notes []models.Note
	if err := config.DB.Where("user_id = ?", user.(uint)).Find(&notes).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(notes)
}

// Update a note
func UpdateNote(c *fiber.Ctx) error {
	id := c.Params("id")
	var note models.Note
	if err := config.DB.First(&note, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Note not found"})
	}

	user := c.Locals("user")
	if note.UserID != user.(uint) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "You are not authorized"})
	}

	if err := c.BodyParser(&note); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	config.DB.Save(&note)
	return c.Status(fiber.StatusOK).JSON(note)
}

// Delete a note
func DeleteNote(c *fiber.Ctx) error {
	id := c.Params("id")
	var note models.Note
	if err := config.DB.First(&note, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Note not found"})
	}

	user := c.Locals("user")
	if note.UserID != user.(uint) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "You are not authorized"})
	}

	if err := config.DB.Delete(&note).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Note deleted"})
}
