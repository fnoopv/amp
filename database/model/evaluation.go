package model

import (
	"github.com/dromara/carbon/v2"
	"gorm.io/gorm"
)

// Evaluation 测评记录
type Evaluation struct {
	// ID 唯一ID
	ID string `json:"id" gorm:"column:id;primaryKey"`
	// FillingID 所属的备案ID
	FillingID string `json:"filling_id" gorm:"column:filling_id;not null"`
	// CompletedAt 测评时间
	CompletedAt carbon.Date `json:"completed_at" gorm:"column:completed_at;not null"`
	// SerialNumber 测评编号
	SerialNumber string `json:"serial_number" gorm:"column:serial_number;not null"`
	// CreatedAt 创建时间
	CreatedAt carbon.DateTime `json:"created_at" gorm:"column:created_at;autoCreateTime;type:datetime"`
	// UpdatedAt 更新时间
	UpdatedAt carbon.DateTime `json:"updated_at" gorm:"column:updated_at;autoUpdateTime;type:datetime"`
	DeletedAt gorm.DeletedAt
}

// Network 备案台账表表名
func (Evaluation) TableName() string {
	return "evaluations"
}
