package v1

import (
	"pool_backend/src/api"
	"pool_backend/src/global"
	"pool_backend/src/model/request"
	"pool_backend/src/service"
	"pool_backend/src/util/validator"

	"github.com/gin-gonic/gin"
)

var filPoolRatioService service.FilPoolRatioService

//FilPoolRatioV1 api
type FilPoolRatioV1 struct {
}

// Create 批量创建股东矿池比例
// @Summary 批量创建股东矿池比例
// @Description 批量创建股东矿池比例
// @Tags 股东矿池比例
// @Accept application/json
// @Produce application/json
// @Param param body request.CreateFilPoolRatio false "创建股东矿池比例 需要的参数"
// @Success 200 {object} response.Resp
// @Router /v1/fil_pool_ratio/create [post]
func (filPoolRatioV1 *FilPoolRatioV1) Create(c *gin.Context) {

	//参数校验
	reqParam := new(request.CreateFilPoolRatio)
	data, err := validator.ParseRequest(c, reqParam)
	if err != nil {
		return
	}
	reqData, ok := data.(*request.CreateFilPoolRatio)
	if !ok {
		global.Logger.Error("创建股东矿池比例 请求参数 类型错误:", data)
		api.ResponseHTTPOK("req_err", "错误请求!", "", c)
		return
	}

	//创建矿池数据
	res := filPoolRatioService.Create(reqData)

	//返回响应
	api.ResponseHTTPOK(res.Code, res.Msg, res.Data, c)
	return

}

// GetList 获取股东矿池比例数据列表
// @Summary 获取股东矿池比例数据列表
// @Description 获取股东矿池比例数据列表
// @Tags 股东矿池比例
// @Accept application/json
// @Produce application/json
// @Param param query request.GetList false "获取股东矿池比例数据列表 需要的参数"
// @Success 200 {object} response.Resp
// @Router /v1/fil_pool_ratio/list [get]
func (filPoolRatioV1 *FilPoolRatioV1) GetList(c *gin.Context) {

	//参数校验
	reqParam := new(request.GetList)
	data, err := validator.ParseRequest(c, reqParam)
	if err != nil {
		return
	}
	reqData, ok := data.(*request.GetList)
	if !ok {
		global.Logger.Error("获取股东矿池比例数据列表 请求参数 类型错误:", data)
		api.ResponseHTTPOK("req_err", "错误请求!", "", c)
		return
	}

	//获取矿池数据
	res := filPoolRatioService.GetList(reqData)

	//返回响应
	api.ResponseHTTPOK(res.Code, res.Msg, res.Data, c)
	return

}

// Detail 获取股东矿池列表
// @Summary 根据矿池id获取股东矿池列表
// @Description 根据矿池id获取股东矿池列表
// @Tags 股东矿池比例
// @Accept application/json
// @Produce application/json
// @Param param query request.Detail false "根据矿池id获取股东矿池列表 需要的参数"
// @Success 200 {object} response.Resp
// @Router /v1/fil_pool_ratio/detail [get]
func (filPoolRatioV1 *FilPoolRatioV1) Detail(c *gin.Context) {

	//参数校验
	reqParam := new(request.Detail)
	data, err := validator.ParseRequest(c, reqParam)
	if err != nil {
		return
	}
	reqData, ok := data.(*request.Detail)
	if !ok {
		global.Logger.Error("根据矿池id获取股东矿池列表 请求参数 类型错误:", data)
		api.ResponseHTTPOK("req_err", "错误请求!", "", c)
		return
	}

	//获取矿池数据
	res := filPoolRatioService.Detail(reqData)

	//返回响应
	api.ResponseHTTPOK(res.Code, res.Msg, res.Data, c)
	return

}
