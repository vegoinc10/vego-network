package services

import (
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/models"
	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/repositories"
)

type StoreService struct {
	storeRepo *repositories.StoreRepository
}

func NewStoreService(storeRepo *repositories.StoreRepository) *StoreService {
	return &StoreService{
		storeRepo: storeRepo,
	}
}

func (s *StoreService) CreateStore(ownerID string, req *models.CreateStoreRequest) (*models.Store, error) {

	slug := strings.ToLower(req.Name)
	slug = strings.ReplaceAll(slug, " ", "-")

	store := &models.Store{
		ID:          uuid.New().String(),
		OwnerID:     ownerID,
		Name:        req.Name,
		Slug:        slug,
		Description: req.Description,
		Email:       req.Email,
		Phone:       req.Phone,
		State:       req.State,
		LGA:         req.LGA,
		Address:     req.Address,
		Status:      "active",
		Verified:    false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Latitude:    req.Latitude,
		Longitude:   req.Longitude,
	}

	err := s.storeRepo.CreateStore(store)
	if err != nil {
		return nil, err
	}

	return store, nil
}

func (s *StoreService) GetStore(id string) (*models.Store, error) {
	return s.storeRepo.GetStoreByID(id)
}

func (s *StoreService) GetStores() ([]models.Store, error) {
	return s.storeRepo.ListStores()
}

func (s *StoreService) UpdateStore(id string, req *models.CreateStoreRequest) (*models.Store, error) {

	store, err := s.storeRepo.GetStoreByID(id)
	if err != nil {
		return nil, err
	}

	store.Name = req.Name
	store.Description = req.Description
	store.Email = req.Email
	store.Phone = req.Phone
	store.State = req.State
	store.LGA = req.LGA
	store.Address = req.Address
	store.UpdatedAt = time.Now()
	store.Latitude = req.Latitude
	store.Longitude = req.Longitude

	err = s.storeRepo.UpdateStore(store)
	if err != nil {
		return nil, err
	}

	return store, nil
}

func (s *StoreService) DeleteStore(id string) error {
	return s.storeRepo.DeleteStore(id)
}
