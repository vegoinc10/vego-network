package models

import "time"

type CommissionSetting struct {
	ID                string    `json:"id"`
	CommissionRate    float64   `json:"commission_rate"`
	MinimumWithdrawal float64   `json:"minimum_withdrawal"`
	WithdrawalFee     float64   `json:"withdrawal_fee"`
	Active            bool      `json:"active"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
