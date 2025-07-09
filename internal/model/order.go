package model

import (
	"time"

	"gorm.io/gorm"

	e "github.com/n3xtchen/gin-3at/internal/domain/entity"
)

type Order struct {
	gorm.Model
	OrderNum    string  // 订单号
	BuyerID     uint    `gorm:"not null"`
	User        User    `gorm:"foreignKey:BuyerID;references:ID"`
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

func (model Order) ToEntity() e.Order {
	orderItems := make([]e.OrderItem, len(model.Items))
	for i, item := range model.Items {
		orderItems[i] = item.ToEntity()

	}
	return e.Order{
		ID:          int(model.ID),
		OrderNumber: model.OrderNum,
		BuyerID:     int(model.BuyerID),
		AddressID:   int(model.AddressID),
		Address:     model.Address.ToEntity(), // Address should be set separately
		Items:       orderItems,
		Amount:      model.Amount,
		Status:      e.OrderStatus(model.Status),
		Remark:      model.Remark,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
		PaidAt:      model.PaidAt,
		ShippedAt:   model.ShippedAt,
		CompletedAt: model.CompletedAt,
		CanceledAt:  model.CanceledAt,
	}
}

func FromEntityOrder(entity *e.Order) *Order {
	// items := make([]OrderItem, len(entity.Items))
	// for i, item := range entity.Items {
	// 	items[i] = FromEntityOrderItem(item)
	// }
	return &Order{
		Model:    gorm.Model{ID: uint(entity.ID)},
		OrderNum: entity.OrderNumber,
		BuyerID:  uint(entity.BuyerID),
		// Address:     FromEntityAddress(entity.Address), // Address should be set separately
		// Items:       items,
		Amount:      entity.Amount,
		Status:      uint(entity.Status),
		Remark:      entity.Remark,
		PaidAt:      entity.PaidAt,
		ShippedAt:   entity.ShippedAt,
		CompletedAt: entity.CompletedAt,
		CanceledAt:  entity.CanceledAt,
	}
}
