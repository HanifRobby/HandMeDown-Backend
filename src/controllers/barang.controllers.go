package controllers

import (
	"handmedown-backend/src/config"
	"handmedown-backend/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type productRequest struct {
// }

func GetAllProducts(context *gin.Context) {

	db := config.DB

	var products []models.Barang

	// Query untuk mencari data produk
	err := db.Find(&products)
	if err.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error getting data"})
		return
	}

	// Creating HTTP Response
	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    products,
	})
}
