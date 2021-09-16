package delayer

import (
	"errors"

	"github.com/go-redis/redis/v8"
)

// initSettings 创建一个初始化配置
func initSettings() *ClientSetting {
	return &ClientSetting{}
}

type ClientSetting struct {
	Rdc       *redis.Client
	RdcPrefix string
}

func (cs ClientSetting) Validate() error {
	if rdc == nil {
		return errors.New("请调用 delayer.WithRedis 方法")
	}

	if snowflakeNode == nil {
		return errors.New("请调用 delayer.WithSnowflakeNode 方法")
	}

	return nil
}
