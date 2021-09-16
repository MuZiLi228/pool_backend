package v1

import (
	"pool_backend/src/api"
	"pool_backend/src/global"
	"pool_backend/src/model/request"
	"pool_backend/src/service"
	"pool_backend/src/util/validator"

	"github.com/gin-gonic/gin"
)

var filPoolDailyIncomeService service.FilPoolDailyIncomeService

//FilPoolDailyIncomeV1 api
type FilPoolDailyIncomeV1 struct {
}

// Create 创建每日矿池收益
// @Summary 创建每日矿池收益
// @Description 创建每日矿池收益
// @Tags 矿池
// @Accept application/json
// @Produce application/json
// @Param param body request.CreateFilPoolDailyIncome false "创建每日矿池收益 需要的参数"
// @Success 200 {object} response.Resp
// @Router /v1/fil_pool_daily/create [post]
func (filPoolDailyIncomeV1 *FilPoolDailyIncomeV1) Create(c *gin.Context) {

	//参数校验
	reqParam := new(request.CreateFilPoolDailyIncome)
	data, err := validator.ParseRequest(c, reqParam)
	if err != nil {
		return
	}
	reqData, ok := data.(*request.CreateFilPoolDailyIncome)
	if !ok {
		global.Logger.Error("创建矿池 请求参数 类型错误:", data)
		api.ResponseHTTPOK("req_err", "错误请求!", "", c)
		return
	}

	//创建每日矿池收益数据
	res := filPoolDailyIncomeService.Create(reqData)

	//返回响应
	api.ResponseHTTPOK(res.Code, res.Msg, res.Data, c)
	return

}

// GetList 获取每日矿池收益数据列表
// @Summary 获取每日矿池收益数据列表
// @Description 获取每日矿池收益数据列表
// @Tags 矿池
// @Accept application/json
// @Produce application/json
// @Param param query request.GetList false "获取每日矿池收益数据列表 需要的参数"
// @Success 200 {object} response.Resp
// @Router /v1/fil_pool_daily/list [get]
func (filPoolDailyIncomeV1 *FilPoolDailyIncomeV1) GetList(c *gin.Context) {

	//参数校验
	reqParam := new(request.GetList)
	data, err := validator.ParseRequest(c, reqParam)
	if err != nil {
		return
	}
	reqData, ok := data.(*request.GetList)
	if !ok {
		global.Logger.Error("获取每日矿池收益数据列表 请求参数 类型错误:", data)
		api.ResponseHTTPOK("req_err", "错误请求!", "", c)
		return
	}

	//获取每日矿池收益数据列表
	res := filPoolDailyIncomeService.GetList(reqData)

	//返回响应
	api.ResponseHTTPOK(res.Code, res.Msg, res.Data, c)
	return

}

// Detail 获取每日矿池收益数据详情
// @Summary 获取每日矿池收益数据详情
// @Description 获取每日矿池收益数据详情
// @Tags 矿池
// @Accept application/json
// @Produce application/json
// @Param param query request.Detail false "获取每日矿池收益数据详情 需要的参数"
// @Success 200 {object} response.Resp
// @Router /v1/fil_pool_daily/detail [get]
func (filPoolDailyIncomeV1 *FilPoolDailyIncomeV1) Detail(c *gin.Context) {

	//参数校验
	reqParam := new(request.Detail)
	data, err := validator.ParseRequest(c, reqParam)
	if err != nil {
		return
	}
	reqData, ok := data.(*request.Detail)
	if !ok {
		global.Logger.Error("获取每日矿池收益数据详情 请求参数 类型错误:", data)
		api.ResponseHTTPOK("req_err", "错误请求!", "", c)
		return
	}

	//获取每日矿池收益数据详情
	res := filPoolDailyIncomeService.Detail(reqData)

	//返回响应
	api.ResponseHTTPOK(res.Code, res.Msg, res.Data, c)
	return

}
