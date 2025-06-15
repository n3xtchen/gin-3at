package model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderNum    string  // 订单号
	BuyerID     uint    `gorm:"not null"`
	AddressID   uint    `gorm:"not null"`
	Address     Address `gorm:"foreignKey:AddressID;references:ID"`
	Amount      float64 `gorm:"not null"`
	Items       []OrderItem
	Status      uint   // 1 未支付  2 已支付
	Remark      string // 备注
	PaidAt      *time.Time
	ShippedAt   *time.Time
	CompletedAt *time.Time
	CanceledAt  *time.Time
}
