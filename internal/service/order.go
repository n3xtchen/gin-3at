package service

import (
	"context"

	"gorm.io/gorm"

	e "github.com/n3xtchen/gin-3at/internal/domain/entity"
	repo "github.com/n3xtchen/gin-3at/internal/domain/repository"
	"github.com/n3xtchen/gin-3at/internal/pkg"
)

type OrderService interface {
	// CreateOrder creates a new order with the given details.
	CreateOrder(order *e.Order) error
	// GetOrderDetail retrieves the details of an order by its ID.
	GetOrderDetail(orderID int) (*e.Order, error)
	// You can add more methods as needed, such as UpdateOrder, DeleteOrder, etc.
}

type OrderServiceImp struct {
	db        *gorm.DB
	OrderRepo repo.OrderRepository
}

func NewOrderService(db *gorm.DB, orderRepo repo.OrderRepository) OrderService {
	return &OrderServiceImp{
		db:        db,
		OrderRepo: orderRepo,
	}
}

func (s *OrderServiceImp) CreateOrder(order *e.Order) error {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		ctx := pkg.WithTxToContext(context.Background(), tx)
		return s.OrderRepo.Save(ctx, order)
	})

	return err
}

func (s *OrderServiceImp) GetOrderDetail(orderID int) (*e.Order, error) {
	return s.OrderRepo.GetDetailByID(orderID)
}
