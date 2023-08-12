package internal

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func SetRedis(key string, val string) bool {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		Username: "",
	})
	var ctx = context.Background()
	res, _ := rdb.Set(ctx, key, val, 0).Result()

	return res != ""
}
