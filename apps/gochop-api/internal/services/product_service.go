package services

import "github.com/vegoinc10/vego-network/apps/gochop-api/internal/repositories"

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}
