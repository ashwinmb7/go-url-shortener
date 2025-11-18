package models

import "time"

// URL represents a shortened URL in the database
// This is your database model - it maps to a table called "urls"
type URL struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Code      string    `gorm:"uniqueIndex;not null" json:"code"`        // Short code (e.g., "abc123")
	LongURL   string    `gorm:"not null" json:"long_url"`                // Original long URL
	Visits    int64     `gorm:"default:0" json:"visits"`                 // Visit counter
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TODO: You'll use GORM to auto-migrate this model to create the table
// This happens in your database initialization code

