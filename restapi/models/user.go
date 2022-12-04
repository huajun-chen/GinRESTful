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
	Name      string         `json:"name" gorm:"size:8;comment:'名字'"`
	Password  string         `json:"password" gorm:"size:20;comment:'密码'"`
	HeadUrl   string         `json:"head_url" gorm:"size:512;comment:'网址'"`
	Birthday  time.Time      `json:"birthday" gorm:"comment:'生日'"`
	Address   string         `json:"address" gorm:"size:256;comment:'地址'"`
	Desc      string         `json:"desc" gorm:"size:256;comment:'描述'"`
	Gender    int            `json:"gender" gorm:"size:4;comment:'性别'"`
	Role      int            `json:"role" gorm:"size:4;comment:'角色'"`
	Mobile    string         `json:"mobile" gorm:"size:11;comment:'电话'"`
}
