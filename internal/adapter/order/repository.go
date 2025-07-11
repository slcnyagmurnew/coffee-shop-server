package order

import "coffee-server/internal/domain"

type Repository interface {
	Save(order *domain.Order) error
	FindById(id string) (*domain.Order, error)
}
