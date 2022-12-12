package v1

import (
	"GinRESTful/restapi/forms"
	"GinRESTful/restapi/response"
	"GinRESTful/restapi/service"
	"GinRESTful/restapi/utils"
	"github.com/gin-gonic/gin"
)

// ConRegister 控制层：注册用户
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func ConRegister(c *gin.Context) {
	// 获取注册需要的参数
	registerForm := forms.RegisterForm{}
	if err := c.ShouldBindJSON(&registerForm); err != nil {
		parErrStr := utils.HandleValidatorError(err)
		response.Response(c, parErrStr)
		return
	}
	resStruct := service.SerRegister(registerForm, c)
	response.Response(c, resStruct)
}

// ConLogin 控制层：用户登录
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func ConLogin(c *gin.Context) {
	// 获取登录时需要的参数
	loginForm := forms.LoginForm{}
	if err := c.ShouldBindJSON(&loginForm); err != nil {
		parErrStr := utils.HandleValidatorError(err)
		response.Response(c, parErrStr)
		return
	}
	resStruct := service.SerLogin(loginForm, c)
	response.Response(c, resStruct)
}

// ConLogout 控制层：用户登出
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func ConLogout(c *gin.Context) {
	resStruct := service.SerLogout(c)
	response.Response(c, resStruct)
}

// ConGetMyselfInfo 控制层：获取用户自己的信息
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func ConGetMyselfInfo(c *gin.Context) {
	// 从uri参数中获取用户ID
	userId := forms.IdForm{}
	if err := c.ShouldBindUri(&userId); err != nil {
		parErrStr := utils.HandleValidatorError(err)
		response.Response(c, parErrStr)
		return
	}
	resStruct := service.SerGetMyselfInfo(userId, c)
	response.Response(c, resStruct)
}

// ConGetUserList 控制层：获取用户列表
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func ConGetUserList(c *gin.Context) {
	// 查看用户列表时需要的参数
	userListForm := forms.UserListForm{}
	if err := c.ShouldBindQuery(&userListForm); err != nil {
		parErrStr := utils.HandleValidatorError(err)
		response.Response(c, parErrStr)
		return
	}
	resStruct := service.SerGetUserList(userListForm)
	response.Response(c, resStruct)
}

// ConModifyUserInfo 控制层：修改用户信息
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func ConModifyUserInfo(c *gin.Context) {
	// 从uri中获取需要修改的用户ID
	userId := forms.IdForm{}
	if err := c.ShouldBindUri(&userId); err != nil {
		parErrStr := utils.HandleValidatorError(err)
		response.Response(c, parErrStr)
		return
	}
	// 获取需要修改的字段的参数
	modUserInfoForm := forms.ModifyUserInfoForm{}
	if err := c.ShouldBindJSON(&modUserInfoForm); err != nil {
		parErrStr := utils.HandleValidatorError(err)
		response.Response(c, parErrStr)
		return
	}
	resStruct := service.SerModifyUserInfo(userId, modUserInfoForm, c)
	response.Response(c, resStruct)
}

// ConDelUser 控制层：删除用户信息（需要权限）
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func ConDelUser(c *gin.Context) {
	// 从uri参数中获取用户ID
	userId := forms.IdForm{}
	if err := c.ShouldBindUri(&userId); err != nil {
		parErrStr := utils.HandleValidatorError(err)
		response.Response(c, parErrStr)
		return
	}
	resStruct := service.SerDelUser(userId)
	response.Response(c, resStruct)
}
