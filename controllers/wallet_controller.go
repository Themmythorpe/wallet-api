package controllers

import (
    "net/http"
    "time"
    "wallet-api-go/models"
    "wallet-api-go/db"
    "github.com/gin-gonic/gin"
    "strconv"


)

// CreateWallet handles wallet creation
func CreateWallet(c *gin.Context) {
    var wallet models.Wallet
    if err := c.ShouldBindJSON(&wallet); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Validate wallet input
    if wallet.UserID == 0 || wallet.Currency == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "UserID and Currency are required"})
        return
    }

    // Create wallet
    if err := db.DB.Create(&wallet).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create wallet"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Wallet created successfully"})
}

// CreditWallet handles crediting a wallet
func CreditWallet(c *gin.Context) {
    // Parse wallet ID from path parameter
    walletID := c.Param("wallet_id")

    // Parse credit amount from request body
    var credit models.Credit
    if err := c.ShouldBindJSON(&credit); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Validate credit input
    if credit.Amount <= 0 || credit.Description == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credit input"})
        return
    }

    // Get the wallet from the database
    var wallet models.Wallet
    if err := db.DB.Where("id = ?", walletID).First(&wallet).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Wallet not found"})
        return
    }

    // Start a database transaction
    tx := db.DB.Begin()

    // Credit the wallet
    wallet.Balance += credit.Amount

    // Update the wallet in the database
    if err := tx.Save(&wallet).Error; err != nil {
        tx.Rollback() // Rollback the transaction if there's an error
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to credit wallet"})
        return
    }

    // Create a new transaction record
    transaction := models.Transaction{
        WalletID:    wallet.ID,
        Amount:      credit.Amount,
        Description: credit.Description,
        Type:        "credit",
        CreatedAt:   time.Now().Format(time.RFC3339),
    }

    // Insert the transaction record into the database
    if err := tx.Create(&transaction).Error; err != nil {
        tx.Rollback() // Rollback the transaction if there's an error
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
        return
    }

    // Commit the transaction
    tx.Commit()

    c.JSON(http.StatusOK, gin.H{"message": "Wallet credited successfully"})
}

// DebitWallet handles debiting a wallet
func DebitWallet(c *gin.Context) {
    // Parse wallet ID from path parameter
    walletID := c.Param("wallet_id")

    // Parse debit amount from request body
    var debit models.Debit
    if err := c.ShouldBindJSON(&debit); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Validate debit input
    if debit.Amount <= 0 || debit.Description == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid debit input"})
        return
    }

    // Get the wallet from the database
    var wallet models.Wallet
    if err := db.DB.Where("id = ?", walletID).First(&wallet).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Wallet not found"})
        return
    }

    // Check if the wallet has sufficient balance for the debit
    if wallet.Balance < debit.Amount {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient balance"})
        return
    }

    // Start a database transaction
    tx := db.DB.Begin()

    // Debit the wallet
    wallet.Balance -= debit.Amount

    // Update the wallet in the database
    if err := tx.Save(&wallet).Error; err != nil {
        tx.Rollback() // Rollback the transaction if there's an error
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to debit wallet"})
        return
    }

    // Create a new transaction record
    transaction := models.Transaction{
        WalletID:    wallet.ID,
        Amount:      debit.Amount,
        Description: debit.Description,
        Type:        "debit",
        CreatedAt:   time.Now().Format(time.RFC3339),
    }

    // Insert the transaction record into the database
    if err := tx.Create(&transaction).Error; err != nil {
        tx.Rollback() // Rollback the transaction if there's an error
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
        return
    }

    // Commit the transaction
    tx.Commit()

    c.JSON(http.StatusOK, gin.H{"message": "Wallet debited successfully"})
}

// GetWalletTransactions handles retrieving transaction history of a wallet
func GetWalletTransactions(c *gin.Context) {
    // Parse wallet ID from path parameter
    walletIDStr := c.Param("wallet_id")
    walletID, err := strconv.Atoi(walletIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wallet ID"})
        return
    }

    // Get wallet transactions
    transactions, err := getTransactions(walletID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve transactions"})
        return
    }

    // Return transaction history
    c.JSON(http.StatusOK, gin.H{"transactions": transactions})
}

// Function to retrieve transactions from the database
func getTransactions(walletID int) ([]models.Transaction, error) {
    // Declare a variable to store transactions
    var transactions []models.Transaction

    // Retrieve transactions from the database
    if err := db.DB.Where("wallet_id = ?", walletID).Find(&transactions).Error; err != nil {
        return nil, err
    }

    return transactions, nil
}