package organization

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

// Repository 组织仓库
type Repository interface {
	Paginate(ctx context.Context, request *filter.Request) (*database.Paginator[*model.Organization], error)
	Create(ctx context.Context, organization *model.Organization) error
	Update(ctx context.Context, id string, organization *model.Organization) error
	Delete(ctx context.Context, id string) error
	FindByID(ctx context.Context, id string) (*model.Organization, error)
}

// Service 组织服务
type Service struct {
	repository Repository
}

// NewService 返回新的组织服务
func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

// Paginate 获取组织列表
func (se *Service) Paginate(ctx context.Context, request *filter.Request) (*database.PaginatorDTO[*dto.Organization], error) {
	paginator, err := se.repository.Paginate(ctx, request)

	return typeutil.MustConvert[*database.PaginatorDTO[*dto.Organization]](paginator), errors.New(err)
}

// Create 新增组织
func (se *Service) Create(ctx context.Context, organization *dto.OrganizationCreate) error {
	modelOrg := typeutil.Copy(&model.Organization{}, organization)

	uid, err := uuid.NewV7()
	if err != nil {
		return errors.New(err)
	}
	modelOrg.ID = uid.String()

	err = se.repository.Create(ctx, modelOrg)

	return errors.New(err)
}

// Update 更新组织信息
func (se *Service) Update(ctx context.Context, id string, organization *dto.OrganizationUpdate) error {
	modelOrg := typeutil.Copy(&model.Organization{}, organization)

	err := se.repository.Update(ctx, id, modelOrg)

	return errors.New(err)
}

// Delete 删除组织
func (se *Service) Delete(ctx context.Context, id string) error {
	err := se.repository.Delete(ctx, id)

	return errors.New(err)
}

// FindByID 根据ID获取组织
func (se *Service) FindByID(ctx context.Context, id string) (*dto.Organization, error) {
	org, err := se.repository.FindByID(ctx, id)
	if err != nil {
		return nil, errors.New(err)
	}

	return typeutil.MustConvert[*dto.Organization](org), nil
}

// Name 返回服务名称,框架使用
func (s *Service) Name() string {
	return service.Organization
}
