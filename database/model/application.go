package model

import (
	"github.com/dromara/carbon/v2"
	"github.com/guregu/null/v6"
	"gorm.io/gorm"
)

// Application 应用表
type Application struct {
	// ID 唯一ID
	ID string `json:"id" gorm:"column:id;primaryKey"`
	// Name 名称
	Name string `json:"name" gorm:"column:name;not null"`
	// OrganizationID 所属组织ID
	OrganizationID null.String `json:"organization_id" gorm:"column:organization_id;default:null"`
	// LaunchedAt 上线时间
	LaunchedAt *carbon.Date `json:"launched_at" gorm:"column:launched_at;default:null"`
	// Description 描述
	Description string `json:"description" gorm:"column:description;default:null"`
	// CreatedAt 创建时间
	CreatedAt carbon.DateTime `json:"created_at" gorm:"column:created_at;autoCreateTime;type:datetime"`
	// UpdatedAt 更新时间
	UpdatedAt carbon.DateTime `json:"updated_at" gorm:"column:updated_at;autoUpdateTime;type:datetime"`
	DeletedAt gorm.DeletedAt

	Organization Organization `gorm:"foreignKey:OrganizationID"`
}

// Application 应用表表名
func (Application) TableName() string {
	return "applications"
}
