package service

import (
	"pool_backend/src/model"
	"pool_backend/src/model/request"
	"pool_backend/src/model/response"
	"pool_backend/src/util"
	"time"
)

var shareholderModel model.Shareholder

// ShareholderService service
type ShareholderService struct {
}

// CheckPwd 账号检查
func (shareholderService *ShareholderService) CheckPwd(reqData *request.PwdLogin) *response.Resp {
	//检查手机号码是否存在
	shareholder, err := shareholderModel.GetByMobile(reqData.Mobile)
	if err != nil {
		return &response.Resp{Code: "sys_err", Msg: "系统出错!", Data: nil}
	}
	if shareholder == nil {
		return &response.Resp{Code: "not_exist", Msg: "账号未注册!", Data: nil}
	}
	if shareholder.IsEnable == false {
		return &response.Resp{Code: "not_auth", Msg: "账号被禁用!", Data: nil}
	}
	//检查账号密码是否匹配
	isOk := util.VerifyPwd(shareholder.LoginPwd, reqData.Pwd)
	if !isOk {
		return &response.Resp{Code: "fail", Msg: "账号密码错误!", Data: nil}
	}

	data := map[string]interface{}{
		"id":             shareholder.ID,
		"role":           shareholder.Role,
		"recommend_code": shareholder.RecommendCode,
	}
	return &response.Resp{Code: "ok", Msg: "请求成功!", Data: data}

}

//Register 股东账号注册
func (shareholderService *ShareholderService) Register(reqData *request.RegisterAccount) *response.Resp {

	//检查手机号码是否已经注册过
	shareholder, err := shareholderModel.GetByMobile(reqData.Mobile)
	if err != nil {
		return &response.Resp{Code: "sys_err", Msg: "系统出错!", Data: nil}
	}
	if shareholder != nil && shareholder.IsEnable == true {
		return &response.Resp{Code: "is_exist", Msg: "账号已存在!", Data: nil}
	}

	//检查短信验证码是否存在
	checkSmsCodeRes := CheckSmsCode(reqData.SmsCode, reqData.Mobile)
	if checkSmsCodeRes.Code != "ok" {
		return &response.Resp{Code: checkSmsCodeRes.Code, Msg: checkSmsCodeRes.Msg, Data: checkSmsCodeRes.Data}
	}

	if shareholder != nil && shareholder.IsEnable == false {
		//后台软删除用户重新注册
		err = shareholderModel.SetState(shareholder.ID, true)
		if err != nil {
			return &response.Resp{Code: "sys_err", Msg: "系统出错!", Data: nil}
		}
	} else {
		//密码加密
		pwdHash, _ := util.EncryptPwd(reqData.Pwd)

		//注册用户
		createShareholderParam := &model.Shareholder{
			ID:            util.GenerateID(1).String(),
			Mobile:        reqData.Mobile,
			LoginPwd:      string(pwdHash),
			RecommendCode: util.GenRandomNumber(),
			IsEnable:      true,
			CreateAt:      time.Now(),
		}
		//没有邀请码填系统管理员推荐码id
		if reqData.RecommendCode == "" {
			shareholderAdmin, err := shareholderModel.GetBySysAdmin()
			createShareholderParam.RecommendShareholderID = shareholderAdmin.ID
			if err != nil {
				return &response.Resp{Code: "sys_err", Msg: "系统出错!", Data: nil}
			}
		} else {
			//有邀请码则填邀请码的推荐者id
			shareholder, err := shareholderModel.GetByRecommendCode(reqData.RecommendCode)
			if err != nil {
				return &response.Resp{Code: "sys_err", Msg: "系统出错!", Data: nil}
			}
			createShareholderParam.RecommendShareholderID = shareholder.ID
			recommendNum := shareholder.RecommendNum + 1
			//更新邀请者推荐人数
			err = shareholderModel.Update(shareholder.ID, &model.Shareholder{RecommendNum: recommendNum})
			if err != nil {
				return &response.Resp{Code: "sys_err", Msg: "系统出错!", Data: nil}
			}
		}
		//创建新注册用户
		err = shareholderModel.Create(createShareholderParam)
		if err != nil {
			return &response.Resp{Code: "sys_err", Msg: "创建账号失败!", Data: nil}
		}

	}

	return &response.Resp{Code: "ok", Msg: "请求成功!", Data: nil}

}

//GetList 获取股东分页数据
func (shareholderService *ShareholderService) GetList(reqData *request.GetList) *response.Resp {

	shareholderInfoList, err := shareholderModel.GetList(uint64(reqData.Num), uint64(reqData.Page))
	if err != nil {
		return &response.Resp{Code: "sys_err", Msg: "获取股东列表 失败!", Data: nil}
	}
	shareholderList, err := shareholderModel.List()
	if err != nil {
		return &response.Resp{Code: "sys_err", Msg: "获取股东列表 失败!", Data: nil}
	}
	//数据处理 将推荐者手机号码返回
	for i, _ := range shareholderInfoList.Data {
		for _, userList := range shareholderList {
			if shareholderInfoList.Data[i].RecommendShareholderID == userList.ID {
				//此处为值拷贝 无法修改  需使用指针或者下标
				shareholderInfoList.Data[i].RecommendMobile = userList.Mobile
			}
		}
	}
	return &response.Resp{Code: "ok", Msg: "请求成功!", Data: shareholderInfoList}

}

