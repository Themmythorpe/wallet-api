package main

import (
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "wallet-api-go/db"
    "wallet-api-go/router"
)

func main() {
    // Load environment variables from .env file
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Initialize database connection
    _, err := db.InitDB()
    if err != nil {
        log.Fatalf("Error initializing database: %v", err)
    }
    defer db.CloseDB()

    // Set Gin to release mode if running in production
    if os.Getenv("GIN_MODE") == "release" {
        gin.SetMode(gin.ReleaseMode)
    }

    // Create a Gin router
    r := gin.Default()

    // Setup routes
    router.SetupRouter(r)

    // Run the server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default port if not specified
    }
    err = r.Run(":" + port)
    if err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}
