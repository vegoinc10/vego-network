package services

import (
	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/models"
	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/repositories"
)

type CommissionService struct {
	walletRepo      *repositories.WalletRepository
	transactionRepo *repositories.WalletTransactionRepository
}

func NewCommissionService(
	walletRepo *repositories.WalletRepository,
	transactionRepo *repositories.WalletTransactionRepository,
) *CommissionService {

	return &CommissionService{
		walletRepo:      walletRepo,
		transactionRepo: transactionRepo,
	}
}

// CalculateCommission calculates the platform commission and vendor earnings.
func (s *CommissionService) CalculateCommission(
	amount float64,
	rate float64,
) (commission float64, vendorAmount float64) {

	commission = amount * (rate / 100)
	vendorAmount = amount - commission

	return
}

func (s *CommissionService) CreditVendor(
	userID string,
	amount float64,
) error {

	return s.walletRepo.CreditWallet(userID, amount)
}

func (s *CommissionService) RecordWalletTransaction(
	walletID string,
	orderID string,
	amount float64,
	description string,
) error {

	tx := &models.WalletTransaction{
		WalletID:    walletID,
		OrderID:     orderID,
		Type:        "credit",
		Amount:      amount,
		Description: description,
	}

	return s.transactionRepo.Create(tx)
}

func (s *CommissionService) SaveRevenue() error {
	return nil
}
