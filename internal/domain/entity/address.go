package entity

import (
	"time"
)

type Address struct {
	ID        int    // 业务唯一标识
	UserID    int    // 用户ID
	Name      string // 收件人姓名
	Phone     string // 收件人电话
	Address   string // 收件人地址
	IsDefault bool   // 是否为默认地址
	CreatedAt time.Time
	UpdatedAt time.Time
}
