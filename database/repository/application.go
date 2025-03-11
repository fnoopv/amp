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

func (ap *Application) Paginate(ctx context.Context, request *filter.Request) (*database.Paginator[*model.Application], error) {
	apps := []*model.Application{}

	paginator, err := filter.Scope(session.DB(ctx, ap.db).Preload("Organization"), request, &apps)

	return paginator, errors.New(err)
}

func (ap *Application) Create(ctx context.Context, app *model.Application) error {
	db := ap.db.Create(app)

	return errors.New(db.Error)
}

func (ap *Application) Update(ctx context.Context, app *model.Application) error {
	db := ap.db.Model(&model.Application{ID: app.ID}).Select("Name", "OrganizationID", "LaunchedAt", "Description").Updates(app)

	return errors.New(db.Error)
}

func (ap *Application) Delete(ctx context.Context, ids []string) error {
	db := ap.db.Where("id in ?", ids).Delete(&model.Application{})

	return errors.New(db.Error)
}
