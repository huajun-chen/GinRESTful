package main

import (
	"GinRESTful/restapi/initialize"
	"github.com/fatih/color"
)

func main() {
	// 1.初始化YAML配置
	initialize.InitConfig()
	// 2.初始化日志信息
	initialize.InitLogger()
	// 3.初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}
	// 4.初始化MySQL
	initialize.InitMysqlDB()
	// 5.初始化Redis
	initialize.InitRedis()
	// 6.初始化admin账户
	initialize.InitAdminAccount()
	color.Cyan("Gin服务开始了...")
	// 7.启动Gin服务（优雅关闭）
	initialize.RunServer()
}
