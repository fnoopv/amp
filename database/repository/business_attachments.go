package repository

import (
	"context"

	"github.com/fnoopv/amp/database/model"
	"gorm.io/gorm"
	"goyave.dev/goyave/v5/util/errors"
	"goyave.dev/goyave/v5/util/session"
)

type BusinessAttachment struct {
	db *gorm.DB
}

func NewBusinessAttachment(db *gorm.DB) *BusinessAttachment {
	return &BusinessAttachment{
		db: db,
	}
}

// Create 新增附件业务关联
func (bu *BusinessAttachment) Create(ctx context.Context, bas []*model.BusinessAttachment) error {
	db := session.DB(ctx, bu.db).Create(bas)

	return errors.New(db.Error)
}

// Delete 删除附件业务关联
func (bu *BusinessAttachment) Delete(
	ctx context.Context,
	businessType, businessID string,
	attachmentTypes []string,
) error {
	db := session.DB(ctx, bu.db).
		Where("business_type = ?", businessType).
		Where("business_id = ?", businessID).
		Where("attachment_type in ?", attachmentTypes).
		Delete(&model.BusinessAttachment{})

	return errors.New(db.Error)
}

// FindAttachmentIDs 查找业务关联的附件ID
func (bu *BusinessAttachment) FindAttachmentIDs(
	ctx context.Context,
	businessType, businessID string,
	attachmentType []string,
) ([]string, error) {
	var ids []string

	db := session.DB(ctx, bu.db).
		Model(&model.BusinessAttachment{}).
		Where("business_type = ?", businessType).
		Where("business_id = ?", businessID).
		Where("attachment_type in ?", attachmentType).
		Pluck("attachment_id", &ids)

	return ids, errors.New(db.Error)
}
