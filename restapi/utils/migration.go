package utils

import (
	"GinRESTful/restapi/global"
	"GinRESTful/restapi/models"
)

// Migration MYSQL迁移，建表，更新表
// 参数：
//		无
// 返回值：
//		无
func Migration() {
	_ = global.DB.AutoMigrate(&models.User{}) // 用户
}
