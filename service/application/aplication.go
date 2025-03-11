package application

import (
	"context"

	"github.com/fnoopv/amp/database/model"
	"github.com/fnoopv/amp/dto"
	"github.com/fnoopv/amp/service"
	"github.com/google/uuid"
	"goyave.dev/filter"
	"goyave.dev/goyave/v5/database"
	"goyave.dev/goyave/v5/util/errors"
	"goyave.dev/goyave/v5/util/typeutil"
)

type Repository interface {
	Paginate(ctx context.Context, request *filter.Request) (*database.Paginator[*model.Application], error)
	Create(ctx context.Context, app *model.Application) error
	Update(ctx context.Context, app *model.Application) error
	Delete(ctx context.Context, id []string) error
}

type Service struct {
	appRepository Repository
}

func NewService(appRepository Repository) *Service {
	return &Service{
		appRepository: appRepository,
	}
}

func (se *Service) Paginate(ctx context.Context, request *filter.Request) (*database.PaginatorDTO[*dto.Application], error) {
	paginator, err := se.appRepository.Paginate(ctx, request)

	return typeutil.MustConvert[*database.PaginatorDTO[*dto.Application]](paginator), errors.New(err)
}

func (se *Service) Create(ctx context.Context, app *dto.ApplicationCreate) error {
	modelApp := typeutil.Copy(&model.Application{}, app)

	uid, err := uuid.NewV7()
	if err != nil {
		return errors.New(err)
	}
	modelApp.ID = uid.String()

	err = se.appRepository.Create(ctx, modelApp)

	return errors.New(err)
}

func (se *Service) Update(ctx context.Context, app *dto.ApplicationUpdate) error {
	modelApp := typeutil.Copy(&model.Application{}, app)

	err := se.appRepository.Update(ctx, modelApp)

	return errors.New(err)
}

func (se *Service) Delete(ctx context.Context, ids []string) error {
	err := se.appRepository.Delete(ctx, ids)

	return errors.New(err)
}

// Name 返回服务名称,框架使用
func (se *Service) Name() string {
	return service.Application
}
