package cron

import (
	"pool_backend/src/global"

	cron2 "github.com/robfig/cron/v3"
)

var C *cron2.Cron

type Logger struct {
}

//Error
func (Logger) Error(err error, msg string, keysAndValues ...interface{}) {
	global.Logger.Error("定时器报错:", err, "响应数据:", msg, "keysAndValues:", keysAndValues)

}

//Info
func (Logger) Info(msg string, keysAndValues ...interface{}) {
	global.Logger.Debug("定时器响应数据:", msg, "keysAndValues:", keysAndValues)
}

func init() {
	C = cron2.New(
		cron2.WithSeconds(),
		// cron2.WithLogger(Logger{}),  //日志记录
		// cron2.WithChain(cron2.SkipIfStillRunning(Logger{})),
	)
}
