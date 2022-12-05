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

// FindUserInfo 根据用户名查询用户是否存在，并返回用户信息
func FindUserInfo(userName string) (*models.User, bool) {
	var userInfo models.User
	// 查询用户
	rows := global.DB.Where(&models.User{UserName: userName}).Find(&userInfo)
	if rows.RowsAffected < 1 {
		return &userInfo, false
	}
	return &userInfo, true
}

// RegisterUser 用户注册
func RegisterUser(insterUserInfo models.User) (uint, error) {
	err := global.DB.Create(&insterUserInfo).Error
	if err != nil {
		return 0, err
	}
	return insterUserInfo.ID, nil
}
