package services

import (
	"errors"

	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/models"
	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/repositories"
)

type WithdrawalService struct {
	walletRepo      *repositories.WalletRepository
	transactionRepo *repositories.WalletTransactionRepository
	withdrawalRepo  *repositories.WithdrawalRepository
}

func NewWithdrawalService(
	walletRepo *repositories.WalletRepository,
	transactionRepo *repositories.WalletTransactionRepository,
	withdrawalRepo *repositories.WithdrawalRepository,
) *WithdrawalService {

	return &WithdrawalService{
		walletRepo:      walletRepo,
		transactionRepo: transactionRepo,
		withdrawalRepo:  withdrawalRepo,
	}
}

func (s *WithdrawalService) RequestWithdrawal(
	userID string,
	req *models.WithdrawalRequest,
) error {

	// Load wallet
	wallet, err := s.walletRepo.GetWalletByUserID(userID)
	if err != nil {
		return err
	}

	// Check balance
	if wallet.AvailableBalance < req.Amount {
		return errors.New("insufficient wallet balance")
	}

	// Debit wallet
	err = s.walletRepo.DebitWallet(userID, req.Amount)
	if err != nil {
		return err
	}

	// Save withdrawal request
	withdrawal := &models.Withdrawal{
		WalletID:      wallet.ID,
		Amount:        req.Amount,
		BankName:      req.BankName,
		AccountName:   req.AccountName,
		AccountNumber: req.AccountNumber,
		Status:        "pending",
	}

	err = s.withdrawalRepo.Create(withdrawal)
	if err != nil {
		return err
	}

	// Record wallet transaction
	tx := &models.WalletTransaction{
		WalletID:    wallet.ID,
		Type:        "withdrawal",
		Amount:      req.Amount,
		Description: "Withdrawal Request",
	}

	return s.transactionRepo.Create(tx)
}
