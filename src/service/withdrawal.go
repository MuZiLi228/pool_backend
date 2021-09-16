package service

import (
	"pool_backend/src/enum"
	"pool_backend/src/global"
	"pool_backend/src/model"
	"pool_backend/src/model/request"
	"pool_backend/src/model/response"
	"pool_backend/src/util"
	"time"

	"gorm.io/gorm"
)

//withdrawalModel model
var withdrawalModel model.Withdrawal

var shareholderDailyIncomeService ShareholderDailyIncomeService

//WithdrawalService service
type WithdrawalService struct {
}

//Application 创建提现申请模块
func (withdrawalService *WithdrawalService) Application(reqData *request.ApplicationWithdrawal) *response.Resp {

	//校验股东是否存在
	shareholder, err := shareholderModel.GetByID(reqData.ShareholderID)
	if err != nil {
		return &response.Resp{Code: "sys_err", Msg: "系统出错!", Data: nil}
	}
	if shareholder == nil {
		return &response.Resp{Code: "not_exist", Msg: "账号不存在!", Data: nil}
	}
	//检查提现密码
	isOk := util.VerifyPwd(shareholder.WithdrawalPwd, reqData.WithdrawalPwd)
	if !isOk {
		return &response.Resp{Code: "fail", Msg: "密码错误!", Data: nil}
	}

	//创建提现申请
	createWithdrawalParam := &model.Withdrawal{
		ID:            util.GenerateID(1).String(),
		ShareholderID: reqData.ShareholderID,
		Hash:          reqData.Hash,
		State:         enum.WithdrawalApplication,
		Amount:        float64(reqData.Amount),
		CreateAt:      time.Now(),
	}
	err = withdrawalModel.Create(createWithdrawalParam)
	if err != nil {
		return &response.Resp{Code: "sys_err", Msg: "创建提现数据失败!", Data: nil}
	}
	return &response.Resp{Code: "ok", Msg: "请求成功!", Data: createWithdrawalParam.ID}

}

//GetList 获取矿池分页数据
func (withdrawalService *WithdrawalService) GetList(reqData *request.GetList) *response.Resp {

	res, err := withdrawalModel.GetList(uint64(reqData.Num), uint64(reqData.Page))
	if err != nil {
		return &response.Resp{Code: "sys_err", Msg: "获取提现申请数据 失败!", Data: nil}
	}
	return &response.Resp{Code: "ok", Msg: "请求成功!", Data: res}

}

//Detail 获取提现详情
func (withdrawalService *WithdrawalService) Detail(reqData *request.Detail) *response.Resp {

	res, err := withdrawalModel.GetByID(reqData.ID)
	if err != nil {
		return &response.Resp{Code: "sys_err", Msg: "获取提现数据 失败!", Data: nil}
	}
	return &response.Resp{Code: "ok", Msg: "请求成功!", Data: res}

}

//ListOfShareholder 股东提现申请列表
func (withdrawalService *WithdrawalService) ListOfShareholder(reqData *request.Detail) *response.Resp {

	res, err := withdrawalModel.GetByShareholderID(reqData.ID)
	if err != nil {
		return &response.Resp{Code: "sys_err", Msg: "获取提现申请 失败!", Data: nil}
	}
	return &response.Resp{Code: "ok", Msg: "请求成功!", Data: res}

}

//Pass 后台通过提现申请列表
func (withdrawalService *WithdrawalService) Pass(reqData *request.Detail) *response.Resp {

	withdrawal, err := withdrawalModel.GetByID(reqData.ID)
	if err != nil {
		return &response.Resp{Code: "sys_err", Msg: "系统出错!", Data: nil}
	}
	//获取用户账户信息
	result := shareholderDailyIncomeService.Info(&request.Detail{
		ID: withdrawal.ShareholderID,
	})
	//将Interface 转换类型
	shareholderInfo, _ := result.Data.(model.IncomeInfo)
	if shareholderInfo.Balance < 0 {
		return &response.Resp{Code: "not_enough", Msg: "余额不足!", Data: nil}
	}

	//审批通过的余额
	balance := shareholderInfo.Income - withdrawal.Amount

	//事务操作
	err = global.DB.GetDbR().Transaction(func(tx *gorm.DB) error {
		state := enum.WithdrawalSuccess
		//更新审批状态
		if err := tx.Model(&model.Withdrawal{}).Where("id = ?", reqData.ID).Update("state", state).Error; err != nil {
			// db处理错误...
			global.Logger.Error("service withdrawal Pass Transaction 更新审批状态 报错:", err.Error)

			return err
		}
		//扣除余额
		if err := tx.Model(&model.Shareholder{}).Where("id = ?", withdrawal.ShareholderID).Update("income", balance).Error; err != nil {
			// db处理错误...
			global.Logger.Error("service withdrawal Pass Transaction 扣除余额 报错:", err.Error)

			return err
		}
		return nil
	})

	if err != nil {
		return &response.Resp{Code: "sys_err", Msg: "系统出错!", Data: nil}
	}
	return &response.Resp{Code: "ok", Msg: "请求成功!", Data: withdrawal}

}

//Reject 后台驳回提现申请列表
func (withdrawalService *WithdrawalService) Reject(reqData *request.WithdrawalReject) *response.Resp {

	param := &model.Withdrawal{
		ID:      reqData.ID,
		State:   enum.WithdrawalReject,
		Content: reqData.Content,
	}
	err := withdrawalModel.Update(param)
	if err != nil {
		return &response.Resp{Code: "sys_err", Msg: "系统出错!", Data: nil}
	}

	return &response.Resp{Code: "ok", Msg: "请求成功!", Data: nil}

}
