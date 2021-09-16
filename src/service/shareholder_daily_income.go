package service

import (
	"pool_backend/src/model"
	"pool_backend/src/model/request"
	"pool_backend/src/model/response"
	"pool_backend/src/util"
)

var shareholderDailyIncomeModel model.ShareholderDailyIncome

//ShareholderDailyIncomeService service
type ShareholderDailyIncomeService struct {
}

//Detail 获取股东每日收益详情
func (shareholderDailyIncomeService *ShareholderDailyIncomeService) Detail(reqData *request.Detail) *response.Resp {

	res, err := shareholderDailyIncomeModel.GetByID(reqData.ID)
	if err != nil {
		return &response.Resp{Code: "sys_err", Msg: "获取股东每日收益详情 失败!", Data: nil}
	}
	return &response.Resp{Code: "ok", Msg: "请求成功!", Data: res}

}

//Info 获取挖矿天数和分配总量
func (shareholderDailyIncomeService *ShareholderDailyIncomeService) Info(reqData *request.Detail) *response.Resp {

	//余额 = 总收益 - 提现通过金额
	//获取已提现金额
	withdrawalSum, err := withdrawalModel.WithdrawalSumByShareholder(reqData.ID)
	if err != nil {
		return &response.Resp{Code: "sys_err", Msg: "系统出错!", Data: nil}
	}
	//获取总收益
	incomeInfo, err := shareholderDailyIncomeModel.Info(reqData.ID)
	balance := incomeInfo.Income - withdrawalSum
	res := &model.IncomeInfo{
		Income:        incomeInfo.Income,
		Balance:       util.Decimal(balance),
		WithdrawalSum: withdrawalSum,
		Day:           incomeInfo.Day,
	}
	return &response.Resp{Code: "ok", Msg: "请求成功!", Data: res}

}
