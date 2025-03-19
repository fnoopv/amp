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

type Domain struct {
	db *gorm.DB
}

// NewDomain 创建新的存储库
func NewDomain(db *gorm.DB) *Domain {
	return &Domain{
		db: db,
	}
}

// Paginate 返回分页器
func (ne *Domain) Paginate(ctx context.Context, request *filter.Request) (*database.Paginator[*model.Domain], error) {
	settings := &filter.Settings[*model.Domain]{
		DisableJoin:   true,
		DisableFields: true,
		// 搜索设置
		FieldsSearch:   []string{"domain"},
		SearchOperator: filter.Operators["$cout"],

		// 排序设置
		DefaultSort: []*filter.Sort{{Field: "updated_at", Order: filter.SortDescending}},
	}
	records := []*model.Domain{}

	paginator, err := settings.Scope(
		session.DB(ctx, ne.db).Preload("Organization").Preload("Fillings"),
		request,
		&records,
	)

	return paginator, errors.New(err)
}

// Create 创建
func (ne *Domain) Create(ctx context.Context, modelRecord *model.Domain) error {
	err := session.DB(ctx, ne.db).Omit("Fillings").Create(modelRecord).Error
	if err != nil {
		return errors.New(err)
	}

	// 更新备案关联
	err = session.DB(ctx, ne.db).
		Model(&model.Domain{ID: modelRecord.ID}).
		Omit("Fillings.*").
		Association("Fillings").
		Append(modelRecord.Fillings)

	return errors.New(err)
}

// Update 更新
func (ne *Domain) Update(ctx context.Context, modelRecord *model.Domain) error {
	err := session.DB(ctx, ne.db).
		Model(&model.Domain{ID: modelRecord.ID}).
		Select(
			"Dame",
			"OrganizationID",
			"Description",
		).
		Updates(modelRecord).Error
	if err != nil {
		return errors.New(err)
	}

	// 更新备案关联
	err = session.DB(ctx, ne.db).
		Model(&model.Domain{ID: modelRecord.ID}).
		Omit("Fillings.*").
		Association("Fillings").
		Replace(modelRecord.Fillings)

	return errors.New(err)
}

// Delete 删除
func (ne *Domain) Delete(ctx context.Context, ids []string) error {
	db := ne.db.WithContext(ctx).Where("id in ?", ids).Delete(&model.Domain{})

	return errors.New(db.Error)
}

// FindByIDs 根据ID获取所有域名
func (ne *Domain) FindByIDs(ctx context.Context, ids []string) ([]*model.Domain, error) {
	networks := []*model.Domain{}

	db := session.DB(ctx, ne.db).Where("id in ?", ids).Find(&networks)

	return networks, errors.New(db.Error)
}

// Option 获取枚举
func (ne *Domain) Option(ctx context.Context) ([]*model.Domain, error) {
	records := []*model.Domain{}
	db := ne.db.WithContext(ctx).Find(&records)

	return records, errors.New(db.Error)
}
