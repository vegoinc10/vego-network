package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/config"
)

func Connect(cfg *config.Config) *pgx.Conn {
	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	log.Println("✅ Connected to PostgreSQL")

	return conn
}
