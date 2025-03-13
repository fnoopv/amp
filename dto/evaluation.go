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
}

type EvaluationDelete struct {
	IDs []string `json:"ids"`
}
