package models

import "time"

type PlatformRevenue struct {
	ID string `json:"id"`

	OrderID string `json:"order_id"`

	VendorID string `json:"vendor_id"`

	GrossAmount float64 `json:"gross_amount"`

	CommissionRate float64 `json:"commission_rate"`

	CommissionAmount float64 `json:"commission_amount"`

	VendorAmount float64 `json:"vendor_amount"`

	CreatedAt time.Time `json:"created_at"`
}
