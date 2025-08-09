package pkg

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func InitCache(host string, port int, password string, db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + strconv.Itoa(port),
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	// Test the connection
	if err := client.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	return client
}
