package db

import (
    "fmt"
    "os"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB // Exported variable

// InitDB initializes a connection to the MySQL database using GORM
func InitDB() (*gorm.DB, error) {
    // Print out the values of environment variables
    fmt.Println("DB_USERNAME:", os.Getenv("DB_USERNAME"))
    fmt.Println("DB_PASSWORD:", os.Getenv("DB_PASSWORD"))
    fmt.Println("DB_HOST:", os.Getenv("DB_HOST"))
    fmt.Println("DB_PORT:", os.Getenv("DB_PORT"))
    fmt.Println("DB_NAME:", os.Getenv("DB_NAME"))

    // Construct the dataSourceName using environment variables
    dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
        os.Getenv("DB_USERNAME"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"))

    // Connect to the database using GORM
    db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    DB = db // Assign to exported DB variable

    return db, nil
}

// CloseDB closes the database connection
func CloseDB() {
    sqlDB, err := DB.DB()
    if err != nil {
        // Handle error
    }
    sqlDB.Close()
}
