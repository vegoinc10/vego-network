package models

type WithdrawalRequest struct {
	Amount float64 `json:"amount" binding:"required"`

	BankName string `json:"bank_name" binding:"required"`

	AccountName string `json:"account_name" binding:"required"`

	AccountNumber string `json:"account_number" binding:"required"`
}
