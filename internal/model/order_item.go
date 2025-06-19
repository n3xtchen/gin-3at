package model

import (
	"gorm.io/gorm"

	e "github.com/n3xtchen/gin-3at/internal/domain/entity"
)

type OrderItem struct {
	gorm.Model
	OrderID   uint    `gorm:"not null"`
	ProductID uint    `gorm:"not null"`
	Product   Product `gorm:"foreignKey:ProductID;references:ID"` // 商品信息
	Price     float64
	Quantity  int // 数量
}

func (model OrderItem) ToEntity() e.OrderItem {
	return e.OrderItem{
		ID:          int(model.ID),
		ProductID:   int(model.ProductID),
		ProductName: model.Product.Name, // Assuming Product has a Name field
		Quantity:    model.Quantity,
		Price:       model.Price,
		Total:       model.Price * float64(model.Quantity),
	}
}

func FromEntityOrderItem(entity e.OrderItem) OrderItem {
	return OrderItem{
		ProductID: uint(entity.ProductID),
		Price:     entity.Price,
		Quantity:  entity.Quantity,
	}
}
