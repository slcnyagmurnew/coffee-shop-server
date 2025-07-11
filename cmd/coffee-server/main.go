package main

import (
	"coffee-server/internal/adapter/order"
	"coffee-server/internal/handler"
	"coffee-server/internal/router"
	"coffee-server/internal/service"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"log"
	"net/http"
	"time"
)

func main() {
	// init order repo
	orderRepo := order.NewCacheRepository()
	// init order service
	orderService := service.NewOrderService(orderRepo)
	// init order handler
	orderHandler := handler.NewOrderHandler(orderService)
	//setup router
	app := router.Router(orderHandler)
	// build HTTP server instead of r.Run()
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      adaptor.FiberApp(app),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Printf("🚀 Server is starting on %s\n", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
