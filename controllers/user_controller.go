package controllers

import (
    "net/http"
	"time"
	"wallet-api-go/models"
	"wallet-api-go/db"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// RegisterUser handles user registration
func RegisterUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Validate user input
    if user.Email == "" || user.Password == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
        return
    }

    // Check if user already exists
    if userExists(user.Email) {
        c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
        return
    }

    // Create user
    if err := createUser(&user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Function to check if user already exists
func userExists(email string) bool {
    var existingUser models.User
    if err := db.DB.Where("email = ?", email).First(&existingUser).Error; err != nil {
        return false // User does not exist
    }
    return true // User exists
}

// Function to create a new user
func createUser(user *models.User) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    user.Password = string(hashedPassword)

    if err := db.DB.Create(&user).Error; err != nil {
        return err
    }

    return nil
}

// LoginUser handles user login
func LoginUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate user input
	if user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	// Authenticate user
	authenticated, err := authenticateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to authenticate user"})
		return
	}

	if !authenticated {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate JWT token
	token, err := generateToken(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// authenticateUser authenticates the user by verifying the email and password
func authenticateUser(user *models.User) (bool, error) {
	var existingUser models.User
	// Query the database to retrieve the user with the provided email
	if err := db.DB.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
		// User not found or error occurred
		return false, err
	}

	// Compare the provided password with the hashed password from the database
	err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		// Passwords don't match
		return false, nil
	}

	// Passwords match, user is authenticated
	return true, nil
}

// generateToken generates a JWT token for the provided email
func generateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
