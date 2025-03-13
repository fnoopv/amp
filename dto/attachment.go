package dto

import (
	"github.com/dromara/carbon/v2"
)

type Attachment struct {
	// ID 唯一ID
	ID string `json:"id"`
	// Name 原始文件名
	Name string `json:"name"`
	// Mime 文件类型
	Mime string `json:"mime"`
	// Size 文件大小,字节
	Size int64 `json:"size"`
	// SHA256Sum SHA256校验值
	SHA256Sum string `json:"sha256_sum"`
}

type AttachmentInternal struct {
	// ID 唯一ID
	ID string `json:"id"`
	// Name 原始文件名
	Name string `json:"name"`
	// Mime 文件类型
	Mime string `json:"mime"`
	// Size 文件大小,字节
	Size int64 `json:"size"`
	// StoragePath 存储路径
	StoragePath string `json:"storage_path"`
	// UploaderID 上传人ID
	UploaderID string `json:"uploader_id"`
	// UploadAt 上传时间
	UploadAt carbon.DateTime `json:"upload_at"`
	// SHA256Sum SHA256校验值
	SHA256Sum string `json:"sha256_sum"`
	// CreatedAt 创建时间
	CreatedAt carbon.DateTime `json:"created_at" gorm:"column:created_at;autoCreateTime;type:datetime"`
	// UpdatedAt 更新时间
	UpdatedAt carbon.DateTime `json:"updated_at" gorm:"column:updated_at;autoUpdateTime;type:datetime"`
}
type AttachmentCreate struct {
	// ID 唯一ID
	ID string `json:"id"`
	// Name 原始文件名
	Name string `json:"name"`
	// Mime 文件类型
	Mime string `json:"mime"`
	// Size 文件大小,字节
	Size int64 `json:"size"`
	// StoragePath 存储路径
	StoragePath string `json:"storage_path"`
	// UploaderID 上传人ID
	UploaderID string `json:"uploader_id"`
	// UploadAt 上传时间
	UploadAt carbon.DateTime `json:"upload_at"`
	// SHA256Sum SHA256校验值
	SHA256Sum string `json:"sha256_sum"`
}
