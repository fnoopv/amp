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

type UserLoginLog struct {
	db *gorm.DB
}

func NewUserLoginLog(db *gorm.DB) *UserLoginLog {
	return &UserLoginLog{
		db: db,
	}
}

func (us *UserLoginLog) Paginate(ctx *context.Context, request *filter.Request) (*database.Paginator[*model.UserLoginLog], error) {
	records := []*model.UserLoginLog{}

	paginator, err := filter.Scope(session.DB(*ctx, us.db), request, &records)

	return paginator, errors.New(err)
}

func (us *UserLoginLog) Create(ctx context.Context, record *model.UserLoginLog) error {
	db := us.db.Create(record)

	return errors.New(db.Error)
}
