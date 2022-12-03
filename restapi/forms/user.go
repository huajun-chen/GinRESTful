package forms

// LoginForm 用户名+密码登录
type LoginForm struct {
	UserName string `json:"user_name" binding:"required,min=3,max=20"` // 用户名
	Password string `json:"password" binding:"required,min=8,max=64"`  // 密码
}
