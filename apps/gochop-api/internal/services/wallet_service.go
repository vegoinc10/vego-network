package services

import (
	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/models"
	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/repositories"
)

type WalletService struct {
	walletRepo *repositories.WalletRepository
}

func NewWalletService(walletRepo *repositories.WalletRepository) *WalletService {
	return &WalletService{
		walletRepo: walletRepo,
	}
}

func (s *WalletService) CreateWallet(userID string) error {
	return s.walletRepo.CreateWallet(userID)
}

func (s *WalletService) GetWallet(userID string) (*models.Wallet, error) {
	return s.walletRepo.GetWalletByUserID(userID)
}
