package dto

import "github.com/dromara/carbon/v2"

// UserLoginLog 用户登录日志
type UserLoginLog struct {
	// ID 唯一ID
	ID string `json:"id"`
	// UserID 登录用户ID
	UserID string `json:"user_id"`
	// LoginAt 登录时间
	LoginAt carbon.DateTime `json:"login_at"`
	// IsSuccess 登录是否成功
	IsSuccess bool `json:"is_success"`
	// IPAddress 登录IP
	IPAddress string `json:"ip_address"`
	// UserAgent 浏览器/设备原始字符串
	UserAgent string `json:"user_agent,omitempty"`
	// FailureReason 登录失败原因
	FailureReason string `json:"failure_reason,omitempty"`
	// CreatedAt 创建时间
	CreatedAt carbon.DateTime `json:"created_at"`
	// UpdatedAt 更新时间
	UpdatedAt carbon.DateTime `json:"updated_at"`
}
