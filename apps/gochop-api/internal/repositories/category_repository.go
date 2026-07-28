package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/models"
)

type CategoryRepository struct {
	db *pgx.Conn
}

func NewCategoryRepository(db *pgx.Conn) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (r *CategoryRepository) GetCategories() ([]models.Category, error) {

	rows, err := r.db.Query(
		context.Background(),
		`
		SELECT
			id,
			name,
			slug,
			COALESCE(description, '') AS description,
			created_at
		FROM categories
		ORDER BY name ASC
		`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var categories []models.Category

	for rows.Next() {

		var c models.Category

		err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.Slug,
			&c.Description,
			&c.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		categories = append(categories, c)
	}

	return categories, nil
}
