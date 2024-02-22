// db/transaction.go

package db

import (
    "wallet-api-go/models"
    // "github.com/google/uuid"
)

// GetTransactionsByWalletID retrieves transactions by wallet ID using GORM
func GetTransactionsByWalletID(walletID string) ([]models.Transaction, error) {
    var transactions []models.Transaction
    if err := DB.Where("wallet_id = ?", walletID).Find(&transactions).Error; err != nil {
        return nil, err
    }
    return transactions, nil
}
