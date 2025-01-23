package controllers

import (
    "net/http"
    "github.com/gofiber/fiber/v2"
    "github.com/ayushsarode/SnapNotes/server/models"
    "github.com/ayushsarode/SnapNotes/server/utils"
)

func CreateNote(c *fiber.Ctx) error {
    var note models.Note
    if err := c.BodyParser(&note); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    userID, _ := utils.GetUserIDFromToken(c)
    note.UserID = userID

    if err := models.CreateNote(&note); err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create note"})
    }

    return c.Status(http.StatusOK).JSON(note)
}

func GetNotes(c *fiber.Ctx) error {
    userID, _ := utils.GetUserIDFromToken(c)
    notes, err := models.GetNotesByUser(userID)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch notes"})
    }

    return c.Status(http.StatusOK).JSON(notes)
}

func UpdateNote(c *fiber.Ctx) error {
    noteID := c.Params("id")
    var note models.Note
    if err := c.BodyParser(&note); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    userID, _ := utils.GetUserIDFromToken(c)
    updatedNote, err := models.UpdateNoteByID(userID, noteID, &note)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update note"})
    }

    return c.Status(http.StatusOK).JSON(updatedNote)
}

func DeleteNote(c *fiber.Ctx) error {
    noteID := c.Params("id")
    userID, _ := utils.GetUserIDFromToken(c)
    err := models.DeleteNoteByID(userID, noteID)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete note"})
    }

    return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Note deleted successfully"})
}
