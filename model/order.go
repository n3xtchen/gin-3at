package model

import "gorm.io/gorm"

// 支付状态枚举
const (
	OrderStatusUnpaid    = 1 // 未支付
	OrderStatusPaid      = 2 // 已支付
	OrderStatusCancelled = 3 // 已取消
	OrderStatusRefunded  = 4 // 已退款
	OrderStatusCompleted = 5 // 已完成
)

type Order struct {
	gorm.Model
	OrderNum   string  // 订单号
	UserID     uint    `gorm:"not null"`
	AddressID  uint    `gorm:"not null"`
	Address    Address `gorm:"foreignKey:AddressID;references:ID"`
	Amount     float64 `gorm:"not null"`
	OrderItems []OrderItem
	Status     uint // 1 未支付  2 已支付
}
