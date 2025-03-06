package model

import "github.com/dromara/carbon/v2"

// Role 角色表
type Role struct {
	// ID 唯一ID
	ID string `json:"id" gorm:"column:id;primaryKey"`
	// Name 角色名称
	Name string `json:"name" gorm:"column:name;not null"`
	// Description 描述
	Description string `json:"description" gorm:"column:description;default:null"`
	// IsBuiltin 是否系统内置,系统内置不允许更新和删除
	IsBuiltin bool `json:"is_builtin" gorm:"column:is_builtin;not null;default:false"`
	// CreatedAt 创建时间
	CreatedAt carbon.DateTime `json:"created_at" gorm:"column:created_at;autoCreateTime;type:datetime"`
	// UpdatedAt 更新时间
	UpdatedAt carbon.DateTime `json:"updated_at" gorm:"column:updated_at;autoUpdateTime;type:datetime"`

	Users    []User    `json:"users" gorm:"many2many:user_users;foreignKey:ID;joinForeignKey:RoleID;References:ID;joinReferences:UserID"`
	Features []Feature `json:"eatures" gorm:"many2many:feature_role_permissions;foreignKey:ID;joinForeignKey:RoleID;References:ID;joinReferences:FeatureID"`
}

// Role 角色表表名
func (Role) TableName() string {
	return "roles"
}