//List 获取股东数据
func (shareholderService *ShareholderService) List() *response.Resp {

	res, err := shareholderModel.List()
	if err != nil {
		return &response.Resp{Code: "sys_err", Msg: "获取股东列表 失败!", Data: nil}
	}
	return &response.Resp{Code: "ok", Msg: "请求成功!", Data: res}

}

//SubordinateList 获取股东下级列表
func (shareholderService *ShareholderService) SubordinateList(ID string) *response.Resp {

	res, err := shareholderModel.SubordinateList(ID)
	if err != nil {
		return &response.Resp{Code: "sys_err", Msg: "获取下级列表 失败!", Data: nil}
	}
	return &response.Resp{Code: "ok", Msg: "请求成功!", Data: res}

}

//UpdatePercentShareholderID  股东更新再分配占比
func (shareholderService *ShareholderService) UpdatePercentShareholderID(reqData *request.UpdatePercentShareholderID) *response.Resp {

	param := &model.Shareholder{
		PercentThreeShareholderID: reqData.PercentThreeShareholderID,
		PercentFiveShareholderID:  reqData.PercentFiveShareholderID,
	}

	err := shareholderModel.Update(reqData.ShareholderID, param)
	if err != nil {
		return &response.Resp{Code: "sys_err", Msg: "更新股份占比失败!", Data: nil}
	}
	return &response.Resp{Code: "ok", Msg: "请求成功!", Data: nil}

}

//Detail 获取股东详细信息
func (shareholderService *ShareholderService) Detail(reqData *request.Detail) *response.Resp {

	shareholder, err := shareholderModel.GetByID(reqData.ID)
	if err != nil {
		return &response.Resp{Code: "sys_err", Msg: "获取股东信息 失败!", Data: nil}
	}
	isPayHash := true
	if shareholder.WithdrawalPwd == "" {
		isPayHash = false
	}
	data := map[string]interface{}{
		"id":                           shareholder.ID,
		"role":                         shareholder.Role,
		"recommend_code":               shareholder.RecommendCode,
		"mobile":                       shareholder.Mobile,
		"income":                       shareholder.Income,
		"is_enable":                    shareholder.IsEnable,
		"withdrawal_limit":             shareholder.WithdrawalLimit,
		"recommend_shareholder_id":     shareholder.RecommendShareholderID,
		"recommend_mobile":             "",
		"recommend_num":                shareholder.RecommendNum,
		"recommend_allocation_ratio":   shareholder.RecommendAllocationRatio,
		"fil_pool_num":                 shareholder.FilPoolNum,
		"percent_three_shareholder_id": shareholder.PercentThreeShareholderID,
		"percent_five_shareholder_id":  shareholder.PercentFiveShareholderID,
		"recent_withdrawal_account":    shareholder.RecentWithdrawalAccount,
		"create_at":                    shareholder.CreateAt,
		"is_pay_hash":                  isPayHash,
	}
	return &response.Resp{Code: "ok", Msg: "请求成功!", Data: data}

}

//SetWithdrawalPwd  股东更新再分配占比
func (shareholderService *ShareholderService) SetWithdrawalPwd(reqData *request.SetWithdrawalPwd) *response.Resp {

	//根据id获取股东hash密码 与 传过来的登录密码进行比较
	shareholder, err := shareholderModel.GetByID(reqData.ShareholderID)
	if err != nil {
		return &response.Resp{Code: "sys_err", Msg: "系统出错!", Data: nil}
	}

	//检查账号密码是否匹配
	isOk := util.VerifyPwd(shareholder.LoginPwd, reqData.LoginPwd)
	if !isOk {
		return &response.Resp{Code: "fail", Msg: "账号密码错误!", Data: nil}
	}

	//校验登录密码才可以设置提现密码
	//提现密码加密
	WithdrawalPwdHash, _ := util.EncryptPwd(reqData.WithdrawalPwd)
	param := &model.Shareholder{
		WithdrawalPwd: string(WithdrawalPwdHash),
	}
	err = shareholderModel.Update(reqData.ShareholderID, param)
	if err != nil {
		return &response.Resp{Code: "set_err", Msg: "设置提现密码失败!", Data: nil}
	}
	return &response.Resp{Code: "ok", Msg: "请求成功!", Data: nil}

}

//DailyIncomeList 获取股东每日收益数据列表分页
func (shareholderService *ShareholderService) DailyIncomeList(reqData *request.GetList) *response.Resp {

	dailyIncomeList, err := shareholderDailyIncomeModel.GetList(reqData)
	if err != nil {
		return &response.Resp{Code: "sys_err", Msg: "系统出错!", Data: nil}
	}

	return &response.Resp{Code: "ok", Msg: "请求成功!", Data: dailyIncomeList}

}

//SetEnable 禁用
func (shareholderService *ShareholderService) SetEnable(reqData *request.SetEnable) *response.Resp {

	//检查id是否为管理员id
	admin, err := shareholderModel.GetByID(reqData.ID)
	if err != nil {
		return &response.Resp{Code: "sys_err", Msg: "系统出错!", Data: nil}
	}
	if admin.Role != "admin" {
		return &response.Resp{Code: "not_auth", Msg: "暂无权限!", Data: nil}
	}
	err = shareholderModel.SetState(reqData.ShareholderID, false)
	if err != nil {
		return &response.Resp{Code: "sys_err", Msg: "系统出错!", Data: nil}
	}

	return &response.Resp{Code: "ok", Msg: "请求成功!", Data: nil}

}
