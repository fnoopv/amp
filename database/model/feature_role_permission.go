package model

import "github.com/dromara/carbon/v2"

// FeatureRolePermission 角色功能权限关联表
type FeatureRolePermission struct {
	// RoleID 角色ID
	RoleID string `json:"role_id" gorm:"column:role_id;primaryKey"`
	// FeatureID 功能ID
	FeatureID string `json:"feature_id" gorm:"column:feature_id;primaryKey"`
	// Read 读权限, GET请求属于读权限
	Read bool `json:"read" gorm:"column:read;not null;default:false"`
	// Write 写权限
	Write bool `json:"write" gorm:"column:write;not null;default:false"`
	// CreatedAt 创建时间
	CreatedAt carbon.DateTime `json:"created_at" gorm:"column:created_at;autoCreateTime;type:datetime"`
	// UpdatedAt 更新时间
	UpdatedAt carbon.DateTime `json:"updated_at" gorm:"column:updated_at;autoUpdateTime;type:datetime"`
}

// RoleFeaturePermission 角色功能权限表表名
func (FeatureRolePermission) TableName() string {
	return "feature_role_permissions"
}
