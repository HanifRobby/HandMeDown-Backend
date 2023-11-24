package controllers

import (
	"net/http"

	"handmedown-backend/src/middleware"

	"github.com/gin-gonic/gin"
)

// LoginHandler handles the login request
func LoginHandler(c *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Your authentication logic here
	if credentials.Username == "user" && credentials.Password == "pass" {
		// Create JWT token
		tokenString, err := middleware.CreateToken(credentials.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}
