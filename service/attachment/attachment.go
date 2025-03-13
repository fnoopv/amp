package attachment

import (
	"context"

	"github.com/fnoopv/amp/database/model"
	"github.com/fnoopv/amp/dto"
	"github.com/fnoopv/amp/service"
	"goyave.dev/goyave/v5/util/errors"
	"goyave.dev/goyave/v5/util/typeutil"
)

type Repository interface {
	Create(ctx context.Context, att *model.Attachment) error
	Update(ctx context.Context, att *model.Attachment) error
	FindByID(ctx context.Context, id string) (*model.Attachment, error)
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (se *Service) Create(ctx context.Context, att *dto.AttachmentCreate) error {
	modelAtt := typeutil.MustConvert[*model.Attachment](att)

	err := se.repository.Create(ctx, modelAtt)

	return errors.New(err)
}

func (se *Service) FindByID(ctx context.Context, id string) (*dto.AttachmentInternal, error) {
	att, err := se.repository.FindByID(ctx, id)

	return typeutil.MustConvert[*dto.AttachmentInternal](att), errors.New(err)
}

// Name 返回服务名称,框架使用
func (s *Service) Name() string {
	return service.Attachment
}
