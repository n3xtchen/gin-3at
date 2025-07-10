package model

import (
	"gorm.io/gorm"

	e "github.com/n3xtchen/gin-3at/internal/domain/entity"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20) not null;unique"` // 用户名
	Password string `gorm:"type:varchar(255) not null"`       // 密码
	Email    string `gorm:"type:varchar(50) not null;unique"` // 邮箱
}

func FromEntityUser(entity *e.User) *User {
	return &User{
		Model:    gorm.Model{ID: uint(entity.ID)},
		Username: entity.Username,
		Password: entity.Password,
		Email:    entity.Email,
	}
}

func (model User) ToEntity() e.User {
	return e.User{
		ID:       int(model.ID),
		Username: model.Username,
		Password: model.Password,
		Email:    model.Email,
	}
}
