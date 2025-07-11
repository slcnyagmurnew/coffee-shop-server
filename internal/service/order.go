package service

import (
	"coffee-server/internal/adapter/order"
	"coffee-server/internal/domain"
)

type OrderService struct {
	orderRepo order.Repository
}

func NewOrderService(repository order.Repository) *OrderService {
	return &OrderService{
		orderRepo: repository,
	}
}

func (orderService *OrderService) CreateOrder(customerId string, orders []domain.OrderItem) error {
	newOrder := domain.NewOrder(customerId)
	for _, item := range orders {
		newOrder.AddItem(item)
	}
	return orderService.orderRepo.Save(newOrder)
}

func (orderService *OrderService) GetOrder(orderId string) (*domain.Order, error) {
	existingOrder, err := orderService.orderRepo.FindById(orderId)
	if err != nil {
		return nil, err
	}
	return existingOrder, nil
}
