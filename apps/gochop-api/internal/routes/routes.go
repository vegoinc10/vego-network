package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"

	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/handlers"
	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/middleware"
	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/repositories"
	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/services"
)

func SetupRoutes(router *gin.Engine, db *pgx.Conn) {

	// Authentication
	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	// Products
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	api := router.Group("/api/v1")

	// Authentication routes
	auth := api.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	// Public product routes
	publicProducts := api.Group("/products")
	{
		publicProducts.GET("", productHandler.GetProducts)
		publicProducts.GET("/:id", productHandler.GetProduct)
	}

	// Protected product routes
	protectedProducts := api.Group("/products")
	protectedProducts.Use(middleware.AuthMiddleware())
	{
		protectedProducts.POST("", productHandler.CreateProduct)
		protectedProducts.PUT("/:id", productHandler.UpdateProduct)
		protectedProducts.DELETE("/:id", productHandler.DeleteProduct)
	}
	// Categories
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	categories := api.Group("/categories")
	{
		categories.GET("", categoryHandler.GetCategories)
	}
	//cart
	cartRepo := repositories.NewCartRepository(db)
	cartService := services.NewCartService(cartRepo)
	cartHandler := handlers.NewCartHandler(cartService)
	cart := api.Group("/cart")
	cart.Use(middleware.AuthMiddleware())
	{
		cart.POST("", cartHandler.AddToCart)
	}
}
