package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/models"
)

type StoreRepository struct {
	db *pgx.Conn
}

func NewStoreRepository(db *pgx.Conn) *StoreRepository {
	return &StoreRepository{
		db: db,
	}
}

func (r *StoreRepository) CreateStore(store *models.Store) error {

	query := `
	INSERT INTO stores (
    id,
    owner_id,
    name,
    slug,
    description,
    email,
    phone,
    state,
    lga,
    address,
    latitude,
    longitude,
    verified,
    status,
    created_at,
    updated_at
)
VALUES (
    $1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14
)
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		store.ID,
		store.OwnerID,
		store.Name,
		store.Slug,
		store.Description,
		store.Email,
		store.Phone,
		store.State,
		store.LGA,
		store.Address,
		store.Latitude,
		store.Longitude,
		store.Verified,
		store.Status,
		store.CreatedAt,
		store.UpdatedAt,
	)

	return err
}

func (r *StoreRepository) GetStoreByID(id string) (*models.Store, error) {

	query := `
	SELECT
		id,
		owner_id,
		name,
		slug,
		description,
		logo_url,
		banner_url,
		email,
		phone,
		state,
		lga,
		address,
		latitude,
        longitude,
		verified,
		status,
		created_at,
		updated_at
	FROM stores
	WHERE id = $1
	`

	var store models.Store

	err := r.db.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(
		&store.ID,
		&store.OwnerID,
		&store.Name,
		&store.Slug,
		&store.Description,
		&store.LogoURL,
		&store.BannerURL,
		&store.Email,
		&store.Phone,
		&store.State,
		&store.LGA,
		&store.Address,
		&store.Latitude,
		&store.Longitude,
		&store.Verified,
		&store.Status,
		&store.CreatedAt,
		&store.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &store, nil
}

func (r *StoreRepository) GetStoreByOwner(ownerID string) (*models.Store, error) {

	query := `
	SELECT
		id,
		owner_id,
		name,
		slug,
		description,
		logo_url,
		banner_url,
		email,
		phone,
		state,
		lga,
		address,
		latitude,
        longitude,
		verified,
		status,
		created_at,
		updated_at
	FROM stores
	WHERE owner_id = $1
	LIMIT 1
	`

	var store models.Store

	err := r.db.QueryRow(
		context.Background(),
		query,
		ownerID,
	).Scan(
		&store.ID,
		&store.OwnerID,
		&store.Name,
		&store.Slug,
		&store.Description,
		&store.LogoURL,
		&store.BannerURL,
		&store.Email,
		&store.Phone,
		&store.State,
		&store.LGA,
		&store.Address,
		&store.Latitude,
		&store.Longitude,
		&store.Verified,
		&store.Status,
		&store.CreatedAt,
		&store.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &store, nil
}

func (r *StoreRepository) ListStores() ([]models.Store, error) {

	rows, err := r.db.Query(
		context.Background(),
		`
		SELECT
			id,
			owner_id,
			name,
			slug,
			description,
			logo_url,
			banner_url,
			email,
			phone,
			state,
			lga,
			address,
			latitude,
            longitude,
			verified,
			status,
			created_at,
			updated_at
		FROM stores
		ORDER BY created_at DESC
		`,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stores []models.Store

	for rows.Next() {

		var store models.Store

		err := rows.Scan(
			&store.ID,
			&store.OwnerID,
			&store.Name,
			&store.Slug,
			&store.Description,
			&store.LogoURL,
			&store.BannerURL,
			&store.Email,
			&store.Phone,
			&store.State,
			&store.LGA,
			&store.Address,
			&store.Latitude,
			&store.Longitude,
			&store.Verified,
			&store.Status,
			&store.CreatedAt,
			&store.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		stores = append(stores, store)
	}

	return stores, nil
}

func (r *StoreRepository) UpdateStore(store *models.Store) error {

	query := `
	UPDATE stores
	SET
    name = $1,
    description = $2,
    email = $3,
    phone = $4,
    state = $5,
    lga = $6,
    address = $7,
    latitude = $8,
    longitude = $9,
    updated_at = $10
WHERE id = $11
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		store.Name,
		store.Description,
		store.Email,
		store.Phone,
		store.State,
		store.LGA,
		store.Address,
		store.Latitude,
		store.Longitude,
		store.UpdatedAt,
		store.ID,
	)

	return err
}

func (r *StoreRepository) DeleteStore(id string) error {

	_, err := r.db.Exec(
		context.Background(),
		`DELETE FROM stores WHERE id = $1`,
		id,
	)

	return err
}

func (r *StoreRepository) GetStoreByProductID(productID string) (*models.Store, error) {

	query := `
	SELECT
		s.id,
		s.owner_id,
		s.name,
		s.slug,
		s.description,
		s.logo_url,
		s.banner_url,
		s.email,
		s.phone,
		s.state,
		s.lga,
		s.address,
		latitude,
        longitude,
		s.verified,
		s.status,
		s.created_at,
		s.updated_at
	FROM stores s
	INNER JOIN products p
		ON p.store_id = s.id
	WHERE p.id = $1
	`

	var store models.Store

	err := r.db.QueryRow(
		context.Background(),
		query,
		productID,
	).Scan(
		&store.ID,
		&store.OwnerID,
		&store.Name,
		&store.Slug,
		&store.Description,
		&store.LogoURL,
		&store.BannerURL,
		&store.Email,
		&store.Phone,
		&store.State,
		&store.LGA,
		&store.Address,
		&store.Latitude,
		&store.Longitude,
		&store.Verified,
		&store.Status,
		&store.CreatedAt,
		&store.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &store, nil
}
