package db

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client
var RedisCtx context.Context = context.Background()

// InitializeRedis Initializes a Redis client
func InitializeRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	fmt.Println("Redis Client has been initialized.")
}
