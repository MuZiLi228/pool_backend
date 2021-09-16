package v1

import (
	"context"
	"encoding/json"
	"pool_backend/src/api"
	"pool_backend/src/global"

	"github.com/gin-gonic/gin"
)

//NewsV1 api
type NewsV1 struct {
}

// GetWebExchange 获取交易所排行数据
// @Summary 获取交易所排行数据
// @Description 获取交易所排行数据
// @Tags 资讯模块
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Resp
// @Router /v1/news/web_exchange [get]
func (NewsV1 *NewsV1) GetWebExchange(c *gin.Context) {

	key := "feixiaohao:web-exchange:list"
	val, err := global.Cache.Get(context.Background(), key).Result()
	if err != nil {
		global.Logger.Error("api NewsV1 GetWebExchange 获取redis数据 报错:", err.Error())
		api.ResponseHTTPOK("sys_err", "系统出错!", nil, c)
		return
	}
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		global.Logger.Error("api NewsV1 GetWebExchange 解析数据 报错:", err.Error())
		api.ResponseHTTPOK("sys_err", "系统出错!", nil, c)
		return
	}
	api.ResponseHTTPOK("ok", "请求成功!", data, c)
	return

}



// GetDailyNews 24小说交易排行数据
// @Summary 24小说交易排行数据
// @Description 24小说交易排行数据
// @Tags 资讯模块
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Resp
// @Router /v1/news/daily_news [get]
func (NewsV1 *NewsV1) GetDailyNews(c *gin.Context) {

	key := "feixiaohao:daily-news:list"
	val, err := global.Cache.Get(context.Background(), key).Result()
	if err != nil {
		global.Logger.Error("api NewsV1 GetDailyNews 获取redis数据 报错:", err.Error())
		api.ResponseHTTPOK("sys_err", "系统出错!", nil, c)
		return
	}
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		global.Logger.Error("api NewsV1 GetDailyNews 解析数据 报错:", err.Error())
		api.ResponseHTTPOK("sys_err", "系统出错!", nil, c)
		return
	}
	api.ResponseHTTPOK("ok", "请求成功!", data, c)
	return
}




// GetFuturesStat 24小时爆仓统计
// @Summary 24小时爆仓统计
// @Description 24小时爆仓统计
// @Tags 资讯模块
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Resp
// @Router /v1/news/futures_stat [get]
func (NewsV1 *NewsV1) GetFuturesStat(c *gin.Context) {

	key := "feixiaohao:futures:stat"
	val, err := global.Cache.Get(context.Background(), key).Result()
	if err != nil {
		global.Logger.Error("api NewsV1 GetFuturesStat 获取redis数据 报错:", err.Error())
		api.ResponseHTTPOK("sys_err", "系统出错!", nil, c)
		return
	}
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		global.Logger.Error("api NewsV1 GetFuturesStat 解析数据 报错:", err.Error())
		api.ResponseHTTPOK("sys_err", "系统出错!", nil, c)
		return
	}
	api.ResponseHTTPOK("ok", "请求成功!", data, c)
	return

}


// GetFuturesmarketBitcoin 24小时多空比
// @Summary 24小时多空比
// @Description 24小时多空比
// @Tags 资讯模块
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Resp
// @Router /v1/news/futuresmarket_bitcoin [get]
func (NewsV1 *NewsV1) GetFuturesmarketBitcoin(c *gin.Context) {

	key := "feixiaohao:futuresmarket:bitcoin"
	val, err := global.Cache.Get(context.Background(), key).Result()
	if err != nil {
		global.Logger.Error("api NewsV1 GetFuturesmarketBitcoin 获取redis数据 报错:", err.Error())
		api.ResponseHTTPOK("sys_err", "系统出错!", nil, c)
		return
	}
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		global.Logger.Error("api NewsV1 GetFuturesmarketBitcoin 解析数据 报错:", err.Error())
		api.ResponseHTTPOK("sys_err", "系统出错!", nil, c)
		return
	}
	api.ResponseHTTPOK("ok", "请求成功!", data, c)
	return

}

// GetFuturesmarketExchange 交易所期货数据
// @Summary 交易所期货数据
// @Description 交易所期货数据
// @Tags 资讯模块
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Resp
// @Router /v1/news/futuresmarket_exchange [get]
func (NewsV1 *NewsV1) GetFuturesmarketExchange(c *gin.Context) {

	key := "feixiaohao:futuresmarket:exchange"
	val, err := global.Cache.Get(context.Background(), key).Result()
	if err != nil {
		global.Logger.Error("api NewsV1 GetFuturesmarketExchange 获取redis数据 报错:", err.Error())
		api.ResponseHTTPOK("sys_err", "系统出错!", nil, c)
		return
	}
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		global.Logger.Error("api NewsV1 GetFuturesmarketExchange 解析数据 报错:", err.Error())
		api.ResponseHTTPOK("sys_err", "系统出错!", nil, c)
		return
	}
	api.ResponseHTTPOK("ok", "请求成功!", data, c)
	return

}


