package controllers

import (
	"handmedown-backend/src/config"
	"handmedown-backend/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var requestAddCartItem struct {
	ProductID uint `json:"product_id"`
}

func AddToCart(context *gin.Context) {

	if err := context.ShouldBindJSON(&requestAddCartItem); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Dapatkan informasi pengguna dari konteks (setelah melewati middleware autorisasi)
	userID, exists := context.Get("userID")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Konversi userID ke uint
	userIDUint, ok := userID.(uint)
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Cari keranjang berdasarkan UserID
	var cart models.Keranjang
	if err := config.DB.Where("user_id = ?", userIDUint).First(&cart).Error; err != nil {
		// Jika keranjang belum ada, buat keranjang baru
		cart.UserID = userIDUint
		if err := config.DB.Create(&cart).Error; err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
	}

	// Tambahkan barang ke keranjang_barang
	cartItem := models.KeranjangBarang{
		KeranjangID: cart.ID,
		BarangID:    requestAddCartItem.ProductID,
	}

	if err := config.DB.Create(&cartItem).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Kirim respons sukses
	context.JSON(http.StatusOK, gin.H{"message": "Product added to cart successfully"})

}

var requestDeleteCart struct {
	IDBarang uint `json:"id_barang"`
}

func DeleteCartItem(context *gin.Context) {
	if err := context.ShouldBindJSON(&requestDeleteCart); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Dapatkan informasi pengguna dari konteks (setelah melewati middleware autorisasi)
	userID, exists := context.Get("userID")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Konversi userID ke uint
	userIDUint, ok := userID.(uint)
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Cari keranjang berdasarkan UserID
	var cart models.Keranjang
	if err := config.DB.Where("user_id = ?", userIDUint).First(&cart).Error; err != nil {
		cart.UserID = userIDUint
		if err := config.DB.Create(&cart).Error; err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
	}

	// Tambahkan barang ke keranjang_barang
	cartItem := models.KeranjangBarang{
		KeranjangID: cart.ID,
		BarangID:    requestDeleteCart.IDBarang,
	}

	if err := config.DB.Where("keranjang_id = ? AND barang_id = ?", cartItem.KeranjangID, cartItem.BarangID).Delete(&cartItem).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Kirim respons sukses
	context.JSON(http.StatusOK, gin.H{"message": "Delete Item from Cart Success"})
}
