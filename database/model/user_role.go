package model

import (
	"github.com/dromara/carbon/v2"
)

// UserRole 用户角色表
type UserRole struct {
	// UserID 用户ID
	UserID string `json:"user_id" gorm:"column:user_id;not null;primaryKey"`
	// RoleID 角色ID
	RoleID string `json:"role_id" gorm:"column:role_id;not null:primaryKey"`
	// CreatedAt 创建时间
	CreatedAt carbon.DateTime `json:"created_at" gorm:"column:created_at;autoCreateTime;type:datetime"`
	// UpdatedAt 更新时间
	UpdatedAt carbon.DateTime `json:"updated_at" gorm:"column:updated_at;autoUpdateTime;type:datetime"`
}

// TableName 用户角色表表名
func (UserRole) TableName() string {
	return "user_roles"
}
