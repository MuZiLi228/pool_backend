package cache

import (
	"context"
	"pool_backend/configs"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

//New 实例化redis
func New() (*redis.Client, error) {
	client, err := redisConnect()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func redisConnect() (*redis.Client, error) {
	cfg := configs.Get().Redis
	client := redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		Password:     cfg.Pass,
		DB:           cfg.Db,
		MaxRetries:   cfg.MaxRetries,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, errors.Wrap(err, "ping redis err")
	}

	return client, nil
}
