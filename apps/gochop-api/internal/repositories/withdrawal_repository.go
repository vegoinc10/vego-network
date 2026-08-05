package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/models"
)

type WithdrawalRepository struct {
	db *pgx.Conn
}

func NewWithdrawalRepository(db *pgx.Conn) *WithdrawalRepository {
	return &WithdrawalRepository{
		db: db,
	}
}

func (r *WithdrawalRepository) Create(
	withdrawal *models.Withdrawal,
) error {

	query := `
	INSERT INTO withdrawals (
		wallet_id,
		amount,
		bank_name,
		account_name,
		account_number,
		status,
		reference,
		failure_reason,
		created_at,
		processed_at,
		updated_at
	)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		withdrawal.WalletID,
		withdrawal.Amount,
		withdrawal.BankName,
		withdrawal.AccountName,
		withdrawal.AccountNumber,
		withdrawal.Status,
		withdrawal.Reference,
		withdrawal.FailureReason,
		withdrawal.CreatedAt,
		withdrawal.ProcessedAt,
		withdrawal.UpdatedAt,
	)

	return err
}

func (r *WithdrawalRepository) GetByWalletID(
	walletID string,
) ([]models.Withdrawal, error) {

	query := `
	SELECT
		id,
		wallet_id,
		amount,
		bank_name,
		account_name,
		account_number,
		status,
		reference,
		failure_reason,
		created_at,
		processed_at,
		updated_at
	FROM withdrawals
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

	var withdrawals []models.Withdrawal

	for rows.Next() {

		var w models.Withdrawal

		err := rows.Scan(
			&w.ID,
			&w.WalletID,
			&w.Amount,
			&w.BankName,
			&w.AccountName,
			&w.AccountNumber,
			&w.Status,
			&w.Reference,
			&w.FailureReason,
			&w.CreatedAt,
			&w.ProcessedAt,
			&w.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		withdrawals = append(withdrawals, w)
	}

	return withdrawals, nil
}

func (r *WithdrawalRepository) UpdateStatus(
	id string,
	status string,
	reference string,
	failureReason string,
) error {

	query := `
	UPDATE withdrawals
	SET
		status=$1,
		reference=$2,
		failure_reason=$3,
		processed_at=NOW(),
		updated_at=NOW()
	WHERE id=$4
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		status,
		reference,
		failureReason,
		id,
	)

	return err
}
