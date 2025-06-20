package repository

import (
	"context"

	e "github.com/n3xtchen/gin-3at/internal/domain/entity"
)

type OrderRepository interface {

	// Get Order Detail
	GetDetailByID(orderID int) (*e.Order, error)

	// CreateOrder creates a new order with the given details.
	Save(ctx context.Context, order *e.Order) error
}
