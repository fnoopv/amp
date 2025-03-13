package model

import (
	"github.com/dromara/carbon/v2"
	"github.com/guregu/null/v6"
	"gorm.io/gorm"
)

// Filling 备案台账
type Filling struct {
	// ID 唯一ID
	ID string `json:"id" gorm:"column:id;primaryKey"`
	// Name 名称
	Name string `json:"name" gorm:"column:name;not null"`
	// OrganizationID 所属组织ID
	OrganizationID null.String `json:"organization_id" gorm:"column:organization_id"`
	// KindPrimary 备案大类
	KindPrimary string `json:"kind_primary" gorm:"column:kind_primary;not null"`
	// KindSecondary 备案小类
	KindSecondary null.String `json:"kind_secondary" gorm:"column:kind_secondary"`
	// SerialNumber 备案编号
	SerialNumber string `json:"serial_number" gorm:"column:serial_number;not null"`
	// CompletedAt 备案时间
	CompletedAt carbon.Date `json:"completed_at" gorm:"column:completed_at;not null"`
	// GradeLevel 等保等级
	GradeLevel string `json:"grade_level" gorm:"column:grade_level"`
	// Description 描述
	Description string `json:"description" gorm:"column:description;default:null"`
	// CreatedAt 创建时间
	CreatedAt carbon.DateTime `json:"created_at" gorm:"column:created_at;autoCreateTime;type:datetime"`
	// UpdatedAt 更新时间
	UpdatedAt carbon.DateTime `json:"updated_at" gorm:"column:updated_at;autoUpdateTime;type:datetime"`
	DeletedAt gorm.DeletedAt

	Organization Organization `gorm:"foreignKey:OrganizationID"`
	Evaluations  []Evaluation `gorm:"foreignKey:FillingID"`
}

// Network 备案台账表表名
func (Filling) TableName() string {
	return "fillings"
}
