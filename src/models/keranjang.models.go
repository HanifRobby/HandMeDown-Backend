package models

import (
	"gorm.io/gorm"
)

type Keranjang struct {
	gorm.Model
	UserID uint
	User   User `gorm:"foreignKey:UserID"`
}
