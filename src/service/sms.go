package service

import (
	"context"
	"fmt"
	"pool_backend/configs"
	"pool_backend/src/global"
	"pool_backend/src/model/request"
	"pool_backend/src/model/response"
	"pool_backend/src/util"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/go-redis/redis/v8"
)

//SmsService service
type SmsService struct {
}

//Send 发送短信
func (smsService *SmsService) Send(reqData *request.SmsSend) *response.Resp {

	//配置信息
	regionID := configs.Get().Sms.RegionID
	accessKeyID := configs.Get().Sms.AccessKeyID
	accessKeySecret := configs.Get().Sms.AccessKeySecret
	SignName := configs.Get().Sms.SignName
	TemplateCode := configs.Get().Sms.TemplateCode

	//生成6位短信验证码
	code := util.GenRandomNumber()

	//实例化
	client, err := dysmsapi.NewClientWithAccessKey(regionID, accessKeyID, accessKeySecret)
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = reqData.Mobile
	request.SignName = SignName
	request.TemplateCode = TemplateCode
	request.TemplateParam = fmt.Sprintf("{\"code\":\"%s\"}", code)

	//调用发送短信
	smsRes, err := client.SendSms(request)
	if err != nil {
		global.Logger.Error("发送短信出错:", err.Error)
		return &response.Resp{Code: "fail", Msg: "发送短信失败!", Data: nil}
	}
	if smsRes.Code != "OK" && smsRes.Message != "OK" {
		global.Logger.Error("发送阿里短信 报错:", err.Error)
		return &response.Resp{Code: "fail", Msg: "发送短信失败!", Data: nil}
	}

	//成功后设置redis缓存 5分钟内有效
	redisKey := fmt.Sprintf("phone_sms_code:%s", reqData.Mobile)
	err = global.Cache.Set(context.Background(), redisKey, code, time.Minute*5).Err()
	if err != nil {
		global.Logger.Error("发送阿里短信 缓存redis key 报错:", err.Error)
		return &response.Resp{Code: "sys_err", Msg: "系统出错!", Data: nil}
	}
	return &response.Resp{Code: "ok", Msg: "请求成功!", Data: nil}
}

//CheckSmsCode 校验短信验证码
func CheckSmsCode(smsCode, mobile string) *response.Resp {
	//成功后设置code缓存 5分钟内有效
	redisKey := fmt.Sprintf("phone_sms_code:%s", mobile)

	value, err := global.Cache.Get(context.Background(), redisKey).Result()
	if err == redis.Nil {
		return &response.Resp{Code: "not_exist", Msg: "验证码无效!", Data: nil}
	} else if err != nil {
		global.Logger.Error("获取redis验证码 出错:", err.Error)
		return &response.Resp{Code: "sys_err", Msg: "系统出错!", Data: nil}
	}
	//请求通过后,将redis key删掉
	if value == smsCode {
		err := global.Cache.Del(context.Background(), redisKey).Err()
		if err != nil {
			global.Logger.Error("短信验证码, 将redis key删掉 报错:", err.Error)
		}
	}
	return &response.Resp{Code: "ok", Msg: "验证码通过!", Data: mobile}

}
