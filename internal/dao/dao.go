package dao

import (
	"context"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dbContextKey = "db"

func InitMySQL(user, password, host string, port int, database string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8mb4",
		user, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

// 把事务 db 放入 context
func WithTxToContext(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, dbContextKey, tx)
}

// 从 context 获取 db，如果没传 tx 就返回默认 db
func GetDBFromContext(ctx context.Context, defaultDB *gorm.DB) *gorm.DB {
	tx, ok := ctx.Value(dbContextKey).(*gorm.DB)
	if ok && tx != nil {
		return tx
	}
	return defaultDB
}
