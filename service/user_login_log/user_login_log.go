package userloginlog

import (
	"context"

	"github.com/fnoopv/amp/database/model"
	"github.com/fnoopv/amp/dto"
	"github.com/google/uuid"
	"goyave.dev/filter"
	"goyave.dev/goyave/v5/database"
	"goyave.dev/goyave/v5/util/errors"
	"goyave.dev/goyave/v5/util/typeutil"
)

type Repository interface {
	Paginate(ctx context.Context, request *filter.Request) (*database.Paginator[*model.UserLoginLog], error)
	Create(ctx context.Context, record *model.UserLoginLog) error
}

type Service struct {
	userLoginLogRepository Repository
}

func NewService(userLoginLogRepository Repository) *Service {
	return &Service{
		userLoginLogRepository: userLoginLogRepository,
	}
}

func (se *Service) Paginate(ctx context.Context, request *filter.Request) (*database.PaginatorDTO[*dto.UserLoginLog], error) {
	paginator, err := se.userLoginLogRepository.Paginate(ctx, request)

	return typeutil.MustConvert[*database.PaginatorDTO[*dto.UserLoginLog]](paginator), errors.New(err)
}

func (se *Service) Create(ctx context.Context, record *dto.UserLoginLog) error {
	modelRecord := typeutil.Copy(&model.UserLoginLog{}, record)

	uid, err := uuid.NewV7()
	if err != nil {
		return errors.New(err)
	}
	modelRecord.ID = uid.String()

	err = se.userLoginLogRepository.Create(ctx, modelRecord)

	return errors.New(err)
}
