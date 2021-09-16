package model

import (
	"pool_backend/src/enum"
	"pool_backend/src/global"
	"time"
)

// Withdrawal model
type Withdrawal struct {
	ID            string    `gorm:"column:id;primary_key" json:"id" form:"id"`
	ShareholderID string    `gorm:"column:shareholder_id" json:"shareholder_id" form:"shareholder_id"`
	Amount        float64   `gorm:"column:amount" json:"amount" form:"amount"`
	State         int64     `gorm:"column:state" json:"state" form:"state"`
	Content       string    `gorm:"column:content" json:"content" form:"content"`
	WithLog       string    `gorm:"column:with_log" json:"with_log" form:"with_log"`
	Hash          string    `gorm:"column:hash" json:"hash" form:"hash"`
	CreateAt      time.Time `gorm:"column:create_at" json:"create_at" form:"create_at"`
	EndAt         time.Time `gorm:"column:end_at" json:"end_at" form:"end_at" `
}

//UserWithdrawal model
type UserWithdrawal struct {
	ID                      string    `gorm:"column:id;primary_key" json:"id" form:"id"`
	ShareholderID           string    `gorm:"column:shareholder_id" json:"shareholder_id" form:"shareholder_id"`
	Amount                  float64   `gorm:"column:amount" json:"amount" form:"amount"`
	State                   int64     `gorm:"column:state" json:"state" form:"state"`
	Content                 string    `gorm:"column:content" json:"content" form:"content"`
	WithLog                 string    `gorm:"column:with_log" json:"with_log" form:"with_log"`
	Hash                    string    `gorm:"column:hash" json:"hash" form:"hash"`
	Mobile                  string    `gorm:"column:mobile" json:"mobile" form:"mobile" `
	Income                  float64   `gorm:"column:income" json:"income" form:"income" `
	WithdrawalLimit         float64   `gorm:"column:withdrawal_limit" json:"withdrawal_limit" form:"withdrawal_limit" `
	RecentWithdrawalAccount string    `gorm:"column:recent_withdrawal_account" json:"recent_withdrawal_account" form:"recent_withdrawal_account" `
	IsEnable                bool      `gorm:"column:is_enable" json:"is_enable" form:"is_enable" `
	CreateAt                time.Time `gorm:"column:create_at" json:"create_at" form:"create_at"`
	EndAt                   time.Time `gorm:"column:end_at" json:"end_at" form:"end_at" `
}

//WithdrawalList 提现数据分页列表
type WithdrawalList struct {
	PageInfo PageInfo         `json:"page_info"`
	Data     []UserWithdrawal `json:"data"`
}

//TableName 表名
func (Withdrawal) TableName() string {
	return "withdrawal"
}

//Update 更新withdrawal
func (withdrawal *Withdrawal) Update(withdrawalParam *Withdrawal) error {
	if result := global.DB.GetDbR().Table(Withdrawal{}.TableName()).Where("id", withdrawalParam.ID).Save(&withdrawalParam); result.Error != nil {
		// 查询db处理错误...
		global.Logger.Error("更新 Withdrawal 出错:", result.Error)
		return result.Error
	}
	return nil
}

//Create 创建withdrawal
func (withdrawal *Withdrawal) Create(withdrawalParam *Withdrawal) error {
	if result := global.DB.GetDbR().Table(Withdrawal{}.TableName()).Create(&withdrawalParam); result.Error != nil {
		// 查询db处理错误...
		global.Logger.Error("创建 Withdrawal 出错:", result.Error)
		return result.Error
	}
	return nil
}

//GetByID 根据id获取信息
func (withdrawal *Withdrawal) GetByID(ID string) (*Withdrawal, error) {
	withdrawalRes := new(Withdrawal)
	if result := global.DB.GetDbR().Table(Withdrawal{}.TableName()).Where("id = ?", ID).First(&withdrawalRes); result.Error != nil {
		// 查询db处理错误...
		global.Logger.Error("model Withdrawal GetByID 报错:", result.Error)
		return nil, result.Error
	}
	return withdrawal, nil
}

//GetList 获取提现申请列表
func (withdrawal *Withdrawal) GetList(num uint64, page uint64) (*WithdrawalList, error) {
	withdrawalList := new(WithdrawalList)
	withdrawalList.PageInfo.Page = page
	withdrawalList.PageInfo.Size = num
	withdrawalList.Data = make([]UserWithdrawal, 0)

	tx := global.DB.GetDbR().Table(withdrawal.TableName() + " AS w").
		Joins("left join shareholder as s on s.id=w.shareholder_id")

	// 统计总数
	if err := tx.Count(&withdrawalList.PageInfo.Total).Error; err != nil {
		global.Logger.Error("获取 提现申请 列表 统计条数 报错:", err.Error())
		return nil, err
	}

	// 排序
	tx = tx.Order("w.create_at DESC")

	// 分页
	tx = tx.Limit(int(num)).Offset(int(page*num - num))

	// 筛选字段
	tx = tx.Select("w.*,s.mobile,s.withdrawal_limit,s.is_enable")

	// 查询数据
	if err := tx.Find(&withdrawalList.Data).Error; err != nil {
		global.Logger.Error("获取 提现申请 列表 报错:", err.Error())
		return nil, err
	}

	withdrawalList.PageInfo.TotalPage = withdrawalList.PageInfo.Total / int64(withdrawalList.PageInfo.Size)
	if withdrawalList.PageInfo.Total%int64(withdrawalList.PageInfo.Size) > 0 {
		withdrawalList.PageInfo.TotalPage++
	}

	return withdrawalList, nil
}

//GetByShareholderID 根据id获取信息
func (withdrawal *Withdrawal) GetByShareholderID(ShareholderID string) ([]Withdrawal, error) {
	withdrawalList := make([]Withdrawal, 0)
	if result := global.DB.GetDbR().Table(Withdrawal{}.TableName()).Where("shareholder_id = ?", ShareholderID).Order("create_at DESC").Find(&withdrawalList); result.Error != nil {
		// 查询db处理错误...
		global.Logger.Error("model Withdrawal GetByShareholderID 报错:", result.Error)
		return nil, result.Error
	}
	return withdrawalList, nil
}

//WithdrawalSumByShareholder 获取股东提现金额总额
func (withdrawal *Withdrawal) WithdrawalSumByShareholder(ShareholderID string) (float64, error) {
	var withdrawalSum float64
	var count int64
	tx := global.DB.GetDbR().Table(Withdrawal{}.TableName())
	tx = tx.Where("shareholder_id", ShareholderID).Where("state", enum.WithdrawalSuccess)

	// 统计总数
	if err := tx.Count(&count).Error; err != nil {
		return 0.0, err
	}
	if count != 0 {
		//统计总收益
		err := tx.Select("SUM(amount) as income").Pluck("amount", &withdrawalSum).Error
		if err != nil {
			global.Logger.Error("model FilPoolDailyIncome Info 报错:", err.Error())
			return 0.0, err
		}
	} else {
		withdrawalSum = 0.0
	}

	return withdrawalSum, nil
}
