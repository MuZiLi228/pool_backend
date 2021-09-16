package model

import (
	"errors"
	"pool_backend/src/global"
	"time"

	"gorm.io/gorm"
)

// Shareholder model
type Shareholder struct {
	ID     string `gorm:"column:id;primary_key" json:"id" form:"id" `
	Mobile string `gorm:"column:mobile" json:"mobile" form:"mobile"`
	//角色权限 默认为空字符串,只有管理员才为admin
	Role string `gorm:"column:role" json:"role" form:"role"`
	//登录密码
	LoginPwd string `gorm:"column:login_pwd" json:"login_pwd" form:"login_pwd" `
	//提现密码
	WithdrawalPwd string `gorm:"column:withdrawal_pwd" json:"withdrawal_pwd" form:"withdrawal_pwd" `
	//总收益
	Income float64 `gorm:"column:income" json:"income" form:"income" `
	//是否可用
	IsEnable bool `gorm:"column:is_enable" json:"is_enable" form:"is_enable" `
	//可提现额度
	WithdrawalLimit int64 `gorm:"column:withdrawal_limit" json:"withdrawal_limit" form:"withdrawal_limit" `
	//推荐者ID
	RecommendShareholderID string `gorm:"column:recommend_shareholder_id" json:"recommend_shareholder_id" form:"recommend_shareholder_id" `
	//推荐人数
	RecommendNum int64 `gorm:"column:recommend_num" json:"recommend_num" form:"recommend_num" `
	//本账号推荐码
	RecommendCode string `gorm:"column:recommend_code" json:"recommend_code" form:"recommend_code" `
	//VIP级别 对应分配比例
	RecommendAllocationRatio int64 `gorm:"column:recommend_allocation_ratio" json:"recommend_allocation_ratio" form:"recommend_allocation_ratio" `
	//矿池数量
	FilPoolNum int64 `gorm:"column:fil_pool_num" json:"fil_pool_num" form:"fil_pool_num" `
	//最近的提现目标账号
	RecentWithdrawalAccount string `gorm:"column:recent_withdrawal_account" json:"recent_withdrawal_account" form:"recent_withdrawal_account" `
	//矿池分配占比3% 分配者id
	PercentThreeShareholderID string `gorm:"column:percent_three_shareholder_id" json:"percent_three_shareholder_id" form:"percent_three_shareholder_id" `
	//矿池分配占比5% 分配者id
	PercentFiveShareholderID string `gorm:"column:percent_five_shareholder_id" json:"percent_five_shareholder_id" form:"percent_five_shareholder_id" `

	CreateAt time.Time `gorm:"column:create_at" json:"create_at" form:"create_at" `
}

// ShareholderInfo model
type ShareholderInfo struct {
	ID     string `gorm:"column:id;primary_key" json:"id" form:"id" `
	Mobile string `gorm:"column:mobile" json:"mobile" form:"mobile"`
	//角色权限 默认为空字符串,只有管理员才为admin
	Role string `gorm:"column:role" json:"role" form:"role"`
	//总收益
	Income int64 `gorm:"column:income" json:"income" form:"income" `
	//是否可用
	IsEnable bool `gorm:"column:is_enable" json:"is_enable" form:"is_enable" `
	//可提现额度
	WithdrawalLimit int64 `gorm:"column:withdrawal_limit" json:"withdrawal_limit" form:"withdrawal_limit" `
	//推荐者ID
	RecommendShareholderID string `gorm:"column:recommend_shareholder_id" json:"recommend_shareholder_id" form:"recommend_shareholder_id" `
	//推荐者手机号码
	RecommendMobile string `gorm:"column:recommend_mobile" json:"recommend_mobile" form:"recommend_mobile"`
	//推荐人数
	RecommendNum int64 `gorm:"column:recommend_num" json:"recommend_num" form:"recommend_num" `
	//本账号推荐码
	RecommendCode string `gorm:"column:recommend_code" json:"recommend_code" form:"recommend_code" `
	//VIP级别 对应分配比例
	RecommendAllocationRatio int64 `gorm:"column:recommend_allocation_ratio" json:"recommend_allocation_ratio" form:"recommend_allocation_ratio" `
	//矿池数量
	FilPoolNum int64 `gorm:"column:fil_pool_num" json:"fil_pool_num" form:"fil_pool_num" `
	//最近的提现目标账号
	RecentWithdrawalAccount string `gorm:"column:recent_withdrawal_account" json:"recent_withdrawal_account" form:"recent_withdrawal_account" `
	//矿池分配占比3% 分配者id
	PercentThreeShareholderID string `gorm:"column:percent_three_shareholder_id" json:"percent_three_shareholder_id" form:"percent_three_shareholder_id" `
	//矿池分配占比5% 分配者id
	PercentFiveShareholderID string `gorm:"column:percent_five_shareholder_id" json:"percent_five_shareholder_id" form:"percent_five_shareholder_id" `

	CreateAt time.Time `gorm:"column:create_at" json:"create_at" form:"create_at" `
}

