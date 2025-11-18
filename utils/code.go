package utils

import (
	"math/rand"
	"time"
)

// GenerateCode creates a random short code for URL shortening
// Returns a string of random alphanumeric characters
//
// How it works:
// 1. Creates a charset with all letters (a-z, A-Z) and numbers (0-9)
// 2. Creates a byte slice of the desired length
// 3. For each position, picks a random index in the charset
// 4. Converts the byte slice to a string and returns it
func GenerateCode(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Create a new random source with current time as seed
	// (rand.Seed is deprecated in Go 1.20+, so we use rand.New instead)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	resultBytes := make([]byte, length)

	for i := range resultBytes {
		randomIndex := r.Intn(len(charset))
		resultBytes[i] = charset[randomIndex]
	}

	return string(resultBytes)
}
