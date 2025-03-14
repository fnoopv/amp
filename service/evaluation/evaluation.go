package evaluation

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
	BusinessType         = "evaluation"
	EvaluationAttachment = "evaluation"
	RepairAttachment     = "rapair"
)

type evaluationRepository interface {
	Paginate(ctx context.Context, request *filter.Request) (*database.Paginator[*model.Evaluation], error)
	Create(ctx context.Context, evaluation *model.Evaluation) error
	Update(ctx context.Context, evaluation *model.Evaluation) error
	Delete(ctx context.Context, ids []string) error
}

type businessAttachmentRepository interface {
	Create(ctx context.Context, bas []*model.BusinessAttachment) error
	Delete(ctx context.Context, businessType, businessID string, attachmentType []string) error
}

type attachmentRepository interface {
	FindByIDs(ctx context.Context, ids []string) ([]*model.Attachment, error)
}

type Service struct {
	session                      session.Session
	evaluationRepository         evaluationRepository
	businessAttachmentRepository businessAttachmentRepository
	attachmentRepository         attachmentRepository
}

func NewService(
	session session.Session,
	evaluationRepository evaluationRepository,
	businessAttbusinessAttachmentRepository businessAttachmentRepository,
	attattachmentRepository attachmentRepository,
) *Service {
	return &Service{
		session:                      session,
		evaluationRepository:         evaluationRepository,
		businessAttachmentRepository: businessAttbusinessAttachmentRepository,
		attachmentRepository:         attattachmentRepository,
	}
}

func (se *Service) Paginate(ctx context.Context, request *filter.Request) (
	*database.PaginatorDTO[*dto.Evaluation],
	error,
) {
	evaluations, err := se.evaluationRepository.Paginate(ctx, request)

	return typeutil.MustConvert[*database.PaginatorDTO[*dto.Evaluation]](evaluations), errors.New(err)
}

func (se *Service) Create(ctx context.Context, evaluation *dto.EvaluationCreate) error {
	modelEvaluation := typeutil.Copy(&model.Evaluation{}, evaluation)

	id, err := uid.Generate()
	if err != nil {
		return errors.New(err)
	}

	modelEvaluation.ID = id

	err = se.evaluationRepository.Create(ctx, modelEvaluation)

	return errors.New(err)
}

func (se *Service) Update(ctx context.Context, evaluation *dto.EvaluationUpdate) error {
	modelEvaluation := typeutil.Copy(&model.Evaluation{}, evaluation)

	err := se.session.Transaction(ctx, func(ctx context.Context) error {
		var err error

		// 更新测评
		err = se.evaluationRepository.Update(ctx, modelEvaluation)
		if err != nil {
			return errors.New(err)
		}

		// 检查附件是否存在
		attIDs := lo.Uniq(append(evaluation.EvaluationAttachmentIDs, evaluation.RepairAttachmentIDs...))
		atts, err := se.attachmentRepository.FindByIDs(ctx, attIDs)
		if err != nil {
			return errors.New(err)
		}
		if len(atts) != len(attIDs) {
			var existsIDs []string
			for _, v := range atts {
				existsIDs = append(existsIDs, v.ID)
			}
			withoutIDs := lo.Without(attIDs, existsIDs...)
			return errors.New(fmt.Errorf("some attachment not exists,id: %v", withoutIDs))
		}

		// 删除关联表记录
		err = se.businessAttachmentRepository.Delete(
			ctx,
			BusinessType,
			modelEvaluation.ID,
			[]string{EvaluationAttachment, RepairAttachment},
		)
		if err != nil {
			return errors.New(err)
		}

		// 插入关联表记录
		bas := []*model.BusinessAttachment{}
		for _, v := range evaluation.EvaluationAttachmentIDs {
			bas = append(bas, &model.BusinessAttachment{
				BusinessType:   BusinessType,
				BusinessID:     modelEvaluation.ID,
				AttachmentType: EvaluationAttachment,
				AttachmentID:   v,
			})
		}
		for _, v := range evaluation.RepairAttachmentIDs {
			bas = append(bas, &model.BusinessAttachment{
				BusinessType:   BusinessType,
				BusinessID:     modelEvaluation.ID,
				AttachmentType: RepairAttachment,
				AttachmentID:   v,
			})
		}
		err = se.businessAttachmentRepository.Create(ctx, bas)

		return errors.New(err)
	})

	return errors.New(err)
}

func (se *Service) Delete(ctx context.Context, ids []string) error {
	err := se.evaluationRepository.Delete(ctx, ids)

	return errors.New(err)
}

// Name 返回服务名称,框架使用
func (s *Service) Name() string {
	return service.Evaluation
}
