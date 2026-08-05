package models

import "time"

type CartItem struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	ProductID string    `json:"product_id"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AddToCartRequest struct {
	ProductID string `json:"product_id" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required"`
}

type UpdateCartRequest struct {
	Quantity int `json:"quantity" binding:"required"`
}
type CartResponse struct {
	CartID      string  `json:"cart_id"`
	ProductID   string  `json:"product_id"`
	StoreID     string  `json:"store_id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Subtotal    float64 `json:"subtotal"`
}
