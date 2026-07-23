package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"

	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/handlers"
	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/repositories"
	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/services"
)

func SetupRoutes(router *gin.Engine, db *pgx.Conn) {

	// Authentication
	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	// Product
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	api := router.Group("/api/v1")

	// Authentication routes
	auth := api.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)

	}

	// Product routes
	products := api.Group("/products")
	{
		products.POST("", productHandler.CreateProduct)
		products.GET("", productHandler.GetProducts)
		products.GET("/:id", productHandler.GetProduct)
		products.PUT("/:id", productHandler.UpdateProduct)
		products.DELETE("/:id", productHandler.DeleteProduct)
	}
}
