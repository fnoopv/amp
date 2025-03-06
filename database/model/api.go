package model

import "github.com/dromara/carbon/v2"

// API 接口表
type API struct {
	// ID 唯一ID
	ID string `json:"id" gorm:"column:id;primaryKey"`
	// FeatureID 所属功能ID
	FeatureID string `json:"feature_id" gorm:"column:feature_id;not null"`
	// Method 请求方法
	Method string `json:"method" gorm:"column:method;not null"`
	// Path 请求路径
	Path string `json:"path" gorm:"column:method;not null"`
	// Description 描述
	Description string `json:"description" gorm:"column:description;default:null"`
	// CreatedAt 创建时间
	CreatedAt carbon.DateTime `json:"created_at" gorm:"column:created_at;autoCreateTime;type:datetime"`
	// UpdatedAt 更新时间
	UpdatedAt carbon.DateTime `json:"updated_at" gorm:"column:updated_at;autoUpdateTime;type:datetime"`
}

// API 接口表表名
func (API) TableName() string {
	return "apis"
}
