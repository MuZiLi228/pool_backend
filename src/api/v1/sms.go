package v1

import (
	"pool_backend/src/api"
	"pool_backend/src/global"
	"pool_backend/src/model/request"
	"pool_backend/src/service"
	"pool_backend/src/util/validator"

	"github.com/gin-gonic/gin"
)

var smsService service.SmsService

//SmsV1 api
type SmsV1 struct {
}

// Send 发送短信验证码
// @Summary 发送短信验证码
// @Description 发送短信验证码
// @Tags common
// @Accept application/json
// @Produce application/json
// @Param param body request.SmsSend false "发送短信验证码 需要的参数"
// @Success 200 {object} response.Resp
// @Router /v1/sms/send [post]
func (smsV1 *SmsV1) Send(c *gin.Context) {
	//参数校验
	reqParam := new(request.SmsSend)
	data, err := validator.ParseRequest(c, reqParam)
	if err != nil {
		return
	}
	reqData, ok := data.(*request.SmsSend)
	if !ok {
		global.Logger.Error("发送短信验证码 请求参数 类型错误:", data)
		api.ResponseHTTPOK("req_err", "错误请求!", "", c)
		return
	}

	//发短信验证码
	res := smsService.Send(reqData)

	//返回响应
	api.ResponseHTTPOK(res.Code, res.Msg, res.Data, c)
	return

}

