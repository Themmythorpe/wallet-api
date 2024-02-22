package router

import (
    "github.com/gin-gonic/gin"
    "wallet-api-go/controllers"
)

// SetupRouter initializes the routes for the application
func SetupRouter(r *gin.Engine) {
    // User routes
    userRoutes := r.Group("/users")
    {
        userRoutes.POST("/register", controllers.RegisterUser)
        userRoutes.POST("/login", controllers.LoginUser)
    }

    // Wallet routes
    walletRoutes := r.Group("/wallets")
    {
        walletRoutes.POST("/create", controllers.CreateWallet)
        walletRoutes.POST("/:wallet_id/credit", controllers.CreditWallet)
        walletRoutes.POST("/:wallet_id/debit", controllers.DebitWallet)
        walletRoutes.GET("/:wallet_id/transactions", controllers.GetWalletTransactions)
    }
}
