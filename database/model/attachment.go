package model

import (
	"github.com/dromara/carbon/v2"
	"github.com/guregu/null/v6"
)

// Attachment 附件表
type Attachment struct {
	// ID 唯一ID
	ID string `json:"id" gorm:"column:id;primaryKey"`
	// Name 原始文件名
	Name string `json:"name" gorm:"column:name;not null"`
	// Mime 文件类型
	Mime string `json:"mime" gorm:"column:mime;not null"`
	// Size 文件大小,字节
	Size int64 `json:"size" gorm:"column:size;not null"`
	// StoragePath 存储路径
	StoragePath string `json:"storage_path" gorm:"column:storage_path;not null"`
	// UploaderID 上传人ID
	UploaderID string `json:"uploader_id" gorm:"column:uploader_id;not null"`
	// UploadAt 上传时间
	UploadAt carbon.DateTime `json:"upload_at" gorm:"column:upload_at;not null"`
	// BusinessKind 业务类型
	BusinessKind null.String `json:"business_kind" gorm:"column:business_kind"`
	// BusinessID 业务ID
	BusinessID null.String `json:"business_id" gorm:"column:business_id"`
	// SHA256Sum SHA256校验值
	SHA256Sum string `json:"sha256_sum" gorm:"column:sha256_sum;not null"`
	// CreatedAt 创建时间
	CreatedAt carbon.DateTime `json:"created_at" gorm:"column:created_at;autoCreateTime;type:datetime"`
	// UpdatedAt 更新时间
	UpdatedAt carbon.DateTime `json:"updated_at" gorm:"column:updated_at;autoUpdateTime;type:datetime"`

	User User `gorm:"foreignKey:UploaderID"`
}

// TableName 附件表表名
func (Attachment) TableName() string {
	return "attachments"
}
