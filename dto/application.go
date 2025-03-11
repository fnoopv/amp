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
	OrganizationID string `json:"organization_id,omitempty,omitzero"`
	// Method 请求方法
	LaunchedAt *carbon.Date `json:"launched_at,omitempty,omitzero"`
	// Description 描述
	Description string `json:"description,omitempty,omitzero"`
	// CreatedAt 创建时间
	CreatedAt carbon.DateTime `json:"created_at"`
	// UpdatedAt 更新时间
	UpdatedAt carbon.DateTime `json:"updated_at"`
}

type ApplicationCreate struct {
	// Name 名称
	Name string `json:"name"`
	// OrganizationID 所属组织ID
	OrganizationID string `json:"organization_id,omitempty,omitzero"`
	// Method 请求方法
	LaunchedAt *carbon.Date `json:"launched_at,omitempty,omitzero"`
	// Description 描述
	Description string `json:"description,omitempty,omitzero"`
}

type ApplicationUpdate struct {
	// Name 名称
	Name string `json:"name"`
	// OrganizationID 所属组织ID
	OrganizationID string `json:"organization_id,omitempty,omitzero"`
	// Method 请求方法
	LaunchedAt *carbon.Date `json:"launched_at,omitempty,omitzero"`
	// Description 描述
	Description string `json:"description,omitempty,omitzero"`
}
