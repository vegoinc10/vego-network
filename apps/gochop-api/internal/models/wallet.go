package models

import "time"

type Wallet struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`

	AvailableBalance float64 `json:"available_balance"`
	PendingBalance   float64 `json:"pending_balance"`

	TotalEarned    float64 `json:"total_earned"`
	TotalWithdrawn float64 `json:"total_withdrawn"`

	Currency string `json:"currency"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
