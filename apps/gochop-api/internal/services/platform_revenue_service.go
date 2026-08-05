package services

import (
	"time"

	"github.com/google/uuid"

	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/models"
	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/repositories"
)

const DefaultCommissionRate = 5.0

type PlatformRevenueService struct {
	repo *repositories.PlatformRevenueRepository
}

func NewPlatformRevenueService(
	repo *repositories.PlatformRevenueRepository,
) *PlatformRevenueService {

	return &PlatformRevenueService{
		repo: repo,
	}
}

func (s *PlatformRevenueService) RecordRevenue(
	orderID string,
	vendorID string,
	grossAmount float64,
	commission float64,
	vendorAmount float64,
) error {

	revenue := &models.PlatformRevenue{
		ID:               uuid.New().String(),
		OrderID:          orderID,
		VendorID:         vendorID,
		GrossAmount:      grossAmount,
		CommissionRate:   DefaultCommissionRate,
		CommissionAmount: commission,
		VendorAmount:     vendorAmount,
		CreatedAt:        time.Now(),
	}

	return s.repo.Create(revenue)
}
