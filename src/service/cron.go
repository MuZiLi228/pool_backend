package service

import (
	"context"
	"pool_backend/src/global"
	"pool_backend/src/model"
	"pool_backend/src/util"
	"pool_backend/src/util/curl"
	"time"
)

var updatedFilPoolRatioModel model.UpdatedFilPoolRatio

//CronService 定时器service
type CronService struct {
}

//ShareholderFirstIncome 初次分配占比
type ShareholderFirstIncome struct {
	ShareholderID string
	FistIncome    float64
}

//-------------------------------------------资讯模块------------------------------------------------------

//GetWebExchange 非小号交易所排行数据
func (cronService *CronService) GetWebExchange() {
	url := "https://dncapi.bqrank.net/api/v2/exchange/web-exchange?page=1&pagesize=100&sort_type=exrank&asc=1&isinnovation=1&type=all&webp=1"
	res, err := curl.Get(url, nil)
	if err != nil {
		global.Logger.Error("cron 定时器 GetWebExchange 抓取数据 报错:", err.Error())
	}
	key := "feixiaohao:web-exchange:list"
	err = global.Cache.Set(context.Background(), key, res, time.Hour*24).Err()
	if err != nil {
		global.Logger.Error("cron 定时器 GetWebExchange 缓存redis 报错:", err.Error())
	}
}

//GetDailyNews 24小说交易排行数据
func (cronService *CronService) GetDailyNews() {
	url := "https://dncapi.bqrank.net/api/v4/news/news?channelid=24&direction=1&webp=1"
	res, err := curl.Get(url, nil)
	if err != nil {
		global.Logger.Error("cron 定时器 GetDailyNews 抓取数据 报错:", err.Error())
	}
	key := "feixiaohao:daily-news:list"
	err = global.Cache.Set(context.Background(), key, res, time.Hour*24).Err()
	if err != nil {
		global.Logger.Error("cron 定时器 GetDailyNews 缓存redis 报错:", err.Error())
	}
}

//GetFuturesStat 24小时爆仓统计
func (cronService *CronService) GetFuturesStat() {
	url := "https://dncapi.bqrank.net/api/v3/futures/stat/liquidation?timetype=24H&webp=1"
	res, err := curl.Get(url, nil)
	if err != nil {
		global.Logger.Error("cron 定时器 GetFuturesStat 抓取数据 报错:", err.Error())
	}
	key := "feixiaohao:futures:stat"
	err = global.Cache.Set(context.Background(), key, res, time.Hour*24).Err()
	if err != nil {
		global.Logger.Error("cron 定时器 GetFuturesStat 缓存redis 报错:", err.Error())
	}
}

//GetFuturesmarketBitcoin 24小时多空比
func (cronService *CronService) GetFuturesmarketBitcoin() {
	url := "https://dncapi.bqrank.net/api/v3/futuresmarket/bitcoin/longshort?webp=1"
	res, err := curl.Get(url, nil)
	if err != nil {
		global.Logger.Error("cron 定时器 GetFuturesmarketBitcoin 抓取数据 报错:", err.Error())
	}
	key := "feixiaohao:futuresmarket:bitcoin"
	err = global.Cache.Set(context.Background(), key, res, time.Hour*24).Err()
	if err != nil {
		global.Logger.Error("cron 定时器 GetFuturesmarketBitcoin 缓存redis 报错:", err.Error())
	}
}

//GetFuturesmarketExchange 交易所期货数据
func (cronService *CronService) GetFuturesmarketExchange() {
	url := "https://dncapi.bqrank.net/api/v3/futuresmarket/exchange/list?webp=1"
	res, err := curl.Get(url, nil)
	if err != nil {
		global.Logger.Error("cron 定时器 GetFuturesmarketExchange 抓取数据 报错:", err.Error())
	}
	key := "feixiaohao:futuresmarket:exchange"
	err = global.Cache.Set(context.Background(), key, res, time.Hour*24).Err()
	if err != nil {
		global.Logger.Error("cron 定时器 GetFuturesmarketExchange 缓存redis 报错:", err.Error())
	}
}

//----------------------------------分账模块-----------------初版 可迭代-----------------------

