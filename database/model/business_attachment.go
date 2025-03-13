package model

// BusinessAttachment 附件业务关联表
type BusinessAttachment struct {
	ID uint `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	// BusinessType 业务类型
	BusinessType string `json:"business_type" gorm:"column:business_type;not null"`
	// BusinessID 业务ID
	BusinessID string `json:"business_id" gorm:"column:business_id;not null"`
	// AttachmentType 附件在业务中的类型
	AttachmentType string `json:"attachment_type" gorm:"column:attachment_type;not null"`
	// AttachmentID 关联附件ID
	AttachmentID string `json:"attachment_id" gorm:"column:attachment_id;not null"`
}

// TableName 附件业务关联表表名
func (BusinessAttachment) TableName() string {
	return "business_attachments"
}
