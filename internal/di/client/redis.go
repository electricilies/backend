package client

import (
	"backend/config"

	"github.com/redis/go-redis/v9"
)

func NewRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: config.Cfg.RedisAddr,
	})
}
