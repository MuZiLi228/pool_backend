package service

import (
	"pool_backend/src/model"
	"pool_backend/src/model/request"
	"pool_backend/src/model/response"
	"time"
)

var filPoolRatioModel model.FilPoolRatio

//FilPoolRatioService service
type FilPoolRatioService struct {
}

//Create 批量创建股东矿池数据
func (fillPoolService *FilPoolRatioService) Create(reqData *request.CreateFilPoolRatio) *response.Resp {

	//根据矿池id删除掉之前的记录
	err := filPoolRatioModel.DeleteByFilPoolID(reqData.FilPoolID)
	if err != nil {
		return &response.Resp{Code: "sys_err", Msg: "创建矿池数据失败!", Data: nil}
	}

	filPoolRatioList := make([]model.FilPoolRatio, 0)
	for _, value := range reqData.ShareholerShareList {
		createFilPoolRatioParam := model.FilPoolRatio{
			FilPoolID:          reqData.FilPoolID,
			ProportionOfShares: value.ProportionOfShares,
			ShareholderID:      value.ShareholderID,
			CreateAt:           time.Now(),
		}
		filPoolRatioList = append(filPoolRatioList, createFilPoolRatioParam)
	}
	err = filPoolRatioModel.Create(&filPoolRatioList)
	if err != nil {
		return &response.Resp{Code: "sys_err", Msg: "创建矿池数据失败!", Data: nil}
	}
	return &response.Resp{Code: "ok", Msg: "请求成功!", Data: nil}

}

//GetList 获取股东矿池分页数据
func (fillPoolService *FilPoolRatioService) GetList(reqData *request.GetList) *response.Resp {

	res, err := filPoolRatioModel.GetList(uint64(reqData.Num), uint64(reqData.Page))
	if err != nil {
		return &response.Resp{Code: "sys_err", Msg: "获取矿池数据 失败!", Data: nil}
	}
	return &response.Resp{Code: "ok", Msg: "请求成功!", Data: res}

}

//Detail 获取股东矿池列表
func (fillPoolService *FilPoolRatioService) Detail(reqData *request.Detail) *response.Resp {

	res, err := filPoolRatioModel.GetByFilPoolID(reqData.ID)
	if err != nil {
		return &response.Resp{Code: "sys_err", Msg: "获取矿池数据 失败!", Data: nil}
	}
	return &response.Resp{Code: "ok", Msg: "请求成功!", Data: res}

}
