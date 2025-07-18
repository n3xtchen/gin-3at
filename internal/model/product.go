package model

import (
	e "github.com/n3xtchen/gin-3at/internal/domain/entity"
	"gorm.io/gorm"
)

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

func (model Product) ToEntity() e.Product {
	return e.Product{
		ID:          int(model.ID),
		Name:        model.Name,
		Description: model.Description,
		Price:       model.Price,
		ImageURL:    model.ImageURL,
		Stock:       model.Stock,
		CategoryID:  int(model.CategoryID),
		Category:    model.Category.ToEntity(), // Category should be set separately
	}
}

func FromEntityProduct(entity *e.Product) *Product {
	return &Product{
		Model:       gorm.Model{ID: uint(entity.ID)},
		Name:        entity.Name,
		Description: entity.Description,
		Price:       entity.Price,
		ImageURL:    entity.ImageURL,
		Stock:       entity.Stock,
		CategoryID:  uint(entity.CategoryID),
		// Category will be set separately if needed
	}
}
