package repository

import (
	"context"

	"github.com/fnoopv/amp/database/model"
	"gorm.io/gorm"
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

// FindByFillingID 根据备案ID查找搜索测评记录
func (ev *Evaluation) FindByFillingID(ctx context.Context, fillingID string) (
	[]*model.Evaluation,
	error,
) {
	evaluations := []*model.Evaluation{}

	db := session.DB(ctx, ev.db).Where("filling_id = ?", fillingID).Find(&evaluations)

	return evaluations, errors.New(db.Error)
}

// Create 创建
func (ev *Evaluation) Create(ctx context.Context, evaluation *model.Evaluation) error {
	db := session.DB(ctx, ev.db).Create(evaluation)
	return errors.New(db.Error)
}

// Update 更新
func (ev *Evaluation) Update(ctx context.Context, evaluation *model.Evaluation) error {
	db := session.DB(ctx, ev.db).
		Model(&model.Evaluation{ID: evaluation.ID}).
		Select("FillingID", "CompletedAt", "SerialNumber").
		Updates(evaluation)

	return errors.New(db.Error)
}

// Delete 删除
func (ev *Evaluation) Delete(ctx context.Context, ids []string) error {
	db := session.DB(ctx, ev.db).Where("id in ?", ids).Delete(&model.Evaluation{})

	return errors.New(db.Error)
}
