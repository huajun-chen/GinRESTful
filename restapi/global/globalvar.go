package global

import (
	"GinRESTful/restapi/config"
	ut "github.com/go-playground/universal-translator"
	"go.uber.org/zap"
)

var (
	Lg           *zap.Logger
	Settings     config.ServerConfig
	Trans        ut.Translator
	ParameterErr = "参数不正确"
)
