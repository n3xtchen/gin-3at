package model

import (
	"gorm.io/gorm"

	e "github.com/n3xtchen/gin-3at/internal/domain/entity"
)

type Category struct {
	gorm.Model
	Name        string `gorm:"type:varchar(50) not null"`  // 分类名称
	Description string `gorm:"type:varchar(255) not null"` // 分类描述
	// ParentID    uint      `gorm:"not null"`                            // 父分类ID，0表示顶级分类
	// Parent      *Category `gorm:"foreignKey:ParentID;references:ID"`   // 父分类
	Products []Product `gorm:"foreignKey:CategoryID;references:ID"` // 该分类下的商品
}

func (model Category) ToEntity() e.Category {
	return e.Category{
		ID:          int(model.ID),
		Name:        model.Name,
		Description: model.Description,
	}
}

func FromEntityCategory(entity *e.Category) *Category {
	return &Category{
		Model:       gorm.Model{ID: uint(entity.ID)},
		Name:        entity.Name,
		Description: entity.Description,
	}
}
