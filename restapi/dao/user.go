package dao

import (
	"GinRESTful/restapi/global"
	"GinRESTful/restapi/models"
	"GinRESTful/restapi/utils"
)

// GetUserListDao 获取用户列表
func GetUserListDao(page, pageSize int) (int, []models.User, error) {
	var usersCount int64
	var users []models.User
	// 查询用户总数量
	global.DB.Find(&users).Count(&usersCount)

	offset := utils.OffsetResult(page, pageSize)
	limit := utils.LimitResult(pageSize)
	// 根据条件获取用户数据
	err := global.DB.Offset(offset).Limit(limit).Order("id desc").Find(&users).Error
	if err != nil {
		return 0, nil, err
	}

	return int(usersCount), users, nil
}

// FindUserInfo 通过username查找用户信息
func FindUserInfo(username, password string) (*models.User, bool) {
	var user models.User
	// 查询用户
	rows := global.DB.Where(&models.User{Name: username, Password: password}).Find(&user)
	if rows.RowsAffected < 1 {
		return &user, false
	}
	return &user, true
}
