package domain

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

type domainRepository interface {
	Paginate(ctx context.Context, request *filter.Request) (*database.Paginator[*model.Domain], error)
	Create(ctx context.Context, record *model.Domain) error
	Update(ctx context.Context, record *model.Domain) error
	Delete(ctx context.Context, id []string) error
	Option(ctx context.Context) ([]*model.Domain, error)
}

type fillingRepository interface {
	FindByIDs(ctx context.Context, ids []string) ([]*model.Filling, error)
}

type Service struct {
	session           session.Session
	domainRepository  domainRepository
	fillingRepository fillingRepository
}

func NewService(session session.Session, networkRepository domainRepository, fillingRepository fillingRepository) *Service {
	return &Service{
		session:           session,
		domainRepository:  networkRepository,
		fillingRepository: fillingRepository,
	}
}

func (se *Service) Paginate(ctx context.Context, request *filter.Request) (*database.PaginatorDTO[*dto.Domain], error) {
	paginator, err := se.domainRepository.Paginate(ctx, request)
	if err != nil {
		return nil, errors.New(err)
	}

	dtoPaginator := typeutil.MustConvert[*database.PaginatorDTO[*dto.Domain]](paginator)
	for _, v := range dtoPaginator.Records {
		v.FillingIDs = lo.Map(v.Fillings, func(item dto.Filling, _ int) string {
			return item.ID
		})
	}

	return dtoPaginator, errors.New(err)
}

func (se *Service) Create(ctx context.Context, createRecord *dto.DomainCreate) error {
	modelRecord := typeutil.Copy(&model.Domain{}, createRecord)

	id, err := uid.Generate()
	if err != nil {
		return errors.New(err)
	}
	modelRecord.ID = id
	err = se.session.Transaction(ctx, func(ctx context.Context) error {
		// 处理关联备案
		if len(createRecord.FillingIDs) > 0 {
			fillings, err := se.fillingRepository.FindByIDs(ctx, createRecord.FillingIDs)
			if err != nil {
				return errors.New(err)
			}

			if len(fillings) != len(createRecord.FillingIDs) {
				return errors.New(fmt.Errorf("some filling is not exist: %v", createRecord.FillingIDs))
			}

			modelRecord.Fillings = lo.Map(createRecord.FillingIDs, func(item string, index int) model.Filling {
				return model.Filling{
					ID: item,
				}
			})
		}

		// 创建
		err := se.domainRepository.Create(ctx, modelRecord)
		if err != nil {
			return errors.New(err)
		}

		return nil
	})

	return errors.New(err)
}

func (se *Service) Update(ctx context.Context, updareRecord *dto.DomainUpdate) error {
	modelRecord := typeutil.Copy(&model.Domain{}, updareRecord)

	err := se.session.Transaction(ctx, func(ctx context.Context) error {
		// 处理关联备案
		if len(updareRecord.FillingIDs) > 0 {
			fillings, err := se.fillingRepository.FindByIDs(ctx, updareRecord.FillingIDs)
			if err != nil {
				return errors.New(err)
			}

			if len(fillings) != len(updareRecord.FillingIDs) {
				return errors.New(fmt.Errorf("some filling is not exist: %v", updareRecord.FillingIDs))
			}

			modelRecord.Fillings = lo.Map(updareRecord.FillingIDs, func(item string, index int) model.Filling {
				return model.Filling{
					ID: item,
				}
			})
		}

		// 更新
		err := se.domainRepository.Update(ctx, modelRecord)
		if err != nil {
			return errors.New(err)
		}

		return nil
	})

	return errors.New(err)
}

func (se *Service) Delete(ctx context.Context, ids []string) error {
	err := se.domainRepository.Delete(ctx, ids)

	return errors.New(err)
}

// Option 获取枚举
func (se *Service) Option(ctx context.Context) ([]*dto.Domain, error) {
	records, err := se.domainRepository.Option(ctx)

	return typeutil.MustConvert[[]*dto.Domain](records), errors.New(err)
}

// Name 返回服务名称,框架使用
func (se *Service) Name() string {
	return service.Domain
}
