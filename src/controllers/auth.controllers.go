package controllers

import (
	"handmedown-backend/src/config"
	"handmedown-backend/src/models"
	"net/http"

	"handmedown-backend/src/middleware"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// RegisterHandlers handles the registration request
func RegisterHandler(c *gin.Context) {
	var registerData models.Akun
	if err := c.BindJSON(&registerData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Hash password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerData.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Update Akun model with hashed password
	registerData.Password = string(hashedPassword)

	// Create a new user account
	if err := config.DB.Create(&registerData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user account"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

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

	var akun models.Akun
	result := config.DB.Where("username = ?", credentials.Username).Preload("User").First(&akun)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Verifikasi password menggunakan bcrypt
	err := bcrypt.CompareHashAndPassword([]byte(akun.Password), []byte(credentials.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Create JWT token
	tokenString, err := middleware.CreateToken(credentials.Username, akun.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})

}
