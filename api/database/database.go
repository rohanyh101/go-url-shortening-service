package database

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

func CreateClient(dbNo int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDR"),
		Password: os.Getenv("DB_PASS"),
		DB:       dbNo,
	})

	// // Check if the client is connected successfully
	// _, err := rdb.Ping(Ctx).Result()
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	// }

	// return rdb, nil

	return rdb
}
