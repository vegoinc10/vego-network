package models

import "time"

type Product struct {
	ID       string `json:"id"`
	SellerID string `json:"seller_id"`

	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`

	Price    float64 `json:"price"`
	Currency string  `json:"currency"`

	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`

	MarketName string `json:"market_name"`
	State      string `json:"state"`
	LGA        string `json:"lga"`

	ImageURL string `json:"image_url"`

	Status string `json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