//ShareholderList 股东手机列表
type ShareholderList struct {
	ID     string `gorm:"column:id;primary_key" json:"id" form:"id" `
	Mobile string `gorm:"column:mobile" json:"mobile" form:"mobile"`
}

//IncomeInfo 收益数据
type IncomeInfo struct {
	Income        float64 `json:"income"`
	Balance       float64 `json:"balance"`
	WithdrawalSum float64 `json:"withdrawal_sum"`
	Day           int64   `json:"day"`
}

//ShareholderPageList 矿池数据分页列表
type ShareholderPageList struct {
	PageInfo PageInfo          `json:"page_info"`
	Data     []ShareholderInfo `json:"data"`
}

//TableName 表名
func (Shareholder) TableName() string {
	return "shareholder"
}

//GetByID 根据id获取信息
func (shareholder *Shareholder) GetByID(ID string) (*Shareholder, error) {
	shareholderRes := new(Shareholder)
	if result := global.DB.GetDbR().Table(Shareholder{}.TableName()).Where("id = ?", ID).Where("is_enable", true).Find(&shareholderRes).Limit(1); result.Error != nil {
		// 查询db处理错误...
		global.Logger.Error("根据id获取 shareholder 详情 报错:", result.Error)
		return nil, result.Error
	}
	return shareholderRes, nil
}

//GetByMobile 根据Moblie获取信息
func (shareholder *Shareholder) GetByMobile(Moblie string) (*Shareholder, error) {
	shareholderRes := new(Shareholder)
	if result := global.DB.GetDbR().Table(Shareholder{}.TableName()).Where("mobile = ?", Moblie).First(&shareholderRes); result.Error != nil {
		// 查询db处理错误...
		isNotFound := errors.Is(result.Error, gorm.ErrRecordNotFound)
		if !isNotFound {
			global.Logger.Error("根据Moblie获取 shareholder 详情 报错:", result.Error)
			return nil, result.Error
		}
		return nil, nil //not found handle
	}
	return shareholderRes, nil
}

//GetByRecommendCode 根据推荐码获取推荐者id
func (shareholder *Shareholder) GetByRecommendCode(recommendCode string) (*Shareholder, error) {
	shareholderRes := new(Shareholder)
	if result := global.DB.GetDbR().Table(Shareholder{}.TableName()).Where("recommend_code = ?", recommendCode).Where("is_enable", true).First(&shareholderRes); result.Error != nil {
		// 查询db处理错误...
		global.Logger.Error("根据recommendCode获取 shareholder 详情 报错:", result.Error)
		return nil, result.Error
	}
	return shareholderRes, nil
}

