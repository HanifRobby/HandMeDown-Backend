package models

import (
	"gorm.io/gorm"
)

	type KeranjangBarang struct {
		gorm.Model
		KeranjangID uint
		BarangID    uint
		Keranjang   Keranjang `gorm:"foreignKey:KeranjangID"`
		Barang      Barang    `gorm:"foreignKey:BarangID"`
	}
