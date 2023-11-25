package controllers

import (
	"handmedown-backend/src/config"
	"handmedown-backend/src/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductResponse struct {
	ID          uint    `json:"ID"`
	NamaBarang  string  `json:"NamaBarang"`
	Harga       float64 `json:"Harga"`
	Terjual     bool    `json:"Terjual"`
	URLGambar   string  `json:"URLGambar"`
	PenjualID   uint    `json:"PenjualID"`
	NamaPenjual string  `json:"NamaPenjual"`
}

func GetAllProducts(context *gin.Context) {

	db := config.DB

	var products []models.Barang

	// Query untuk mencari data produk
	err := db.Preload("Penjual").Find(&products)
	if err.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error getting data"})
		return
	}

	var productResponses []ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ProductResponse{
			ID:          product.ID,
			NamaBarang:  product.NamaBarang,
			Harga:       product.Harga,
			Terjual:     product.Terjual,
			PenjualID:   product.PenjualID,
			URLGambar:   product.URLGambar,
			NamaPenjual: product.Penjual.Nama,
		})
	}

	// Creating HTTP Response
	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    productResponses,
	})
}

type ProductDetailResponse struct {
	ID          uint    `json:"ID"`
	NamaBarang  string  `json:"NamaBarang"`
	Harga       float64 `json:"Harga"`
	Deskripsi   string  `json:"Deskripsi"`
	Terjual     bool    `json:"Terjual"`
	URLGambar   string  `json:"URLGambar"`
	PenjualID   uint    `json:"PenjualID"`
	NamaPenjual string  `json:"NamaPenjual"`
	NoTelp      string  `json:"NoTelp"`
	Alamat      string  `json:"Alamat"`
}

func GetProductDetail(context *gin.Context) {
	db := config.DB

	// Mendapatkan ID barang dari parameter URL
	productID, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var product models.Barang
	// Query untuk mencari data produk berdasarkan ID
	err = db.Preload("Penjual").First(&product, productID).Error
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Membuat respons HTTP
	productResponse := ProductDetailResponse{
		ID:          product.ID,
		NamaBarang:  product.NamaBarang,
		Harga:       product.Harga,
		Deskripsi:   product.Deskripsi,
		Terjual:     product.Terjual,
		PenjualID:   product.PenjualID,
		URLGambar:   product.URLGambar,
		NamaPenjual: product.Penjual.Nama,
		NoTelp:      product.Penjual.NoTelp,
		Alamat:      product.Penjual.Alamat,
	}

	context.JSON(http.StatusOK, gin.H{
		"data": productResponse,
	})
}
