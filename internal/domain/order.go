package domain

type Order struct {
	Id          string `json:"id"`
	CustomerId  string `json:"customerId"`
	OrderItems  []OrderItem
	Status      string
	TotalAmount float64
}

type OrderRequest struct {
	CustomerId string `json:"customerId"`
	OrderItems []OrderItem
}

func NewOrder(customerId string) *Order {
	return &Order{
		Id:         customerId,
		CustomerId: customerId,
		Status:     "Pending",
	}
}

func (order *Order) AddItem(item OrderItem) {
	order.OrderItems = append(order.OrderItems, item)
	order.TotalAmount += item.Price * float64(item.Quantity)
}
