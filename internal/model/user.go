package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20) not null;unique"` // 用户名
	Password string `gorm:"type:varchar(255) not null"`       // 密码
	Email    string `gorm:"type:varchar(50) not null;unique"` // 邮箱
	Phone    string `gorm:"type:varchar(11) not null;unique"` // 手机号
}
