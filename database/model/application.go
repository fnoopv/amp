package model

import "github.com/dromara/carbon/v2"

// Application 应用表
type Application struct {
	// ID 唯一ID
	ID string `json:"id" gorm:"column:id;primaryKey"`
	// Name 名称
	Name string `json:"name" gorm:"column:name;not null"`
	// OrganizationID 所属组织ID
	OrganizationID string `json:"organization_id" gorm:"column:organization_id"`
	// Method 请求方法
	LaunchedAt *carbon.DateTime `json:"launched_at" gorm:"column:launched_at;default:null"`
	// Description 描述
	Description string `json:"description" gorm:"column:description;default:null"`
	// CreatedAt 创建时间
	CreatedAt carbon.DateTime `json:"created_at" gorm:"column:created_at;autoCreateTime;type:datetime"`
	// UpdatedAt 更新时间
	UpdatedAt carbon.DateTime `json:"updated_at" gorm:"column:updated_at;autoUpdateTime;type:datetime"`
}

// Application 应用表表名
func (Application) TableName() string {
	return "applications"
}
