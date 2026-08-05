package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/models"
)

type UserRepository struct {
	DB *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) Create(user *models.User) error {

	query := `
	INSERT INTO users
	(
		full_name,
		email,
		phone,
		password_hash,
		role
	)
	VALUES
	($1,$2,$3,$4,$5)
	RETURNING id, created_at;
	`

	return r.DB.QueryRow(
		context.Background(),
		query,
		user.FullName,
		user.Email,
		user.Phone,
		user.PasswordHash,
		user.Role,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)
}
func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {

	query := `
SELECT
    id,
    full_name,
    email,
    password_hash,
    role,
    created_at,
    updated_at
FROM users
WHERE email = $1
`

	var user models.User

	err := r.DB.QueryRow(
		context.Background(),
		query,
		email,
	).Scan(
		&user.ID,
		&user.FullName,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
