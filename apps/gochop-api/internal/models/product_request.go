package models

type CreateProductRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Category    string `json:"category" binding:"required"`

	Price    float64 `json:"price" binding:"required"`
	Currency string  `json:"currency"`

	Quantity int    `json:"quantity" binding:"required"`
	Unit     string `json:"unit" binding:"required"`

	MarketName string `json:"market_name"`
	State      string `json:"state"`
	LGA        string `json:"lga"`

	ImageURL string `json:"image_url"`
}
