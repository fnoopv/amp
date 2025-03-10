package model

import "github.com/dromara/carbon/v2"

// Feature 功能表
type Feature struct {
	// ID 唯一ID
	ID string `json:"id" gorm:"column:id;primaryKey"`
	// Name 功能名称
	Name string `json:"name" gorm:"column:name;not null"`
	// Description 描述
	Description string `json:"description" gorm:"column:description;default:null"`
	// CreatedAt 创建时间
	CreatedAt carbon.DateTime `json:"created_at" gorm:"column:created_at;autoCreateTime;type:datetime"`
	// UpdatedAt 更新时间
	UpdatedAt carbon.DateTime `json:"updated_at" gorm:"column:updated_at;autoUpdateTime;type:datetime"`

	Roles []Role `json:"roles" gorm:"many2many:feature_role_permissions;foreignKey:ID;joinForeignKey:FeatureID;References:ID;joinReferences:RoleID"`
}

// Feature 功能表表名
func (Feature) TableName() string {
	return "features"
}
