package dto

import (
	"time"

	"github.com/dromara/carbon/v2"
)

// UserInternal 系统内用户信息
type UserInternal struct {
	User
	// Password 密码
	Password string `json:"password"`
	// IsMFAVerified 是否已经二次验证
	IsMFAVerified bool `json:"is_mfa_verified"`
}

type User struct {
	// ID 唯一ID
	ID string `json:"id"`
	// Email 邮箱
	Email string `json:"email,omitempty"`
	// NickName 显示名
	NickName string `json:"nick_name"`
	// UserName 用户名
	UserName string `json:"username"`
	// Status 账户状态 active-正常,inactive-未激活,disabled-禁用,banned-封禁
	Status string `json:"status"`
	// IsMFAActive MFA是否已经绑定
	IsMFAActive bool `json:"is_mfa_active"`
	// PasswordUpdatedAt 密码最后修改时间
	PasswordUpdatedAt *time.Time `json:"password_updated_at,omitempty"`
	// CreatedAt 创建时间
	CreatedAt carbon.DateTime `json:"created_at"`
	// UpdatedAt 更新时间
	UpdatedAt carbon.DateTime `json:"updated_at"`
}

// UserIndex 用户列表
type UserIndex struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

// UserCreate 用户创建
type UserCreate struct {
	// Email 邮箱
	Email string `json:"email"`
	// NickName 显示名
	NickName string `json:"nick_name"`
	// UserName 用户名
	UserName string `json:"username"`
	// Status 账户状态 active-正常,inactive-未激活,disabled-禁用,banned-封禁
	Status string `json:"status"`
}

// UserCreateResponse 用户创建成功响应
type UserCreateResponse struct {
	Password string `json:"password"`
}

// UserResetPasswordResponse 重置密码成功响应
type UserResetPasswordResponse struct {
	Password string `json:"password"`
}

// UserChangePassword 更改密码
type UserChangePassword struct {
	// OldPassword 原密码
	OldPassword string `json:"old_password"`
	// NewPassword 新密码
	NewPassword string `json:"new_password"`
	// ConfirmPassword 确认新密码
	ConfirmPassword string `json:"confirm_password"`
}

// UserUpdate 更新用户信息
type UserUpdate struct {
	// Email 邮箱
	Email string `json:"email"`
	// NickName 显示名
	NickName string `json:"nick_name"`
	// UserName 用户名
	UserName string `json:"username"`
}
