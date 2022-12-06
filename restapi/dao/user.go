package dao

import (
	"GinRESTful/restapi/global"
	"GinRESTful/restapi/models"
	"GinRESTful/restapi/utils"
)

// DaoGetUserList 获取用户列表
// 参数：
//		page：页数/第几页
//		pageSize：每页的数量
// 返回值：
//		int：查询用户的总数量
//		[]models.User：用户信息列表
//		error：错误信息
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
// 参数：
//		userId：用户ID
// 返回值：
//		*models.User：用户信息的指针
//		bool：查询是否成功
func DaoFindUserInfoToId(userId uint) (*models.User, bool) {
	var userInfo models.User
	rows := global.DB.Where(&models.User{ID: userId}).Find(&userInfo)
	if rows.RowsAffected < 1 {
		return &userInfo, false
	}
	return &userInfo, true
}

// DaoFindUserInfoToUserName 根据用户名查询用户是否存在，并返回用户信息
// 参数：
//		userName：用户名
// 返回值：
//		*models.User：用户信息的指针
//		bool：查询是否成功
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
// 参数：
//		insterUserInfo：需要插入的用户信息
// 返回值：
//		uint：插入的用户ID
//		error：错误信息
func DaoRegisterUser(insterUserInfo models.User) (uint, error) {
	err := global.DB.Create(&insterUserInfo).Error
	if err != nil {
		return 0, err
	}
	return insterUserInfo.ID, nil
}

// DaoModifyUserInfo 修改用户信息
// 参数：
//		userId：用户ID
//		userMod：需要修改的用户消息
// 返回值：
//		error：错误信息
func DaoModifyUserInfo(userId uint, userMod models.User) error {
	err := global.DB.Model(&models.User{}).Where("id = ?", userId).Updates(userMod).Error
	if err != nil {
		return err
	}
	return nil
}

// DaoDelUserToPriKey 根据用户主键删除用户（假删除，正常获取不到）
// 参数：
//		userMod：需要删除的用户主键
// 返回值：
//		error：错误信息
func DaoDelUserToPriKey(userMod models.User) error {
	err := global.DB.Delete(&userMod).Error
	return err
}
