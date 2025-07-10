package coffee

import (
	"coffee-server/internal/model"
	"log"
	"time"
)

// CreateOrder add new coffee to channel from request
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
			log.Printf("transferred coffee %d", o.Id)
		default:
			log.Println("waiting for orders in wait queue...")
		}
	}
}
