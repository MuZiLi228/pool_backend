package v1

import (
	"pool_backend/src/api"
	"pool_backend/src/global"
	"pool_backend/src/model/request"
	"pool_backend/src/service"
	"pool_backend/src/util/validator"

	"github.com/gin-gonic/gin"
)

var shareholderDailyIncomeService service.ShareholderDailyIncomeService

// ShareholderDailyIncomeV1 api
type ShareholderDailyIncomeV1 struct {
}


// DailyIncomeList 获取股东每日收益数据列表分页
// @Summary 获取股东每日收益数据列表分页
// @Description 获取股东每日收益数据列表分页 后台用户管理
// @Tags 股东
// @Accept application/json
// @Produce application/json
// @Param param query request.GetList false "获取股东每日收益数据列表分页 需要的参数"
// @Success 200 {object} response.Resp
// @Router /v1/shareholder/income_list [get]
func (shareholderDailyIncomeV1 *ShareholderDailyIncomeV1) DailyIncomeList(c *gin.Context) {

	//参数校验
	reqParam := new(request.GetList)
	data, err := validator.ParseRequest(c, reqParam)
	if err != nil {
		return
	}
	reqData, ok := data.(*request.GetList)
	if !ok {
		global.Logger.Error("获取股东每日收益数据列表分页 请求参数 类型错误:", data)
		api.ResponseHTTPOK("req_err", "错误请求!", "", c)
		return
	}

	//获取股东每日收益数据
	res := shareholderService.DailyIncomeList(reqData)

	//返回响应
	api.ResponseHTTPOK(res.Code, res.Msg, res.Data, c)
	return

}


// Info 获取收益统计数据
// @Summary 获取收益统计数据
// @Description  获取收益统计数据 app
// @Tags 股东
// @Accept application/json
// @Produce application/json
// @Param param query request.Detail false " 获取收益统计数据 需要的参数"
// @Success 200 {object} response.Resp
// @Router /v1/shareholder/info [get]
func (shareholderDailyIncomeV1 *ShareholderDailyIncomeV1) Info(c *gin.Context) {

	//参数校验
	reqParam := new(request.Detail)
	data, err := validator.ParseRequest(c, reqParam)
	if err != nil {
		return
	}
	reqData, ok := data.(*request.Detail)
	if !ok {
		global.Logger.Error("获取收益统计数据 请求参数 类型错误:", data)
		api.ResponseHTTPOK("req_err", "错误请求!", "", c)
		return
	}

	//获取收益统计数据
	res := shareholderDailyIncomeService.Info(reqData)

	//返回响应
	api.ResponseHTTPOK(res.Code, res.Msg, res.Data, c)
	return

}
