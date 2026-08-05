package services

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/models"
	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/repositories"
)

type OrderService struct {
	orderRepo             *repositories.OrderRepository
	cartRepo              *repositories.CartRepository
	productRepo           *repositories.ProductRepository
	storeRepo             *repositories.StoreRepository
	walletRepo            *repositories.WalletRepository
	walletTransactionRepo *repositories.WalletTransactionRepository

	commissionService *CommissionService
}

func NewOrderService(
	orderRepo *repositories.OrderRepository,
	cartRepo *repositories.CartRepository,
	productRepo *repositories.ProductRepository,
	storeRepo *repositories.StoreRepository,
	walletRepo *repositories.WalletRepository,
	walletTransactionRepo *repositories.WalletTransactionRepository,
	commissionService *CommissionService,
) *OrderService {

	return &OrderService{
		orderRepo:             orderRepo,
		cartRepo:              cartRepo,
		productRepo:           productRepo,
		storeRepo:             storeRepo,
		walletRepo:            walletRepo,
		walletTransactionRepo: walletTransactionRepo,
		commissionService:     commissionService,
	}
}

func (s *OrderService) Checkout(userID string) (*models.Order, error) {

	items, err := s.cartRepo.GetCart(userID)
	if err != nil {
		return nil, err
	}

	total := 0.0

	// Check stock and calculate total
	for _, item := range items {

		available, err := s.productRepo.GetProductQuantity(item.ProductID)
		if err != nil {
			return nil, err
		}

		if available < item.Quantity {
			return nil, errors.New("insufficient stock")
		}

		total += item.Subtotal
	}

	order := &models.Order{
		ID:             uuid.New().String(),
		BuyerID:        userID,
		Status:         "pending",
		TotalAmount:    total,
		PaymentStatus:  "pending",
		DeliveryStatus: "pending",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	// Save order
	err = s.orderRepo.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	// Save order items and deduct stock
	for _, item := range items {

		orderItem := &models.OrderItem{
			ID:        uuid.New().String(),
			OrderID:   order.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
			Subtotal:  item.Subtotal,
		}

		err = s.orderRepo.CreateOrderItem(orderItem)
		if err != nil {
			return nil, err
		}

		err = s.productRepo.UpdateStock(item.ProductID, item.Quantity)
		if err != nil {
			return nil, err
		}
	}

	// Clear cart
	err = s.cartRepo.ClearCart(userID)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderService) UpdateOrderStatus(orderID, status string) error {

	// Update the order status first
	err := s.orderRepo.UpdateOrderStatus(orderID, status)
	if err != nil {
		return err
	}

	// Only settle funds when the order is delivered
	if status != "delivered" {
		return nil
	}

	if err != nil {
		return err
	}

	items, err := s.orderRepo.GetOrderItems(orderID)
	if err != nil {
		return err
	}
	for _, item := range items {

		store, err := s.storeRepo.GetStoreByProductID(item.ProductID)
		if err != nil {
			return err
		}

		vendorID := store.OwnerID

		commission, vendorAmount :=
			s.commissionService.CalculateCommission(
				item.Subtotal,
				DefaultCommissionRate,
			)

		err = s.walletRepo.CreditWallet(
			vendorID,
			vendorAmount,
		)

		if err != nil {
			return err
		}

		_ = commission
	}
	// TODO:
	// 1. Load the order
	// 2. Find the vendor(s)
	// 3. Calculate commission
	// 4. Credit vendor wallet
	// 5. Record wallet transaction
	// 6. Record platform revenue

	return nil
}

func (s *OrderService) GetOrders(userID string) ([]models.Order, error) {
	return s.orderRepo.GetOrdersByBuyer(userID)
}