//GetBySysAdmin 获取系统最高账号   目前产品默认只有一个
func (shareholder *Shareholder) GetBySysAdmin() (*Shareholder, error) {
	shareholderRes := new(Shareholder)
	if result := global.DB.GetDbR().Table(Shareholder{}.TableName()).Where("role = ?", "admin").First(&shareholderRes); result.Error != nil {
		// 查询db处理错误...
		global.Logger.Error("根据 role 获取 shareholder 详情 报错:", result.Error)
		return nil, result.Error
	}
	return shareholderRes, nil
}

//GetList 获取股东账号列表
func (shareholder *Shareholder) GetList(num uint64, page uint64) (*ShareholderPageList, error) {
	shareholderList := new(ShareholderPageList)
	shareholderList.PageInfo.Page = page
	shareholderList.PageInfo.Size = num
	shareholderList.Data = make([]ShareholderInfo, 0)

	tx := global.DB.GetDbR().Table(Shareholder{}.TableName() + " AS s")

	// 排序
	tx = tx.Order("create_at desc").Where("is_enable", true)

	// 统计总数
	if err := tx.Count(&shareholderList.PageInfo.Total).Error; err != nil {
		global.Logger.Error("获取 获取股东账号 列表 统计条数 报错:", err.Error())
		return nil, err
	}

	// 分页
	tx = tx.Limit(int(num)).Offset(int(page*num - num))

	// 筛选字段
	tx = tx.Select("s.*")

	// 查询数据
	if err := tx.Find(&shareholderList.Data).Error; err != nil {
		global.Logger.Error("获取 获取股东账号 列表 报错:", err.Error())
		return nil, err
	}

	shareholderList.PageInfo.TotalPage = shareholderList.PageInfo.Total / int64(shareholderList.PageInfo.Size)
	if shareholderList.PageInfo.Total%int64(shareholderList.PageInfo.Size) > 0 {
		shareholderList.PageInfo.TotalPage++
	}

	return shareholderList, nil
}

//Create 创建 shareholder
func (shareholder *Shareholder) Create(shareholderParam *Shareholder) error {
	if result := global.DB.GetDbR().Table(Shareholder{}.TableName()).Create(&shareholderParam); result.Error != nil {
		// 查询db处理错误...
		global.Logger.Error("创建 shareholder 出错:", result.Error)
		return result.Error
	}
	return nil
}

//List 获取股东列表
func (shareholder *Shareholder) List() ([]ShareholderList, error) {
	shareholderList := make([]ShareholderList, 0)
	if result := global.DB.GetDbR().Table(Shareholder{}.TableName()).Where("is_enable", true).Find(&shareholderList); result.Error != nil {
		// 查询db处理错误...
		global.Logger.Error("获取 Shareholder list 报错:", result.Error)
		return nil, result.Error
	}
	return shareholderList, nil
}

//SubordinateList 获取股东下级列表
func (shareholder *Shareholder) SubordinateList(ID string) ([]ShareholderList, error) {
	shareholderList := make([]ShareholderList, 0)
	if result := global.DB.GetDbR().Table(Shareholder{}.TableName()).Where("is_enable", true).Where("recommend_shareholder_id", ID).Find(&shareholderList); result.Error != nil {
		// 查询db处理错误...
		global.Logger.Error("model Shareholder SubordinateList 报错:", result.Error)
		return nil, result.Error
	}
	return shareholderList, nil
}

//Update 更新    updates 更新非0字段
func (shareholder *Shareholder) Update(ID string, shareholderParam *Shareholder) error {

	if result := global.DB.GetDbR().Table(Shareholder{}.TableName()).Where("id", ID).Updates(shareholderParam); result.Error != nil {
		// 查询db处理错误...
		global.Logger.Error("model shareholder update 出错:", result.Error)
		return result.Error
	}
	return nil
}


//SetState 更新    updates 更新非0字段
func (shareholder *Shareholder) SetState(ID string,state bool) error {

	if result := global.DB.GetDbR().Table(Shareholder{}.TableName()).Where("id", ID).Update("is_enable",state); result.Error != nil {
		// 查询db处理错误...
		global.Logger.Error("model shareholder update 出错:", result.Error)
		return result.Error
	}
	return nil
}

