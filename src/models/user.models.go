package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama   string
	Email  string
	NoTelp string
	Alamat string
}
