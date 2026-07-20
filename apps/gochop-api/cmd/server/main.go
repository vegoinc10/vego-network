package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/config"
	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/database"
)

func main() {
	cfg := config.Load()

	db := database.Connect(cfg)
	defer db.Close(nil)

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": cfg.AppName,
			"version": "1.0.0",
		})
	})

	router.Run(":" + cfg.Port)
}
