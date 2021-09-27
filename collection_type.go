package wlc

// BTType 游戏用户行为类型
type BTType int

const (
	BTTypeLogout BTType = 0 // 下线
	BTTypeLogin  BTType = 1 // 上线
)

// CTType 用户行为数据上报类型
type CTType int

const (
	CTTypeUser  = 0 // 认证用户
	CTTypeGuest = 2 // 游客
)

// LoginTraceParam 用户行为数据上报请求参数
type LoginTraceParam struct {
	Collections []*LoginTrace `json:"collections"`
}

func (this *LoginTraceParam) Add(trace *LoginTrace) {
	if trace == nil {
		return
	}
	trace.No = int8(len(this.Collections) + 1)
	this.Collections = append(this.Collections, trace)
}

func (this *LoginTraceParam) AddUser(session string, bType BTType, opTime int64, identifier string) {
	var t = &LoginTrace{}
	t.SI = session
	t.BT = bType
	t.OT = opTime
	t.CT = CTTypeUser
	t.PI = identifier
	this.Add(t)
}

func (this *LoginTraceParam) AddUserLogin(session string, opTime int64, identifier string) {
	this.AddUser(session, BTTypeLogin, opTime, identifier)
}

func (this *LoginTraceParam) AddUserLogout(session string, opTime int64, identifier string) {
	this.AddUser(session, BTTypeLogout, opTime, identifier)
}

func (this *LoginTraceParam) AddGuest(session string, bType BTType, opTime int64, device string) {
	var t = &LoginTrace{}
	t.SI = session
	t.BT = bType
	t.OT = opTime
	t.CT = CTTypeGuest
	t.DI = device
	this.Add(t)
}

func (this *LoginTraceParam) AddGuestLogin(session string, opTime int64, device string) {
	this.AddGuest(session, BTTypeLogin, opTime, device)
}

func (this *LoginTraceParam) AddGuestLogout(session string, opTime int64, device string) {
	this.AddGuest(session, BTTypeLogout, opTime, device)
}

type LoginTrace struct {
	No int8   `json:"no"` // 在批量模式中标识一条行为数据，取值范围 1-128
	SI string `json:"si"` // 长度 32， 一个会话标识只能对应唯一的实名用户，一个实名用户可以拥有多个会话标识；同一用户单次游戏会话中，上下线动作必须使用同一会话标识上报 备注：会话标识仅标识一次用户会话，生命周期仅为一次上线和与之匹配的一次下线，不会对生命周期之外的任何业务有任何影响
	BT BTType `json:"bt"` // 游戏用户行为类型 0：下线 1：上线
	OT int64  `json:"ot"` // 行为发生时间戳，单位秒
	CT CTType `json:"ct"` // 用户行为数据上报类型 0：已认证通过用户 2：游客用户
	DI string `json:"di"` // 长度 32，游客模式设备标识，由游戏运营单位生成，游客用户下必填
	PI string `json:"pi"` // 长度 38，已通过实名认证用户的唯一标识，已认证通过用户必填
}

// LoginTraceRsp 用户行为数据上报返回数据
type LoginTraceRsp struct {
	Error
	Data *LoginTraceData `json:"data"`
}

type LoginTraceData struct {
	Results []*LoginTraceResult `json:"results"`
}

type LoginTraceResult struct {
	No int8 `json:"no"`
	Error
}
