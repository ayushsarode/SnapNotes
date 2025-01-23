package controllers

import (
    "net/http"
    "github.com/ayushsarode/SnapNotes/server/models"
    "github.com/ayushsarode/SnapNotes/server/utils"
    "github.com/gofiber/fiber/v2"
    "golang.org/x/crypto/bcrypt"
)

// Register user
func Register(c *fiber.Ctx) error {
    var user models.User
    if err := c.BodyParser(&user); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    // Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
    }
    user.Password = string(hashedPassword)

    // Create user in the database
    if err := models.CreateUser(&user); err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
    }

    return c.Status(http.StatusOK).JSON(fiber.Map{"message": "User registered successfully"})
}

// Login user and generate JWT
func Login(c *fiber.Ctx) error {
    var user models.User
    var input models.User

    if err := c.BodyParser(&input); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    // Find user by email
    if err := models.GetUserByEmail(&user, input.Email); err != nil {
        return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
    }

    // Compare password
    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
    if err != nil {
        return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
    }

    // Generate JWT token
    token, err := utils.GenerateJWT(user.ID)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
    }

    return c.Status(http.StatusOK).JSON(fiber.Map{"token": token})
}
