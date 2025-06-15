package entity

import (
	"time"
)

type OrderStatus uint

// 支付状态枚举
const (
	OrderStatusUnpaid    OrderStatus = 1 // 未支付
	OrderStatusPaid      OrderStatus = 2 // 已支付
	OrderStatusCancelled OrderStatus = 3 // 已取消
	OrderStatusRefunded  OrderStatus = 4 // 已退款
	OrderStatusCompleted OrderStatus = 5 // 已完成
)

type Order struct {
	ID          string      // 业务唯一标识
	OrderNumber string      // 订单编号
	BuyerID     string      // 买家ID
	Items       []OrderItem // 订单项
	Amount      float64     // 总金额
	Status      OrderStatus // 订单状态
	Remark      string      // 备注
	CreatedAt   time.Time
	UpdatedAt   time.Time
	PaidAt      *time.Time
	ShippedAt   *time.Time
	CompletedAt *time.Time
	CanceledAt  *time.Time
}

type OrderItem struct {
	ProductID   string
	ProductName string
	Quantity    int
	Price       float64
	Total       float64
}
