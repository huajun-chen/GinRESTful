package dao

import (
	"GinRESTful/restapi/global"
	"GinRESTful/restapi/models"
	"GinRESTful/restapi/utils"
)

// DaoGetUserList 获取用户列表
func DaoGetUserList(page, pageSize int) (int, []models.User, error) {
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

// DaoFindUserInfoToId 根据用户ID查询用户信息
func DaoFindUserInfoToId(userId uint) (*models.User, bool) {
	var userInfo models.User
	rows := global.DB.Where(&models.User{ID: userId}).Find(&userInfo)
	if rows.RowsAffected < 1 {
		return &userInfo, false
	}
	return &userInfo, true
}

// DaoFindUserInfoToUserName 根据用户名查询用户是否存在，并返回用户信息
func DaoFindUserInfoToUserName(userName string) (*models.User, bool) {
	var userInfo models.User
	// 查询用户
	rows := global.DB.Where(&models.User{UserName: userName}).Find(&userInfo)
	if rows.RowsAffected < 1 {
		return &userInfo, false
	}
	return &userInfo, true
}

// DaoRegisterUser 用户注册
func DaoRegisterUser(insterUserInfo models.User) (uint, error) {
	err := global.DB.Create(&insterUserInfo).Error
	if err != nil {
		return 0, err
	}
	return insterUserInfo.ID, nil
}

// DaoModifyUserInfo 修改用户信息
func DaoModifyUserInfo(userId uint, userMod models.User) error {
	err := global.DB.Model(&models.User{}).Where("id = ?", userId).Updates(userMod).Error
	if err != nil {
		return err
	}
	return nil
}
