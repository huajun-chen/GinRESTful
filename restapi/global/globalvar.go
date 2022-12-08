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
	I18nMap  map[string]string
)

const (
	SettingFile = "./setting-dev.yaml"
)
