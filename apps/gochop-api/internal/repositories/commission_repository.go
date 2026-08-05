package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/models"
)

type CommissionRepository struct {
	db *pgx.Conn
}

func NewCommissionRepository(db *pgx.Conn) *CommissionRepository {
	return &CommissionRepository{
		db: db,
	}
}

func (r *CommissionRepository) GetActiveSetting() (*models.CommissionSetting, error) {

	ctx := context.Background()

	var setting models.CommissionSetting

	query := `
		SELECT
			id,
			commission_rate,
			minimum_withdrawal,
			withdrawal_fee,
			active,
			created_at,
			updated_at
		FROM commission_settings
		WHERE active = TRUE
		ORDER BY created_at DESC
		LIMIT 1
	`

	err := r.db.QueryRow(
		ctx,
		query,
	).Scan(
		&setting.ID,
		&setting.CommissionRate,
		&setting.MinimumWithdrawal,
		&setting.WithdrawalFee,
		&setting.Active,
		&setting.CreatedAt,
		&setting.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &setting, nil
}
