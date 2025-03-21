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

type Filling struct {
	db *gorm.DB
}

// NewFilling 创建新的存储库
func NewFilling(db *gorm.DB) *Filling {
	return &Filling{
		db: db,
	}
}

// Paginate 返回分页器
func (fi *Filling) Paginate(ctx context.Context, request *filter.Request) (*database.Paginator[*model.Filling], error) {
	settings := &filter.Settings[*model.Filling]{
		DisableJoin:   true,
		DisableFields: true,
		// 搜索设置
		FieldsSearch:   []string{"name"},
		SearchOperator: filter.Operators["$cout"],

		// 排序设置
		DefaultSort: []*filter.Sort{{Field: "updated_at", Order: filter.SortDescending}},
	}
	fillings := []*model.Filling{}

	paginator, err := settings.Scope(
		session.DB(ctx, fi.db).
			Preload("Organization").
			Preload("Evaluations").
			Preload("Applications").
			Preload("Networks").
			Preload("Domains"),
		request,
		&fillings,
	)

	return paginator, errors.New(err)
}

// Create 创建
func (fi *Filling) Create(ctx context.Context, filling *model.Filling) error {
	db := session.DB(ctx, fi.db).Create(filling)
	return errors.New(db.Error)
}

// Update 更新
func (fi *Filling) Update(ctx context.Context, filling *model.Filling) error {
	db := fi.db.WithContext(ctx).
		Model(&model.Filling{ID: filling.ID}).
		Select(
			"Name",
			"OrganizationID",
			"KindPrimary",
			"KindSecondary",
			"SerialNumber",
			"CompletedAt",
			"GradeLevel",
			"Description",
		).
		Updates(filling)

	return errors.New(db.Error)
}

// Delete 删除
func (fi *Filling) Delete(ctx context.Context, ids []string) error {
	db := session.DB(ctx, fi.db).Where("id in ?", ids).Delete(&model.Filling{})

	return errors.New(db.Error)
}

// FindByIDs 根据ID获取所有备案
func (fi *Filling) FindByIDs(ctx context.Context, ids []string) ([]*model.Filling, error) {
	fillings := []*model.Filling{}

	db := session.DB(ctx, fi.db).Where("id in ?", ids).Find(&fillings)

	return fillings, errors.New(db.Error)
}

// Option 获取所有备案
func (fi *Filling) Option(ctx context.Context) ([]*model.Filling, error) {
	fillings := []*model.Filling{}
	db := fi.db.WithContext(ctx).Find(&fillings)

	return fillings, errors.New(db.Error)
}
