package model

import (
	"github.com/dromara/carbon/v2"
)

// User 用户
// TODO: 关联角色
type User struct {
	// ID 唯一ID
	ID string `json:"id" gorm:"column:id;primaryKey"`
	// Email 邮箱
	Email string `json:"email" gorm:"column:email"`
	// NickName 显示名
	NickName string `json:"nick_name" gorm:"column:nick_name;not null"`
	// UserName 用户名
	UserName string `json:"username" gorm:"column:username;not null"`
	// Password 密码
	Password string `json:"password" gorm:"column:password;not null;type:char(60)"`
	// Status 账户状态 active-正常,inactive-未激活,disabled-禁用,banned-封禁
	Status string `json:"status" gorm:"column:status;not null"`
	// MFAKey 多因素认证密钥
	MFAKey string `json:"-" gorm:"column:mfa_key"`
	// IsMFAActive MFA是否已经设置
	IsMFAActive bool `json:"is_mfa_active" gorm:"column:is_mfa_active;not null;default:false"`
	// PasswordUpdatedAt 密码最后修改时间
	PasswordUpdatedAt *carbon.DateTime `json:"password_updated_at" gorm:"column:password_updated_at;type:datetime;default:null"`
	// CreatedAt 创建时间
	CreatedAt carbon.DateTime `json:"created_at" gorm:"column:created_at;autoCreateTime;type:datetime"`
	// UpdatedAt 更新时间
	UpdatedAt carbon.DateTime `json:"updated_at" gorm:"column:updated_at;autoUpdateTime;type:datetime"`
}

// TableName 用户表表名
func (User) TableName() string {
	return "users"
}
