package router

import (
	"coffee-server/internal/handler"
	"github.com/gofiber/fiber/v3"
)

// Router create routing group for coffee application
func Router() *fiber.App {
	// init fiber app
	app := fiber.New()
	// init api group
	api := app.Group("/api")
	// add v1 group
	v1 := api.Group("/v1")

	// add index
	v1.Get("/", handler.Index)
	// add healthcheck
	v1.Get("/health", handler.Health)

	return app
}
