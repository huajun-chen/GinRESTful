package controller

import (
	"GinRESTful/restapi/dao"
	"GinRESTful/restapi/forms"
	"GinRESTful/restapi/global"
	"GinRESTful/restapi/models"
	"GinRESTful/restapi/response"
	"GinRESTful/restapi/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Register 注册用户
func Register(c *gin.Context) {
	registerForm := forms.RegisterForm{}
	if err := c.ShouldBindJSON(&registerForm); err != nil {
		// 参数异常
		utils.HandleValidatorError(c, err)
		return
	}
	// 验证码
	userSet := global.Settings.UserInfo
	// 判断是否开启验证码登录
	if userSet.CaptchaLogin {
		if !store.Verify(registerForm.CaptchaId, registerForm.Captcha, true) {
			response.Response(c, response.ResponseStruct{
				Code: global.CaptchaIncorCode,
				Msg:  global.CaptchaIncor,
			})
			return
		}
	}
	// 判断用户名是否存在
	_, ok := dao.DaoFindUserInfoToUserName(registerForm.UserName)
	if ok {
		response.Response(c, response.ResponseStruct{
			Code: global.UserNameExistsCode,
			Msg:  global.UserNameExists,
		})
		return
	}
	// 两次密码是否一致
	if registerForm.Password != registerForm.Password2 {
		response.Response(c, response.ResponseStruct{
			Code: global.PassWordDiffCode,
			Msg:  global.PassWordDiff,
		})
		return
	}
	// 密码加密
	pwd, _ := utils.SetPassword(registerForm.Password)
	// 添加用户
	insterUserInfo := models.User{
		UserName: registerForm.UserName,
		Password: pwd,
	}
	userId, err := dao.DaoRegisterUser(insterUserInfo)
	if err != nil {
		response.Response(c, response.ResponseStruct{
			Code: global.RegisterFailCode,
			Msg:  global.RegisterFail,
		})
		return
	}
	// 获取Token
	token := utils.CreateToken(c, userId, 2, insterUserInfo.UserName)
	data := make(map[string]interface{})
	data["id"] = userId
	data["name"] = insterUserInfo.UserName
	data["token"] = token
	response.Response(c, response.ResponseStruct{
		Code: http.StatusOK,
		Msg:  global.RegisterSucc,
		Data: data,
	})
}

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
	userInfo, ok := dao.DaoFindUserInfoToUserName(loginForm.UserName)
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
		Msg:  global.LoginSucc,
		Data: data,
	})
}

// GetUserList 获取用户列表
func GetUserList(c *gin.Context) {
	// 获取参数
	userListForm := forms.UserListForm{}
	if err := c.ShouldBindQuery(&userListForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}
	// 获取数据
	page, pageSize := utils.PageZero(userListForm.Page, userListForm.PageSize)
	total, userList, err := dao.DaoGetUserList(page, pageSize)
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
	// 过滤用户列表，只返回需要的数据
	var values []forms.NeedsUserInfo
	for _, u := range userList {
		needUserInfo := forms.NeedsUserInfo{
			ID:        int(u.ID),
			CreatedAt: u.CreatedAt.Format("2006-01-02"),
			UserName:  u.UserName,
			Gender:    strconv.Itoa(u.Gender),
			Desc:      u.Desc,
			Role:      strconv.Itoa(u.Role),
			Mobile:    u.Mobile,
			Email:     u.Email,
		}
		values = append(values, needUserInfo)
	}
	// 获取数据正常
	data := make(map[string]interface{})
	data["total"] = total
	data["values"] = values
	response.Response(c, response.ResponseStruct{
		Code: http.StatusOK,
		Data: data,
	})
}

// ModifyUserInfo 修改用户信息
func ModifyUserInfo(c *gin.Context) {
	// 获取需要修改的用户ID
	userId := forms.IdForm{}
	if err := c.ShouldBindUri(&userId); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}
	// 获取需要修改的字段的参数
	modUserInfoForm := forms.ModifyUserInfoForm{}
	if err := c.ShouldBindJSON(&modUserInfoForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}

	// 判断是否是本人，只能修改自己的信息
	tokenUserId, _ := c.Get("userId")
	if tokenUserId != userId.ID {
		response.Response(c, response.ResponseStruct{
			Code: global.UpdateDBErrCode,
			Msg:  global.UpdateDBErr,
		})
		return
	}
	// 定义models.User接收需要修改的字段和对应的值
	userMod := models.User{}

	// 如果有修改密码，需要对密码进行判断
	if modUserInfoForm.PasswordOld != "" && modUserInfoForm.Password != "" {
		// 判断旧密码是否正确
		userInfo, _ := dao.DaoFindUserInfoToId(userId.ID)
		pwdBool := utils.CheckPassword(userInfo.Password, modUserInfoForm.PasswordOld)
		if !pwdBool {
			response.Response(c, response.ResponseStruct{
				Code: global.PwdOldErrCode,
				Msg:  global.PwdOldErr,
			})
			return
		}
		fmt.Println("哈哈", "旧密码", modUserInfoForm.PasswordOld, "新密码：", modUserInfoForm.Password)
		// 判断旧密码与新密码是否一致
		if modUserInfoForm.PasswordOld == modUserInfoForm.Password {
			response.Response(c, response.ResponseStruct{
				Code: global.PwdOldNewSameCode,
				Msg:  global.PwdOldNewSame,
			})
			return
		}
		// 判断修改后的两个密码密码是否一致
		if modUserInfoForm.Password != modUserInfoForm.Password2 {
			response.Response(c, response.ResponseStruct{
				Code: global.PassWordDiffCode,
				Msg:  global.PassWordDiff,
			})
			return
		}
		// 密码加密
		pwdStr, _ := utils.SetPassword(modUserInfoForm.Password)
		// 如果有密码修改，才将密码更新，否则不更新（json的参数没对密码做太多的限制，避免恶意传参）
		userMod.Password = pwdStr
	}

	// 更新
	userMod.Gender = modUserInfoForm.Gender
	userMod.Desc = modUserInfoForm.Desc
	userMod.Mobile = modUserInfoForm.Mobile
	userMod.Email = modUserInfoForm.Email
	if err := dao.DaoModifyUserInfo(userId.ID, userMod); err != nil {
		response.Response(c, response.ResponseStruct{
			Code: global.UpdateDBErrCode,
			Msg:  global.UpdateDBErr,
		})
		return
	}
	response.Response(c, response.ResponseStruct{
		Code: http.StatusOK,
		Msg:  global.UpdateDBSucc,
	})
}
