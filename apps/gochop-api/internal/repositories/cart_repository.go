package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/models"
)

type CartRepository struct {
	db *pgx.Conn
}

func NewCartRepository(db *pgx.Conn) *CartRepository {
	return &CartRepository{
		db: db,
	}
}

func (r *CartRepository) AddToCart(userID string, req *models.AddToCartRequest) error {

	query := `
	INSERT INTO cart_items (
		id,
		user_id,
		product_id,
		quantity
	)
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (user_id, product_id)
	DO UPDATE SET
		quantity = cart_items.quantity + EXCLUDED.quantity,
		updated_at = NOW()
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		uuid.New().String(),
		userID,
		req.ProductID,
		req.Quantity,
	)

	return err
}
