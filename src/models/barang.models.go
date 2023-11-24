package models

import (
	"gorm.io/gorm"
)

type Barang struct {
	gorm.Model
	NamaBarang string
	Harga      float64
	Deskripsi  string
	Terjual    bool
	PenjualID  uint
	URLGambar  string
	Penjual    User `gorm:"foreignKey:PenjualID"`
}
