package filling

import (
	"context"

	"github.com/fnoopv/amp/database/model"
	"github.com/fnoopv/amp/dto"
	"github.com/fnoopv/amp/pkg/uid"
	"github.com/fnoopv/amp/service"
	"goyave.dev/filter"
	"goyave.dev/goyave/v5/database"
	"goyave.dev/goyave/v5/util/errors"
	"goyave.dev/goyave/v5/util/typeutil"
)

type Repository interface {
	Paginate(ctx context.Context, request *filter.Request) (*database.Paginator[*model.Filling], error)
	Create(ctx context.Context, filling *model.Filling) error
	Update(ctx context.Context, filling *model.Filling) error
	Delete(ctx context.Context, ids []string) error
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (se *Service) Paginate(ctx context.Context, request *filter.Request) (*database.PaginatorDTO[*dto.Filling], error) {
	fillings, err := se.repository.Paginate(ctx, request)

	return typeutil.MustConvert[*database.PaginatorDTO[*dto.Filling]](fillings), errors.New(err)
}

func (se *Service) Create(ctx context.Context, filling *dto.FillingCreate) error {
	modelFilling := typeutil.Copy(&model.Filling{}, filling)

	id, err := uid.Generate()
	if err != nil {
		return errors.New(err)
	}

	modelFilling.ID = id

	err = se.repository.Create(ctx, modelFilling)

	return errors.New(err)
}

func (se *Service) Update(ctx context.Context, filling *dto.FillingUpdate) error {
	modelFilling := typeutil.Copy(&model.Filling{}, filling)

	err := se.repository.Update(ctx, modelFilling)

	return errors.New(err)
}

func (se *Service) Delete(ctx context.Context, ids []string) error {
	err := se.repository.Delete(ctx, ids)

	return errors.New(err)
}

// Name 返回服务名称,框架使用
func (s *Service) Name() string {
	return service.Filling
}
