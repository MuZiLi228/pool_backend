package model

import (
	"errors"
	"pool_backend/src/global"
	"pool_backend/src/model/request"
	"time"

	"gorm.io/gorm"
)

// ShareholderDailyIncome model
type ShareholderDailyIncome struct {
	ID string `gorm:"column:id;primary_key" json:"id" form:"id"`

	ShareholderID string `gorm:"column:shareholder_id" json:"shareholder_id" form:"shareholder_id"`
	//收入类型
	IncomeType string `gorm:"column:income_type" json:"income_type" form:"income_type"`
	//矿池每日收益id
	FilPoolDailyIncomeID string `gorm:"column:fil_pool_daily_income_id" json:"fil_pool_daily_income_id" form:"fil_pool_daily_income_id"`
	//可提现收益
	CashableIncome float64 `gorm:"column:cashable_income" json:"cashable_income" form:"cashable_income"`
	//该项名义收益
	TheNominalIncome float64 `gorm:"column:the_nominal_income" json:"the_nominal_income" form:"the_nominal_income"`

	CreateAt time.Time `gorm:"column:create_at" json:"create_at" form:"create_at"`
}

//ShareholderDailyIncomeList 股东每日收益数据分页列表
type ShareholderDailyIncomeList struct {
	PageInfo PageInfo                 `json:"page_info"`
	Data     []ShareholderDailyIncome `json:"data"`
}

//ShareholderIncome 统计收益数据
type ShareholderIncome struct {
	Income float64 `json:"income"`
	Day    int64   `json:"day"`
}

//TableName 表名
func (ShareholderDailyIncome) TableName() string {
	return "shareholder_daily_income"
}

//GetByID 根据ID获取信息
func (shareholderDailyIncome *ShareholderDailyIncome) GetByID(ID string) (*ShareholderDailyIncome, error) {
	shareholderRes := new(ShareholderDailyIncome)
	if result := global.DB.GetDbR().Table(ShareholderDailyIncome{}.TableName()).Where("id = ?", ID).Find(&shareholderRes).Limit(1); result.Error != nil {
		// 查询db处理错误...
		isNotFound := errors.Is(result.Error, gorm.ErrRecordNotFound)
		if !isNotFound {
			global.Logger.Error("model ShareholderDailyIncome GetByID 报错:", result.Error)
			return nil, result.Error
		}
		return nil, nil //not found handle
	}
	return shareholderRes, nil
}

//Create 创建
func (shareholderDailyIncome *ShareholderDailyIncome) Create(shareholderDailyIncomeParam *ShareholderDailyIncome) error {
	if result := global.DB.GetDbR().Table(ShareholderDailyIncome{}.TableName()).Create(&shareholderDailyIncomeParam); result.Error != nil {
		// 查询db处理错误...
		global.Logger.Error("model ShareholderDailyIncome Create 出错:", result.Error)
		return result.Error
	}
	return nil
}

//BatchCreate 批量创建
func (shareholderDailyIncome *ShareholderDailyIncome) BatchCreate(shareholderDailyIncomeParam []ShareholderDailyIncome) error {
	if result := global.DB.GetDbR().Table(ShareholderDailyIncome{}.TableName()).Create(&shareholderDailyIncomeParam); result.Error != nil {
		// 查询db处理错误...
		global.Logger.Error("model ShareholderDailyIncome BatchCreate 出错:", result.Error)
		return result.Error
	}
	return nil
}

//GetList 获取股东每日收益列表
func (shareholderDailyIncome *ShareholderDailyIncome) GetList(reqData *request.GetList) (*ShareholderDailyIncomeList, error) {
	page := uint64(reqData.Page)
	num := uint64(reqData.Num)
	shareholderDailyIncomeList := new(ShareholderDailyIncomeList)
	shareholderDailyIncomeList.PageInfo.Page = page
	shareholderDailyIncomeList.PageInfo.Size = num
	shareholderDailyIncomeList.Data = make([]ShareholderDailyIncome, 0)

	tx := global.DB.GetDbR().Table(ShareholderDailyIncome{}.TableName())

	tx = tx.Where("shareholder_id", reqData.ID)

	// 排序
	tx = tx.Order("create_at desc")

	//查询
	err := tx.Limit(int(num)).Offset(int(page*num - num)).Find(&shareholderDailyIncomeList.Data).Error
	if err != nil {
		global.Logger.Error("model FilPoolDailyIncome GetList 报错:", err.Error())
		return nil, err
	}

	// 统计总数
	if err := tx.Count(&shareholderDailyIncomeList.PageInfo.Total).Error; err != nil {
		return nil, err
	}

	shareholderDailyIncomeList.PageInfo.TotalPage = shareholderDailyIncomeList.PageInfo.Total / int64(shareholderDailyIncomeList.PageInfo.Size)
	if shareholderDailyIncomeList.PageInfo.Total%int64(shareholderDailyIncomeList.PageInfo.Size) > 0 {
		shareholderDailyIncomeList.PageInfo.TotalPage++
	}

	return shareholderDailyIncomeList, nil
}

//Info 获取统计数据
func (shareholderDailyIncome *ShareholderDailyIncome) Info(ID string) (*ShareholderIncome, error) {

	info := new(ShareholderIncome)
	tx := global.DB.GetDbR().Table(ShareholderDailyIncome{}.TableName())
	tx = tx.Where("shareholder_id", ID)
	// 统计总数
	if err := tx.Count(&info.Day).Error; err != nil {
		global.Logger.Error("model FilPoolDailyIncome Info 报错:", err.Error())
		return nil, err
	}
	//如果表数据没有该用户数据 返回0
	if info.Day != 0 {   
		//统计总收益
		err := tx.Select("SUM(cashable_income) as income").Pluck("cashable_income", &info.Income).Error
		if err != nil {
			global.Logger.Error("model FilPoolDailyIncome Info 报错:", err.Error())
			return nil, err
		}
	}else{
		info.Income = 0
	}
	return info, nil
}
