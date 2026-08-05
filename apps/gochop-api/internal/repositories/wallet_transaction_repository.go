package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/models"
)

type WalletTransactionRepository struct {
	db *pgx.Conn
}

func NewWalletTransactionRepository(db *pgx.Conn) *WalletTransactionRepository {
	return &WalletTransactionRepository{
		db: db,
	}
}

func (r *WalletTransactionRepository) Create(tx *models.WalletTransaction) error {

	query := `
	INSERT INTO wallet_transactions
	(
		wallet_id,
		order_id,
		type,
		amount,
		description
	)
	VALUES ($1,$2,$3,$4,$5)
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		tx.WalletID,
		tx.OrderID,
		tx.Type,
		tx.Amount,
		tx.Description,
	)

	return err
}

func (r *WalletTransactionRepository) GetTransactionsByWallet(
	walletID string,
) ([]models.WalletTransaction, error) {

	query := `
	SELECT
		id,
		wallet_id,
		order_id,
		type,
		amount,
		description,
		created_at
	FROM wallet_transactions
	WHERE wallet_id = $1
	ORDER BY created_at DESC
	`

	rows, err := r.db.Query(
		context.Background(),
		query,
		walletID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []models.WalletTransaction

	for rows.Next() {

		var tx models.WalletTransaction

		err := rows.Scan(
			&tx.ID,
			&tx.WalletID,
			&tx.OrderID,
			&tx.Type,
			&tx.Amount,
			&tx.Description,
			&tx.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		transactions = append(transactions, tx)
	}

	return transactions, nil
}
