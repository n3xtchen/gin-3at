package model

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderID   uint    `gorm:"not null"`
	ProductID uint    `gorm:"not null"`
	Product   Product `gorm:"foreignKey:ProductID;references:ID"` // 商品信息
	Price     float64
	Quantity  int // 数量
}
