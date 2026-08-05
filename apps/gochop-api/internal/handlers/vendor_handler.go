package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/services"
)

type VendorHandler struct {
	service *services.VendorService
}

func NewVendorHandler(service *services.VendorService) *VendorHandler {
	return &VendorHandler{
		service: service,
	}
}

func (h *VendorHandler) GetOrders(c *gin.Context) {

	vendorID := c.GetString("userID")

	orders, err := h.service.GetVendorOrders(vendorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, orders)
}
