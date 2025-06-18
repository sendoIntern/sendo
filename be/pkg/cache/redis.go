package cache

import (
	"be/pkg/config"
	"context"
	"log"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	client *redis.Client
	once   sync.Once
	ctx    = context.Background()
)

func InitRedis() {
	once.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr:     config.GetEnv("REDIS_ADDR", "localhost:6379"), // service name in docker-compose
			Password: "",                                            // let empty if do not have set password
			DB:       0,
		})
	})
}

func GetRedis() *redis.Client {
	InitRedis()
	return client
}

func SetCache(key string, value interface{}, ttl time.Duration) error {
	return GetRedis().Set(ctx, key, value, ttl).Err()
}

func GetCache(key string) (string, error) {
	return GetRedis().Get(ctx, key).Result()
}

func ClearCacheByKey(key string) {
	keys, err := GetRedis().Keys(ctx, key).Result()
	if err != nil {
		log.Println("Failed to find keys to clear:", err)
		return
	}
	if len(keys) > 0 {
		if err := GetRedis().Del(ctx, keys...).Err(); err != nil {
			log.Println("Failed to delete keys:", err)
		}
	}
}
