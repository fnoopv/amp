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
	OrganizationID string `json:"organization_id,omitempty"`
	// Organization 所属组织
	Organization Organization `json:"organization,omitzero"`
	// Evaluations 关联测评
	Evaluations []Evaluation `json:"evaluations,omitempty"`
	// Applications 关联应用
	Applications []Application `json:"applications,omitempty"`
	// Networks 关联网络
	Networks []Network `json:"networks,omitempty"`
	// Domains 关联域名
	Domains []Domain `json:"domains,omitempty"`
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
	// ProofAttachmentID 备案证明ID
	ProofAttachmentIDs []string `json:"proof_attachment_ids,omitempty"`
	// ProofAttachments 备案证明附件
	ProofAttachments []Attachment `json:"proof_attachments,omitempty"`
	// Description 描述
	Description string `json:"description,omitempty"`
	// CreatedAt 创建时间
	CreatedAt carbon.DateTime `json:"created_at"`
	// UpdatedAt 更新时间
	UpdatedAt carbon.DateTime `json:"updated_at"`
}

type FillingCreate struct {
	// Name 名称
	Name string `json:"name"`
	// OrganizationID 所属组织ID
	OrganizationID string `json:"organization_id"`
	// KindPrimary 备案大类
	KindPrimary string `json:"kind_primary"`
	// KindSecondary 备案小类
	KindSecondary string `json:"kind_secondary"`
	// SerialNumber 备案编号
	SerialNumber string `json:"serial_number"`
	// CompletedAt 备案时间
	CompletedAt carbon.Date `json:"completed_at"`
	// GradeLevel 等保等级
	GradeLevel string `json:"grade_level"`
	// ProofAttachmentID 备案证明ID
	ProofAttachmentIDs []string `json:"proof_attachment_ids"`
	// Description 描述
	Description string `json:"description"`
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
	KindSecondary string `json:"kind_secondary"`
	// SerialNumber 备案编号
	SerialNumber string `json:"serial_number"`
	// CompletedAt 备案时间
	CompletedAt carbon.Date `json:"completed_at"`
	// GradeLevel 等保等级
	GradeLevel string `json:"grade_level"`
	// ProofAttachmentID 备案证明ID
	ProofAttachmentIDs []string `json:"proof_attachment_ids"`
	// Description 描述
	Description string `json:"description"`
}

type FillingDelete struct {
	IDs []string `json:"ids"`
}
