package main

import (
	"coffee-server/internal/router"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// signal catcher for sigterm
	_, cancel := context.WithCancel(context.Background())

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigs
		log.Println("shutdown signal received")
		cancel()
	}()

	//setup router
	r := router.Router()
	//start listen
	log.Fatal(r.Listen(":3000"))
}
