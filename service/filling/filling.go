package filling

import (
	"context"
	"fmt"

	"github.com/fnoopv/amp/database/model"
	"github.com/fnoopv/amp/dto"
	"github.com/fnoopv/amp/pkg/uid"
	"github.com/fnoopv/amp/service"
	"github.com/samber/lo"
	"goyave.dev/filter"
	"goyave.dev/goyave/v5/database"
	"goyave.dev/goyave/v5/util/errors"
	"goyave.dev/goyave/v5/util/session"
	"goyave.dev/goyave/v5/util/typeutil"
)

const (
	businessType   = "filling"
	attachmentType = "proof"
)

type fillingRepository interface {
	Paginate(ctx context.Context, request *filter.Request) (*database.Paginator[*model.Filling], error)
	Create(ctx context.Context, filling *model.Filling) error
	Update(ctx context.Context, filling *model.Filling) error
	Delete(ctx context.Context, ids []string) error
	Option(ctx context.Context) ([]*model.Filling, error)
}

type businessAttachmentRepository interface {
	Create(ctx context.Context, bas []*model.BusinessAttachment) error
	Delete(ctx context.Context, businessType, businessID string, attachmentType []string) error
	FindAttachmentIDs(ctx context.Context, businessType, businessID string, attachmentType []string) ([]string, error)
}

type attachmentRepository interface {
	FindByIDs(ctx context.Context, ids []string) ([]*model.Attachment, error)
}

type Service struct {
	session                      session.Session
	fillingRepository            fillingRepository
	businessAttachmentRepository businessAttachmentRepository
	attachmentRepository         attachmentRepository
}

func NewService(
	session session.Session,
	fillingRepository fillingRepository,
	businessAttachmentRepository businessAttachmentRepository,
	attachmentRepository attachmentRepository,
) *Service {
	return &Service{
		session:                      session,
		fillingRepository:            fillingRepository,
		businessAttachmentRepository: businessAttachmentRepository,
		attachmentRepository:         attachmentRepository,
	}
}

func (se *Service) Paginate(ctx context.Context, request *filter.Request) (
	*database.PaginatorDTO[*dto.Filling],
	error,
) {
	var fillings *database.PaginatorDTO[*dto.Filling]
	err := se.session.Transaction(ctx, func(ctx context.Context) error {
		var err error
		modelFillings, err := se.fillingRepository.Paginate(ctx, request)
		if err != nil {
			return errors.New(err)
		}

		fillings, err = typeutil.Convert[*database.PaginatorDTO[*dto.Filling]](modelFillings)
		if err != nil {
			return errors.New(err)
		}

		for _, filling := range fillings.Records {
			ids, err := se.businessAttachmentRepository.FindAttachmentIDs(
				ctx,
				businessType,
				filling.ID,
				[]string{attachmentType},
			)
			if err != nil {
				return errors.New(err)
			}
			filling.ProofAttachmentIDs = ids

			modelAttachments, err := se.attachmentRepository.FindByIDs(ctx, ids)
			if err != nil {
				return errors.New(err)
			}
			attachments, err := typeutil.Convert[*[]dto.Attachment](&modelAttachments)
			if err != nil {
				return errors.New(err)
			}

			filling.ProofAttachments = *attachments
		}

		return err
	})

	return fillings, errors.New(err)
}

func (se *Service) Create(ctx context.Context, filling *dto.FillingCreate) error {
	modelFilling := typeutil.Copy(&model.Filling{}, filling)

	id, err := uid.Generate()
	if err != nil {
		return errors.New(err)
	}
	modelFilling.ID = id

	err = se.session.Transaction(ctx, func(ctx context.Context) error {
		err := se.fillingRepository.Create(ctx, modelFilling)
		if err != nil {
			return errors.New(err)
		}

		if len(filling.ProofAttachmentIDs) > 0 {
			// 检查附件是否存在
			atts, err := se.attachmentRepository.FindByIDs(ctx, filling.ProofAttachmentIDs)
			if err != nil {
				return errors.New(err)
			}
			if len(atts) != len(filling.ProofAttachmentIDs) {
				var existsIDs []string
				for _, v := range atts {
					existsIDs = append(existsIDs, v.ID)
				}
				withoutIDs := lo.Without(filling.ProofAttachmentIDs, existsIDs...)
				return errors.New(fmt.Errorf("some attachment not exists,id: %v", withoutIDs))
			}

			// 设置附件关联
			var bas []*model.BusinessAttachment
			for _, ba := range filling.ProofAttachmentIDs {
				bas = append(bas, &model.BusinessAttachment{
					BusinessType:   businessType,
					BusinessID:     id,
					AttachmentType: attachmentType,
					AttachmentID:   ba,
				})
			}
			err = se.businessAttachmentRepository.Create(ctx, bas)
			if err != nil {
				return errors.New(err)
			}
		}

		return nil
	})

	return errors.New(err)
}

func (se *Service) Update(ctx context.Context, filling *dto.FillingUpdate) error {
	modelFilling := typeutil.Copy(&model.Filling{}, filling)

	err := se.session.Transaction(ctx, func(ctx context.Context) error {
		err := se.fillingRepository.Update(ctx, modelFilling)
		if err != nil {
			return errors.New(err)
		}

		err = se.businessAttachmentRepository.Delete(ctx, businessType, filling.ID, []string{attachmentType})
		if err != nil {
			return errors.New(err)
		}

		if len(filling.ProofAttachmentIDs) > 0 {
			// 检查附件是否存在
			atts, err := se.attachmentRepository.FindByIDs(ctx, filling.ProofAttachmentIDs)
			if err != nil {
				return errors.New(err)
			}
			if len(atts) != len(filling.ProofAttachmentIDs) {
				var existsIDs []string
				for _, v := range atts {
					existsIDs = append(existsIDs, v.ID)
				}
				withoutIDs := lo.Without(filling.ProofAttachmentIDs, existsIDs...)
				return errors.New(fmt.Errorf("some attachment not exists,id: %v", withoutIDs))
			}

			// 设置附件关联
			var bas []*model.BusinessAttachment
			for _, ba := range filling.ProofAttachmentIDs {
				bas = append(bas, &model.BusinessAttachment{
					BusinessType:   businessType,
					BusinessID:     filling.ID,
					AttachmentType: attachmentType,
					AttachmentID:   ba,
				})
			}
			err = se.businessAttachmentRepository.Create(ctx, bas)
			if err != nil {
				return errors.New(err)
			}
		}

		return nil
	})

	return errors.New(err)
}

// Delete 删除备案
func (se *Service) Delete(ctx context.Context, ids []string) error {
	err := se.session.Transaction(ctx, func(ctx context.Context) error {
		err := se.fillingRepository.Delete(ctx, ids)
		if err != nil {
			return errors.New(err)
		}

		for _, id := range ids {
			err = se.businessAttachmentRepository.Delete(ctx, businessType, id, []string{attachmentType})
			if err != nil {
				return errors.New(err)
			}
		}
		return errors.New(err)
	})

	return errors.New(err)
}

// Option 获取所有备案
func (se *Service) Option(ctx context.Context) ([]*dto.Filling, error) {
	fillings, err := se.fillingRepository.Option(ctx)

	return typeutil.MustConvert[[]*dto.Filling](fillings), errors.New(err)
}

// Name 返回服务名称,框架使用
func (s *Service) Name() string {
	return service.Filling
}
