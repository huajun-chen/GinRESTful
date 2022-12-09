package models

import (
	"gorm.io/gorm"
	"time"
)

// User user结构体
type User struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	UserName  string         `json:"user_name" gorm:"size:16;comment:'用户名'"`
	Password  string         `json:"password" gorm:"size:64;comment:'密码'"`
	Gender    int            `json:"gender" gorm:"size:4;default:3;comment:'性别（1：男；2：女；3：未知）'"`
	Desc      string         `json:"desc" gorm:"size:256;default:'这个人很懒，什么都没留下...';comment:'描述'"`
	Role      int            `json:"role" gorm:"size:4;default:2;comment:'角色（1：管理员；2：普通用户）'"`
	Mobile    string         `json:"mobile" gorm:"size:11;comment:'电话'"`
	Email     string         `json:"email" gorm:"size:32;comment:'邮箱'"`
}

// TableName 自定义表名
// 参数：
//		无
// 返回值：
//		string：表名
func (User) TableName() string {
	return "user"
}
