package coffee

import (
	"coffee-server/internal/model"
	"context"
	"log"
	"sync"
	"time"
)

type Shop struct {
	Id          int
	Name        string
	Location    string
	WorkerCount int
	IsActive    bool

	wg        sync.WaitGroup
	stop      chan struct{}
	done      chan struct{}
	orders    chan model.Order
	waitQueue chan model.Order
}

func New(id int, name string, location string, workerCount int, isActive bool) (*Shop, error) {
	// todo check active shop ids if exists, then warn (from shop repo)
	return &Shop{
		Id:          id,
		Name:        name,
		Location:    location,
		WorkerCount: workerCount,
		IsActive:    isActive,

		wg:        sync.WaitGroup{},
		stop:      make(chan struct{}),
		done:      make(chan struct{}),
		orders:    make(chan model.Order, workerCount),
		waitQueue: make(chan model.Order),
	}, nil

}

func (shop *Shop) Worker(ctx context.Context) {
	defer shop.wg.Done()

	for {
		select {
		case <-ctx.Done():
			log.Println("termination signal received")
			return
		case <-shop.stop:
			log.Printf("worker %d: stop signal received, exiting\n", shop.Id)
			return
		case o, ok := <-shop.orders:
			if !ok {
				log.Println("channel closed, no more orders")
				return
			}
			log.Printf("coffee %d preparing in Shop %d...", o.Id, shop.Id)
			time.Sleep(10 * time.Second)
		default:
			log.Println("no orders in queue, channel is empty")
			time.Sleep(2 * time.Second)
		}
	}
}

func (shop *Shop) Start(ctx context.Context) {
	if shop.stop != nil {
		// already running
	}
	for i := 0; i < shop.WorkerCount; i++ {
		shop.wg.Add(1)
		go shop.Worker(ctx)
	}
	go func() {
		shop.wg.Wait()   // wait for all workers to return
		close(shop.done) // signal completion
	}()
}

func (shop *Shop) Stop() {
	close(shop.stop)
	close(shop.orders)
}

func (shop *Shop) Wait() {
	<-shop.done
}
