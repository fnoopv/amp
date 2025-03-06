package model

import "github.com/dromara/carbon/v2"

// Menu 菜单表
type Menu struct {
	// ID 唯一ID
	ID string `json:"id" gorm:"column:id;primaryKey"`
	// ParentID 上级菜单ID
	ParentID string `json:"parent_id" gorm:"column:parent_id;default null"`
	// Icon 图标
	Icon string `json:"icon" gorm:"column:icon;not null"`
	// Path 路径
	Path string `json:"path" gorm:"column:path;not null"`
	// Order 排序, 越小越靠前
	Order int `json:"order" gorm:"column:order;default null"`
	// IsHidden 是否隐藏
	IsHidden string `json:"is_hidden" gorm:"column:is_hidden;not null;default:false"`
	// CreatedAt 创建时间
	CreatedAt carbon.DateTime `json:"created_at" gorm:"column:created_at;autoCreateTime;type:datetime"`
	// UpdatedAt 更新时间
	UpdatedAt carbon.DateTime `json:"updated_at" gorm:"column:updated_at;autoUpdateTime;type:datetime"`
}

// Menu 菜单表表名
func (Menu) TableName() string {
	return "menus"
}
