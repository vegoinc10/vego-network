package models

import "time"

type Commission struct {
	ID               string    `json:"id"`
	OrderID          string    `json:"order_id"`
	VendorID         string    `json:"vendor_id"`
	CommissionRate   float64   `json:"commission_rate"`
	CommissionAmount float64   `json:"commission_amount"`
	VendorAmount     float64   `json:"vendor_amount"`
	Status           string    `json:"status"`
	CreatedAt        time.Time `json:"created_at"`
}
