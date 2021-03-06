package db

import (
	"context"
	"github.com/Swan/Nameless/config"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
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

	log.Info("Successfully connected to Redis!")
}
