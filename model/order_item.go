package model

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderID   uint `gorm:"not null"`
	ProductID uint `gorm:"not null"`
	Price     float64
	Quantity  int // 数量
}
