package global

import (
	"GinRESTful/restapi/config"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Lg       *zap.Logger
	Settings config.ServerConfig
	Trans    ut.Translator
	DB       *gorm.DB
	Redis    *redis.Client
)

const (
	SettingFile      = "./setting-dev.yaml"
	ParameterErr     = "参数不正确"
	InsertDBErr      = "数据添加失败"
	DeleteDBErr      = "数据删除失败"
	UpdateDBErr      = "数据修改失败"
	SelectDBErr      = "数据查询失败"
	DataEmpty        = "获取数据为空"
	SystemErr        = "系统内部错误"
	CaptchaErr       = "生成验证码错误"
	CaptchaIncor     = "验证码不正确"
	NoToken          = "请求未携带token，无权访问，请先登录"
	AuthExpired      = "Token已过期"
	InvalidToken     = "无效的Token"
	CreateTokenFail  = "生成Token失败"
	NotRegistered    = "该用户未注册"
	AuthInsufficient = "权限不足"

	ParameterErrCode     = 10000
	InsertDBErrCode      = 10001
	DeleteDBErrCode      = 10002
	UpdateDBErrCode      = 10003
	SelectDBErrCode      = 10004
	DataEmptyCode        = 10005
	SystemErrCode        = 10006
	CaptchaErrCode       = 10007
	CaptchaIncorCode     = 10008
	NoTokenCode          = 10009
	AuthExpiredCode      = 10010
	InvalidTokenCode     = 10011
	CreateTokenFailCode  = 10012
	NotRegisteredCode    = 10013
	AuthInsufficientCode = 10014

	Page            = 1           // 页数，第几页，默认值
	PageSize        = 20          // 每页的数量，默认值
	TokenExpiration = 60 * 60 * 6 // Token过期时间(默认6小时)
)
