package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/config"
	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/database"
	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/routes"
)

func main() {
	cfg := config.Load()

	db := database.Connect(cfg)
	defer db.Close(context.Background())

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": cfg.AppName,
			"version": "1.0.0",
		})
	})

	routes.SetupRoutes(router, db)

	router.Run(":" + cfg.Port)
}
