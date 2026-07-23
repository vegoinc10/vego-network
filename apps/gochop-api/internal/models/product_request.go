package models

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Category    string  `json:"category" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Unit        string  `json:"unit" binding:"required"`
	Quantity    int     `json:"quantity" binding:"required"`
	MarketName  string  `json:"market_name"`
	Location    string  `json:"location"`
	ImageURL    string  `json:"image_url"`
}
