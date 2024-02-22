// models/transaction.go

package models

import "time"

// Credit represents a credit transaction in the system
type Credit struct {
    Amount      float64 `json:"amount"`
    Description string  `json:"description"`
}

// Debit represents a debit transaction in the system
type Debit struct {
    Amount      float64 `json:"amount"`
    Description string  `json:"description"`
}

// Transaction represents a transaction in the system
type Transaction struct {
    ID          int    `json:"id"`
    WalletID    int    `json:"wallet_id"`
    Amount      float64   `json:"amount"`
    Description string    `json:"description"`
    Type        string    `json:"type"` // "credit" or "debit"
    CreatedAt   string `json:"created_at"` // Change the type to string
}

// NewTransaction creates and returns a new Transaction instance
func NewTransaction(walletID int, amount float64, description, t string) *Transaction {
    return &Transaction{
        WalletID:    walletID,
        Amount:      amount,
        Description: description,
        Type:        t,
        CreatedAt:   time.Now().Format(time.RFC3339), // Convert time to string using a specific format
    }
}
