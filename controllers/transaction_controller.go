package controllers

import (
    "net/http"
    "wallet-api-go/models"
    "wallet-api-go/db"
    "github.com/gin-gonic/gin"
)

// CreateTransaction handles creating a new transaction
func CreateTransaction(c *gin.Context) {
    var transaction models.Transaction
    if err := c.ShouldBindJSON(&transaction); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Validate transaction input
    if transaction.Amount <= 0 || transaction.Description == "" || transaction.WalletID == 0 { // Change comparison to 0
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction input"})
        return
    }

    // Create transaction
    if err := db.DB.Create(&transaction).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Transaction created successfully"})
}

// GetTransactionByID handles retrieving a transaction by ID
func GetTransactionByID(c *gin.Context) {
    // Parse transaction ID from path parameter
    transactionID := c.Param("transaction_id")

    // Get transaction by ID
    transaction, err := getTransactionByID(transactionID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve transaction"})
        return
    }

    // Return transaction details
    c.JSON(http.StatusOK, gin.H{"transaction": transaction})
}

// Function to retrieve a transaction by its ID
func getTransactionByID(transactionID string) (*models.Transaction, error) {
    // Retrieve transaction from the database based on the transaction ID
    var transaction models.Transaction
    if err := db.DB.Where("id = ?", transactionID).First(&transaction).Error; err != nil {
        return nil, err
    }

    return &transaction, nil
}

// GetAllTransactions handles retrieving all transactions
func GetAllTransactions(c *gin.Context) {
    // Get all transactions (implement this logic according to your needs)

    // Return list of transactions
    c.JSON(http.StatusOK, gin.H{"message": "All transactions retrieved successfully"})
}
