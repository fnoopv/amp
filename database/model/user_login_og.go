package model

// UserLoginLog 用户登录记录表
type UserLoginLog struct {
	// ID 唯一ID
	ID string `json:"id" gorm:"column:id;primaryKey"`
	// UserID 登录用户ID
	UserID string `json:"user_id" gorm:"column:user_id;not null"`
	// LoginAt 登录时间
	LoginAt int64 `json:"login_at" gorm:"column:login_at;not null"`
	// IsSuccess 是否登录成功
	IsSuccess bool `gorm:"column:is_success;not null"`
	// IPAddress 登录IP
	IPAddress string `json:"ip_address" gorm:"column:ip_address;not null"`
	// UserAgent 浏览器/设备原始字符串
	UserAgent string `json:"user_agent" gorm:"column:user_agent;default:null"`
	// FailureReason 登录失败原因
	FailureReason string `json:"failure_reason" gorm:"column:failure_reason;defailt:null"`
	// CreatedAt 创建时间
	CreatedAt int64 `json:"created_at" gorm:"column:created_at"`
	// UpdatedAt 更新时间
	UpdatedAt int64 `json:"updated_at" gorm:"column:updated_at"`
}

// TableName 用户表表名
func (UserLoginLog) TableName() string {
	return "user_login_logs"
}
