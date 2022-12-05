package forms

// LoginForm 用户名+密码登录
type LoginForm struct {
	UserName  string `json:"user_name" binding:"required,min=3,max=20"` // 用户名
	Password  string `json:"password" binding:"required,min=8,max=64"`  // 密码
	Captcha   string `json:"captcha" binding:"required,len=5"`          // 验证码
	CaptchaId string `json:"captcha_id" binding:"required"`             //验证码ID
}

// UserListForm 用户列表参数
type UserListForm struct {
	PageForm
}

// RegisterForm 用户注册
type RegisterForm struct {
	LoginForm
	Password2 string `json:"password2" binding:"required,min=8,max=64"` // 重复密码
}
