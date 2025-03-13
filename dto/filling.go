package dto

import (
	"github.com/dromara/carbon/v2"
)

type Filling struct {
	// ID 唯一ID
	ID string `json:"id"`
	// Name 名称
	Name string `json:"name"`
	// OrganizationID 所属组织ID
	OrganizationID string `json:"organization_id"`
	// KindPrimary 备案大类
	KindPrimary string `json:"kind_primary"`
	// KindSecondary 备案小类
	KindSecondary string `json:"kind_secondary,omitempty"`
	// SerialNumber 备案编号
	SerialNumber string `json:"serial_number"`
	// CompletedAt 备案时间
	CompletedAt carbon.Date `json:"completed_at"`
	// GradeLevel 等保等级
	GradeLevel string `json:"grade_level,omitempty"`
	// Description 描述
	Description string `json:"description,omitempty"`
	// CreatedAt 创建时间
	CreatedAt carbon.DateTime `json:"created_at"`
	// UpdatedAt 更新时间
	UpdatedAt carbon.DateTime `json:"updated_at"`

	Organization Organization `json:"organization,omitzero"`
	Evaluations  []Evaluation `json:"evaluations,omitzero,omitempty"`
}

type FillingCreate struct {
	// Name 名称
	Name string `json:"name"`
	// OrganizationID 所属组织ID
	OrganizationID string `json:"organization_id"`
	// KindPrimary 备案大类
	KindPrimary string `json:"kind_primary"`
	// KindSecondary 备案小类
	KindSecondary string `json:"kind_secondary,omitempty"`
	// SerialNumber 备案编号
	SerialNumber string `json:"serial_number"`
	// CompletedAt 备案时间
	CompletedAt carbon.Date `json:"completed_at"`
	// GradeLevel 等保等级
	GradeLevel string `json:"grade_level,omitempty"`
	// Description 描述
	Description string `json:"description,omitempty"`
}

type FillingUpdate struct {
	// ID 唯一ID
	ID string `json:"id"`
	// Name 名称
	Name string `json:"name"`
	// OrganizationID 所属组织ID
	OrganizationID string `json:"organization_id"`
	// KindPrimary 备案大类
	KindPrimary string `json:"kind_primary"`
	// KindSecondary 备案小类
	KindSecondary string `json:"kind_secondary,omitempty"`
	// SerialNumber 备案编号
	SerialNumber string `json:"serial_number"`
	// CompletedAt 备案时间
	CompletedAt carbon.Date `json:"completed_at"`
	// GradeLevel 等保等级
	GradeLevel string `json:"grade_level,omitempty"`
	// Description 描述
	Description string `json:"description,omitempty"`
}

type FillingDelete struct {
	IDs []string `json:"ids"`
}
