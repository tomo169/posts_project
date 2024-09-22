package config

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client
var Ctx = context.Background()

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	// Test the connection
	err := Rdb.Ping(Ctx).Err()
	if err != nil {
		panic(err)
	}
}
