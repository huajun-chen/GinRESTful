package initialize

import (
	"GinRESTful/restapi/config"
	"GinRESTful/restapi/global"
	"github.com/spf13/viper"
)

func InitConfig() {
	// 实例化viper
	v := viper.New()
	// 文件的路径设置
	v.SetConfigFile(global.SettingFile)
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	serverConfig := config.ServerConfig{}
	// 给serverConfig初始化值
	err = v.Unmarshal(&serverConfig)
	if err != nil {
		panic(err)
	}
	// 传递全局变量
	global.Settings = serverConfig
}
