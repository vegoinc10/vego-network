package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/services"
)

type ProductHandler struct {
	productService *services.ProductService
}

func NewProductHandler(productService *services.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Create Product endpoint",
	})
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get Products endpoint",
	})
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get Product endpoint",
	})
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Product endpoint",
	})
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Product endpoint",
	})
}