//UpdatePoolRatio 更新矿池表分配比例
//矿池表 需每天凌晨0点更新记录矿池股东分配比例fil_pool_ratio表数据 到新表update_fil_pool_ratio
func (cronService *CronService) UpdatePoolRatio() {
	global.Logger.Debug("cron 定时器 UpdatePoolRatio 获取矿池分配比例表 开始执行!")
	//查询fil_pool_ratio数据
	filPoolRatioList, err := filPoolRatioModel.All()
	if err != nil {
		global.Logger.Error("cron 定时器 UpdatePoolRatio 获取矿池分配比例表 报错:", err.Error())
	}

	//将updated_fil_pool_ratio表之前的记录删除
	updatedFilPoolRatioModel.Delete()

	//更新记录到updated_fil_pool_ratio表  股东矿池初分配以此表数据为准
	updateFilPoolRatioList := make([]model.UpdatedFilPoolRatio, 0)
	for _, value := range filPoolRatioList {
		createFilPoolRatioParam := model.UpdatedFilPoolRatio{
			FilPoolID:          value.FilPoolID,
			ProportionOfShares: value.ProportionOfShares,
			ShareholderID:      value.ShareholderID,
			CreateAt:           time.Now(),
		}
		updateFilPoolRatioList = append(updateFilPoolRatioList, createFilPoolRatioParam)
	}
	updatedFilPoolRatioModel.BatchCreate(updateFilPoolRatioList)
	global.Logger.Debug("cron 定时器 UpdatePoolRatio 获取矿池分配比例表 更新完成!")

}

//DistributePoolIncome 矿池分配收益
//一个矿池里面有多个股东进行分配
// 此定时器是在凌晨12点半计算分配,此时应该获取到昨天录入的矿池收入数据
// 记录分配完的股东每日收益值  更改矿池每日收益分配状态为已分配
func (cronService *CronService) DistributePoolIncome() {

	//获取时间为当前时间的前一天,且状态为未进行分配的数据
	filPoolIncomeList, err := filPoolDailyIncomeModel.GetIncomeList()
	if err != nil {
		global.Logger.Error("cron 定时器 FirstDistribute 获取update_fil_pool_ratio表数据 报错:", err.Error())
	}
	global.Logger.Debug("cron 定时器 DistributePoolIncome 进入分配计算 开始! 可计算矿池收益数量为:", len(filPoolIncomeList))
	if len(filPoolIncomeList) > 0 {
		//遍历获取所有矿池股东占比表
		for _, filPoolIncomeData := range filPoolIncomeList {
			//1、初分配 按股份占比进行分配
			shareholderFirstIncomeList := FirstDistribute(filPoolIncomeData.FilPoolID, float64(filPoolIncomeData.AssignedVal))
			//2、再分配 股东表里面 3% 5% 记录两个字段是分配者id,80%分配给本人
			Redistribute(shareholderFirstIncomeList, &filPoolIncomeData)
		}
		global.Logger.Debug("cron 定时器 DistributePoolIncome 进入分配计算 结束! over")
	} else {
		global.Logger.Debug("cron 定时器 DistributePoolIncome 进入分配计算 结束! 暂无符合条件 over")
	}

}

//FirstDistribute 初分配计算
func FirstDistribute(filPoolID string, filCoin float64) []ShareholderFirstIncome {
	global.Logger.Debug("cron 定时器 FirstDistribute 初次分配计算 开始!")

	//根据矿池id 获取update_fil_pool_ratio表数据
	filPoolRatioList, err := updatedFilPoolRatioModel.GetByFilPoolID(filPoolID)
	if err != nil {
		global.Logger.Error("cron 定时器 FirstDistribute 获取update_fil_pool_ratio表数据 报错:", err.Error())
	}

	//遍历矿池有多少个股东占股
	shareholderFirstIncomeList := make([]ShareholderFirstIncome, 0)
	for _, filPoolRatioData := range filPoolRatioList {
		//对每个股东占比进行分配计算    后面考虑访问量需加锁
		fistIncome := filCoin * float64(filPoolRatioData.ProportionOfShares) / 100
		// filCoin = fistIncome //分配完币也要对应减
		shareholderFirstIncome := ShareholderFirstIncome{
			ShareholderID: filPoolRatioData.ShareholderID,
			FistIncome:    fistIncome,
		}
		shareholderFirstIncomeList = append(shareholderFirstIncomeList, shareholderFirstIncome)
	}
	global.Logger.Debug("cron 定时器 FirstDistribute 初次分配计算完成!", shareholderFirstIncomeList)

	return shareholderFirstIncomeList

}

