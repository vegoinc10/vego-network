package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/services"
)

type OrderHandler struct {
	service *services.OrderService
}

func NewOrderHandler(service *services.OrderService) *OrderHandler {
	return &OrderHandler{
		service: service,
	}
}

func (h *OrderHandler) Checkout(c *gin.Context) {

	userID := c.GetString("userID")

	order, err := h.service.Checkout(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Order created successfully",
		"order":   order,
	})
}

func (h *OrderHandler) GetOrders(c *gin.Context) {

	userID := c.GetString("userID")

	orders, err := h.service.GetOrders(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) UpdateOrderStatus(c *gin.Context) {

	orderID := c.Param("id")
	status := c.Param("status")

	err := h.service.UpdateOrderStatus(orderID, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order status updated successfully",
	})
}
