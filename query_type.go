package wlc

// QueryResult 实名认证查询返回数据
type QueryResult struct {
	Status CheckStatus `json:"status"` // 认证结果
	PI     string      `json:"pi"`     // 已通过实名认证用户的唯一标识
}
