package response

//Resp 响应参数
type Resp struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//AccountLoinLog 响应参数
type AccountLoinLog struct {
	ID         string `json:"id" form:"id"`
	AccountID  string `json:"account_id" form:"account_id" binding:"required"`
	NftAccount string `json:"nft_account" form:"nft_account" binding:"required"`
}
