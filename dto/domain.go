package dto

import (
	"github.com/dromara/carbon/v2"
)

type Domain struct {
	// ID 唯一ID
	ID string `json:"id"`
	// Domain 域名
	Domain string `json:"domain"`
	// OrganizationID 所属组织ID
	OrganizationID string `json:"organization_id,omitempty"`
	// Organization 所属组织
	Organization Organization `json:"Organization,omitzero"`
	// FillingIDs 关联备案ID
	FillingIDs []string `json:"filling_ids,omitempty"`
	// Fillings 关联备案
	Fillings []Filling `json:"fillings,omitempty"`
	// Description 描述
	Description string `json:"description,omitempty"`
	// CreatedAt 创建时间
	CreatedAt carbon.DateTime `json:"created_at"`
	// UpdatedAt 更新时间
	UpdatedAt carbon.DateTime `json:"updated_at"`
}

type DomainCreate struct {
	// Domain 域名
	Domain string `json:"domain"`
	// OrganizationID 所属组织ID
	OrganizationID string `json:"organization_id"`
	// FillingIDs 关联备案ID
	FillingIDs []string `json:"filling_ids"`
	// Description 描述
	Description string `json:"description"`
}

type DomainUpdate struct {
	// ID 唯一ID
	ID string `json:"id"`
	// Domain 域名
	Domain string `json:"domain"`
	// OrganizationID 所属组织ID
	OrganizationID string `json:"organization_id"`
	// FillingIDs 关联备案ID
	FillingIDs []string `json:"filling_ids"`
	// Description 描述
	Description string `json:"description"`
}

type DomainDelete struct {
	IDs []string `json:"ids"`
}
