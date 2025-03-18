package application

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

type appRepository interface {
	Paginate(ctx context.Context, request *filter.Request) (*database.Paginator[*model.Application], error)
	Create(ctx context.Context, app *model.Application) error
	Update(ctx context.Context, app *model.Application) error
	Delete(ctx context.Context, id []string) error
}

type fillingRepository interface {
	FindByIDs(ctx context.Context, ids []string) ([]*model.Filling, error)
}

type networkRepository interface {
	FindByIDs(ctx context.Context, ids []string) ([]*model.Network, error)
}

type Service struct {
	session           session.Session
	appRepository     appRepository
	fillingRepository fillingRepository
	networkRepository networkRepository
}

func NewService(
	session session.Session,
	appRepository appRepository,
	fillingRepository fillingRepository,
	networkRepository networkRepository,
) *Service {
	return &Service{
		session:           session,
		appRepository:     appRepository,
		fillingRepository: fillingRepository,
		networkRepository: networkRepository,
	}
}

func (se *Service) Paginate(ctx context.Context, request *filter.Request) (*database.PaginatorDTO[*dto.Application], error) {
	paginator, err := se.appRepository.Paginate(ctx, request)

	dtoPaginator := typeutil.MustConvert[*database.PaginatorDTO[*dto.Application]](paginator)
	for _, v := range dtoPaginator.Records {
		v.FillingIDs = lo.Map(v.Fillings, func(item dto.Filling, _ int) string {
			return item.ID
		})
		v.NetworkIDs = lo.Map(v.Networks, func(item dto.Network, index int) string {
			return item.ID
		})
	}

	return dtoPaginator, errors.New(err)
}

func (se *Service) Create(ctx context.Context, app *dto.ApplicationCreate) error {
	modelApp := typeutil.Copy(&model.Application{}, app)

	id, err := uid.Generate()
	if err != nil {
		return errors.New(err)
	}
	modelApp.ID = id
	err = se.session.Transaction(ctx, func(ctx context.Context) error {
		// 处理关联备案
		if len(app.FillingIDs) > 0 {
			fillings, err := se.fillingRepository.FindByIDs(ctx, app.FillingIDs)
			if err != nil {
				return errors.New(err)
			}

			if len(fillings) != len(app.FillingIDs) {
				return errors.New(fmt.Errorf("some filling is not exist: %v", app.FillingIDs))
			}

			modelApp.Fillings = lo.Map(app.FillingIDs, func(item string, index int) model.Filling {
				return model.Filling{
					ID: item,
				}
			})
		}
		// 处理关联网络
		if len(app.NetworkIDs) > 0 {
			networks, err := se.networkRepository.FindByIDs(ctx, app.NetworkIDs)
			if err != nil {
				return errors.New(err)
			}

			if len(networks) != len(app.NetworkIDs) {
				return errors.New(fmt.Errorf("some network is not exist: %v", app.NetworkIDs))
			}

			modelApp.Networks = lo.Map(app.NetworkIDs, func(item string, index int) model.Network {
				return model.Network{
					ID: item,
				}
			})
		}

		// 创建应用
		err := se.appRepository.Create(ctx, modelApp)
		if err != nil {
			return errors.New(err)
		}

		return nil
	})

	return errors.New(err)
}

func (se *Service) Update(ctx context.Context, app *dto.ApplicationUpdate) error {
	modelApp := typeutil.Copy(&model.Application{}, app)

	err := se.session.Transaction(ctx, func(ctx context.Context) error {
		// 处理关联备案
		if len(app.FillingIDs) > 0 {
			fillings, err := se.fillingRepository.FindByIDs(ctx, app.FillingIDs)
			if err != nil {
				return errors.New(err)
			}

			if len(fillings) != len(app.FillingIDs) {
				return errors.New(fmt.Errorf("some filling is not exist: %v", app.FillingIDs))
			}

			modelApp.Fillings = lo.Map(app.FillingIDs, func(item string, index int) model.Filling {
				return model.Filling{
					ID: item,
				}
			})
		}
		// 处理关联网络
		if len(app.NetworkIDs) > 0 {
			networks, err := se.networkRepository.FindByIDs(ctx, app.NetworkIDs)
			if err != nil {
				return errors.New(err)
			}

			if len(networks) != len(app.NetworkIDs) {
				return errors.New(fmt.Errorf("some network is not exist: %v", app.NetworkIDs))
			}

			modelApp.Networks = lo.Map(app.NetworkIDs, func(item string, index int) model.Network {
				return model.Network{
					ID: item,
				}
			})
		}

		// 更新应用
		err := se.appRepository.Update(ctx, modelApp)
		if err != nil {
			return errors.New(err)
		}

		return nil
	})

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
