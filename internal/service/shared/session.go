package shared

import (
	"context"

	e "github.com/n3xtchen/gin-3at/internal/domain/entity"
)

type SessionContext interface {
	UserID() string
	save(ctx context.Context, user *e.User) error
	clear(ctx context.Context) error
}
