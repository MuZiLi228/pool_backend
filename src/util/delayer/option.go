package delayer

import (
	"github.com/bwmarrin/snowflake"
	"github.com/go-redis/redis/v8"
)

type ClientOption interface {
	Apply(settings *ClientSetting)
}

// WithRedis 设置redis客户端连接
func WithRedis(rdc *redis.Client) ClientOption {
	return withRedis{rdc}
}

type withRedis struct{ Rdc *redis.Client }

func (w withRedis) Apply(*ClientSetting) {
	rdc = w.Rdc
}

// WithRedisPrefix 设置缓存指定前缀
func WithRedisPrefix(prefix string) ClientOption {
	return withRedisPrefix{prefix}
}

type withRedisPrefix struct{ rdcPrefix string }

func (w withRedisPrefix) Apply(*ClientSetting) {
	rdcPrefix = w.rdcPrefix
}

func WithSnowflakeNode(node int64) ClientOption {
	return withSnowflakeNode{node}
}

type withSnowflakeNode struct{ id int64 }

func (w withSnowflakeNode) Apply(*ClientSetting) {
	var err error
	snowflakeNode, err = snowflake.NewNode(w.id)
	if err != nil {
		panic(err.Error())
	}
}
