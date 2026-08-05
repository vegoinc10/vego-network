package models

import "time"

type Withdrawal struct {
	ID            string     `json:"id"`
	WalletID      string     `json:"wallet_id"`
	Amount        float64    `json:"amount"`
	BankName      string     `json:"bank_name"`
	AccountName   string     `json:"account_name"`
	AccountNumber string     `json:"account_number"`
	Status        string     `json:"status"`
	Reference     string     `json:"reference"`
	FailureReason string     `json:"failure_reason"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	ProcessedAt   *time.Time `json:"processed_at,omitempty"`
}
