package service

import (
	"pool_backend/src/model"
	"pool_backend/src/model/request"
	"pool_backend/src/model/response"
	"pool_backend/src/util"
	"time"
)

var filPoolModel model.FilPool

//FilPoolService service
type FilPoolService struct {
}

//Create 创建矿池数据
func (fillPoolService *FilPoolService) Create(reqData *request.CreateFilPool) *response.Resp {
	createFilPoolParam := &model.FilPool{
		ID:                      util.GenerateID(1).String(),
		Name:                    reqData.Name,
		Miner:                   reqData.Miner,
		MinerBalance:            reqData.MinerBalance,
		SectorSize:              reqData.SectorSize,
		MinerAvailableBalance:   reqData.MinerAvailableBalance,
		EffectiveComputingPower: reqData.EffectiveComputingPower,
		OriginalComputingPower:  reqData.OriginalComputingPower,
		NodeID:                  reqData.NodeID,
		CreateAt:                time.Now(),
	}
	err := filPoolModel.Create(createFilPoolParam)
	if err != nil {
		return &response.Resp{Code: "sys_err", Msg: "创建矿池数据失败!", Data: nil}
	}
	return &response.Resp{Code: "ok", Msg: "请求成功!", Data: createFilPoolParam.ID}

}

//GetList 获取矿池分页数据
func (fillPoolService *FilPoolService) GetList(reqData *request.GetList) *response.Resp {

	res, err := filPoolModel.GetList(uint64(reqData.Num), uint64(reqData.Page))
	if err != nil {
		return &response.Resp{Code: "sys_err", Msg: "获取矿池数据 失败!", Data: nil}
	}
	return &response.Resp{Code: "ok", Msg: "请求成功!", Data: res}

}

//Detail 获取矿池详情
func (fillPoolService *FilPoolService) Detail(reqData *request.Detail) *response.Resp {

	res, err := filPoolModel.GetByID(reqData.ID)
	if err != nil {
		return &response.Resp{Code: "sys_err", Msg: "获取矿池数据 失败!", Data: nil}
	}
	return &response.Resp{Code: "ok", Msg: "请求成功!", Data: res}

}

//Update 更新矿池数据
func (fillPoolService *FilPoolService) Update(reqData *request.UpadteFilPool) *response.Resp {

	updateFilPoolParam := &model.FilPool{
		ID:                      reqData.ID,
		Name:                    reqData.Name,
		Miner:                   reqData.Miner,
		MinerBalance:            reqData.MinerBalance,
		SectorSize:              reqData.SectorSize,
		MinerAvailableBalance:   reqData.MinerAvailableBalance,
		EffectiveComputingPower: reqData.EffectiveComputingPower,
		OriginalComputingPower:  reqData.OriginalComputingPower,
		NodeID:                  reqData.NodeID,
	}
	err := filPoolModel.Update(updateFilPoolParam)
	if err != nil {
		return &response.Resp{Code: "sys_err", Msg: "更新矿池数据 失败!", Data: nil}
	}
	return &response.Resp{Code: "ok", Msg: "请求成功!", Data: nil}

}
