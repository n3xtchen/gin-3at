package model

import (
	"gorm.io/gorm"

	e "github.com/n3xtchen/gin-3at/internal/domain/entity"
)

type Address struct {
	gorm.Model
	UserID    int    `gorm:"not null"`
	User      User   `gorm:"foreignKey:UserID;references:ID"`
	Name      string `gorm:"type:varchar(20) not null"`
	Phone     string `gorm:"type:varchar(11) not null"`
	Address   string `gorm:"type:varchar(50) not null"`
	IsDefault bool   `gorm:"not null"` // 是否为默认地址
}

func (model Address) ToEntity() e.Address {
	return e.Address{
		ID:        int(model.ID),
		UserID:    int(model.UserID),
		Name:      model.Name,
		Phone:     model.Phone,
		Address:   model.Address,
		IsDefault: model.IsDefault,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func FromEntityAddress(entity e.Address) Address {
	return Address{
		Model:     gorm.Model{ID: uint(entity.ID)},
		UserID:    entity.UserID,
		Name:      entity.Name,
		Phone:     entity.Phone,
		Address:   entity.Address,
		IsDefault: entity.IsDefault,
	}
}
