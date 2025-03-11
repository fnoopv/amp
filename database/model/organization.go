package model

import (
	"github.com/dromara/carbon/v2"
	"github.com/guregu/null/v6"
	"gorm.io/gorm"
)

// Organization 组织架构
type Organization struct {
	// ID 唯一ID
	ID string `json:"id" gorm:"column:id;primaryKey;not null"`
	// ParentID 上级组织id, 为空时是顶级组织
	ParentID null.String `json:"parent_id" gorm:"column:parent_id;default:null"`
	// Name 组织名称
	Name string `json:"name" gorm:"column:name;not null"`
	// Kind 组织类型 company-公司,department-部门
	Kind string `json:"kind" gorm:"column:kind;not null"`
	// Order 组织排序
	Order int `json:"order" gorm:"column:order;default:null"`
	// Description 描述
	Description null.String `json:"description" gorm:"column:description"`
	// CreatedAt 创建时间
	CreatedAt carbon.DateTime `json:"created_at" gorm:"column:created_at;autoCreateTime;type:datetime"`
	// UpdatedAt 更新时间
	UpdatedAt carbon.DateTime `json:"updated_at" gorm:"column:updated_at;autoUpdateTime;type:datetime"`
	DeletedAt gorm.DeletedAt

	Children []Organization `json:"children" gorm:"foreignKey:ParentID;constraint:OnUpdate:OnDelete:CASCADE"`
}

// TableName 组织架构表表名
func (Organization) TableName() string {
	return "organizations"
}
