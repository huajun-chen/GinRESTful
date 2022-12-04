package utils

import (
	"GinRESTful/restapi/global"
	"GinRESTful/restapi/models"
)

// Migration MYSQL迁移，建表，更新表
func Migration() {
	_ = global.DB.AutoMigrate(&models.User{}) // 用户
}
