package dto

import (
	"github.com/dromara/carbon/v2"
)

type Application struct {
	// ID 唯一ID
	ID string `json:"id"`
	// Name 名称
	Name string `json:"name"`
	// OrganizationID 所属组织ID
	OrganizationID string `json:"organization_id,omitempty"`
	// Method 请求方法
	LaunchedAt *carbon.Date `json:"launched_at,omitempty"`
	// Description 描述
	Description string `json:"description,omitempty"`
	// CreatedAt 创建时间
	CreatedAt carbon.DateTime `json:"created_at"`
	// UpdatedAt 更新时间
	UpdatedAt carbon.DateTime `json:"updated_at"`

	Organization Organization `json:"organization,omitzero"`
	Fillings     []Filling    `json:"fillings,omitempty"`
}

type ApplicationCreate struct {
	// Name 名称
	Name string `json:"name"`
	// OrganizationID 所属组织ID
	OrganizationID string `json:"organization_id"`
	// Method 请求方法
	LaunchedAt *carbon.Date `json:"launched_at"`
	// FillingIDs 关联备案ID
	FillingIDs []string `json:"filling_ids"`
	// Description 描述
	Description string `json:"description"`
}

type ApplicationUpdate struct {
	// ID 唯一ID
	ID string `json:"id"`
	// Name 名称
	Name string `json:"name"`
	// OrganizationID 所属组织ID
	OrganizationID string `json:"organization_id"`
	// Method 请求方法
	LaunchedAt *carbon.Date `json:"launched_at"`
	// FillingIDs 关联备案ID
	FillingIDs []string `json:"filling_ids"`
	// Description 描述
	Description string `json:"description"`
}

type ApplicationDelete struct {
	IDs []string `json:"ids"`
}
