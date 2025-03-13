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

type Evaluation struct {
	db *gorm.DB
}

// Evaluation 创建新的存储库
func NewEvaluation(db *gorm.DB) *Evaluation {
	return &Evaluation{
		db: db,
	}
}

// Paginate 返回分页器
func (ev *Evaluation) Paginate(ctx context.Context, request *filter.Request) (
	*database.Paginator[*model.Evaluation],
	error,
) {
	evaluations := []*model.Evaluation{}

	paginator, err := filter.Scope(session.DB(ctx, ev.db), request, &evaluations)

	return paginator, errors.New(err)
}

// Create 创建
func (ev *Evaluation) Create(ctx context.Context, evaluation *model.Evaluation) error {
	db := ev.db.WithContext(ctx).Create(evaluation)
	return errors.New(db.Error)
}

// Update 更新
func (ev *Evaluation) Update(ctx context.Context, evaluation *model.Evaluation) error {
	db := ev.db.WithContext(ctx).
		Model(&model.Evaluation{ID: evaluation.ID}).
		Select("FillingID", "CompletedAt", "SerialNumber").
		Updates(evaluation)

	return errors.New(db.Error)
}

// Delete 删除
func (ev *Evaluation) Delete(ctx context.Context, ids []string) error {
	db := ev.db.WithContext(ctx).Where("id in ?", ids).Delete(&model.Evaluation{})

	return errors.New(db.Error)
}
