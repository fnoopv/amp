package model

import (
	"github.com/dromara/carbon/v2"
	"github.com/guregu/null/v6"
	"gorm.io/gorm"
)

// Domain 域名表
type Domain struct {
	// ID 唯一ID
	ID string `json:"id" gorm:"column:id;primaryKey"`
	// Domain 域名
	Domain string `json:"domain" gorm:"column:domain;not null"`
	// OrganizationID 所属组织ID
	OrganizationID null.String `json:"organization_id" gorm:"column:organization_id;default:null"`
	// Description 描述
	Description string `json:"description" gorm:"column:description;default:null"`
	// CreatedAt 创建时间
	CreatedAt carbon.DateTime `json:"created_at" gorm:"column:created_at;autoCreateTime;type:datetime"`
	// UpdatedAt 更新时间
	UpdatedAt carbon.DateTime `json:"updated_at" gorm:"column:updated_at;autoUpdateTime;type:datetime"`
	DeletedAt gorm.DeletedAt

	Organization Organization `gorm:"foreignKey:OrganizationID"`
	Fillings     []Filling    `gorm:"many2many:domain_fillings;foreignKey:ID;joinForeignKey:DomainID;References:ID;joinReferences:FillingID"`
}

// TableName 域名表表名
func (Domain) TableName() string {
	return "domains"
}
