package model

import (
	"pool_backend/src/global"
	"time"
)

// FilPoolRatio model
type FilPoolRatio struct {
	FilPoolID          string    `gorm:"column:fil_pool_id" json:"fil_pool_id" form:"fil_pool_id"`
	ShareholderID      string    `gorm:"column:shareholder_id" json:"shareholder_id" form:"shareholder_id"`
	ProportionOfShares int64     `gorm:"column:proportion_of_shares" json:"proportion_of_shares" form:"proportion_of_shares"`
	CreateAt           time.Time `gorm:"column:create_at" json:"create_at" form:"create_at"`
	EndAt              time.Time `gorm:"column:end_at" json:"end_at" form:"end_at"`
	UpdateAt           time.Time `gorm:"column:update_at" json:"update_at" form:"update_at"`
}

//FilPoolRatioInfo 返回
type FilPoolRatioInfo struct {
	ShareholderID      string `gorm:"column:shareholder_id" json:"shareholder_id" form:"shareholder_id"`
	ProportionOfShares int64  `gorm:"column:proportion_of_shares" json:"proportion_of_shares" form:"proportion_of_shares"`
}

//FilPoolRatioList 股东矿池数据分页列表
type FilPoolRatioList struct {
	PageInfo PageInfo       `json:"page_info"`
	Data     []FilPoolRatio `json:"data"`
}

//TableName 表名
func (FilPoolRatio) TableName() string {
	return "fil_pool_ratio"
}

//Create 创建filPool
func (filPoolRatio *FilPoolRatio) Create(filPoolParam *[]FilPoolRatio) error {
	if result := global.DB.GetDbR().Table(FilPoolRatio{}.TableName()).Create(&filPoolParam); result.Error != nil {
		// 查询db处理错误...
		global.Logger.Error("创建 FilPoolRatio 出错:", result.Error)
		return result.Error
	}
	return nil
}

//GetByFilPoolID 根据id获取信息
func (filPoolRatio *FilPoolRatio) GetByFilPoolID(ID string) (*[]FilPoolRatioInfo, error) {
	filPoolRatioList := make([]FilPoolRatioInfo, 0)
	if result := global.DB.GetDbR().Table(FilPoolRatio{}.TableName()).Where("fil_pool_id = ?", ID).Find(&filPoolRatioList); result.Error != nil {
		// 查询db处理错误...
		global.Logger.Error("根据id获取 FilPoolRatio list 报错:", result.Error)
		return nil, result.Error
	}
	return &filPoolRatioList, nil
}

//GetList 获取矿池列表
func (filPoolRatio *FilPoolRatio) GetList(num uint64, page uint64) (*FilPoolRatioList, error) {
	filPoolList := new(FilPoolRatioList)
	filPoolList.PageInfo.Page = page
	filPoolList.PageInfo.Size = num
	filPoolList.Data = make([]FilPoolRatio, 0)

	tx := global.DB.GetDbR().Table(FilPoolRatio{}.TableName())

	// 排序
	tx = tx.Order("create_at desc")

	//查询
	err := tx.Limit(int(num)).Offset(int(page*num - num)).Find(&filPoolList.Data).Error
	if err != nil {
		global.Logger.Error("获取 矿池 列表 报错:", err.Error())
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

//DeleteByFilPoolID 根据矿池id删除数据
func (filPoolRatio *FilPoolRatio) DeleteByFilPoolID(ID string) error {
	if result := global.DB.GetDbR().Table(FilPoolRatio{}.TableName()).Where("fil_pool_id = ?", ID).Delete(&filPoolRatio); result.Error != nil {
		// 查询db处理错误...
		global.Logger.Error("根据矿池id删除 FilPoolRatio list 报错:", result.Error)
		return result.Error
	}
	return nil
}

//All 获取矿池分配所有数据
func (filPoolRatio *FilPoolRatio) All() ([]FilPoolRatio, error) {
	filPoolRatioList := make([]FilPoolRatio, 0)

	if result := global.DB.GetDbR().Table(FilPoolRatio{}.TableName()).Find(&filPoolRatioList); result.Error != nil {
		// 查询db处理错误...
		global.Logger.Error("model FilPoolRatio All 报错:", result.Error)
		return nil, result.Error
	}
	return filPoolRatioList, nil
}
