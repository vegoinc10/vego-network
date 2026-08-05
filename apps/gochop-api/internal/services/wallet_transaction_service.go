package services

import (
	"github.com/google/uuid"

	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/models"
	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/repositories"
)

type WalletTransactionService struct {
	repo *repositories.WalletTransactionRepository
}

func NewWalletTransactionService(repo *repositories.WalletTransactionRepository) *WalletTransactionService {
	return &WalletTransactionService{
		repo: repo,
	}
}

func (s *WalletTransactionService) Record(
	walletID string,
	orderID string,
	txType string,
	amount float64,
	description string,
) error {

	tx := &models.WalletTransaction{
		ID:          uuid.New().String(),
		WalletID:    walletID,
		OrderID:     orderID,
		Type:        txType,
		Amount:      amount,
		Description: description,
	}

	return s.repo.Create(tx)
}
