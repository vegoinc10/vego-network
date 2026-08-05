package models

import "time"

type WalletTransaction struct {
	ID       string `json:"id"`
	WalletID string `json:"wallet_id"`
	OrderID  string `json:"order_id"`

	Type        string  `json:"type"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`

	CreatedAt time.Time `json:"created_at"`
}
