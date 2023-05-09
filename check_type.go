package wlc

// CheckStatus 认证结果
type CheckStatus int

const (
	CheckStatusSuccess CheckStatus = 0 // 认证成功
	CheckStatusProcess CheckStatus = 1 // 认证中
	CheckStatusFailed  CheckStatus = 2 // 认证失败
)

// CheckParam 实名认证请求参数
type CheckParam struct {
	AI    string `json:"ai"`    // 长度 32，游戏内部成员标识
	Name  string `json:"name"`  // 用户姓名
	IdNum string `json:"idNum"` // 用户身份证号码
}

// CheckResult 实名认证返回数据
type CheckResult struct {
	Status CheckStatus `json:"status"` // 认证结果
	PI     string      `json:"pi"`     // 已通过实名认证用户的唯一标识
}
