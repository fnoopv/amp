package repository

import (
	"context"

	"github.com/fnoopv/amp/database/model"
	"gorm.io/gorm"
	"goyave.dev/filter"
	"goyave.dev/goyave/v5/database"
	"goyave.dev/goyave/v5/util/errors"
	"goyave.dev/goyave/v5/util/session"
)

type Network struct {
	db *gorm.DB
}

// NewNetwork 创建新的存储库
func NewNetwork(db *gorm.DB) *Network {
	return &Network{
		db: db,
	}
}

// Paginate 返回分页器
func (ne *Network) Paginate(ctx context.Context, request *filter.Request) (*database.Paginator[*model.Network], error) {
	settings := &filter.Settings[*model.Network]{
		DisableJoin:   true,
		DisableFields: true,
		// 搜索设置
		FieldsSearch:   []string{"name"},
		SearchOperator: filter.Operators["$cout"],

		// 排序设置
		DefaultSort: []*filter.Sort{{Field: "updated_at", Order: filter.SortDescending}},
	}
	networks := []*model.Network{}

	paginator, err := settings.Scope(
		session.DB(ctx, ne.db).Preload("Organization").Preload("Filling"),
		request,
		&networks,
	)

	return paginator, errors.New(err)
}

// Create 创建
func (ne *Network) Create(ctx context.Context, modelRecord *model.Network) error {
	db := session.DB(ctx, ne.db).Omit("Filling").Create(modelRecord)

	return errors.New(db.Error)
}

// Update 更新
func (ne *Network) Update(ctx context.Context, modelRecord *model.Network) error {
	db := session.DB(ctx, ne.db).
		Model(&model.Network{ID: modelRecord.ID}).
		Select(
			"Name",
			"OrganizationID",
			"FillingID",
			"Description",
		).
		Updates(modelRecord)

	return errors.New(db.Error)
}

// Delete 删除
func (ne *Network) Delete(ctx context.Context, ids []string) error {
	db := ne.db.WithContext(ctx).Where("id in ?", ids).Delete(&model.Network{})

	return errors.New(db.Error)
}

// FindByIDs 根据ID获取所有备案
func (ne *Network) FindByIDs(ctx context.Context, ids []string) ([]*model.Network, error) {
	networks := []*model.Network{}

	db := session.DB(ctx, ne.db).Where("id in ?", ids).Find(&networks)

	return networks, errors.New(db.Error)
}

// Option 获取所有备案
func (ne *Network) Option(ctx context.Context) ([]*model.Network, error) {
	networks := []*model.Network{}
	db := ne.db.WithContext(ctx).Find(&networks)

	return networks, errors.New(db.Error)
}
