package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID        uint `gorm:"primary_key"`
	Code      string
	Price     uint
	Type      string
	Warehouse uint
}
