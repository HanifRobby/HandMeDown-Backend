package config

import (
	"fmt"
	"handmedown-backend/src/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB connects go to mysql database
func InitializeDB() (*gorm.DB, error) {
	errorENV := godotenv.Load()
	if errorENV != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=true&loc=Local", dbUser, dbPass, dbHost, dbName)
	DB, errorDB := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errorDB != nil {
		panic("Gagal terhubung ke database" + errorDB.Error())
	}

	// Set MaxIdleConns and MaxOpenConns for connection pooling
	sqlDB, err := DB.DB()
	if err != nil {
		panic("Gagal mendapatkan objek DB")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return DB, nil
}

// DisconnectDB is stopping your connection to mysql database
func DisconnectDB(db *gorm.DB) {
	if db != nil {
		dbSQL, err := db.DB()
		if err != nil {
			panic("Gagal untuk mematikan koneksi ke database")
		}
		dbSQL.Close()
	}
}

// Migrate Database
func MigrateDB(db *gorm.DB) {
	if db == nil {
		panic("DB is nil")
	}

	// AutoMigrate akan membuat tabel atau memperbarui struktur tabel sesuai dengan definisi model
	err := db.AutoMigrate(&models.User{}, &models.Akun{}, &models.Barang{}, &models.Keranjang{}, &models.OrderList{}, &models.KeranjangBarang{}, &models.Ulasan{})
	if err != nil {
		panic("Error during migration: " + err.Error())
	}
}
