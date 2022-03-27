package cache

import (
	"github.com/go-redis/redis/v8"
)

func New() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "cache:6379",
		Password: "",
	})
}
