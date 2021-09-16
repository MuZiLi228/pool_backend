package model

import (
	"pool_backend/src/global"
	"time"
)

// UpdatedFilPoolRatio model 更新后的矿池分配占比  以此为分配
type UpdatedFilPoolRatio struct {
	FilPoolID          string    `gorm:"column:fil_pool_id" json:"fil_pool_id" form:"fil_pool_id"`
	ShareholderID      string    `gorm:"column:shareholder_id" json:"shareholder_id" form:"shareholder_id"`
	ProportionOfShares int64     `gorm:"column:proportion_of_shares" json:"proportion_of_shares" form:"proportion_of_shares"`
	CreateAt           time.Time `gorm:"column:create_at" json:"create_at" form:"create_at"`
}

//TableName 表名
func (UpdatedFilPoolRatio) TableName() string {
	return "updated_fil_pool_ratio"
}



//Delete 删除数据表所有数据
func (updatedFilPoolRatio *UpdatedFilPoolRatio) Delete() error {
	if result := global.DB.GetDbR().Exec("delete from updated_fil_pool_ratio"); result.Error != nil {
		// 查询db处理错误...
		global.Logger.Error("model FilPoolRatio Delete 报错:", result.Error)
		return result.Error
	}
	return nil
}



//BatchCreate 创建filPool
func (updatedFilPoolRatio *UpdatedFilPoolRatio) BatchCreate(updatedFilPoolRatioParam []UpdatedFilPoolRatio) error {
	if result := global.DB.GetDbR().Table(UpdatedFilPoolRatio{}.TableName()).Create(&updatedFilPoolRatioParam); result.Error != nil {
		// 查询db处理错误...
		global.Logger.Error("model UpdatedFilPoolRatio BatchCreate 报错:", result.Error)
		return result.Error
	}
	return nil
}



//GetByFilPoolID 根据id获取信息
func (updatedFilPoolRatio *UpdatedFilPoolRatio) GetByFilPoolID(ID string) ([]FilPoolRatioInfo, error) {
	filPoolRatioList := make([]FilPoolRatioInfo, 0)
	if result := global.DB.GetDbR().Table(FilPoolRatio{}.TableName()).Where("fil_pool_id = ?", ID).Find(&filPoolRatioList); result.Error != nil {
		// 查询db处理错误...
		global.Logger.Error("model UpdatedFilPoolRatio GetByFilPoolID 报错:", result.Error)
		return nil, result.Error
	}
	return filPoolRatioList, nil
}