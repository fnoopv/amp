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

type Application struct {
	db *gorm.DB
}

func NewApplication(db *gorm.DB) *Application {
	return &Application{
		db: db,
	}
}

// Paginate 分页
func (ap *Application) Paginate(ctx context.Context, request *filter.Request) (
	*database.Paginator[*model.Application],
	error,
) {
	settings := &filter.Settings[*model.Application]{
		DisableJoin:   true,
		DisableFields: true,
		// 搜索设置
		FieldsSearch:   []string{"name"},
		SearchOperator: filter.Operators["$cout"],

		// 排序设置
		DefaultSort: []*filter.Sort{{Field: "updated_at", Order: filter.SortDescending}},
	}
	apps := []*model.Application{}

	paginator, err := settings.Scope(
		session.DB(ctx, ap.db).Preload("Organization").Preload("Fillings").Preload("Networks"),
		request,
		&apps,
	)

	return paginator, errors.New(err)
}

// Create 创建
func (ap *Application) Create(ctx context.Context, app *model.Application) error {
	err := session.DB(ctx, ap.db).Omit("Fillings").Omit("Networks").Create(app).Error
	if err != nil {
		return errors.New(err)
	}

	// 更新备案关联
	err = session.DB(ctx, ap.db).
		Model(&model.Application{ID: app.ID}).
		Omit("Fillings.*").
		Association("Fillings").
		Replace(app.Fillings)
	if err != nil {
		return errors.New(err)
	}

	// 更新网络关联
	err = session.DB(ctx, ap.db).
		Model(&model.Application{ID: app.ID}).
		Omit("Networks.*").
		Association("Networks").
		Replace(app.Networks)

	return errors.New(err)
}

// Update 更新
func (ap *Application) Update(ctx context.Context, app *model.Application) error {
	err := session.DB(ctx, ap.db).
		Model(&model.Application{ID: app.ID}).
		Select("Name", "OrganizationID", "LaunchedAt", "Description").
		Updates(app).Error
	if err != nil {
		return errors.New(err)
	}

	// 更新备案关联
	err = session.DB(ctx, ap.db).
		Model(&model.Application{ID: app.ID}).
		Omit("Fillings.*").
		Association("Fillings").
		Replace(app.Fillings)
	if err != nil {
		return errors.New(err)
	}

	// 更新网络关联
	err = session.DB(ctx, ap.db).
		Model(&model.Application{ID: app.ID}).
		Omit("Networks.*").
		Association("Networks").
		Replace(app.Networks)

	return errors.New(err)
}

// Delete 删除
func (ap *Application) Delete(ctx context.Context, ids []string) error {
	db := ap.db.WithContext(ctx).Where("id in ?", ids).Delete(&model.Application{})

	return errors.New(db.Error)
}
