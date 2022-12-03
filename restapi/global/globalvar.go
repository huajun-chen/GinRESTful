package global

import (
	"GinRESTful/restapi/config"
	"go.uber.org/zap"
)

var (
	Lg       *zap.Logger
	Settings config.ServerConfig
)
