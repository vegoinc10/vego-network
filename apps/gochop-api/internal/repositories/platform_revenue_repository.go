package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/models"
)

type PlatformRevenueRepository struct {
	db *pgx.Conn
}

func NewPlatformRevenueRepository(db *pgx.Conn) *PlatformRevenueRepository {
	return &PlatformRevenueRepository{
		db: db,
	}
}

func (r *PlatformRevenueRepository) Create(
	revenue *models.PlatformRevenue,
) error {

	query := `
	INSERT INTO platform_revenue
	(
		order_id,
		vendor_id,
		gross_amount,
		commission_rate,
		commission_amount,
		vendor_amount
	)
	VALUES ($1,$2,$3,$4,$5,$6)
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		revenue.OrderID,
		revenue.VendorID,
		revenue.GrossAmount,
		revenue.CommissionRate,
		revenue.CommissionAmount,
		revenue.VendorAmount,
	)

	return err
}
