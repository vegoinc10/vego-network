package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/models"
)

type OrderRepository struct {
	db *pgx.Conn
}

func NewOrderRepository(db *pgx.Conn) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) CreateOrder(order *models.Order) error {

	query := `
	INSERT INTO orders (
		id,
		buyer_id,
		status,
		total_amount,
		payment_status,
		delivery_status,
		created_at,
		updated_at
	)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		order.ID,
		order.BuyerID,
		order.Status,
		order.TotalAmount,
		order.PaymentStatus,
		order.DeliveryStatus,
		order.CreatedAt,
		order.UpdatedAt,
	)

	return err
}

func (r *OrderRepository) CreateOrderItem(item *models.OrderItem) error {

	query := `
	INSERT INTO order_items (
		id,
		order_id,
		product_id,
		quantity,
		unit_price,
		total_price
	)
	VALUES ($1,$2,$3,$4,$5,$6)
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		item.ID,
		item.OrderID,
		item.ProductID,
		item.Quantity,
		item.Price,
		item.Subtotal,
	)

	return err
}

func (r *OrderRepository) GetOrdersByBuyer(userID string) ([]models.Order, error) {

	rows, err := r.db.Query(
		context.Background(),
		`
		SELECT
			id,
			buyer_id,
			status,
			total_amount,
			payment_status,
			delivery_status,
			created_at,
			updated_at
		FROM orders
		WHERE buyer_id = $1
		ORDER BY created_at DESC
		`,
		userID,
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

func NewOrder() *models.Order {
	return &models.Order{
		ID: uuid.New().String(),
	}
}
func (r *OrderRepository) UpdateOrderStatus(orderID, status string) error {

	_, err := r.db.Exec(
		context.Background(),
		`
		UPDATE orders
		SET
			status = $1,
			updated_at = NOW()
		WHERE id = $2
		`,
		status,
		orderID,
	)

	return err
}

func (r *OrderRepository) GetOrderByID(orderID string) (*models.Order, error) {

	query := `
	SELECT
		id,
		buyer_id,
		status,
		total_amount,
		payment_status,
		delivery_status,
		created_at,
		updated_at
	FROM orders
	WHERE id = $1
	`

	var order models.Order

	err := r.db.QueryRow(
		context.Background(),
		query,
		orderID,
	).Scan(
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

	return &order, nil
}

func (r *OrderRepository) GetOrderItems(orderID string) ([]models.OrderItem, error) {

	query := `
	SELECT
		id,
		order_id,
		product_id,
		quantity,
		unit_price,
		total_price
	FROM order_items
	WHERE order_id = $1
	`

	rows, err := r.db.Query(
		context.Background(),
		query,
		orderID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.OrderItem

	for rows.Next() {

		var item models.OrderItem

		err := rows.Scan(
			&item.ID,
			&item.OrderID,
			&item.ProductID,
			&item.Quantity,
			&item.Price,
			&item.Subtotal,
		)

		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

func (r *OrderRepository) GetVendorByProduct(productID string) (string, error) {

	var ownerID string

	query := `
	SELECT s.owner_id
	FROM products p
	JOIN stores s
		ON s.id = p.store_id
	WHERE p.id = $1
	`

	err := r.db.QueryRow(
		context.Background(),
		query,
		productID,
	).Scan(&ownerID)

	return ownerID, err
}

func (r *OrderRepository) GetOrders(ctx context.Context) ([]models.Order, error) {

	rows, err := r.db.Query(
		ctx,
		`
		SELECT 
			id,
			buyer_id,
			status,
			total_amount,
			created_at
		FROM orders
		ORDER BY created_at DESC
		`,
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
			&order.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}
