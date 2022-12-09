package controller

import (
	"GinRESTful/restapi/dao"
	"GinRESTful/restapi/forms"
	"GinRESTful/restapi/global"
	"GinRESTful/restapi/models"
	"GinRESTful/restapi/response"
	"GinRESTful/restapi/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

// Register 注册用户
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
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
				Code: 10008,
				Msg:  global.I18nMap["10008"],
			})
			return
		}
	}
	// 判断用户名是否存在
	_, ok := dao.DaoFindUserInfoToUserName(registerForm.UserName)
	if ok {
		response.Response(c, response.ResponseStruct{
			Code: 10017,
			Msg:  global.I18nMap["10017"],
		})
		return
	}
	// 两次密码是否一致
	if registerForm.Password != registerForm.Password2 {
		response.Response(c, response.ResponseStruct{
			Code: 10016,
			Msg:  global.I18nMap["10016"],
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
			Code: 10018,
			Msg:  global.I18nMap["10018"],
		})
		return
	}
	// 生成新的Token
	token := utils.CreateToken(c, userId, 2, insterUserInfo.UserName)
	data := forms.RegLogReturn{
		ID:    userId,
		Name:  insterUserInfo.UserName,
		Token: token,
	}
	response.Response(c, response.ResponseStruct{
		Code: http.StatusOK,
		Msg:  global.I18nMap["2000"],
		Data: data,
	})
}

// Login 用户登录
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
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
				Code: 10008,
				Msg:  global.I18nMap["10008"],
			})
			return
		}
	}
	// 查询用户是否存在
	userInfo, ok := dao.DaoFindUserInfoToUserName(loginForm.UserName)
	if !ok {
		response.Response(c, response.ResponseStruct{
			Code: 10013,
			Msg:  global.I18nMap["10013"],
		})
		return
	}
	// 判断密码是否正确
	pwdBool := utils.CheckPassword(userInfo.Password, loginForm.Password)
	if !pwdBool {
		response.Response(c, response.ResponseStruct{
			Code: 10015,
			Msg:  global.I18nMap["10015"],
		})
		return
	}
	token := utils.CreateToken(c, userInfo.ID, userInfo.Role, userInfo.UserName)
	data := forms.RegLogReturn{
		ID:    userInfo.ID,
		Name:  userInfo.UserName,
		Token: token,
	}
	response.Response(c, response.ResponseStruct{
		Code: http.StatusOK,
		Msg:  global.I18nMap["2001"],
		Data: data,
	})
}

// Logout 用户登出
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func Logout(c *gin.Context) {
	// 获取Token
	tokenStr, _ := c.Get("token")
	// 获取用户ID
	userId, _ := c.Get("userId")
	// 获取Token到期时间
	tokenExpiresAt, _ := c.Get("tokenExpiresAt")
	// 计算Token剩余的时间（Token到期时间戳 - 当前时间戳）
	timeLeft := time.Duration(tokenExpiresAt.(int64)-time.Now().Unix()) * time.Second
	// 计算Token MD5值
	tokenMD5 := utils.MD5(tokenStr.(string))
	// 将Key（Token MD5值），value（用户ID），到期时间（Token剩余的时间）加入Redis
	// 延迟10秒执行，避免此用户的其他请求还未返回Token就失效
	go func() {
		time.Sleep(10 * time.Second)
		err := utils.RedisSetStr(tokenMD5, userId, timeLeft)
		if err != nil {
			// Set Redis 错误的话，只记录日志
			zap.L().Error("token Set Redis faild", zap.String("Redis Set", tokenStr.(string)))
		}
	}()

	response.Response(c, response.ResponseStruct{
		Code: http.StatusOK,
		Msg:  global.I18nMap["2002"],
	})
}

// GetMyselfInfo 获取用户自己的信息
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func GetMyselfInfo(c *gin.Context) {
	// 从参数中获取用户ID
	userId := forms.IdForm{}
	if err := c.ShouldBindUri(&userId); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}
	// 判断是本人，不能获取别人的用户信息
	tokenUserId, _ := c.Get("userId")
	if tokenUserId != userId.ID {
		response.Response(c, response.ResponseStruct{
			Code: 10014,
			Msg:  global.I18nMap["10014"],
		})
		return
	}

	// 通过用户ID获取用户信息
	myselfInfo, _ := dao.DaoFindUserInfoToId(userId.ID)
	data := forms.UserInfoReturn{
		ID:        myselfInfo.ID,
		CreatedAt: myselfInfo.CreatedAt.Format("2006-01-02"),
		UserName:  myselfInfo.UserName,
		Gender:    strconv.Itoa(myselfInfo.Gender),
		Desc:      myselfInfo.Desc,
		Role:      strconv.Itoa(myselfInfo.Role),
		Mobile:    myselfInfo.Mobile,
		Email:     myselfInfo.Email,
	}
	response.Response(c, response.ResponseStruct{
		Code: http.StatusOK,
		Data: data,
	})
}

// GetUserList 获取用户列表
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
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
			Code: 10004,
			Msg:  global.I18nMap["10004"],
		})
		return
	}
	// 获取数据为空
	if total == 0 {
		response.Response(c, response.ResponseStruct{
			Code: 10005,
			Msg:  global.I18nMap["10005"],
		})
		return
	}
	// 过滤用户列表，只返回需要的数据
	var values []forms.UserInfoReturn
	for _, u := range userList {
		needUserInfo := forms.UserInfoReturn{
			ID:        u.ID,
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
	data := forms.UserListReturn{
		Total:  total,
		Values: values,
	}
	response.Response(c, response.ResponseStruct{
		Code: http.StatusOK,
		Data: data,
	})
}

// ModifyUserInfo 修改用户信息
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
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
			Code: 10003,
			Msg:  global.I18nMap["10003"],
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
				Code: 10019,
				Msg:  global.I18nMap["10019"],
			})
			return
		}
		// 判断旧密码与新密码是否一致
		if modUserInfoForm.PasswordOld == modUserInfoForm.Password {
			response.Response(c, response.ResponseStruct{
				Code: 10020,
				Msg:  global.I18nMap["10020"],
			})
			return
		}
		// 判断修改后的两个密码密码是否一致
		if modUserInfoForm.Password != modUserInfoForm.Password2 {
			response.Response(c, response.ResponseStruct{
				Code: 10016,
				Msg:  global.I18nMap["10016"],
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
			Code: 10003,
			Msg:  global.I18nMap["10003"],
		})
		return
	}
	response.Response(c, response.ResponseStruct{
		Code: http.StatusOK,
		Msg:  global.I18nMap["2005"],
	})
}

// DelUser 删除用户信息（需要权限）
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func DelUser(c *gin.Context) {
	// 从参数中获取用户ID
	userId := forms.IdForm{}
	if err := c.ShouldBindUri(&userId); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}

	userMod := models.User{ID: userId.ID}
	// 删除用户
	err := dao.DaoDelUserToPriKey(userMod)
	if err != nil {
		response.Response(c, response.ResponseStruct{
			Code: 10002,
			Msg:  global.I18nMap["10002"],
		})
		return
	}
	response.Response(c, response.ResponseStruct{
		Code: http.StatusOK,
		Msg:  global.I18nMap["2004"],
	})
}
