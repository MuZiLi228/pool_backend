package routers

import (
	"pool_backend/configs"
	api "pool_backend/src/api"
	v1 "pool_backend/src/api/v1"
	"pool_backend/src/middleware"
	"pool_backend/src/util/metrics"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	r                        = gin.Default()
	filPoolV1                v1.FilPoolV1
	shareholderV1            v1.ShareholderV1
	smsV1                    v1.SmsV1
	withdrawalV1             v1.WithdrawalV1
	filPoolRatioV1           v1.FilPoolRatioV1
	newsV1                   v1.NewsV1
	filPoolDailyIncomeV1     v1.FilPoolDailyIncomeV1
	appV1                    v1.AppV1
	shareholderDailyIncomeV1 v1.ShareholderDailyIncomeV1
)

//InitRouter 初始化路由
func InitRouter() *gin.Engine {

	//prometheus
	p := metrics.NewPrometheus("gin")
	p.Use(r)
	r.Use(middleware.CORSMiddleware())

	apis := r.Group("")
	{
		//common
		apis.GET("/health", api.Health)
		apis.GET("/get_id", api.GetID)

		v1 := apis.Group("/v1")
		{
			//news
			v1.GET("/news/web_exchange", newsV1.GetWebExchange)
			v1.GET("/news/daily_news", newsV1.GetDailyNews)
			v1.GET("/news/futures_stat", newsV1.GetFuturesStat)
			v1.GET("/news/futuresmarket_bitcoin", newsV1.GetFuturesmarketBitcoin)
			v1.GET("/news/futuresmarket_exchange", newsV1.GetFuturesmarketExchange)

			//fil_pool_daily_income
			v1.POST("/fil_pool_daily/create", filPoolDailyIncomeV1.Create)
			v1.GET("/fil_pool_daily/detail", filPoolDailyIncomeV1.Detail)
			v1.GET("/fil_pool_daily/list", filPoolDailyIncomeV1.GetList)

			//fil_pool
			v1.POST("/fil_pool/create", filPoolV1.Create)
			v1.GET("/fil_pool/detail", filPoolV1.Detail)
			v1.GET("/fil_pool/list", filPoolV1.GetList)
			v1.POST("/fil_pool/update", filPoolV1.Update)

			//fil_pool_ratio
			v1.POST("/fil_pool_ratio/create", filPoolRatioV1.Create)
			v1.GET("/fil_pool_ratio/detail", filPoolRatioV1.Detail)
			v1.GET("/fil_pool_ratio/list", filPoolRatioV1.GetList)

			//shareholder
			v1.GET("/shareholder/detail", shareholderV1.Detail)
			v1.POST("/account/register", shareholderV1.Register)
			v1.POST("/account/pwd_login", shareholderV1.PwdLogin)
			v1.POST("/shareholder/update_percent", shareholderV1.UpdatePercentShareholderID)
			v1.GET("/account/list", shareholderV1.GetList)
			v1.GET("/shareholder/list", shareholderV1.List)
			v1.POST("/shareholder/set_pwd", shareholderV1.SetWithdrawalPwd)
			v1.POST("/shareholder/set_enable", shareholderV1.SetEnable)
			v1.GET("/shareholder/subordinate_list", shareholderV1.SubordinateList)

			//shareholder_daily_income
			v1.GET("/shareholder/income_list", shareholderDailyIncomeV1.DailyIncomeList)
			v1.GET("/shareholder/info", shareholderDailyIncomeV1.Info)

			//sms
			v1.POST("/sms/send", smsV1.Send)

			//withdrawal
			v1.POST("/withdrawal/application", withdrawalV1.Application)
			v1.POST("/withdrawal/pass", withdrawalV1.Pass)
			v1.POST("/withdrawal/reject", withdrawalV1.Reject)
			v1.GET("/withdrawal/detail", withdrawalV1.Detail)
			v1.GET("/withdrawal/list", withdrawalV1.GetList)
			v1.GET("/withdrawal/account_list", withdrawalV1.ListOfShareholder)

			//app
			v1.POST("/app/uploads", appV1.PutApp)
			v1.GET("/app", appV1.GetAppVersion)
			v1.GET("/app/down/:type/:name", appV1.AppDownload)

		}
	}

	//pprof
	pprof.Register(r, "/sys/pprof")

	//swagger接口文档
	url := ginSwagger.URL(configs.SwaggerURL())
	r.GET("/sys/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
