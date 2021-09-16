package model

import (
	"pool_backend/src/global"
	"time"
)

// FilPool model
type FilPool struct {
	ID string `gorm:"column:id;primary_key" json:"id" form:"id"`
	//矿池名称
	Name string `gorm:"column:name" json:"name" form:"name"`
	//矿工
	Miner string `gorm:"column:miner" json:"miner" form:"miner"`
	//账户余额
	MinerBalance float64 `gorm:"column:miner_balance" json:"miner_balance" form:"miner_balance"`
	//可用余额
	MinerAvailableBalance float64 `gorm:"column:miner_available_balance" json:"miner_available_balance" form:"miner_available_balance"`
	//扇区大小
	SectorSize float64 `gorm:"column:sector_size" json:"sector_size" form:"sector_size" `
	//有效算力
	EffectiveComputingPower string `gorm:"column:effective_computing_power" json:"effective_computing_power" form:"effective_computing_power" `
	//原值算力
	OriginalComputingPower string `gorm:"column:original_computing_power" json:"original_computing_power" form:"original_computing_power" `
	//股东人数
	ShareholdersNum int64 `gorm:"column:shareholders_num" json:"shareholders_num" form:"shareholders_num" `
	//节点id
	NodeID string `gorm:"column:node_id" json:"node_id" form:"node_id"`

	CreateAt time.Time `gorm:"column:create_at" json:"create_at" form:"create_at"`
}

//Balance id
type Balance struct {
	ID           string
	MinerBalance float64
}

//FilPoolList 矿池数据分页列表
type FilPoolList struct {
	PageInfo PageInfo  `json:"page_info"`
	Data     []FilPool `json:"data"`
}

//TableName 表名
func (FilPool) TableName() string {
	return "fil_pool"
}

//Update 更新filPool
func (filPool *FilPool) Update(filPoolParam *FilPool) error {
	if result := global.DB.GetDbR().Table(FilPool{}.TableName()).Where("id", filPoolParam.ID).Updates(&filPoolParam); result.Error != nil {
		// 查询db处理错误...
		global.Logger.Error("更新 filPool 出错:", result.Error)
		return result.Error
	}
	return nil
}

//Create 创建filPool
func (filPool *FilPool) Create(filPoolParam *FilPool) error {
	if result := global.DB.GetDbR().Table(FilPool{}.TableName()).Create(&filPoolParam); result.Error != nil {
		// 查询db处理错误...
		global.Logger.Error("创建 filPool 出错:", result.Error)
		return result.Error
	}
	return nil
}


//GetByID 根据id获取信息
func (filPool *FilPool) GetByID(ID string) (*FilPool, error) {
	filPoolRes := new(FilPool)
	if result := global.DB.GetDbR().Table(FilPool{}.TableName()).Where("id = ?", ID).Find(&filPoolRes).Limit(1); result.Error != nil {
		// 查询db处理错误...
		global.Logger.Error("根据id获取 filPool 详情 报错:", result.Error)
		return nil, result.Error
	}
	return filPoolRes, nil
}

//GetList 获取矿池列表
func (filPool *FilPool) GetList(num uint64, page uint64) (*FilPoolList, error) {
	filPoolList := new(FilPoolList)
	filPoolList.PageInfo.Page = page
	filPoolList.PageInfo.Size = num
	filPoolList.Data = make([]FilPool, 0)

	tx := global.DB.GetDbR().Table(FilPool{}.TableName())

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

