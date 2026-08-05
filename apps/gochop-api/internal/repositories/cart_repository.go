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
func (r *CartRepository) GetCart(userID string) ([]models.CartResponse, error) {

	query := `
	SELECT
	c.id,
	p.id,
	p.store_id,
	p.name,
	p.price,
	c.quantity,
	(p.price * c.quantity) AS subtotal
FROM cart_items c
JOIN products p
	ON c.product_id = p.id
WHERE c.user_id = $1
ORDER BY c.created_at DESC
`
	rows, err := r.db.Query(
		context.Background(),
		query,
		userID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var items []models.CartResponse

	for rows.Next() {

		var item models.CartResponse

		err := rows.Scan(
			&item.CartID,
			&item.ProductID,
			&item.StoreID,
			&item.ProductName,
			&item.Price,
			&item.Quantity,
			&item.Subtotal,
		)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}
func (r *CartRepository) ClearCart(userID string) error {

	_, err := r.db.Exec(
		context.Background(),
		`DELETE FROM cart_items WHERE user_id=$1`,
		userID,
	)

	return err
}
