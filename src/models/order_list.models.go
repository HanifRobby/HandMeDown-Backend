package models

import (
	"gorm.io/gorm"
)

type OrderList struct {
	gorm.Model
	PembeliID uint
	BarangID  uint
	PenjualID uint
	Pembeli   User   `gorm:"foreignKey:PembeliID"`
	Barang    Barang `gorm:"foreignKey:BarangID"`
	Penjual   User   `gorm:"foreignKey:PenjualID"`
}
