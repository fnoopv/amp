package model

// UserLoginLog 用户登录记录表
type UserLoginLog struct {
	// ID 唯一ID
	ID string `gorm:"column:id,primaryKey"`
	// UserID 登录用户ID
	UserID string `gorm:"column:user_id,not null,index:idx_user_id_login_status"`
	// LoginAt 登录时间
	LoginAt int64 `gorm:"column:login_at,not null"`
	// LoginStatus 登录结果 success-登录成功,failure-失败
	LoginStatus string `gorm:"column:login_status,not null,index:idx_user_id_login_status"`
	// IPAddress 登录IP
	IPAddress string `gorm:"column:ip_address,not null"`
	// UserAgent 浏览器/设备原始字符串
	UserAgent string `gorm:"column:user_agent,not null"`
	// FailureReason 登录失败原因
	FailureReason string `gorm:"column:failure_reason"`
	// CreatedAt 创建时间
	CreatedAt int64 `gorm:"column:created_at,autoCreateTime:milli"`
	// UpdatedAt 更新时间
	UpdatedAt int64 `gorm:"column:updated_at,autoUpdateTime:milli"`
}

// TableName 用户表表名
func (UserLoginLog) TableName() string {
	return "user_login_logs"
}
