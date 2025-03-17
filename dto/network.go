package dto

import (
	"github.com/dromara/carbon/v2"
)

type Network struct {
	// ID 唯一ID
	ID string `json:"id"`
	// Name 名称
	Name string `json:"name"`
	// OrganizationID 所属组织ID
	OrganizationID string `json:"organization_id,omitempty"`
	// Organization 所属组织
	Organization Organization `json:"organization,omitzero"`
	// FillingID 关联备案ID
	FillingID string `json:"filling_id,omitempty"`
	// Filling 关联备案
	Filling Filling `json:"filling,omitzero"`
	// Description 描述
	Description string `json:"description,omitempty"`
	// CreatedAt 创建时间
	CreatedAt carbon.DateTime `json:"created_at"`
	// UpdatedAt 更新时间
	UpdatedAt carbon.DateTime `json:"updated_at"`
}

type NetworkCreate struct {
	// Name 名称
	Name string `json:"name"`
	// OrganizationID 所属组织ID
	OrganizationID string `json:"organization_id"`
	// FillingID 关联备案ID
	FillingID string `json:"filling_id"`
	// Description 描述
	Description string `json:"description"`
}

type NetworkUpdate struct {
	// ID 唯一ID
	ID string `json:"id"`
	// Name 名称
	Name string `json:"name"`
	// OrganizationID 所属组织ID
	OrganizationID string `json:"organization_id"`
	// FillingID 关联备案ID
	FillingID string `json:"filling_id"`
	// Description 描述
	Description string `json:"description"`
}

type NetworkDelete struct {
	IDs []string `json:"ids"`
}
