package models

import "time"

type Store struct {
	ID          string `json:"id"`
	OwnerID     string `json:"owner_id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`

	LogoURL   string `json:"logo_url"`
	BannerURL string `json:"banner_url"`

	Email   string `json:"email"`
	Phone   string `json:"phone"`
	State   string `json:"state"`
	LGA     string `json:"lga"`
	Address string `json:"address"`

	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`

	Verified bool   `json:"verified"`
	Status   string `json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateStoreRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`

	Email   string `json:"email"`
	Phone   string `json:"phone"`
	State   string `json:"state"`
	LGA     string `json:"lga"`
	Address string `json:"address"`

	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
