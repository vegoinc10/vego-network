package services

import (
	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/models"
	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/repositories"
)

type VendorService struct {
	repo *repositories.VendorRepository
}

func NewVendorService(repo *repositories.VendorRepository) *VendorService {
	return &VendorService{
		repo: repo,
	}
}

func (s *VendorService) GetVendorOrders(vendorID string) ([]models.Order, error) {
	return s.repo.GetVendorOrders(vendorID)
}
