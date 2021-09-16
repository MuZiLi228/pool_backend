package cache

import (
	"context"
	redisPack "github.com/go-redis/redis/v8"
)

//Subscribe 订阅
func Subscribe(ctx context.Context, channels []string) (client *redisPack.Client, pubsub *redisPack.PubSub, err error) {
	client, err = redisConnect()
	if err != nil {
		return
	}
	// defer client.Close()
	pubsub = client.Subscribe(ctx, channels...)
	return
}

//Publish 发布
func Publish(ctx context.Context, channel, data string) (err error) {
	client, err := redisConnect()
	if err != nil {
		return
	}
	defer client.Close()
	err = client.Publish(ctx, channel, data).Err()
	return
}
