package service

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

// SerRegister 业务层：注册用户
// 参数：
//		registerForm：注册账户时需要的参数
//		c：gin.Context的指针
// 返回值：
//		response.ResStruct：响应的结构体
func SerRegister(registerForm forms.RegisterForm, c *gin.Context) response.ResStruct {
	// 验证码
	userSet := global.Settings.UserInfo
	// 判断是否开启验证码登录
	if userSet.CaptchaLogin {
		if !store.Verify(registerForm.CaptchaId, registerForm.Captcha, true) {
			failStruct := response.ResStruct{
				Code: 10008,
				Msg:  global.I18nMap["10008"],
			}
			return failStruct
		}
	}
	// 判断用户名是否存在
	_, ok := dao.DaoFindUserInfoToUserName(registerForm.UserName)
	if ok {
		failStruct := response.ResStruct{
			Code: 10017,
			Msg:  global.I18nMap["10017"],
		}
		return failStruct
	}
	// 两次密码是否一致
	if registerForm.Password != registerForm.Password2 {
		failStruct := response.ResStruct{
			Code: 10016,
			Msg:  global.I18nMap["10016"],
		}
		return failStruct
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
		zap.S().Errorf("%s：%s", global.I18nMap["10018"], err)
		failStruct := response.ResStruct{
			Code: 10018,
			Msg:  global.I18nMap["10018"],
		}
		return failStruct
	}
	// 生成新的Token
	token := utils.CreateToken(c, userId, 2, insterUserInfo.UserName)
	data := forms.RegLogReturn{
		ID:    userId,
		Name:  insterUserInfo.UserName,
		Token: token,
	}
	succStruct := response.ResStruct{
		Code: http.StatusOK,
		Msg:  global.I18nMap["2000"],
		Data: data,
	}
	return succStruct
}

// SerLogin 业务层：用户登录
// 参数：
//		loginForm：登录时需要的参数
//		c：gin.Context的指针
// 返回值：
//		response.ResStruct：响应的结构体
func SerLogin(loginForm forms.LoginForm, c *gin.Context) response.ResStruct {
	// 验证码
	userSet := global.Settings.UserInfo
	// 判断是否开启验证码登录
	if userSet.CaptchaLogin {
		if !store.Verify(loginForm.CaptchaId, loginForm.Captcha, true) {
			failStruct := response.ResStruct{
				Code: 10008,
				Msg:  global.I18nMap["10008"],
			}
			return failStruct
		}
	}
	// 查询用户是否存在
	userInfo, ok := dao.DaoFindUserInfoToUserName(loginForm.UserName)
	if !ok {
		failStruct := response.ResStruct{
			Code: 10013,
			Msg:  global.I18nMap["10013"],
		}
		return failStruct
	}
	// 判断密码是否正确
	pwdBool := utils.CheckPassword(userInfo.Password, loginForm.Password)
	if !pwdBool {
		failStruct := response.ResStruct{
			Code: 10015,
			Msg:  global.I18nMap["10015"],
		}
		return failStruct
	}
	// 生成新的Token
	token := utils.CreateToken(c, userInfo.ID, userInfo.Role, userInfo.UserName)
	data := forms.RegLogReturn{
		ID:    userInfo.ID,
		Name:  userInfo.UserName,
		Token: token,
	}
	succStruct := response.ResStruct{
		Code: http.StatusOK,
		Msg:  global.I18nMap["2001"],
		Data: data,
	}
	return succStruct
}

// SerLogout 业务层：用户登出
// 参数：
//		c：gin.Context的指针
// 返回值：
//		response.ResStruct：响应的结构体
func SerLogout(c *gin.Context) response.ResStruct {
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
			zap.S().Errorf("token MD5 Set Redis faild：%s", tokenStr)
		}
	}()

	succStruct := response.ResStruct{
		Code: http.StatusOK,
		Msg:  global.I18nMap["2002"],
	}
	return succStruct
}