//Redistribute 再分配计算
func Redistribute(shareholderFirstIncomeList []ShareholderFirstIncome, filPoolDailyIncome *model.FilPoolDailyIncome) {
	global.Logger.Debug("进入再分配 由初次分配计算数量为:", len(shareholderFirstIncomeList))
	for _, shareholderFirstIncomeData := range shareholderFirstIncomeList {
		//查询股东表 查看是否有设置再分配,如果为空则为
		shareholder, err := shareholderModel.GetByID(shareholderFirstIncomeData.ShareholderID)
		if err != nil {
			global.Logger.Error("cron 定时器 Redistribute 获取shareholder表数据 报错:", err.Error())
		}
		global.Logger.Info("此时获取到初次分配后该股东id：%s", shareholderFirstIncomeData.ShareholderID, "收益值为%d:", shareholderFirstIncomeData.FistIncome)

		// 判断3%和5%是否有设置 没有设置则直接分配给该股东
		if shareholder.PercentFiveShareholderID == "" && shareholder.PercentThreeShareholderID == "" {
			global.Logger.Debug("进入再分配条件1: 百分之3和百分之5没有设置")
			income := shareholderFirstIncomeData.FistIncome * 88 / 100
			//没有设置则直接分配给该股东88%
			param := &model.ShareholderDailyIncome{
				ID:                   util.GenerateID(1).String(),
				ShareholderID:        shareholder.ID,
				IncomeType:           "fil_pool_daily",
				FilPoolDailyIncomeID: filPoolDailyIncome.FilPoolID,
				CashableIncome:       income,
				TheNominalIncome:     income,
				CreateAt:             time.Now(),
			}
			shareholderDailyIncomeModel.Create(param)
		}

		//3%没分配5%有分配  没有设置则分配给初分配的股东
		if shareholder.PercentThreeShareholderID == "" && shareholder.PercentFiveShareholderID != "" {
			global.Logger.Debug("进入再分配条件2: 百分之3没有设置 百分之5有设置")
			theNominalIncome := shareholderFirstIncomeData.FistIncome * 88 / 100
			shareholderIncome := shareholderFirstIncomeData.FistIncome * 83 / 100
			percentFiveShareholderIncome := shareholderFirstIncomeData.FistIncome * 5 / 100
			ShareholderDailyIncomeList := make([]model.ShareholderDailyIncome, 0)
			percentThreeShareholderParam := model.ShareholderDailyIncome{
				ID:                   util.GenerateID(1).String(),
				ShareholderID:        shareholder.PercentFiveShareholderID,
				IncomeType:           "fil_pool_daily",
				FilPoolDailyIncomeID: filPoolDailyIncome.FilPoolID,
				CashableIncome:       percentFiveShareholderIncome,
				TheNominalIncome:     shareholderIncome,
				CreateAt:             time.Now(),
			}
			shareholderParam := model.ShareholderDailyIncome{
				ID:                   util.GenerateID(2).String(),
				ShareholderID:        shareholder.ID,
				IncomeType:           "fil_pool_daily",
				FilPoolDailyIncomeID: filPoolDailyIncome.FilPoolID,
				CashableIncome:       shareholderIncome,
				TheNominalIncome:     theNominalIncome,
				CreateAt:             time.Now(),
			}
			ShareholderDailyIncomeList = append(ShareholderDailyIncomeList, percentThreeShareholderParam, shareholderParam)
			shareholderDailyIncomeModel.BatchCreate(ShareholderDailyIncomeList)
		}
		//3%有分配5%没分配
		if shareholder.PercentThreeShareholderID != "" && shareholder.PercentFiveShareholderID == "" {
			global.Logger.Debug("进入再分配条件3: 百分之3有设置 百分之5没有设置")
			theNominalIncome := shareholderFirstIncomeData.FistIncome * 88 / 100
			shareholderIncome := shareholderFirstIncomeData.FistIncome * 85 / 100
			percentThreeShareholderIncome := shareholderFirstIncomeData.FistIncome * 3 / 100
			ShareholderDailyIncomeList := make([]model.ShareholderDailyIncome, 0)
			percentThreeShareholderParam := model.ShareholderDailyIncome{
				ID:                   util.GenerateID(1).String(),
				ShareholderID:        shareholder.PercentThreeShareholderID,
				IncomeType:           "fil_pool_daily",
				FilPoolDailyIncomeID: filPoolDailyIncome.FilPoolID,
				CashableIncome:       percentThreeShareholderIncome,
				TheNominalIncome:     shareholderIncome,
				CreateAt:             time.Now(),
			}
			shareholderParam := model.ShareholderDailyIncome{
				ID:                   util.GenerateID(2).String(),
				ShareholderID:        shareholder.ID,
				IncomeType:           "fil_pool_daily",
				FilPoolDailyIncomeID: filPoolDailyIncome.FilPoolID,
				CashableIncome:       shareholderIncome,
				TheNominalIncome:     theNominalIncome,
				CreateAt:             time.Now(),
			}
			ShareholderDailyIncomeList = append(ShareholderDailyIncomeList, percentThreeShareholderParam, shareholderParam)
			shareholderDailyIncomeModel.BatchCreate(ShareholderDailyIncomeList)
		}

		//3%和5%都有都有分配
		if shareholder.PercentThreeShareholderID != "" && shareholder.PercentFiveShareholderID != "" {
			global.Logger.Debug("进入再分配条件4: 百分之3和5都有设置")
			theNominalIncome := shareholderFirstIncomeData.FistIncome * 88 / 100
			shareholderIncome := shareholderFirstIncomeData.FistIncome * 80 / 100
			percentThreeShareholderIncome := shareholderFirstIncomeData.FistIncome * 3 / 100
			percentFiveShareholderIncome := shareholderFirstIncomeData.FistIncome * 5 / 100
			ShareholderDailyIncomeList := make([]model.ShareholderDailyIncome, 0)
			percentThreeShareholderParam := model.ShareholderDailyIncome{
				ID:                   util.GenerateID(1).String(),
				ShareholderID:        shareholder.PercentThreeShareholderID,
				IncomeType:           "fil_pool_daily",
				FilPoolDailyIncomeID: filPoolDailyIncome.FilPoolID,
				CashableIncome:       percentThreeShareholderIncome,
				TheNominalIncome:     shareholderIncome,
				CreateAt:             time.Now(),
			}
			percentFiveShareholderParam := model.ShareholderDailyIncome{
				ID:                   util.GenerateID(2).String(),
				ShareholderID:        shareholder.PercentFiveShareholderID,
				IncomeType:           "fil_pool_daily",
				FilPoolDailyIncomeID: filPoolDailyIncome.FilPoolID,
				CashableIncome:       percentFiveShareholderIncome,
				TheNominalIncome:     shareholderIncome,
				CreateAt:             time.Now(),
			}
			shareholderParam := model.ShareholderDailyIncome{
				ID:                   util.GenerateID(3).String(),
				ShareholderID:        shareholder.ID,
				IncomeType:           "fil_pool_daily",
				FilPoolDailyIncomeID: filPoolDailyIncome.FilPoolID,
				CashableIncome:       shareholderIncome,
				TheNominalIncome:     theNominalIncome,
				CreateAt:             time.Now(),
			}
			ShareholderDailyIncomeList = append(ShareholderDailyIncomeList, percentThreeShareholderParam, shareholderParam, percentFiveShareholderParam)
			shareholderDailyIncomeModel.BatchCreate(ShareholderDailyIncomeList)
		}

		//分配完后将矿池每日收益状态改为已分配
		filPoolDailyIncomeModel.Update(&model.FilPoolDailyIncome{ID: filPoolDailyIncome.ID, IsAllocated: true})

		global.Logger.Debug("矿池%s:分配完成！", filPoolDailyIncome.FilPoolID)

	}
}
