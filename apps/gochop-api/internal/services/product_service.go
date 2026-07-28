package services

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/models"
	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) CreateProduct(
	sellerID string,
	req *models.CreateProductRequest,
) error {

	if req.Currency == "" {
		req.Currency = "NGN"
	}

	product := &models.Product{
		ID:          uuid.New().String(),
		SellerID:    sellerID,
		Name:        req.Name,
		Description: req.Description,
		Category:    req.Category,

		Price:      req.Price,
		Currency:   req.Currency,
		Quantity:   req.Quantity,
		Unit:       req.Unit,
		MarketName: req.MarketName,
		State:      req.State,
		LGA:        req.LGA,
		ImageURL:   req.ImageURL,
		Status:     "active",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	return s.repo.CreateProduct(product)
}
func (s *ProductService) GetProducts() ([]models.Product, error) {
	return s.repo.GetProducts()
}
func (s *ProductService) GetProductByID(id string) (*models.Product, error) {
	return s.repo.GetProductByID(id)
}
func (s *ProductService) UpdateProduct(
	id string,
	req *models.CreateProductRequest,
) error {

	if req.Currency == "" {
		req.Currency = "NGN"
	}

	return s.repo.UpdateProduct(id, req)
}
func (s *ProductService) DeleteProduct(id string) error {
	return s.repo.DeleteProduct(id)
}
func (s *ProductService) VerifyOwner(productID, userID string) error {

	product, err := s.repo.GetProductByID(productID)
	if err != nil {
		return err
	}

	if product.SellerID != userID {
		return errors.New("unauthorized")
	}

	return nil
}