// SerGetMyselfInfo 业务层：获取用户自己的信息
// 参数：
//		userId：用户ID参数的结构体
//		c：gin.Context的指针
// 返回值：
//		response.ResStruct：响应的结构体
func SerGetMyselfInfo(userId forms.IdForm, c *gin.Context) response.ResStruct {
	// 判断是本人，不能获取别人的用户信息
	tokenUserId, _ := c.Get("userId")
	if tokenUserId != userId.ID {
		failStruct := response.ResStruct{
			Code: 10014,
			Msg:  global.I18nMap["10014"],
		}
		return failStruct
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
	succStruct := response.ResStruct{
		Code: http.StatusOK,
		Data: data,
	}
	return succStruct
}

// SerGetUserList 业务层：获取用户列表
// 参数：
//		userListForm：查看用户列表时需要的参数
//		c：gin.Context的指针
// 返回值：
//		response.ResStruct：响应的结构体
func SerGetUserList(userListForm forms.UserListForm) response.ResStruct {
	// 获取数据
	page, pageSize := utils.PageZero(userListForm.Page, userListForm.PageSize)
	total, userList, err := dao.DaoGetUserList(page, pageSize)
	if err != nil {
		zap.S().Errorf("%s：%s", global.I18nMap["10004"], err)
		failStruct := response.ResStruct{
			Code: 10004,
			Msg:  global.I18nMap["10004"],
		}
		return failStruct
	}
	// 获取数据为空
	if total == 0 {
		failStruct := response.ResStruct{
			Code: 10005,
			Msg:  global.I18nMap["10005"],
		}
		return failStruct
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
	succStruct := response.ResStruct{
		Code: http.StatusOK,
		Data: data,
	}
	return succStruct
}

// SerModifyUserInfo 业务层：修改用户信息
// 参数：
//		userId：用户ID
//		modUserInfoForm：需要修改的信息
//		c：gin.Context的指针
// 返回值：
//		response.ResStruct：响应的结构体
func SerModifyUserInfo(userId forms.IdForm, modUserInfoForm forms.ModifyUserInfoForm, c *gin.Context) response.ResStruct {
	// 判断是否是本人，只能修改自己的信息
	tokenUserId, _ := c.Get("userId")
	if tokenUserId != userId.ID {
		failStruct := response.ResStruct{
			Code: 10003,
			Msg:  global.I18nMap["10003"],
		}
		return failStruct
	}
	// 定义models.User接收需要修改的字段和对应的值
	userMod := models.User{}

	// 如果有修改密码，需要对密码进行判断
	if modUserInfoForm.PasswordOld != "" && modUserInfoForm.Password != "" {
		// 判断旧密码是否正确
		userInfo, _ := dao.DaoFindUserInfoToId(userId.ID)
		pwdBool := utils.CheckPassword(userInfo.Password, modUserInfoForm.PasswordOld)
		if !pwdBool {
			failStruct := response.ResStruct{
				Code: 10019,
				Msg:  global.I18nMap["10019"],
			}
			return failStruct
		}
		// 判断旧密码与新密码是否一致
		if modUserInfoForm.PasswordOld == modUserInfoForm.Password {
			failStruct := response.ResStruct{
				Code: 10020,
				Msg:  global.I18nMap["10020"],
			}
			return failStruct
		}
		// 判断修改后的两个密码密码是否一致
		if modUserInfoForm.Password != modUserInfoForm.Password2 {
			failStruct := response.ResStruct{
				Code: 10016,
				Msg:  global.I18nMap["10016"],
			}
			return failStruct
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
		zap.S().Errorf("%s：%s", global.I18nMap["10003"], err)
		failStruct := response.ResStruct{
			Code: 10003,
			Msg:  global.I18nMap["10003"],
		}
		return failStruct
	}

	succStruct := response.ResStruct{
		Code: http.StatusOK,
		Msg:  global.I18nMap["2005"],
	}
	return succStruct
}

// SerDelUser 业务层：删除用户信息（需要权限）
// 参数：
//		userId：用户ID
// 返回值：
//		response.ResStruct：响应的结构体
func SerDelUser(userId forms.IdForm) response.ResStruct {
	userMod := models.User{ID: userId.ID}
	// 删除用户
	err := dao.DaoDelUserToPriKey(userMod)
	if err != nil {
		zap.S().Errorf("%s：%s", global.I18nMap["10002"], err)
		failStruct := response.ResStruct{
			Code: 10002,
			Msg:  global.I18nMap["10002"],
		}
		return failStruct
	}

	succStruct := response.ResStruct{
		Code: http.StatusOK,
		Msg:  global.I18nMap["2004"],
	}
	return succStruct
}
