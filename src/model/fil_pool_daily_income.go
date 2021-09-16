package model

import (
	"pool_backend/src/global"
	"pool_backend/src/model/request"
	"time"
)

//FilPoolDailyIncome 矿池每日收益
type FilPoolDailyIncome struct {
	ID string `gorm:"column:id;primary_key" json:"id" form:"id"`
	//矿池id
	FilPoolID string `gorm:"column:fil_pool_id" json:"fil_pool_id" form:"fil_pool_id"`
	//是否已分配
	IsAllocated bool `gorm:"column:is_allocated" json:"is_allocated" form:"is_allocated"`
	//分配值
	AssignedVal float64 `gorm:"column:assigned_val" json:"assigned_val" form:"assigned_val"`
	//上次值
	LastTimeVal float64 `gorm:"column:last_time_val" json:"last_time_val" form:"last_time_val"`
	//本日值
	TodayVal float64 `gorm:"column:today_val" json:"today_val" form:"today_val" `
	//有效算力
	Freed float64 `gorm:"column:freed" json:"freed" form:"freed" `

	CreateAt time.Time `gorm:"column:create_at" json:"create_at" form:"create_at"`
}

//FilPoolDailyIncomeList 矿池数据分页列表
type FilPoolDailyIncomeList struct {
	PageInfo PageInfo             `json:"page_info"`
	Data     []FilPoolDailyIncome `json:"data"`
}

//TableName 表名
func (FilPoolDailyIncome) TableName() string {
	return "fil_pool_daily_income"
}

//Update 更新
func (filPoolDailyIncome *FilPoolDailyIncome) Update(filPoolDailyIncomeParam *FilPoolDailyIncome) error {
	if result := global.DB.GetDbR().Table(FilPoolDailyIncome{}.TableName()).Where("id", filPoolDailyIncomeParam.ID).Updates(&filPoolDailyIncomeParam); result.Error != nil {
		// 查询db处理错误...
		global.Logger.Error("model FilPoolDailyIncome Update 出错:", result.Error)
		return result.Error
	}
	return nil
}

//Create 创建
func (filPoolDailyIncome *FilPoolDailyIncome) Create(filPoolDailyIncomeParam *FilPoolDailyIncome) error {
	if result := global.DB.GetDbR().Table(FilPoolDailyIncome{}.TableName()).Create(&filPoolDailyIncomeParam); result.Error != nil {
		// 查询db处理错误...
		global.Logger.Error("model FilPoolDailyIncome Create 出错:", result.Error)
		return result.Error
	}
	return nil
}

//GetByID 根据id获取信息
func (filPoolDailyIncome *FilPoolDailyIncome) GetByID(ID string) (*FilPoolDailyIncome, error) {
	filPoolRes := new(FilPoolDailyIncome)
	if result := global.DB.GetDbR().Table(FilPoolDailyIncome{}.TableName()).Where("id = ?", ID).Find(&filPoolRes).Limit(1); result.Error != nil {
		// 查询db处理错误...
		global.Logger.Error("model FilPoolDailyIncome GetByID 报错:", result.Error)
		return nil, result.Error
	}
	return filPoolRes, nil
}

//GetList 获取矿池列表
func (filPoolDailyIncome *FilPoolDailyIncome) GetList(reqData *request.GetList) (*FilPoolDailyIncomeList, error) {
	page := uint64(reqData.Page)
	num := uint64(reqData.Num)
	filPoolList := new(FilPoolDailyIncomeList)
	filPoolList.PageInfo.Page = page
	filPoolList.PageInfo.Size = num
	filPoolList.Data = make([]FilPoolDailyIncome, 0)

	tx := global.DB.GetDbR().Table(FilPoolDailyIncome{}.TableName())

	tx = tx.Where("fil_pool_id", reqData.ID)

	// 排序
	tx = tx.Order("create_at desc")

	//查询
	err := tx.Limit(int(num)).Offset(int(page*num - num)).Find(&filPoolList.Data).Error
	if err != nil {
		global.Logger.Error("model FilPoolDailyIncome GetList 报错:", err.Error())
		return nil, err
	}

	// 统计总数
	if err := tx.Count(&filPoolList.PageInfo.Total).Error; err != nil {
		return nil, err
	}

	filPoolList.PageInfo.TotalPage = filPoolList.PageInfo.Total / int64(filPoolList.PageInfo.Size)
	if filPoolList.PageInfo.Total%int64(filPoolList.PageInfo.Size) > 0 {
		filPoolList.PageInfo.TotalPage++
	}

	return filPoolList, nil
}

//GetIncomeList 获取矿池收益列表  昨天并且为未分配数据
func (filPoolDailyIncome *FilPoolDailyIncome) GetIncomeList() ([]FilPoolDailyIncome, error) {
	filPoolDailyIncomeList := make([]FilPoolDailyIncome, 0)
	if result := global.DB.GetDbR().Table(FilPoolDailyIncome{}.TableName()).
		Where("is_allocated", false).
		Where("create_at>=current_date-1").
		Where("create_at <current_date").
		Find(&filPoolDailyIncomeList); result.Error != nil {
		// 查询db处理错误...
		global.Logger.Error("model FilPoolDailyIncome GetBalanceList 报错:", result.Error)
		return nil, result.Error
	}
	return filPoolDailyIncomeList, nil
}
