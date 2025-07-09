package ordering

import (
	"coffee-server/internal/model"
	"context"
	"log"
	"sync"
	"time"
)

// MakeCoffee receive order channel and make coffee
// todo add cfg for each coffee with timing
func MakeCoffee(ctx context.Context, c <-chan model.Order, w *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			log.Println("Termination signal received")
			return
		case o, ok := <-c:
			if !ok {
				log.Println("channel closed, no more orders")
				return
			}
			log.Printf("Coffee %d preparing...", o.Id)
			time.Sleep(5 * time.Second)
			w.Done()
		default:
			log.Println("no orders in queue, channel is empty")
			time.Sleep(2 * time.Second)
		}
	}
}

// CreateOrder add new order to channel from request
func CreateOrder(c chan<- model.Order, o model.Order, wc chan<- model.Order) {
	select {
	case c <- o:
		log.Printf("Order %d request received", o.Id)
	case wc <- o:
		log.Printf("Order %d request sent to wait queue", o.Id)
	}
}

// CleanWaitQueue periodically
func CleanWaitQueue(wc chan model.Order, c chan model.Order) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop() // stop ticker eventually

	for {
		select {
		case <-ticker.C:
			o, ok := <-wc
			if !ok {
				log.Println("source channel closed, no more orders")
				return
			}
			c <- o // ← send (blocks until dst has capacity)
			log.Printf("transferred order %d", o.Id)
		default:
			log.Println("waiting for orders in wait queue...")
		}
	}
}
