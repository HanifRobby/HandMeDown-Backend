package controllers

import (
	"fmt"
	"handmedown-backend/src/config"
	"handmedown-backend/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SimplifiedCartResponse struct {
	IDBarang    uint    `json:"IDBarang"`
	NamaBarang  string  `json:"NamaBarang"`
	HargaBarang float64 `json:"HargaBarang"`
	IDPenjual   uint    `json:"IDPenjual"`
	NamaPenjual string  `json:"NamaPenjual"`
}

func GetCart(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Konversi userID ke uint
	userIDUint, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userID type"})
		return
	}

	// Query untuk mendapatkan keranjang berdasarkan userID
	var keranjang models.Keranjang
	if err := config.DB.Where("user_id = ?", userIDUint).First(&keranjang).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting user's cart"})
		return
	}

	var keranjangBarang []models.KeranjangBarang
	if err := config.DB.Preload("Barang").Where("keranjang_id = ?", keranjang.ID).Find(&keranjangBarang).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting user's cart items"})
		return
	}

	// var penjual models.User
	for i, kb := range keranjangBarang {
		// fmt.Printf("PenjualID: %d\n", kb.Barang.PenjualID)
		var penjual models.User
		if err := config.DB.Where("id = ?", kb.Barang.PenjualID).First(&penjual).Error; err != nil {
			fmt.Printf("Error loading seller information: %s\n", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting seller information"})
			return
		}

		// Set informasi penjual ke dalam objek barang yang sesuai
		keranjangBarang[i].Barang.Penjual = penjual
	}

	var simplifiedCart []SimplifiedCartResponse
	for _, item := range keranjangBarang {
		simplifiedCart = append(simplifiedCart, SimplifiedCartResponse{
			IDBarang:    item.Barang.ID,
			NamaBarang:  item.Barang.NamaBarang,
			HargaBarang: item.Barang.Harga,
			IDPenjual:   item.Barang.Penjual.ID,
			NamaPenjual: item.Barang.Penjual.Nama,
		})
	}

	// Lakukan sesuatu dengan data keranjang, misalnya kirim sebagai respons
	c.JSON(http.StatusOK, gin.H{"data": simplifiedCart})
}
