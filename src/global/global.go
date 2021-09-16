package global

import (
	"pool_backend/src/util/db"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	//DB 数据库
	DB db.Repo
	//Cache 缓存
	Cache *redis.Client
	//Viper 配置
	Viper *viper.Viper
	//Logger 日志
	Logger *zap.SugaredLogger
)
