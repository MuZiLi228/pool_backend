package v1

import (
	"pool_backend/src/api"
	"pool_backend/src/global"
	"pool_backend/src/model/request"
	"pool_backend/src/service"
	"pool_backend/src/util/validator"

	"github.com/gin-gonic/gin"
)

var filPoolService service.FilPoolService

//FilPoolV1 api
type FilPoolV1 struct {
}

// Create 创建矿池
// @Summary 创建矿池
// @Description 创建矿池
// @Tags 矿池
// @Accept application/json
// @Produce application/json
// @Param param body request.CreateFilPool false "创建矿池 需要的参数"
// @Success 200 {object} response.Resp
// @Router /v1/fil_pool/create [post]
func (filPoolV1 *FilPoolV1) Create(c *gin.Context) {

	//参数校验
	reqParam := new(request.CreateFilPool)
	data, err := validator.ParseRequest(c, reqParam)
	if err != nil {
		return
	}
	reqData, ok := data.(*request.CreateFilPool)
	if !ok {
		global.Logger.Error("创建矿池 请求参数 类型错误:", data)
		api.ResponseHTTPOK("req_err", "错误请求!", "", c)
		return
	}

	//创建矿池数据
	res := filPoolService.Create(reqData)

	//返回响应
	api.ResponseHTTPOK(res.Code, res.Msg, res.Data, c)
	return

}

// GetList 获取矿池数据列表
// @Summary 获取矿池数据列表
// @Description 获取矿池数据列表
// @Tags 矿池
// @Accept application/json
// @Produce application/json
// @Param param query request.GetList false "获取矿池数据列表 需要的参数"
// @Success 200 {object} response.Resp
// @Router /v1/fil_pool/list [get]
func (filPoolV1 *FilPoolV1) GetList(c *gin.Context) {

	//参数校验
	reqParam := new(request.GetList)
	data, err := validator.ParseRequest(c, reqParam)
	if err != nil {
		return
	}
	reqData, ok := data.(*request.GetList)
	if !ok {
		global.Logger.Error("获取矿池数据 请求参数 类型错误:", data)
		api.ResponseHTTPOK("req_err", "错误请求!", "", c)
		return
	}

	//获取矿池数据
	res := filPoolService.GetList(reqData)

	//返回响应
	api.ResponseHTTPOK(res.Code, res.Msg, res.Data, c)
	return

}

// Detail 获取矿池数据详情
// @Summary 获取矿池数据详情
// @Description 获取矿池数据详情
// @Tags 矿池
// @Accept application/json
// @Produce application/json
// @Param param query request.Detail false "获取矿池数据详情 需要的参数"
// @Success 200 {object} response.Resp
// @Router /v1/fil_pool/detail [get]
func (filPoolV1 *FilPoolV1) Detail(c *gin.Context) {

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

	//获取矿池数据
	res := filPoolService.Detail(reqData)

	//返回响应
	api.ResponseHTTPOK(res.Code, res.Msg, res.Data, c)
	return

}

// Update 更新矿池
// @Summary 更新矿池
// @Description 更新矿池
// @Tags 矿池
// @Accept application/json
// @Produce application/json
// @Param param body request.UpadteFilPool false "更新矿池 需要的参数"
// @Success 200 {object} response.Resp
// @Router /v1/fil_pool/update [post]
func (filPoolV1 *FilPoolV1) Update(c *gin.Context) {

	//参数校验
	reqParam := new(request.UpadteFilPool)
	data, err := validator.ParseRequest(c, reqParam)
	if err != nil {
		return
	}
	reqData, ok := data.(*request.UpadteFilPool)
	if !ok {
		global.Logger.Error("更新矿池 请求参数 类型错误:", data)
		api.ResponseHTTPOK("req_err", "错误请求!", "", c)
		return
	}

	//更新矿池数据
	res := filPoolService.Update(reqData)

	//返回响应
	api.ResponseHTTPOK(res.Code, res.Msg, res.Data, c)
	return

}
