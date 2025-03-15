package dto

import "github.com/dromara/carbon/v2"

type Evaluation struct {
	// ID 唯一ID
	ID string `json:"id"`
	// FillingID 所属的备案ID
	FillingID string `json:"filling_id"`
	// CompletedAt 测评时间
	CompletedAt carbon.Date `json:"completed_at"`
	// SerialNumber 测评编号
	SerialNumber string `json:"serial_number"`
	// EvaluationAttachmentIDs 测评报告ID
	EvaluationAttachmentIDs []string `json:"evaluation_attachment_ids,omitempty"`
	// EvaluationAttachments 测评报告
	EvaluationAttachments []*Attachment `json:"evaluation_attachments,omitempty"`
	// RepairAttachmentIDs 整改报告ID
	RepairAttachmentIDs []string `json:"repair_attachment_ids,omitempty"`
	// RepairAttachments 整改报告
	RepairAttachments []*Attachment `json:"repair_attachments,omitempty"`
	// CreatedAt 创建时间
	CreatedAt carbon.DateTime `json:"created_at"`
	// UpdatedAt 更新时间
	UpdatedAt carbon.DateTime `json:"updated_at"`
}

type EvaluationCreate struct {
	// FillingID 所属的备案ID
	FillingID string `json:"filling_id"`
	// CompletedAt 测评时间
	CompletedAt carbon.Date `json:"completed_at"`
	// SerialNumber 测评编号
	SerialNumber string `json:"serial_number"`
	// EvaluationAttachmentIDs 测评报告ID
	EvaluationAttachmentIDs []string `json:"evaluation_attachment_ids"`
	// RepairAttachmentIDs 整改报告ID
	RepairAttachmentIDs []string `json:"repair_attachment_ids"`
}

type EvaluationUpdate struct {
	// ID 唯一ID
	ID string `json:"id"`
	// FillingID 所属的备案ID
	FillingID string `json:"filling_id"`
	// CompletedAt 测评时间
	CompletedAt carbon.Date `json:"completed_at"`
	// SerialNumber 测评编号
	SerialNumber string `json:"serial_number"`
	// EvaluationAttachmentIDs 测评报告ID
	EvaluationAttachmentIDs []string `json:"evaluation_attachment_ids"`
	// RepairAttachmentIDs 整改报告ID
	RepairAttachmentIDs []string `json:"repair_attachment_ids"`
}

type EvaluationDelete struct {
	IDs []string `json:"ids"`
}
