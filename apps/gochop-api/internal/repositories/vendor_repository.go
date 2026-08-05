package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/models"
)

type VendorRepository struct {
	db *pgx.Conn
}

func NewVendorRepository(db *pgx.Conn) *VendorRepository {
	return &VendorRepository{
		db: db,
	}
}

func (r *VendorRepository) GetVendorOrders(vendorID string) ([]models.Order, error) {

	query := `
	SELECT DISTINCT
		o.id,
		o.buyer_id,
		o.status,
		o.total_amount,
		o.payment_status,
		o.delivery_status,
		o.created_at,
		o.updated_at
	FROM orders o
	INNER JOIN order_items oi ON o.id = oi.order_id
	INNER JOIN products p ON oi.product_id = p.id
	JOIN stores s ON s.id = p.store_id
WHERE s.owner_id = $1
	ORDER BY o.created_at DESC
	`

	rows, err := r.db.Query(
		context.Background(),
		query,
		vendorID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order

	for rows.Next() {

		var order models.Order

		err := rows.Scan(
			&order.ID,
			&order.BuyerID,
			&order.Status,
			&order.TotalAmount,
			&order.PaymentStatus,
			&order.DeliveryStatus,
			&order.CreatedAt,
			&order.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}
