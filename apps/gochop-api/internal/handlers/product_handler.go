package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/models"
	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/services"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {

	var req models.CreateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Temporary seller ID
	// Get seller ID from JWT middleware
	sellerID := c.GetString("userID")

	err := h.service.CreateProduct(sellerID, &req)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Product created successfully",
	})
}

func (h *ProductHandler) GetProducts(c *gin.Context) {

	products, err := h.service.GetProducts()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetProduct(c *gin.Context) {

	id := c.Param("id")

	product, err := h.service.GetProductByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Product not found",
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {

	userID := c.GetString("userID")

	id := c.Param("id")

	// Verify ownership
	err := h.service.VerifyOwner(id, userID)

	if err != nil {

		c.JSON(http.StatusForbidden, gin.H{
			"error": "You cannot modify this product",
		})

		return
	}

	var req models.CreateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = h.service.UpdateProduct(id, &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product updated successfully",
	})
}
func (h *ProductHandler) DeleteProduct(c *gin.Context) {

	userID := c.GetString("userID")

	id := c.Param("id")

	// Verify ownership
	err := h.service.VerifyOwner(id, userID)

	if err != nil {

		c.JSON(http.StatusForbidden, gin.H{
			"error": "You cannot delete this product",
		})

		return
	}

	err = h.service.DeleteProduct(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
	})
}
