package wlc

// QueryRsp 实名认证查询返回数据
type QueryRsp struct {
	Error
	Data *CheckData `json:"data"`
}
