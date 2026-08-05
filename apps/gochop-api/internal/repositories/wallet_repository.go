package repositories

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/models"
)

type WalletRepository struct {
	db *pgx.Conn
}

func NewWalletRepository(db *pgx.Conn) *WalletRepository {
	return &WalletRepository{
		db: db,
	}
}

func (r *WalletRepository) GetWalletByUserID(userID string) (*models.Wallet, error) {

	query := `
	SELECT
		id,
		user_id,
		available_balance,
		pending_balance,
		total_earned,
		total_withdrawn,
		currency,
		created_at,
		updated_at
	FROM wallets
	WHERE user_id = $1
	`

	var wallet models.Wallet

	err := r.db.QueryRow(
		context.Background(),
		query,
		userID,
	).Scan(
		&wallet.ID,
		&wallet.UserID,
		&wallet.AvailableBalance,
		&wallet.PendingBalance,
		&wallet.TotalEarned,
		&wallet.TotalWithdrawn,
		&wallet.Currency,
		&wallet.CreatedAt,
		&wallet.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &wallet, nil
}
func (r *WalletRepository) CreateWallet(userID string) error {

	fmt.Println("Creating wallet for:", userID)

	query := `
	INSERT INTO wallets (
		user_id,
		available_balance,
		pending_balance,
		total_earned,
		total_withdrawn,
		currency
	)
	VALUES ($1,0,0,0,0,'NGN')
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		userID,
	)

	if err != nil {
		fmt.Println("CreateWallet ERROR:", err)
	} else {
		fmt.Println("Wallet created successfully")
	}

	return err
}

func (r *WalletRepository) CreditWallet(userID string, amount float64) error {

	query := `
	UPDATE wallets
	SET
		available_balance = available_balance + $1,
		total_earned = total_earned + $1,
		updated_at = NOW()
	WHERE user_id = $2
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		amount,
		userID,
	)

	return err
}
func (r *WalletRepository) DebitWallet(userID string, amount float64) error {

	query := `
	UPDATE wallets
	SET
		available_balance = available_balance - $1,
		total_withdrawn = total_withdrawn + $1,
		updated_at = NOW()
	WHERE user_id = $2
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		amount,
		userID,
	)

	return err
}

func (r *WalletRepository) GetWalletID(userID string) (string, error) {

	var walletID string

	err := r.db.QueryRow(
		context.Background(),
		`SELECT id FROM wallets WHERE user_id=$1`,
		userID,
	).Scan(&walletID)

	return walletID, err
}
