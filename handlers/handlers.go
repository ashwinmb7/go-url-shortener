package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Shorten handles POST /shorten
// Expected JSON body: {"url": "https://example.com/long-url"}
// Returns: {"code": "abc123", "short_url": "http://localhost:8080/abc123"}
//
// TODO: Implement this handler to:
// 1. Parse the JSON request body to get the long URL
// 2. Validate the URL is valid (use url.Parse or similar)
// 3. Generate a unique code using utils.GenerateCode()
// 4. Check if code already exists in DB (handle collisions)
// 5. Save to database using GORM: db.Create(&urlModel)
// 6. Return the short code and URL in JSON response
func Shorten(c *gin.Context) {
	// TODO: Your implementation here
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented yet"})
}

// Redirect handles GET /:code
// This should:
// 1. Extract the code from URL parameter: c.Param("code")
// 2. Look up the URL in database by code
// 3. If found, increment the visits counter: url.Visits++
// 4. Save the updated visit count: db.Save(&url)
// 5. Redirect to the long URL: c.Redirect(http.StatusFound, url.LongURL)
// 6. If not found, return 404
func Redirect(c *gin.Context) {
	// TODO: Your implementation here
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented yet"})
}

