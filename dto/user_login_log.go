package dto

import "github.com/dromara/carbon/v2"

// UserLoginLog
type UserLoginLog struct {
	// ID 唯一ID
	ID string `json:"id"`
	// UserID 登录用户ID
	UserID string `json:"user_id"`
	// LoginAt 登录时间
	LoginAt carbon.DateTime `json:"login_at"`
	// LoginStatus 登录结果 success-登录成功,failure-失败
	LoginStatus string `json:"login_status"`
	// IPAddress 登录IP
	IPAddress string `json:"ip_address"`
	// UserAgent 浏览器/设备原始字符串
	UserAgent string `json:"user_agent"`
	// FailureReason 登录失败原因
	FailureReason string `json:"failure_reason,omitempty"`
	// CreatedAt 创建时间
	CreatedAt carbon.DateTime `json:"created_at"`
	// UpdatedAt 更新时间
	UpdatedAt carbon.DateTime `json:"updated_at"`
}
