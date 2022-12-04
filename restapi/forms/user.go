package forms

// LoginForm 用户名+密码登录
type LoginForm struct {
	UserName string `json:"user_name" binding:"required,min=3,max=20"` // 用户名
	Password string `json:"password" binding:"required,min=8,max=64"`  // 密码
}

// UserListForm 用户列表参数
type UserListForm struct {
	Page     int `form:"page"`      // 页数，第几页
	PageSize int `form:"page_size"` // 每页的数量
}
