package handler

import "github.com/gofiber/fiber/v2"

// Index page for app
func Index(c *fiber.Ctx) error {
	return c.SendString("Hi from Coffee Shop")
}

// Health check endpoint
func Health(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "healthy",
		"status":  "ok",
	})
}
