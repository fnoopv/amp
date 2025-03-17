package network

import (
	"context"
	"fmt"

	"github.com/fnoopv/amp/database/model"
	"github.com/fnoopv/amp/dto"
	"github.com/fnoopv/amp/pkg/uid"
	"github.com/fnoopv/amp/service"
	"goyave.dev/filter"
	"goyave.dev/goyave/v5/database"
	"goyave.dev/goyave/v5/util/errors"
	"goyave.dev/goyave/v5/util/session"
	"goyave.dev/goyave/v5/util/typeutil"
)

type networkRepository interface {
	Paginate(ctx context.Context, request *filter.Request) (*database.Paginator[*model.Network], error)
	Create(ctx context.Context, record *model.Network) error
	Update(ctx context.Context, record *model.Network) error
	Delete(ctx context.Context, id []string) error
	Option(ctx context.Context) ([]*model.Network, error)
}

type fillingRepository interface {
	FindByIDs(ctx context.Context, ids []string) ([]*model.Filling, error)
}

type Service struct {
	session           session.Session
	networkRepository networkRepository
	fillingRepository fillingRepository
}

func NewService(session session.Session, networkRepository networkRepository, fillingRepository fillingRepository) *Service {
	return &Service{
		session:           session,
		networkRepository: networkRepository,
		fillingRepository: fillingRepository,
	}
}

func (se *Service) Paginate(ctx context.Context, request *filter.Request) (*database.PaginatorDTO[*dto.Network], error) {
	paginator, err := se.networkRepository.Paginate(ctx, request)
	if err != nil {
		return nil, errors.New(err)
	}

	return typeutil.MustConvert[*database.PaginatorDTO[*dto.Network]](paginator), errors.New(err)
}

func (se *Service) Create(ctx context.Context, createRecord *dto.NetworkCreate) error {
	modelRecord := typeutil.Copy(&model.Network{}, createRecord)

	id, err := uid.Generate()
	if err != nil {
		return errors.New(err)
	}
	modelRecord.ID = id
	err = se.session.Transaction(ctx, func(ctx context.Context) error {
		// 检查备案是否存在
		if createRecord.FillingID != "" {
			fillings, err := se.fillingRepository.FindByIDs(ctx, []string{createRecord.FillingID})
			if err != nil {
				return errors.New(err)
			}

			if len(fillings) != 1 {
				return errors.New(fmt.Errorf("filling is not exist: %v", createRecord.FillingID))
			}
		}

		// 创建
		err := se.networkRepository.Create(ctx, modelRecord)
		if err != nil {
			return errors.New(err)
		}

		return nil
	})

	return errors.New(err)
}

func (se *Service) Update(ctx context.Context, updareRecord *dto.NetworkUpdate) error {
	modelRecord := typeutil.Copy(&model.Network{}, updareRecord)

	err := se.session.Transaction(ctx, func(ctx context.Context) error {
		// 检查备案是否存在
		if updareRecord.FillingID != "" {
			fillings, err := se.fillingRepository.FindByIDs(ctx, []string{updareRecord.FillingID})
			if err != nil {
				return errors.New(err)
			}

			if len(fillings) != 1 {
				return errors.New(fmt.Errorf("filling is not exist: %v", updareRecord.FillingID))
			}
		}

		// 更新
		err := se.networkRepository.Update(ctx, modelRecord)
		if err != nil {
			return errors.New(err)
		}

		return nil
	})

	return errors.New(err)
}

func (se *Service) Delete(ctx context.Context, ids []string) error {
	err := se.networkRepository.Delete(ctx, ids)

	return errors.New(err)
}

// Option 获取枚举
func (se *Service) Option(ctx context.Context) ([]*dto.Network, error) {
	records, err := se.networkRepository.Option(ctx)

	return typeutil.MustConvert[[]*dto.Network](records), errors.New(err)
}

// Name 返回服务名称,框架使用
func (se *Service) Name() string {
	return service.Network
}
