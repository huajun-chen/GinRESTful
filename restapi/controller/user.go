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
	userSet := global.Settings.UserInfo
	// 判断是否开启验证码登录
	if userSet.CaptchaLogin {
		if !store.Verify(loginForm.CaptchaId, loginForm.Captcha, true) {
			response.Response(c, response.ResponseStruct{
				Code: global.CaptchaIncorCode,
				Msg:  global.CaptchaIncor,
			})
			return
		}
	}
	// 查询用户是否存在
	userInfo, ok := dao.FindUserInfo(loginForm.UserName)
	if !ok {
		response.Response(c, response.ResponseStruct{
			Code: global.NotRegisteredCode,
			Msg:  global.NotRegistered,
		})
		return
	}
	// 判断密码是否正确
	pwdBool := utils.CheckPassword(userInfo.Password, loginForm.Password)
	if !pwdBool {
		response.Response(c, response.ResponseStruct{
			Code: global.PassWordErrCode,
			Msg:  global.PassWordErr,
		})
		return
	}
	token := utils.CreateToken(c, userInfo.ID, userInfo.Role, userInfo.UserName)
	data := make(map[string]interface{})
	data["id"] = userInfo.ID
	data["name"] = userInfo.UserName
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
