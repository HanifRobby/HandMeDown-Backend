package models

import (
	"gorm.io/gorm"
)

type Akun struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	UserID   uint
	User     User `gorm:"foreignKey:UserID"`
}
