package v1

import (
	"pool_backend/src/api"
	"pool_backend/src/global"
	"pool_backend/src/model/request"
	"pool_backend/src/service"
	"pool_backend/src/util/validator"

	"github.com/gin-gonic/gin"
)

//WithdrawalService service
var WithdrawalService service.WithdrawalService

//WithdrawalV1 api
type WithdrawalV1 struct {
}

// Application 创建提现申请
// @Summary 创建提现申请
// @Description 创建提现申请
// @Tags 提现
// @Accept application/json
// @Produce application/json
// @Param param body request.ApplicationWithdrawal false "创建提现申请 需要的参数"
// @Success 200 {object} response.Resp
// @Router /v1/withdrawal/application [post]
func (withdrawalV1 *WithdrawalV1) Application(c *gin.Context) {

	//参数校验
	reqParam := new(request.ApplicationWithdrawal)
	data, err := validator.ParseRequest(c, reqParam)
	if err != nil {
		return
	}
	reqData, ok := data.(*request.ApplicationWithdrawal)
	if !ok {
		global.Logger.Error("创建提现申请 请求参数 类型错误:", data)
		api.ResponseHTTPOK("req_err", "错误请求!", "", c)
		return
	}

	//创建提现申请
	res := WithdrawalService.Application(reqData)

	//返回响应
	api.ResponseHTTPOK(res.Code, res.Msg, res.Data, c)
	return

}

// GetList 获取提现申请数据列表
// @Summary 获取提现申请数据列表
// @Description 获取提现申请数据列表
// @Tags 提现
// @Accept application/json
// @Produce application/json
// @Param param query request.GetList false "获取提现申请数据列表 需要的参数"
// @Success 200 {object} response.Resp
// @Router /v1/withdrawal/list [get]
func (withdrawalV1 *WithdrawalV1) GetList(c *gin.Context) {

	//参数校验
	reqParam := new(request.GetList)
	data, err := validator.ParseRequest(c, reqParam)
	if err != nil {
		return
	}
	reqData, ok := data.(*request.GetList)
	if !ok {
		global.Logger.Error("获取提现申请数据列表 请求参数 类型错误:", data)
		api.ResponseHTTPOK("req_err", "错误请求!", "", c)
		return
	}

	//获取提现申请数据列表
	res := WithdrawalService.GetList(reqData)

	//返回响应
	api.ResponseHTTPOK(res.Code, res.Msg, res.Data, c)
	return

}

// Detail 获取提现申请详情
// @Summary 获取提现申请详情
// @Description 获取提现申请详情
// @Tags 提现
// @Accept application/json
// @Produce application/json
// @Param param query request.Detail false "获取提现申请详情 需要的参数"
// @Success 200 {object} response.Resp
// @Router /v1/withdrawal/detail [get]
func (withdrawalV1 *WithdrawalV1) Detail(c *gin.Context) {

	//参数校验
	reqParam := new(request.Detail)
	data, err := validator.ParseRequest(c, reqParam)
	if err != nil {
		return
	}
	reqData, ok := data.(*request.Detail)
	if !ok {
		global.Logger.Error("获取提现申请详情 请求参数 类型错误:", data)
		api.ResponseHTTPOK("req_err", "错误请求!", "", c)
		return
	}

	//获取提现申请详情
	res := WithdrawalService.Detail(reqData)

	//返回响应
	api.ResponseHTTPOK(res.Code, res.Msg, res.Data, c)
	return

}

// ListOfShareholder 获取股东提现申请列表
// @Summary 获取股东提现申请列表
// @Description 获取股东提现申请列表
// @Tags 提现
// @Accept application/json
// @Produce application/json
// @Param param query request.Detail false "获取股东提现申请列表 需要的参数"
// @Success 200 {object} response.Resp
// @Router /v1/withdrawal/account_list [get]
func (withdrawalV1 *WithdrawalV1) ListOfShareholder(c *gin.Context) {

	//参数校验
	reqParam := new(request.Detail)
	data, err := validator.ParseRequest(c, reqParam)
	if err != nil {
		return
	}
	reqData, ok := data.(*request.Detail)
	if !ok {
		global.Logger.Error("获取股东提现申请列表 请求参数 类型错误:", data)
		api.ResponseHTTPOK("req_err", "错误请求!", "", c)
		return
	}

	//获取股东提现申请列表
	res := WithdrawalService.ListOfShareholder(reqData)

	//返回响应
	api.ResponseHTTPOK(res.Code, res.Msg, res.Data, c)
	return

}

// Pass 后台通过提现申请
// @Summary 后台通过提现申请
// @Description 后台通过提现申请 1、申请状态改为通过 2、需要扣除对应股东余额
// @Tags 提现
// @Accept application/json
// @Produce application/json
// @Param param body request.Detail false "后台通过提现申请 需要的参数"
// @Success 200 {object} response.Resp
// @Router /v1/withdrawal/pass [post]
func (withdrawalV1 *WithdrawalV1) Pass(c *gin.Context) {

	//参数校验
	reqParam := new(request.Detail)
	data, err := validator.ParseRequest(c, reqParam)
	if err != nil {
		return
	}
	reqData, ok := data.(*request.Detail)
	if !ok {
		global.Logger.Error("后台通过提现申请 请求参数 类型错误:", data)
		api.ResponseHTTPOK("req_err", "错误请求!", "", c)
		return
	}

	//后台通过提现申请
	res := WithdrawalService.Pass(reqData)

	//返回响应
	api.ResponseHTTPOK(res.Code, res.Msg, res.Data, c)
	return

}

// Reject 后台驳回提现申请
// @Summary 后台驳回提现申请
// @Description 后台驳回提现申请 1、申请状态改为驳回 2、需要添加驳回原因
// @Tags 提现
// @Accept application/json
// @Produce application/json
// @Param param body request.WithdrawalReject false "后台驳回提现申请 需要的参数"
// @Success 200 {object} response.Resp
// @Router /v1/withdrawal/reject [post]
func (withdrawalV1 *WithdrawalV1) Reject(c *gin.Context) {

	//参数校验
	reqParam := new(request.WithdrawalReject)
	data, err := validator.ParseRequest(c, reqParam)
	if err != nil {
		return
	}
	reqData, ok := data.(*request.WithdrawalReject)
	if !ok {
		global.Logger.Error("后台驳回提现申请 请求参数 类型错误:", data)
		api.ResponseHTTPOK("req_err", "错误请求!", "", c)
		return
	}

	//后台驳回提现申请
	res := WithdrawalService.Reject(reqData)

	//返回响应
	api.ResponseHTTPOK(res.Code, res.Msg, res.Data, c)
	return

}
