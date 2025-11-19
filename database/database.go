package database

import (
	"go-url-shortener/models"

	"github.com/glebarez/sqlite"
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

	// Connect to SQLite database named "shortener.db"
	DB, err = gorm.Open(sqlite.Dialector{DSN: "shortener.db"}, &gorm.Config{})
	if err != nil {
		return err
	}

	// Run AutoMigrate for the models.URL model
	// This creates the table if it doesn't exist
	err = DB.AutoMigrate(&models.URL{})
	if err != nil {
		return err
	}

	return nil
}

// GetDB returns the global database instance
func GetDB() *gorm.DB {
	return DB
}
