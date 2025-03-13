package repository

import (
	"context"

	"github.com/fnoopv/amp/database/model"
	"gorm.io/gorm"
	"goyave.dev/goyave/v5/util/errors"
)

type Attachment struct {
	db *gorm.DB
}

func NewAttachment(db *gorm.DB) *Attachment {
	return &Attachment{
		db: db,
	}
}

func (at *Attachment) Create(ctx context.Context, att *model.Attachment) error {
	db := at.db.WithContext(ctx).Create(att)

	return errors.New(db.Error)
}

func (at *Attachment) Update(ctx context.Context, att *model.Attachment) error {
	db := at.db.WithContext(ctx).
		Model(&model.Attachment{ID: att.ID}).
		Select("BusinessKind", "BusinessID").
		Updates(att)

	return errors.New(db.Error)
}

func (at *Attachment) FindByID(ctx context.Context, id string) (*model.Attachment, error) {
	var att model.Attachment
	db := at.db.WithContext(ctx).Where("id = ?", id).First(&att)

	return &att, errors.New(db.Error)
}
