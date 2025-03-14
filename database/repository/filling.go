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
	fillings := []*model.Filling{}

	paginator, err := filter.Scope(
		session.DB(ctx, fi.db).Preload("Organization").Preload("Evaluations"),
		request,
		&fillings,
	)

	return paginator, errors.New(err)
}

// Create 创建
func (fi *Filling) Create(ctx context.Context, filling *model.Filling) error {
	db := fi.db.WithContext(ctx).Create(filling)
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
	db := fi.db.WithContext(ctx).Where("id in ?", ids).Delete(&model.Filling{})

	return errors.New(db.Error)
}
