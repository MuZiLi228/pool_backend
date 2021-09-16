package bootstrap

import (
	"pool_backend/src/cron"
	"pool_backend/src/global"
	"pool_backend/src/util/cache"
	"pool_backend/src/util/db"
	"pool_backend/src/util/logger"
	"pool_backend/src/util/validator"
)

//Init 系统初始化
func Init() {

	// 初始化 logger
	loggers := logger.Log()
	global.Logger = loggers
	defer loggers.Sync()

	// 初始化数据库  全部暂时用一个数据库实例
	dbRepo, err := db.New()
	if err != nil {
		loggers.Errorf("new db fail, err:%v", err)
	}
	global.DB = dbRepo

	//	初始化缓存服务
	cacheRepo, err := cache.New()
	if err != nil {
		loggers.Errorf("new redis fail, err:%v", err)
	}
	global.Cache = cacheRepo

	//校验器
	validator.InitVali()

	//定时器
	cron.Register()
	cron.Start()


}
