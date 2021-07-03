package db

import (
	"context"
	"fmt"
	"github.com/Swan/Nameless/src/config"
	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client
var RedisCtx context.Context = context.Background()

// InitializeRedis Initializes a Redis client
func InitializeRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     config.Data.Redis.Address,
		Password: config.Data.Redis.Password,
		DB:       config.Data.Redis.DB,
	})

	fmt.Println("Redis Client has been initialized.")
}
