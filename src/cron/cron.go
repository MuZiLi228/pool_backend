package cron

import (
	"pool_backend/src/service"
	"pool_backend/src/util/cron"
)

var cronService service.CronService

//Register 注册定时任务
func Register() {
	// Every day on midnight  ""   @hourly
	_, _ = cron.C.AddFunc("@hourly", func() {
		//定时器任务
		cronService.GetWebExchange()
		cronService.GetDailyNews()
		cronService.GetFuturesStat()
		cronService.GetFuturesmarketBitcoin()
		cronService.GetFuturesmarketExchange()
	})

	// Every day on midnight "0 */1 * * * ?"  "@daily"
	_, _ = cron.C.AddFunc("@every 3m", func() {
		//定时器任务
		cronService.UpdatePoolRatio()
		cronService.DistributePoolIncome()
	})
}

//Start 启动定时任务
func Start() {
	cron.C.Start()
}

//Stop 停止定时任务
func Stop() {
	cron.C.Stop()
}
