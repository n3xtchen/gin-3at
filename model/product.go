package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string   `gorm:"type:varchar(50) not null"`           // 商品名称
	Description string   `gorm:"type:varchar(255) not null"`          // 商品描述
	Price       float64  `gorm:"not null"`                            // 商品价格
	ImageURL    string   `gorm:"type:varchar(255) not null"`          // 商品图片
	Stock       int      `gorm:"not null"`                            // 商品库存
	CategoryID  uint     `gorm:"not null"`                            // 商品分类ID
	Category    Category `gorm:"foreignKey:CategoryID;references:ID"` // 商品分类
}
