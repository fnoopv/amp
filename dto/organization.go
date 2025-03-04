package dto

import "github.com/dromara/carbon/v2"

// Organization 组织
type Organization struct {
	// ID 唯一ID
	ID string `json:"id"`
	// ParentID 上级组织id, 为空时是顶级组织
	ParentID string `json:"parent_id,omitempty"`
	// Name 组织名称
	Name string `json:"name"`
	// Kind 组织类型 company-公司,department-部门
	Kind string `json:"kind"`
	// Order 组织排序
	Order int `json:"order,omitempty"`
	// CreatedAt 创建时间
	CreatedAt carbon.DateTime `json:"created_at"`
	// UpdatedAt 更新时间
	UpdatedAt carbon.DateTime `json:"updated_at"`
}

// OrganizationIndex 组织列表
type OrganizationIndex struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

// OrganizationCreate 新增组织架构
type OrganizationCreate struct {
	// ParentID 上级组织id, 为空时是顶级组织
	ParentID string `json:"parent_id,omitempty"`
	// Name 组织名称
	Name string `json:"name"`
	// Kind 组织类型 company-公司,department-部门
	Kind string `json:"kind"`
	// Order 组织排序
	Order int `json:"order,omitempty"`
}

// OrganizationUpdate 组织更新
type OrganizationUpdate struct {
	// ParentID 上级组织id, 为空时是顶级组织
	ParentID string `json:"parent_id,omitempty"`
	// Name 组织名称
	Name string `json:"name"`
	// Kind 组织类型 company-公司,department-部门
	Kind string `json:"kind"`
	// Order 组织排序
	Order int `json:"order,omitempty"`
}
