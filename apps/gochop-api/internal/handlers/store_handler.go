package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/models"
	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/services"
)

type StoreHandler struct {
	storeService *services.StoreService
}

func NewStoreHandler(storeService *services.StoreService) *StoreHandler {
	return &StoreHandler{
		storeService: storeService,
	}
}
func (h *StoreHandler) CreateStore(c *gin.Context) {

	var req models.CreateStoreRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := c.GetString("userID")

	store, err := h.storeService.CreateStore(userID, &req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Store created successfully",
		"store":   store,
	})
}
func (h *StoreHandler) GetStore(c *gin.Context) {

	id := c.Param("id")

	store, err := h.storeService.GetStore(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, store)
}
func (h *StoreHandler) GetStores(c *gin.Context) {

	stores, err := h.storeService.GetStores()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, stores)
}
func (h *StoreHandler) UpdateStore(c *gin.Context) {

	id := c.Param("id")

	var req models.CreateStoreRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	store, err := h.storeService.UpdateStore(id, &req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, store)
}
func (h *StoreHandler) DeleteStore(c *gin.Context) {

	id := c.Param("id")

	err := h.storeService.DeleteStore(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Store deleted successfully",
	})
}
