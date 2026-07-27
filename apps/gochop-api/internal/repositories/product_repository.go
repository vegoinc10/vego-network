package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/models"
)

type ProductRepository struct {
	db *pgx.Conn
}

func NewProductRepository(db *pgx.Conn) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) CreateProduct(product *models.Product) error {

	query := `
	INSERT INTO products (
		id,
		seller_id,
		name,
		description,
		category,
		price,
		currency,
		quantity,
		unit,
		market_name,
		state,
		lga,
		image_url,
		status,
		created_at,
		updated_at
	)
	VALUES (
		$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16
	)
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		product.ID,
		product.SellerID,
		product.Name,
		product.Description,
		product.Category,
		product.Price,
		product.Currency,
		product.Quantity,
		product.Unit,
		product.MarketName,
		product.State,
		product.LGA,
		product.ImageURL,
		product.Status,
		product.CreatedAt,
		product.UpdatedAt,
	)

	return err
}
func (r *ProductRepository) GetProducts() ([]models.Product, error) {

	rows, err := r.db.Query(
		context.Background(),
		`SELECT
			id,
			seller_id,
			name,
			description,
			category,
			price,
			currency,
			quantity,
			unit,
			market_name,
			state,
			lga,
			image_url,
			status,
			created_at,
			updated_at
		FROM products
		ORDER BY created_at DESC`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []models.Product

	for rows.Next() {

		var p models.Product

		err := rows.Scan(
			&p.ID,
			&p.SellerID,
			&p.Name,
			&p.Description,
			&p.Category,
			&p.Price,
			&p.Currency,
			&p.Quantity,
			&p.Unit,
			&p.MarketName,
			&p.State,
			&p.LGA,
			&p.ImageURL,
			&p.Status,
			&p.CreatedAt,
			&p.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}
func (r *ProductRepository) GetProductByID(id string) (*models.Product, error) {

	query := `
	SELECT
		id,
		seller_id,
		name,
		description,
		category,
		price,
		currency,
		quantity,
		unit,
		market_name,
		state,
		lga,
		image_url,
		status,
		created_at,
		updated_at
	FROM products
	WHERE id = $1
	`

	var product models.Product

	err := r.db.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(
		&product.ID,
		&product.SellerID,
		&product.Name,
		&product.Description,
		&product.Category,
		&product.Price,
		&product.Currency,
		&product.Quantity,
		&product.Unit,
		&product.MarketName,
		&product.State,
		&product.LGA,
		&product.ImageURL,
		&product.Status,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &product, nil
}
func (r *ProductRepository) UpdateProduct(id string, product *models.CreateProductRequest) error {

	query := `
	UPDATE products
	SET
		name = $1,
		description = $2,
		category = $3,
		price = $4,
		currency = $5,
		quantity = $6,
		unit = $7,
		market_name = $8,
		state = $9,
		lga = $10,
		image_url = $11,
		updated_at = NOW()
	WHERE id = $12
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		product.Name,
		product.Description,
		product.Category,
		product.Price,
		product.Currency,
		product.Quantity,
		product.Unit,
		product.MarketName,
		product.State,
		product.LGA,
		product.ImageURL,
		id,
	)

	return err
}
func (r *ProductRepository) DeleteProduct(id string) error {

	query := `DELETE FROM products WHERE id = $1`

	_, err := r.db.Exec(
		context.Background(),
		query,
		id,
	)

	return err
}
