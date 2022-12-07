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
	PassWordErr      = "密码不正确"
	PassWordDiff     = "密码不一致"
	UserNameExists   = "用户名已存在"
	RegisterFail     = "注册失败"
	PwdOldErr        = "旧密码不正确"
	PwdOldNewSame    = "新密码与旧密码一致"

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
	PassWordErrCode      = 10015
	PassWordDiffCode     = 10016
	UserNameExistsCode   = 10017
	RegisterFailCode     = 10018
	PwdOldErrCode        = 10019
	PwdOldNewSameCode    = 10020

	RegisterSucc = "注册成功"
	LoginSucc    = "登录成功"
	LogoutSucc   = "登出成功"
	InsertDBSucc = "数据添加成功"
	DeleteDBSucc = "数据删除成功"
	UpdateDBSucc = "数据修改成功"
	SelectDBSucc = "数据查询成功"
)
