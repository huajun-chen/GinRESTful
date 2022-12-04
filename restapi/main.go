package main

import (
	"GinRESTful/restapi/global"
	"GinRESTful/restapi/initialize"
	"fmt"
	"github.com/fatih/color"
	"go.uber.org/zap"
)

func main() {
	// 1.初始化YAML配置
	initialize.InitConfig()
	// 2.初始化router
	Router := initialize.InitRouter()
	// 3.初始化日志信息
	initialize.InitLogger()
	// 4.初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}
	// 5.初始化MySQL
	initialize.InitMysqlDB()

	color.Cyan("Gin服务开始了...")
	// 启动Gin，并配置端口
	err := Router.Run(fmt.Sprintf(":%d", global.Settings.Port))
	if err != nil {
		zap.L().Info("This is main.go", zap.String("error", "main启动错误..."))
	}
}
