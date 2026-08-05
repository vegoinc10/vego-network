package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/models"
	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/services"
)

type CartHandler struct {
	service *services.CartService
}

func NewCartHandler(service *services.CartService) *CartHandler {
	return &CartHandler{
		service: service,
	}
}

func (h *CartHandler) AddToCart(c *gin.Context) {

	var req models.AddToCartRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// We'll replace this with the JWT user ID in the next step.
	userID := "02b20e4d-9df5-4ae0-831b-f8dfc0455833"

	err := h.service.AddToCart(userID, &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Product added to cart",
	})
}
func (h *CartHandler) GetCart(c *gin.Context) {

	userID := c.GetString("userID")

	items, err := h.service.GetCart(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	total := 0.0

	for _, item := range items {
		total += item.Subtotal
	}

	c.JSON(http.StatusOK, gin.H{
		"items": items,
		"total": total,
	})
}
