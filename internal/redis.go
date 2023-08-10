package internal

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func SetRedis(val string) bool {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		Username: "",
	})
	var ctx = context.Background()
	res, _ := rdb.Set(ctx, val, "sample", 0).Result()

	return res != ""
}
