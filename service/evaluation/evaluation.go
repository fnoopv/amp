package evaluation

import (
	"context"
	"fmt"

	"github.com/fnoopv/amp/database/model"
	"github.com/fnoopv/amp/dto"
	"github.com/fnoopv/amp/pkg/uid"
	"github.com/fnoopv/amp/service"
	"github.com/samber/lo"
	"goyave.dev/goyave/v5/util/errors"
	"goyave.dev/goyave/v5/util/session"
	"goyave.dev/goyave/v5/util/typeutil"
)

const (
	businessType             = "evaluation"
	attachmentTypeEvaluation = "evaluation"
	attachmentTypeRapire     = "rapair"
)

type evaluationRepository interface {
	FindByFillingID(ctx context.Context, fillingID string) ([]*model.Evaluation, error)
	Create(ctx context.Context, evaluation *model.Evaluation) error
	Update(ctx context.Context, evaluation *model.Evaluation) error
	Delete(ctx context.Context, ids []string) error
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

func (se *Service) FindByFillingID(ctx context.Context, fillingID string) ([]*dto.Evaluation, error) {
	var evaluations []*dto.Evaluation

	err := se.session.Transaction(ctx, func(ctx context.Context) error {
		var err error

		modelEvaluations, err := se.evaluationRepository.FindByFillingID(ctx, fillingID)
		if err != nil {
			return errors.New(err)
		}

		evaluations, err = typeutil.Convert[[]*dto.Evaluation](&modelEvaluations)
		if err != nil {
			return errors.New(err)
		}

		for _, evaluation := range evaluations {
			ids, err := se.businessAttachmentRepository.FindAttachmentIDs(
				ctx,
				businessType,
				evaluation.ID,
				[]string{attachmentTypeEvaluation},
			)
			if err != nil {
				return errors.New(err)
			}
			evaluation.EvaluationAttachmentIDs = ids
			if len(ids) > 0 {
				modelAttachments, err := se.attachmentRepository.FindByIDs(ctx, ids)
				if err != nil {
					return errors.New(err)
				}
				attachments, err := typeutil.Convert[*[]dto.Attachment](&modelAttachments)
				if err != nil {
					return errors.New(err)
				}

				evaluation.EvaluationAttachments = *attachments
			}
			rapireIDs, err := se.businessAttachmentRepository.FindAttachmentIDs(
				ctx,
				businessType,
				evaluation.ID,
				[]string{attachmentTypeRapire},
			)
			if err != nil {
				return errors.New(err)
			}
			evaluation.RepairAttachmentIDs = rapireIDs
			if len(ids) > 0 {
				modelAttachments, err := se.attachmentRepository.FindByIDs(ctx, rapireIDs)
				if err != nil {
					return errors.New(err)
				}
				attachments, err := typeutil.Convert[*[]dto.Attachment](&modelAttachments)
				if err != nil {
					return errors.New(err)
				}

				evaluation.RepairAttachments = *attachments
			}
		}

		return err
	})

	return evaluations, errors.New(err)
}

func (se *Service) Create(ctx context.Context, evaluation *dto.EvaluationCreate) error {
	modelEvaluation := typeutil.Copy(&model.Evaluation{}, evaluation)

	id, err := uid.Generate()
	if err != nil {
		return errors.New(err)
	}
	modelEvaluation.ID = id

	err = se.session.Transaction(ctx, func(ctx context.Context) error {
		err := se.evaluationRepository.Create(ctx, modelEvaluation)
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

		// 插入关联表记录
		bas := []*model.BusinessAttachment{}
		for _, v := range evaluation.EvaluationAttachmentIDs {
			bas = append(bas, &model.BusinessAttachment{
				BusinessType:   businessType,
				BusinessID:     modelEvaluation.ID,
				AttachmentType: attachmentTypeEvaluation,
				AttachmentID:   v,
			})
		}
		for _, v := range evaluation.RepairAttachmentIDs {
			bas = append(bas, &model.BusinessAttachment{
				BusinessType:   businessType,
				BusinessID:     modelEvaluation.ID,
				AttachmentType: attachmentTypeRapire,
				AttachmentID:   v,
			})
		}
		err = se.businessAttachmentRepository.Create(ctx, bas)
		return errors.New(err)
	})

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
			businessType,
			modelEvaluation.ID,
			[]string{attachmentTypeEvaluation, attachmentTypeRapire},
		)
		if err != nil {
			return errors.New(err)
		}

		// 插入关联表记录
		bas := []*model.BusinessAttachment{}
		for _, v := range evaluation.EvaluationAttachmentIDs {
			bas = append(bas, &model.BusinessAttachment{
				BusinessType:   businessType,
				BusinessID:     modelEvaluation.ID,
				AttachmentType: attachmentTypeEvaluation,
				AttachmentID:   v,
			})
		}
		for _, v := range evaluation.RepairAttachmentIDs {
			bas = append(bas, &model.BusinessAttachment{
				BusinessType:   businessType,
				BusinessID:     modelEvaluation.ID,
				AttachmentType: attachmentTypeRapire,
				AttachmentID:   v,
			})
		}
		err = se.businessAttachmentRepository.Create(ctx, bas)
		return errors.New(err)
	})

	return errors.New(err)
}

func (se *Service) Delete(ctx context.Context, ids []string) error {
	err := se.session.Transaction(ctx, func(ctx context.Context) error {
		err := se.evaluationRepository.Delete(ctx, ids)
		if err != nil {
			return errors.New(err)
		}
		for _, id := range ids {
			err = se.businessAttachmentRepository.Delete(
				ctx,
				businessType,
				id,
				[]string{attachmentTypeEvaluation, attachmentTypeRapire},
			)
			if err != nil {
				return errors.New(err)
			}
		}

		return nil
	})

	return errors.New(err)
}

// Name 返回服务名称,框架使用
func (s *Service) Name() string {
	return service.Evaluation
}
