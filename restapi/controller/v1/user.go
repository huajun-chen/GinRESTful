package v1

import (
	"GinRESTful/restapi/response"
	"GinRESTful/restapi/service"
	"github.com/gin-gonic/gin"
)

// ConRegister 控制层：注册用户
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func ConRegister(c *gin.Context) {
	JSONStr := service.SerRegister(c)
	response.Response(c, JSONStr)
}

// ConLogin 控制层：用户登录
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func ConLogin(c *gin.Context) {
	JSONStr := service.SerLogin(c)
	response.Response(c, JSONStr)
}

// ConLogout 控制层：用户登出
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func ConLogout(c *gin.Context) {
	JSONStr := service.SerLogout(c)
	response.Response(c, JSONStr)
}

// ConGetMyselfInfo 控制层：获取用户自己的信息
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func ConGetMyselfInfo(c *gin.Context) {
	JSONStr := service.SerGetMyselfInfo(c)
	response.Response(c, JSONStr)
}

// ConGetUserList 控制层：获取用户列表
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func ConGetUserList(c *gin.Context) {
	JSONStr := service.SerGetUserList(c)
	response.Response(c, JSONStr)
}

// ConModifyUserInfo 控制层：修改用户信息
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func ConModifyUserInfo(c *gin.Context) {
	JSONStr := service.SerModifyUserInfo(c)
	response.Response(c, JSONStr)
}

// ConDelUser 控制层：删除用户信息（需要权限）
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func ConDelUser(c *gin.Context) {
	JSONStr := service.SerDelUser(c)
	response.Response(c, JSONStr)
}
