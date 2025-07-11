package order

import (
	"coffee-server/internal/domain"
	"fmt"
	"sync"
)

type CacheRepository struct {
	orders map[string]*domain.Order
	mu     sync.Mutex
}

func NewCacheRepository() *CacheRepository {
	return &CacheRepository{
		orders: make(map[string]*domain.Order),
	}
}

func (r *CacheRepository) Save(order *domain.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.orders[order.Id] = order
	return nil
}

func (r *CacheRepository) FindById(id string) (order *domain.Order, error error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	o, exists := r.orders[id]
	if !exists {
		return nil, fmt.Errorf("order %s could not found", order.Id)
	}
	return o, nil
}
