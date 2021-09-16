package service

import (
	"pool_backend/src/model"
	"pool_backend/src/model/request"
	"pool_backend/src/model/response"
	"pool_backend/src/util"
	"time"
)

var filPoolDailyIncomeModel model.FilPoolDailyIncome

//FilPoolDailyIncomeService service
type FilPoolDailyIncomeService struct {
}

//Create 创建矿池数据
func (filPoolDailyIncomeService *FilPoolDailyIncomeService) Create(reqData *request.CreateFilPoolDailyIncome) *response.Resp {
	createFilPoolDailyIncomeParam := &model.FilPoolDailyIncome{
		ID:          util.GenerateID(1).String(),
		FilPoolID:   reqData.FilPoolID,
		AssignedVal: reqData.AssignedVal,
		LastTimeVal: reqData.LastTimeVal,
		TodayVal:    reqData.TodayVal,
		Freed:       reqData.Freed,
		CreateAt:    time.Now(),
	}
	err := filPoolDailyIncomeModel.Create(createFilPoolDailyIncomeParam)
	if err != nil {
		return &response.Resp{Code: "sys_err", Msg: "系统出错!", Data: nil}
	}
	return &response.Resp{Code: "ok", Msg: "请求成功!", Data: createFilPoolDailyIncomeParam.ID}

}

//GetList 获取矿池分页数据
func (filPoolDailyIncomeService *FilPoolDailyIncomeService) GetList(reqData *request.GetList) *response.Resp {

	res, err := filPoolDailyIncomeModel.GetList(reqData)
	if err != nil {
		return &response.Resp{Code: "sys_err", Msg: "系统出错!", Data: nil}
	}
	return &response.Resp{Code: "ok", Msg: "请求成功!", Data: res}

}

//Detail 获取矿池详情
func (filPoolDailyIncomeService *FilPoolDailyIncomeService) Detail(reqData *request.Detail) *response.Resp {

	res, err := filPoolDailyIncomeModel.GetByID(reqData.ID)
	if err != nil {
		return &response.Resp{Code: "sys_err", Msg: "系统出错!", Data: nil}
	}
	return &response.Resp{Code: "ok", Msg: "请求成功!", Data: res}

}
