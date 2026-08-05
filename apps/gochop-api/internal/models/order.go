package models

import "time"

type Order struct {
	ID             string    `json:"id"`
	BuyerID        string    `json:"buyer_id"`
	Status         string    `json:"status"`
	TotalAmount    float64   `json:"total_amount"`
	PaymentStatus  string    `json:"payment_status"`
	DeliveryStatus string    `json:"delivery_status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type OrderItem struct {
	ID        string `json:"id"`
	OrderID   string `json:"order_id"`
	ProductID string `json:"product_id"`

	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
	Subtotal float64 `json:"subtotal"`

	// Finance Engine (not stored in order_items table)
	StoreID string `json:"store_id,omitempty"`
	OwnerID string `json:"owner_id,omitempty"`
}
