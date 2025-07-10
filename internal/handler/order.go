package handler

import (
	"coffee-server/internal/model"
	"github.com/gofiber/fiber/v2"
	"log"
)

func OrderCoffee(c *fiber.Ctx) error {
	var orderRequest model.OrderRequest

	if err := c.BodyParser(&orderRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	log.Printf("Got coffee – type=%q, message=%q",
		orderRequest.Type, orderRequest.Message)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "coffee received",
		"status":  "ok",
	})
}
