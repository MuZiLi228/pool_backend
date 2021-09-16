package v1

import (
	"pool_backend/src/api"
	"pool_backend/src/global"
	"pool_backend/src/model/request"
	"pool_backend/src/service"
	"pool_backend/src/util/validator"

	"github.com/gin-gonic/gin"
)

var shareholderService service.ShareholderService

// ShareholderV1 api
type ShareholderV1 struct {
}

// PwdLogin 账号密码登录
// @Summary 账号密码登录
// @Description 账号密码登录
// @Tags 股东
// @Accept application/json
// @Produce application/json
// @Param param body request.PwdLogin false "账号密码登录 需要的参数"
// @Success 200 {object} response.Resp
// @Router /v1/account/pwd_login [post]
func (shareholderV1 *ShareholderV1) PwdLogin(c *gin.Context) {

	//参数校验
	reqParam := new(request.PwdLogin)
	data, err := validator.ParseRequest(c, reqParam)
	if err != nil {
		return
	}
	reqData, ok := data.(*request.PwdLogin)
	if !ok {
		global.Logger.Error("创建账号请求参数 类型错误:", data)
		api.ResponseHTTPOK("req_err", "错误请求!", "", c)
		return
	}

	//账号密码登录
	res := shareholderService.CheckPwd(reqData)

	//返回响应
	api.ResponseHTTPOK(res.Code, res.Msg, res.Data, c)
	return

}

// GetList 获取股东数据列表分页
// @Summary 获取股东数据列表分页
// @Description 获取股东数据列表分页 后台用户管理
// @Tags 股东
// @Accept application/json
// @Produce application/json
// @Param param query request.GetList false "获取股东数据列表 需要的参数"
// @Success 200 {object} response.Resp
// @Router /v1/account/list [get]
func (shareholderV1 *ShareholderV1) GetList(c *gin.Context) {

	//参数校验
	reqParam := new(request.GetList)
	data, err := validator.ParseRequest(c, reqParam)
	if err != nil {
		return
	}
	reqData, ok := data.(*request.GetList)
	if !ok {
		global.Logger.Error("获取股东列表分页数据 请求参数 类型错误:", data)
		api.ResponseHTTPOK("req_err", "错误请求!", "", c)
		return
	}

	//获取股东账号数据
	res := shareholderService.GetList(reqData)

	//返回响应
	api.ResponseHTTPOK(res.Code, res.Msg, res.Data, c)
	return

}

// List 获取股东数据列表
// @Summary 获取股东数据列表
// @Description 矿池分配股东时,需要有股东列表去选择
// @Tags 股东
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Resp
// @Router /v1/shareholder/list [get]
func (shareholderV1 *ShareholderV1) List(c *gin.Context) {

	//获取股东账号数据
	res := shareholderService.List()

	//返回响应
	api.ResponseHTTPOK(res.Code, res.Msg, res.Data, c)
	return

}

// Register 股东账号注册
// @Summary 股东账号注册
// @Description 股东账号注册
// @Tags 股东
// @Accept application/json
// @Produce application/json
// @Param param body request.RegisterAccount false "股东手机验证码登录 需要的参数"
// @Success 200 {object} response.Resp
// @Router /v1/account/register [post]
func (shareholderV1 *ShareholderV1) Register(c *gin.Context) {

	//参数校验
	reqParam := new(request.RegisterAccount)
	data, err := validator.ParseRequest(c, reqParam)
	if err != nil {
		return
	}
	reqData, ok := data.(*request.RegisterAccount)
	if !ok {
		global.Logger.Error("获取股东列表数据 请求参数 类型错误:", data)
		api.ResponseHTTPOK("req_err", "错误请求!", "", c)
		return
	}

	//注册股东账号
	res := shareholderService.Register(reqData)

	//返回响应
	api.ResponseHTTPOK(res.Code, res.Msg, res.Data, c)
	return

}

// SubordinateList 获取股东下级列表
// @Summary 获取股东下级列表
// @Description 获取股东下级列表
// @Tags 股东
// @Accept application/json
// @Produce application/json
// @Param param query request.Detail false "获取股东下级列表 需要的参数"
// @Success 200 {object} response.Resp
// @Router /v1/shareholder/subordinate_list [get]
func (shareholderV1 *ShareholderV1) SubordinateList(c *gin.Context) {
	//参数校验
	reqParam := new(request.Detail)
	data, err := validator.ParseRequest(c, reqParam)
	if err != nil {
		return
	}
	reqData, ok := data.(*request.Detail)
	if !ok {
		global.Logger.Error("获取矿池数据详情 请求参数 类型错误:", data)
		api.ResponseHTTPOK("req_err", "错误请求!", "", c)
		return
	}
	//获取股东下级列表
	res := shareholderService.SubordinateList(reqData.ID)

	//返回响应
	api.ResponseHTTPOK(res.Code, res.Msg, res.Data, c)
	return

}

