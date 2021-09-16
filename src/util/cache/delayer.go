package cache

import (
	"github.com/go-redis/redis/v8"
)

var clients map[string]*redis.Client

func InitRedis(addr, password string, db int, prefix string) {
	if clients == nil {
		clients = make(map[string]*redis.Client)
	}
	if prefix == "" {
		prefix = "default"
	}

	clients[prefix] = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return
}

func GetRedisClient(prefix string) *redis.Client {
	if prefix == "" {
		prefix = "default"
	}

	return clients[prefix]
}
