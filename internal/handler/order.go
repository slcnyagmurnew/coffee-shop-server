package handler

import (
	"coffee-server/internal/domain"
	"coffee-server/internal/service"
	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	orderService *service.OrderService
}

func NewOrderHandler(orderService *service.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

func (orderHandler *OrderHandler) CreateOrder(ctx *fiber.Ctx) error {
	req := new(domain.OrderRequest)

	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request payload"})
	}

	err := orderHandler.orderService.CreateOrder(req.CustomerId, req.OrderItems)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok", "message": "order created"})
}

// Index page for app
func (orderHandler *OrderHandler) Index(ctx *fiber.Ctx) error {
	return ctx.SendString("Hi from Coffee Shop")
}

// Health check endpoint
func (orderHandler *OrderHandler) Health(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "healthy",
		"status":  "ok",
	})
}

// GetOrder check endpoint
func (orderHandler *OrderHandler) GetOrder(ctx *fiber.Ctx) error {
	orderId := ctx.Params("orderId")
	if orderId == "" {
		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "orderId is required"})
	}
	order, err := orderHandler.orderService.GetOrder(orderId)
	if err != nil {
		// maybe translate domain errors to 404 vs 500, etc.
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	if order == nil {
		return ctx.
			Status(fiber.StatusNotFound).
			JSON(fiber.Map{"error": "order not found"})
	}
	return ctx.JSON(order)
}
