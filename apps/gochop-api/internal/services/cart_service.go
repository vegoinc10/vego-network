package services

import (
	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/models"
	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/repositories"
)

type CartService struct {
	repo *repositories.CartRepository
}

func NewCartService(repo *repositories.CartRepository) *CartService {
	return &CartService{
		repo: repo,
	}
}

func (s *CartService) AddToCart(userID string, req *models.AddToCartRequest) error {
	return s.repo.AddToCart(userID, req)
}
func (s *CartService) GetCart(userID string) ([]models.CartResponse, error) {
	return s.repo.GetCart(userID)
}
