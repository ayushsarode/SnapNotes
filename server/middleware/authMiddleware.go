package middleware

import (
    "github.com/gofiber/fiber/v2"
    "github.com/ayushsarode/SnapNotes/server/utils"
    "strings"
)

func AuthRequired(c *fiber.Ctx) error {
    token := c.Get("Authorization")
    token = strings.Replace(token, "Bearer ", "", 1)
    if token == "" {
        return c.Status(401).JSON(fiber.Map{"error": "No token provided"})
    }

    userID, err := utils.ValidateJWT(token)
    if err != nil {
        return c.Status(401).JSON(fiber.Map{"error": "Invalid or expired token"})
    }

    c.Locals("userID", userID)
    return c.Next()
}
