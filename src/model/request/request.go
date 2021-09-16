package request

//RegisterAccount  request
type RegisterAccount struct {
	SmsCode       string `json:"sms_code" form:"sms_code" binding:"required"`   //短信验证码
	RecommendCode string `json:"recommend_code" form:"recommend_code" `         //推荐码 没有，可不填
	Mobile        string `json:"mobile" form:"mobile" binding:"required"`       //手机号码
	Pwd           string `json:"pwd" form:"pwd" binding:"required"`             //密码
	IsEnable      bool   `json:"is_enable" form:"is_enable" binding:"required"` //是否开启
}

//PwdLogin  request
type PwdLogin struct {
	Mobile string `json:"mobile" form:"mobile" binding:"required"`
	Pwd    string `json:"pwd" form:"pwd" binding:"required"`
}

// CreateFilPool request
type CreateFilPool struct {
	Name                    string  `json:"name" form:"name" binding:"required" `                                           //"矿池名称"
	Miner                   string  `json:"miner" form:"miner" binding:"required" `                                         //"矿工"
	MinerBalance            float64 `json:"miner_balance" form:"miner_balance" binding:"required"`                          //"账户余额"
	MinerAvailableBalance   float64 `json:"miner_available_balance" form:"miner_available_balance" binding:"required" `     //"可用余额"
	SectorSize              float64 `json:"sector_size" form:"sector_size" binding:"required" `                             //"扇区大小"
	EffectiveComputingPower string  `json:"effective_computing_power" form:"effective_computing_power" binding:"required" ` //"有效算力"
	OriginalComputingPower  string  `json:"original_computing_power" form:"original_computing_power" binding:"required" `   //"原值算力"
	NodeID                  string  `json:"node_id" form:"node_id" binding:"required" `                                     //"节点id"
}

// CreateFilPoolDailyIncome request
type CreateFilPoolDailyIncome struct {
	//矿池id
	FilPoolID string `json:"fil_pool_id" form:"fil_pool_id" binding:"required"`
	//分配值
	AssignedVal float64 `json:"assigned_val" form:"assigned_val" binding:"required"`
	//上次值
	LastTimeVal float64 `json:"last_time_val" form:"last_time_val" binding:"required"`
	//本日值
	TodayVal float64 `json:"today_val" form:"today_val" binding:"required"`
	//有效算力
	Freed float64 `json:"freed" form:"freed" binding:"required" `
}

// UpadteFilPool request
type UpadteFilPool struct {
	ID                      string  `json:"id" form:"id" binding:"required" `                                               //"矿池名称"
	Name                    string  `json:"name" form:"name" binding:"required" `                                           //"矿池名称"
	Miner                   string  `json:"miner" form:"miner" binding:"required" `                                         //"矿工"
	MinerBalance            float64 `json:"miner_balance" form:"miner_balance" binding:"required"`                          //"账户余额"
	MinerAvailableBalance   float64 `json:"miner_available_balance" form:"miner_available_balance" binding:"required" `     //"可用余额"
	SectorSize              float64 `json:"sector_size" form:"sector_size" binding:"required" `                             //"扇区大小"
	EffectiveComputingPower string  `json:"effective_computing_power" form:"effective_computing_power" binding:"required" ` //"有效算力"
	OriginalComputingPower  string  `json:"original_computing_power" form:"original_computing_power" binding:"required" `   //"原值算力"
	NodeID                  string  `json:"node_id" form:"node_id" binding:"required" `                                     //"节点id"
}

//ApplicationWithdrawal request
type ApplicationWithdrawal struct {
	WithdrawalPwd string  `json:"withdrawal_pwd" form:"withdrawal_pwd" binding:"required"`
	ShareholderID string  ` json:"shareholder_id" form:"shareholder_id" binding:"required"`
	Amount        float64 ` json:"amount" form:"amount" binding:"required"`
	Content       string  ` json:"content" form:"content"`
	WithLog       string  ` json:"with_log" form:"with_log"`
	Hash          string  ` json:"hash" form:"hash" binding:"required"`
}

//Detail  request
type Detail struct {
	ID string `json:"id" form:"id" binding:"required"`
}

//SmsSend  request
type SmsSend struct {
	Mobile string `json:"mobile" form:"mobile" binding:"required"`
}

//GetList  request
type GetList struct {
	ID   string `json:"id" form:"id" `                       //id 可选参数
	Num  uint   `json:"num" form:"num" binding:"required"`   //获取当页条数
	Page uint   `json:"page" form:"page" binding:"required"` //当前页数
}

//PutApp request
type PutApp struct {
	Type        string `json:"type" form:"type" binding:"required"`                 //app类型 android|ios
	VersionName string `json:"version_name" form:"version_name" binding:"required"` //版本名称
	VersionCode int    `json:"version_code" form:"version_code" binding:"required"` //版本号
	VersionDesc string `json:"version_desc" form:"version_desc" binding:"required"` //版本描述
}

//CreateFilPoolRatio 创建矿池股东比例分配
type CreateFilPoolRatio struct {
	FilPoolID           string                `json:"fil_pool_id" form:"fil_pool_id" binding:"required"`
	ShareholerShareList []ShareholerShareList `json:"list" form:"list" binding:"required"`
}

//ShareholerShareList 股东比例
type ShareholerShareList struct {
	ShareholderID      string `json:"shareholder_id" form:"shareholder_id" binding:"required"`
	ProportionOfShares int64  `json:"proportion_of_shares" form:"proportion_of_shares" binding:"required"`
}

//UpdatePercentShareholderID 股东再分配占比
type UpdatePercentShareholderID struct {
	ShareholderID             string `json:"shareholder_id" form:"shareholder_id" binding:"required"`
	PercentThreeShareholderID string `json:"percent_three_shareholder_id" form:"percent_three_shareholder_id"`
	PercentFiveShareholderID  string `json:"percent_five_shareholder_id" form:"percent_five_shareholder_id"`
}

//SetWithdrawalPwd 股东设置提现密码
type SetWithdrawalPwd struct {
	ShareholderID string `json:"shareholder_id" form:"shareholder_id" binding:"required"`
	LoginPwd      string `json:"login_pwd" form:"login_pwd" binding:"required"`
	WithdrawalPwd string `json:"withdrawal_pwd" form:"withdrawal_pwd" binding:"required"`
}

// WithdrawalReject 提现申请驳回
type WithdrawalReject struct {
	ID      string `json:"id" form:"id" binding:"required"`
	Content string `json:"content" form:"content" binding:"required"`
}

//SetEnable  request
type SetEnable struct {
	ID            string `json:"id" form:"id" binding:"required"`
	ShareholderID string `json:"shareholder_id" form:"shareholder_id" binding:"required"`
}
