package controllers

import (
	"handmedown-backend/src/config"
	"handmedown-backend/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrderList(context *gin.Context) {
	db := config.DB

	userID, exists := context.Get("userID")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Konversi userID ke uint
	userIDUint, ok := userID.(uint)
	if !ok {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userID type"})
		return
	}

	// Menggunakan userID untuk mengambil order list
	var orderList []models.OrderList
	result := db.Preload("Barang").Preload("Penjual").Where("pembeli_id = ?", userIDUint).Find(&orderList)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching order list"})
		return
	}

	// Membuat response berdasarkan order list yang ditemukan
	var orderListResponse []gin.H
	for _, order := range orderList {
		orderListResponse = append(orderListResponse, gin.H{
			"IDOrderlist": order.ID,
			"IDPenjual":   order.PenjualID,
			"NamaPenjual": order.Penjual.Nama,
			"NamaBarang":  order.Barang.NamaBarang,
			"HargaBarang": order.Barang.Harga,
		})
	}

	// Mengembalikan response
	context.JSON(http.StatusOK, gin.H{"data": orderListResponse})
}
