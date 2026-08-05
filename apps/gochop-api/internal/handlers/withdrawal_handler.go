package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/models"
	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/services"
)

type WithdrawalHandler struct {
	service *services.WithdrawalService
}

func NewWithdrawalHandler(
	service *services.WithdrawalService,
) *WithdrawalHandler {

	return &WithdrawalHandler{
		service: service,
	}
}

func (h *WithdrawalHandler) RequestWithdrawal(c *gin.Context) {

	userID := c.GetString("userID")

	var req models.WithdrawalRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.service.RequestWithdrawal(
		userID,
		&req,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Withdrawal request submitted successfully",
	})
}
