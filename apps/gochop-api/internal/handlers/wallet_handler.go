package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/services"
)

type WalletHandler struct {
	service *services.WalletService
}

func NewWalletHandler(service *services.WalletService) *WalletHandler {
	return &WalletHandler{
		service: service,
	}
}

func (h *WalletHandler) GetWallet(c *gin.Context) {

	userID := c.GetString("userID")

	wallet, err := h.service.GetWallet(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, wallet)
}
