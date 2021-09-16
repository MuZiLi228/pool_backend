package delayer

import (
	"errors"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/go-redis/redis/v8"
)

const delayerCachePrefix = "delayer"

var (
	snowflakeNode *snowflake.Node
	rdc           *redis.Client
	rdcPrefix     = delayerCachePrefix
)

// NewClient 创建Client
func NewClient(opts ...ClientOption) (*Client, error) {
	settings := initSettings()

	for _, opt := range opts {
		opt.Apply(settings)
	}

	if err := settings.Validate(); err != nil {
		return nil, err
	}

	client := &Client{
		queues: make(map[string]*Queue),
	}

	return client, nil
}

type Client struct {
	queues map[string]*Queue // 队列集
}

// NewQueue 创建新队列
func (c *Client) NewQueue(queueName string, scanInterval time.Duration, poolSize int) (err error) {
	if queueName == "" {
		return errors.New("name 为空")
	} else if poolSize <= 0 {
		return errors.New("poolSize 数值错误")
	}

	if _, ok := c.queues[queueName]; ok {
		return errors.New("该队列已存在")
	}

	if c.queues[queueName], err = newQueue(queueName, scanInterval, poolSize); err != nil {
		return err
	}

	return nil
}

// UseQueue 使用对应的队列
func (c *Client) UseQueue(queueName string) *Queue {
	return c.queues[queueName]
}
