package cache

import (
	"fmt"

	"api/config"

	"github.com/go-redis/redis/v8"
)

func New() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:6379", config.REDIS_HOST),
		Password: config.REDIS_PASSWORD,
	})
}
