package models

// User represents a user in the system
type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"` // Note: You should hash passwords before storing them in the database
    APIKey   string `json:"api_key"`
}

// NewUser creates and returns a new User instance
func NewUser(username, email, password, apiKey string) *User {
    return &User{
        Username: username,
        Email:    email,
        Password: password,
        APIKey:   apiKey,
    }
}
