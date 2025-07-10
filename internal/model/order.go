package model

type Order struct {
	Id      int    `json:"id"`
	Type    string `json:"type"`
	ShopId  int    `json:"shopId"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type OrderRequest struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	ShopId  int    `json:"shopId"`
}
