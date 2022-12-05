package controller

import (
	"GinRESTful/restapi/dao"
	"GinRESTful/restapi/forms"
	"GinRESTful/restapi/global"
	"GinRESTful/restapi/response"
	"GinRESTful/restapi/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login 用户登录
func Login(c *gin.Context) {
	loginForm := forms.LoginForm{}
	if err := c.ShouldBindJSON(&loginForm); err != nil {
		// 参数异常处理
		utils.HandleValidatorError(c, err)
		return
	}
	// 验证码
	if !store.Verify(loginForm.CaptchaId, loginForm.Captcha, true) {
		response.Response(c, response.ResponseStruct{
			Code: global.CaptchaIncorCode,
			Msg:  global.CaptchaIncor,
		})
		return
	}
	user, ok := dao.FindUserInfo(loginForm.UserName, loginForm.Password)
	if !ok {
		response.Response(c, response.ResponseStruct{
			Code: global.NotRegisteredCode,
			Msg:  global.NotRegistered,
		})
		return
	}
	token := utils.CreateToken(c, user.ID, user.Role, user.UserName)
	data := make(map[string]interface{})
	data["id"] = user.ID
	data["name"] = user.UserName
	data["token"] = token
	response.Response(c, response.ResponseStruct{
		Code: http.StatusOK,
		Data: data,
	})
}

func GetUserList(c *gin.Context) {
	// 获取参数
	userListForm := forms.UserListForm{}
	if err := c.ShouldBindQuery(&userListForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}
	// 获取数据
	page, pageSize := utils.PageZero(userListForm.Page, userListForm.PageSize)
	total, userList, err := dao.GetUserListDao(page, pageSize)
	if err != nil {
		response.Response(c, response.ResponseStruct{
			Code: global.SelectDBErrCode,
			Msg:  global.SelectDBErr,
		})
		return
	}
	// 获取数据为空
	if total == 0 {
		response.Response(c, response.ResponseStruct{
			Code: global.DataEmptyCode,
			Msg:  global.DataEmpty,
		})
		return
	}
	// 获取数据正常
	data := make(map[string]interface{})
	data["total"] = total
	data["values"] = userList
	response.Response(c, response.ResponseStruct{
		Code: http.StatusOK,
		Data: data,
	})
}
