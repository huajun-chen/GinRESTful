package forms

// LoginForm 用户登录参数
type LoginForm struct {
	UserName  string `json:"user_name" binding:"required,min=3,max=16"` // 用户名
	Password  string `json:"password" binding:"required,min=8,max=64"`  // 密码
	Captcha   string `json:"captcha" binding:"required,len=5"`          // 验证码
	CaptchaId string `json:"captcha_id" binding:"required"`             //验证码ID
}

// UserListForm 用户列表参数
type UserListForm struct {
	PageForm // 页数，每页数量
}

// RegisterForm 用户注册
type RegisterForm struct {
	LoginForm        // 用户登录需要的参数
	Password2 string `json:"password2" binding:"required,min=8,max=64"` // 重复密码
}

// ModifyUserInfoForm 修改用户信息参数
type ModifyUserInfoForm struct {
	PasswordOld string `json:"password_old" binding:"omitempty,min=8,max=64"` // 旧密码
	Password    string `json:"password" binding:"omitempty,min=8,max=64"`     // 新密码
	Password2   string `json:"password2" binding:"omitempty,min=8,max=64"`    // 新密码2
	Gender      int    `json:"gender" binding:"omitempty,oneof=1 2 3"`        // 性别
	Desc        string `json:"desc" binding:"omitempty,max=256"`              // 描述
	Mobile      string `json:"mobile" binding:"omitempty,len=11"`             // 电话
	Email       string `json:"email" binding:"omitempty,email"`               // 邮箱
}

// UserInfoReturn 定义结构体存储需要返回的用户数据
// values里的数据除了ID是int类型，其他的都是字符串类型，返回的字段不一定全部都是数据库的字段
// 也有可能是数据库字段之间计算之后的值，所以返回的数据结构体单独定义
type UserInfoReturn struct {
	ID        uint   `json:"id"`         // 用户ID
	CreatedAt string `json:"created_at"` // 创建时间
	UserName  string `json:"user_name"`  // 用户名
	Gender    string `json:"gender"`     // 性别
	Desc      string `json:"desc"`       // 描述
	Role      string `json:"role"`       // 角色
	Mobile    string `json:"mobile"`     // 电话
	Email     string `json:"email"`      // 邮箱
}

// RegLogReturn 注册和登录接口返回的数据
type RegLogReturn struct {
	ID    uint   `json:"id"`    // 用户ID
	Name  string `json:"name"`  // 用户名
	Token string `json:"token"` // 用户Token
}

// UserListReturn 用户列表信息
type UserListReturn struct {
	Total  int              `json:"total"`  // 总数
	Values []UserInfoReturn `json:"values"` // 用户列表
}
