package handlers

import (
	"net/http"

	"go-url-shortener/database"
	"go-url-shortener/models"
	"go-url-shortener/utils"

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
	var request struct {
		URL string `json:"url" binding:"required,url"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate unique code (handle collisions)
	db := database.GetDB()
	var code string
	for {
		code = utils.GenerateCode(6)
		var existingURL models.URL
		if err := db.Where("code = ?", code).First(&existingURL).Error; err != nil {
			// Code doesn't exist, we can use it
			break
		}
		// Code exists, generate a new one (loop continues)
	}

	// Create URL model
	urlModel := models.URL{
		Code:    code,
		LongURL: request.URL,
	}

	// Save to database
	if err := db.Create(&urlModel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create short URL"})
		return
	}

	// Return success response
	shortURL := "http://localhost:8080/" + code
	c.JSON(http.StatusCreated, gin.H{
		"code":      code,
		"short_url": shortURL,
	})
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
	// Extract code from URL parameter
	code := c.Param("code")

	// Look up URL in database by code
	var url models.URL
	db := database.GetDB()

	if err := db.Where("code = ?", code).First(&url).Error; err != nil {
		// If not found, return 404
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	// Increment visits counter
	url.Visits++

	// Save the updated visit count
	db.Save(&url)

	// Redirect to the long URL
	c.Redirect(http.StatusFound, url.LongURL)
}
