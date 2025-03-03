package user

import (
	"context"
	"fmt"

	"github.com/fnoopv/amp/database/model"
	"github.com/fnoopv/amp/dto"
	"github.com/fnoopv/amp/pkg/password"
	"github.com/fnoopv/amp/service"
	"github.com/google/uuid"
	"goyave.dev/filter"
	"goyave.dev/goyave/v5/database"
	"goyave.dev/goyave/v5/util/errors"
	"goyave.dev/goyave/v5/util/typeutil"
)

type Repository interface {
	Paginate(ctx context.Context, request *filter.Request) (*database.Paginator[*model.User], error)
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, id string, user *model.User) error
	Delete(ctx context.Context, id string) error
	FindByID(ctx context.Context, id string) (*model.User, error)
	FindByUsername(ctx context.Context, username string) (*model.User, error)
	UpdatePassword(ctx context.Context, id, pwd string) error
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

// Paginate 分页
func (se *Service) Paginate(ctx context.Context, request *filter.Request) (*database.PaginatorDTO[*dto.User], error) {
	paginator, err := se.repository.Paginate(ctx, request)

	return typeutil.MustConvert[*database.PaginatorDTO[*dto.User]](paginator), errors.New(err)
}

// Create 创建用户
func (se *Service) Create(ctx context.Context, user *dto.UserCreate) (string, error) {
	modelUser := typeutil.Copy(&model.User{}, user)

	uid, err := uuid.NewV7()
	if err != nil {
		return "", errors.New(err)
	}
	modelUser.ID = uid.String()

	pwdStr, hashedPwd, err := password.GeneratePasswordAndHash()
	if err != nil {
		return "", errors.New(err)
	}
	modelUser.Password = hashedPwd

	err = se.repository.Create(ctx, modelUser)

	return pwdStr, errors.New(err)
}

// Update 更新用户
func (se *Service) Update(ctx context.Context, id string, user *dto.UserUpdate) error {
	modelUser := typeutil.Copy(&model.User{}, user)
	err := se.repository.Update(ctx, id, modelUser)

	return errors.New(err)
}

// Delete 删除单个用户
func (se *Service) Delete(ctx context.Context, id string) error {
	err := se.repository.Delete(ctx, id)

	return errors.New(err)
}

// UpdatePassword 更新密码
func (se *Service) UpdatePassword(ctx context.Context, id, pwd string) error {
	hashedPwd, err := password.HashPassword(pwd)
	if err != nil {
		return errors.New(err)
	}

	return se.repository.UpdatePassword(ctx, id, hashedPwd)
}

// ResetPassword 重置用户密码
func (se *Service) ResetPassword(ctx context.Context, id string) (string, error) {
	pwdStr, hashedPwd, err := password.GeneratePasswordAndHash()
	if err != nil {
		return "", errors.New(err)
	}

	err = se.repository.UpdatePassword(ctx, id, hashedPwd)

	return pwdStr, errors.New(err)
}

// FindByID 根据ID获取用户信息
func (se *Service) FindByID(ctx context.Context, id string) (*dto.User, error) {
	user, err := se.repository.FindByID(ctx, id)

	return typeutil.MustConvert[*dto.User](user), errors.New(err)
}

// FindByUsername 根据用户名查找用户, 登录认证接口，请勿更改
func (se *Service) FindByUsername(ctx context.Context, username any) (*dto.UserInternal, error) {
	user, err := se.repository.FindByUsername(ctx, fmt.Sprintf("%v", username))

	return typeutil.MustConvert[*dto.UserInternal](user), errors.New(err)
}

// Name 返回服务名称,框架使用
func (se *Service) Name() string {
	return service.User
}
