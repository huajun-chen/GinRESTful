package controller

import (
	"GinRESTful/restapi/dao"
	"GinRESTful/restapi/forms"
	"GinRESTful/restapi/global"
	"GinRESTful/restapi/models"
	"GinRESTful/restapi/response"
	"GinRESTful/restapi/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// NeedsUserInfo 定义结构体存储需要返回的用户数据
// values里的数据除了ID是int类型，其他的都是字符串类型，返回的字段不一定全部都是数据库的字段
// 也有可能是数据库字段之间计算之后的值，所以返回的数据结构体单独定义
type NeedsUserInfo struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	UserName  string `json:"user_name"`
	Gender    string `json:"gender"`
	Desc      string `json:"desc"`
	Role      string `json:"role"`
	Mobile    string `json:"mobile"`
	Email     string `json:"email"`
}

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
	_, ok := dao.FindUserInfo(registerForm.UserName)
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
	userId, err := dao.RegisterUser(insterUserInfo)
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
		Code: global.RegisterSuccCode,
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
		Code: global.LoginSuccCode,
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
	// 过滤用户列表，只返回需要的数据
	var values []NeedsUserInfo
	for _, u := range userList {
		needUserInfo := NeedsUserInfo{
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
