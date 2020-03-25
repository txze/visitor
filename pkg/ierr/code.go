package ierr

const (
	Success = 0

	// 请求相关
	BindDataError          = 10001
	InvalidEmailOrPassword = 10002
	InvalidRequestPhone    = 10003
	InvalidRequestParams   = 10004
	UnexpectedError        = 10010

	// 用户
	InvalidUserId             = 20001
	InvalidToken              = 20002
	InvalidPassword           = 20003
	InvalidUsername           = 20004
	InvalidRoleId             = 20005
	InvalidRoleName           = 20006
	InvalidUsernameOrPassword = 20007
	CannotDeleteYourself      = 20008
	InvalidOpenId             = 20009

	// 数据库错误
	CreateDataFail  = 30001
	UpdateDataFail  = 30002
	QueryDataFail   = 30003
	DeleteDataFail  = 30004
	CountFail       = 30005
	PipeDataError   = 30006
	NothingToUpdate = 30009
	UnexpectDBError = 30010

	// 游戏相关
	InvalidGameCode  = 40001
	GameTimesLimit   = 40002
	InvalidPrizeName = 40003

	// 微信相关
	WxOauthFail   = 50001
	WxConfigError = 50002
)

var ErrorMessage = map[int]string{
	Success: "Success",

	BindDataError:          "解析body数据错误",
	InvalidEmailOrPassword: "无效的邮箱或者密码",
	InvalidRequestPhone:    "无效的电话号码",
	InvalidRequestParams:   "无效的请求参数",
	UnexpectedError:        "未知的错误",

	// 用户
	InvalidUserId:             "无效的用户id",
	InvalidToken:              "无效的token",
	InvalidPassword:           "无效的密码",
	InvalidUsername:           "无效的用户名",
	InvalidRoleId:             "无效的角色id",
	InvalidRoleName:           "无效的角色名称",
	InvalidUsernameOrPassword: "无效的用户名或密码",
	CannotDeleteYourself:      "不能删除自己",
	InvalidOpenId:             "无效的OpenId",

	CreateDataFail:  "创建数据失败",
	UpdateDataFail:  "更新数据失败",
	QueryDataFail:   "查询数据失败",
	DeleteDataFail:  "删除数据失败",
	CountFail:       "统计数据失败",
	PipeDataError:   "管道查询失败",
	NothingToUpdate: "没有数据需要更新",
	UnexpectDBError: "未知的数据库错误",

	InvalidGameCode:  "无效的游戏中奖码",
	GameTimesLimit:   "游戏次数限制",
	InvalidPrizeName: "无效的奖品名字",

	WxOauthFail:   "微信授权失败",
	WxConfigError: "微信配置错误",
}
