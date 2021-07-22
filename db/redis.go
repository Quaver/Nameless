package db

import (
	"context"
	config2 "github.com/Swan/Nameless/config"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

var Redis *redis.Client
var RedisCtx context.Context = context.Background()

// InitializeRedis Initializes a Redis client
func InitializeRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     config2.Data.Redis.Address,
		Password: config2.Data.Redis.Password,
		DB:       config2.Data.Redis.DB,
	})

	log.Info("Successfully connected to Redis!")
}
