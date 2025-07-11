package router

import (
	"coffee-server/internal/handler"
	"github.com/gofiber/fiber/v2"
)

// Router create routing group for coffee application
func Router(orderHandler *handler.OrderHandler) *fiber.App {
	// init fiber app
	app := fiber.New()
	// init api group
	api := app.Group("/api")
	// add v1 group
	v1 := api.Group("/v1")

	// add index
	v1.Get("/", orderHandler.Index)
	// add healthcheck
	v1.Get("/health", orderHandler.Health)
	// create order
	v1.Post("/order", orderHandler.CreateOrder)
	// get existing order
	v1.Get("/order/:orderId", orderHandler.GetOrder)

	return app
}