// UpdatePercentShareholderID 股东更新再分配占比例
// @Summary 股东更新再分配占比例
// @Description 股东更新再分配占比例
// @Tags 股东
// @Accept application/json
// @Produce application/json
// @Param param body request.UpdatePercentShareholderID false "股东更新再分配占比例 需要的参数"
// @Success 200 {object} response.Resp
// @Router /v1/shareholder/update_percent [post]
func (shareholderV1 *ShareholderV1) UpdatePercentShareholderID(c *gin.Context) {
	//参数校验
	reqParam := new(request.UpdatePercentShareholderID)
	data, err := validator.ParseRequest(c, reqParam)
	if err != nil {
		return
	}
	reqData, ok := data.(*request.UpdatePercentShareholderID)
	if !ok {
		global.Logger.Error("获取矿池数据详情 请求参数 类型错误:", data)
		api.ResponseHTTPOK("req_err", "错误请求!", "", c)
		return
	}
	//获取股东下级列表
	res := shareholderService.UpdatePercentShareholderID(reqData)

	//返回响应
	api.ResponseHTTPOK(res.Code, res.Msg, res.Data, c)
	return

}

// Detail 获取股东信息
// @Summary 获取股东信息
// @Description 获取股东信息
// @Tags 股东
// @Accept application/json
// @Produce application/json
// @Param param query request.Detail false "获取股东信息 需要的参数"
// @Success 200 {object} response.Resp
// @Router /v1/shareholder/detail [get]
func (shareholderV1 *ShareholderV1) Detail(c *gin.Context) {

	//参数校验
	reqParam := new(request.Detail)
	data, err := validator.ParseRequest(c, reqParam)
	if err != nil {
		return
	}
	reqData, ok := data.(*request.Detail)
	if !ok {
		global.Logger.Error("获取股东信息 请求参数 类型错误:", data)
		api.ResponseHTTPOK("req_err", "错误请求!", "", c)
		return
	}

	//获取矿池数据
	res := shareholderService.Detail(reqData)

	//返回响应
	api.ResponseHTTPOK(res.Code, res.Msg, res.Data, c)
	return

}



// SetWithdrawalPwd 股东设置提现密码
// @Summary 股东设置提现密码
// @Description 股东设置提现密码 需要提供登录密码 后面可迭代加上短信验证码
// @Tags 股东
// @Accept application/json
// @Produce application/json
// @Param param body request.SetWithdrawalPwd false "股东设置提现密码 需要的参数"
// @Success 200 {object} response.Resp
// @Router /v1/shareholder/set_pwd [post]
func (shareholderV1 *ShareholderV1) SetWithdrawalPwd(c *gin.Context) {

	//参数校验
	reqParam := new(request.SetWithdrawalPwd)
	data, err := validator.ParseRequest(c, reqParam)
	if err != nil {
		return
	}
	reqData, ok := data.(*request.SetWithdrawalPwd)
	if !ok {
		global.Logger.Error("股东设置提现密码 请求参数 类型错误:", data)
		api.ResponseHTTPOK("req_err", "错误请求!", "", c)
		return
	}

	//股东设置提现密码
	res := shareholderService.SetWithdrawalPwd(reqData)

	//返回响应
	api.ResponseHTTPOK(res.Code, res.Msg, res.Data, c)
	return

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
func (shareholderV1 *ShareholderV1) DailyIncomeList(c *gin.Context) {

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



// SetEnable 禁用用户
// @Summary 禁用用户
// @Description 禁用用户 后台用户管理
// @Tags 股东
// @Accept application/json
// @Produce application/json
// @Param param body request.SetEnable false "禁用用户 需要的参数"
// @Success 200 {object} response.Resp
// @Router /v1/shareholder/set_enable [post]
func (shareholderV1 *ShareholderV1) SetEnable(c *gin.Context) {

	//参数校验
	reqParam := new(request.SetEnable)
	data, err := validator.ParseRequest(c, reqParam)
	if err != nil {
		return
	}
	reqData, ok := data.(*request.SetEnable)
	if !ok {
		global.Logger.Error("禁用用户 请求参数 类型错误:", data)
		api.ResponseHTTPOK("req_err", "错误请求!", "", c)
		return
	}

	if reqData.ID == reqData.ShareholderID {
		api.ResponseHTTPOK("req_err", "操作错误!", "", c)
		return
	}

	//禁用用户
	res := shareholderService.SetEnable(reqData)

	//返回响应
	api.ResponseHTTPOK(res.Code, res.Msg, res.Data, c)
	return

}
