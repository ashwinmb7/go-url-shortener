package database

import (
	"go-url-shortener/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init initializes the database connection and runs migrations
// This function:
// 1. Connects to SQLite database (creates if doesn't exist)
// 2. Runs AutoMigrate to create/update the URL table
// 3. Stores the connection in the global DB variable
func Init() error {
	var err error
	
	// TODO: Connect to SQLite database named "shortener.db"
	// Use gorm.Open(sqlite.Open("shortener.db"), &gorm.Config{})
	// Don't forget to handle errors!
	
	// TODO: Run AutoMigrate for the models.URL model
	// This creates the table if it doesn't exist
	
	return err
}

// GetDB returns the global database instance
func GetDB() *gorm.DB {
	return DB
}

