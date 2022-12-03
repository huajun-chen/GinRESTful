package initialize

import (
	"GinRESTful/restapi/global"
	"GinRESTful/restapi/utils"
	"fmt"
	"go.uber.org/zap"
)

// InitLogger 初始化Logger
func InitLogger() {
	// 实例化zap配置
	cfg := zap.NewDevelopmentConfig()
	// 配置日志的输出地址
	cfg.OutputPaths = []string{
		fmt.Sprintf("%s%s.log", global.Settings.LogsAddress, utils.GetNowFormatTodayTime()), "stdout",
	}
	// 创建logger实例
	logg, _ := cfg.Build()
	// 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	zap.ReplaceGlobals(logg)
	// 注册到全局变量中
	global.Lg = logg
}
