package models

import (
	"gorm.io/gorm"
)

type Ulasan struct {
	gorm.Model
	PembeliID uint
	PenjualID uint
	Ulasan    string
	Pembeli   User `gorm:"foreignKey:PembeliID"`
	Penjual   User `gorm:"foreignKey:PenjualID"`
}
