package initialize

import (
	"GinRESTful/restapi/global"
	"GinRESTful/restapi/models"
	"GinRESTful/restapi/utils"
)

// InitAdminAccount 初始化一个admin账户
func InitAdminAccount() {
	// 默认配置的管理员用户名
	adminInfo := global.Settings.AdminInfo
	// 查询admin是否存在
	var isUser models.User
	err := global.DB.Where("user_name = ?", adminInfo.UserName).First(&isUser).Error
	// 不存在，创建
	if err != nil {
		// 加密密码
		pwdBytes, err := utils.SetPassword(adminInfo.Password)
		if err != nil {
			panic(err)
		}
		// 创建管理员账户
		user := models.User{UserName: adminInfo.UserName, Password: string(pwdBytes), Role: 1}
		if err := global.DB.Create(&user).Error; err != nil {
			panic(err)
		}
	}
}
