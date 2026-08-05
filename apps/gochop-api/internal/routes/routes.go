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
	walletRepo := repositories.NewWalletRepository(db)
	storeRepo := repositories.NewStoreRepository(db)

	walletTransactionRepo := repositories.NewWalletTransactionRepository(db)

	commissionService := services.NewCommissionService(
		walletRepo,
		walletTransactionRepo,
	)
	walletService := services.NewWalletService(walletRepo)
	walletHandler := handlers.NewWalletHandler(walletService)

	authService := services.NewAuthService(
		userRepo,
		walletRepo,
	)
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

	// Wallet routes
	wallet := api.Group("/wallet")
	wallet.Use(middleware.AuthMiddleware())
	{
		wallet.GET("", walletHandler.GetWallet)
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
	// Cart
	cartRepo := repositories.NewCartRepository(db)
	cartService := services.NewCartService(cartRepo)
	cartHandler := handlers.NewCartHandler(cartService)

	cart := api.Group("/cart")
	cart.Use(middleware.AuthMiddleware())
	{
		cart.POST("", cartHandler.AddToCart)
		cart.GET("", cartHandler.GetCart)
	}

	// Orders
	orderRepo := repositories.NewOrderRepository(db)
	orderService := services.NewOrderService(
		orderRepo,
		cartRepo,
		productRepo,
		storeRepo,
		walletRepo,
		walletTransactionRepo,
		commissionService,
	)
	orderHandler := handlers.NewOrderHandler(orderService)
	orders := api.Group("/orders")
	orders.Use(middleware.AuthMiddleware())
	{
		orders.POST("/checkout", orderHandler.Checkout)
		orders.GET("", orderHandler.GetOrders)
	}
	//vendor
	vendorRepo := repositories.NewVendorRepository(db)
	vendorService := services.NewVendorService(vendorRepo)
	vendorHandler := handlers.NewVendorHandler(vendorService)
	vendor := api.Group("/vendor")
	vendor.Use(
		middleware.AuthMiddleware(),
		middleware.VendorOnly(),
	)
	{
		vendor.GET("/orders", vendorHandler.GetOrders)

		vendor.PATCH(
			"/orders/:id/:status",
			orderHandler.UpdateOrderStatus,
		)
	}
	//store
	storeService := services.NewStoreService(storeRepo)
	storeHandler := handlers.NewStoreHandler(storeService)

	// Public store routes
	stores := api.Group("/stores")
	{
		// Public
		stores.GET("", storeHandler.GetStores)
		stores.GET("/:id", storeHandler.GetStore)
	}
	// Vendor store management
	vendorStores := api.Group("/stores")
	vendorStores.Use(
		middleware.AuthMiddleware(),
		middleware.VendorOnly(),
	)
	{
		vendorStores.POST("", storeHandler.CreateStore)
		vendorStores.PUT("/:id", storeHandler.UpdateStore)
		vendorStores.DELETE("/:id", storeHandler.DeleteStore)
	}
	// withdrawal
	withdrawalRepo := repositories.NewWithdrawalRepository(db)

	withdrawalService := services.NewWithdrawalService(
		walletRepo,
		walletTransactionRepo,
		withdrawalRepo,
	)

	withdrawalHandler := handlers.NewWithdrawalHandler(
		withdrawalService,
	)
	withdrawals := api.Group("/withdrawals")
	withdrawals.Use(
		middleware.AuthMiddleware(),
		middleware.VendorOnly(),
	)
	{
		withdrawals.POST("", withdrawalHandler.RequestWithdrawal)
	}

}
