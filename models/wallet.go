package models

// Wallet represents a user's wallet in the system
type Wallet struct {
    ID       int     `json:"id"`
    UserID   int     `json:"user_id"`
    Currency string  `json:"currency"`
    Balance  float64 `json:"balance"`
}

// NewWallet creates and returns a new Wallet instance
func NewWallet(userID int, currency string, balance float64) *Wallet {
    return &Wallet{
        UserID:   userID,
        Currency: currency,
        Balance:  balance,
    }
}
