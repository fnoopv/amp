package user

import (
	"context"
	"net/http"

	"github.com/fnoopv/amp/dto"
	"github.com/fnoopv/amp/service"
	"goyave.dev/filter"
	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/database"
	"goyave.dev/goyave/v5/util/typeutil"
)

type Service interface {
	Paginate(ctx context.Context, request *filter.Request) (*database.PaginatorDTO[*dto.User], error)
	Create(ctx context.Context, user *dto.UserCreate) (string, error)
	Update(ctx context.Context, id string, user *dto.UserUpdate) error
	Delete(ctx context.Context, id string) error
	FindByID(ctx context.Context, id string) (*dto.User, error)
	UpdatePassword(ctx context.Context, id, pwd string) error
	ResetPassword(ctx context.Context, id string) (string, error)
}

type Controller struct {
	goyave.Component
	UserService Service
}

func (co *Controller) Init(server *goyave.Server) {
	co.UserService = server.Service(service.User).(Service)
	co.Component.Init(server)
}

func (co *Controller) RegisterRoutes(router *goyave.Router) {
	subRouter := router.Subrouter("/users")
	subRouter.Get("/", co.Index).ValidateQuery(filter.Validation)
	subRouter.Post("/", co.Create).ValidateBody(CreateRequest)

	userRouter := subRouter.Subrouter("/{id}")
	userRouter.Get("/", co.FindByID)
	userRouter.Delete("/", co.Delete)
	userRouter.Patch("/", co.Update).ValidateBody(UpdateRequest)
	userRouter.Post("/pwd", co.UpdatePassword).ValidateBody(UpdatePasswordRequest)
	userRouter.Post("/pwd/reset", co.ResetPassword)
}

func (co *Controller) Index(response *goyave.Response, request *goyave.Request) {
	paginator, err := co.UserService.Paginate(request.Context(), filter.NewRequest(request.Query))
	if response.WriteDBError(err) {
		return
	}

	response.JSON(http.StatusOK, dto.CommonResponse{
		Message: dto.ResponseSuccessMessage,
		Data:    paginator,
	})
}

// Create 创建用户
func (co *Controller) Create(response *goyave.Response, request *goyave.Request) {
	req := typeutil.MustConvert[*dto.UserCreate](request.Data)
	pwd, err := co.UserService.Create(request.Context(), req)
	if err != nil {
		response.Error(err)
	}
	response.JSON(http.StatusOK, dto.CommonResponse{
		Message: dto.ResponseSuccessMessage,
		Data:    dto.UserCreateResponse{Password: pwd},
	})
}

// Update 更新用户信息
func (co *Controller) Update(response *goyave.Response, request *goyave.Request) {
	id := request.RouteParams["id"]
	req := typeutil.MustConvert[*dto.UserUpdate](request.Data)

	if err := co.UserService.Update(request.Context(), id, req); err != nil {
		response.Error(err)
	}

	response.JSON(http.StatusOK, dto.ResponseSuccess)
}

// Delete 删除单个用户
func (co *Controller) Delete(response *goyave.Response, request *goyave.Request) {
	id := request.RouteParams["id"]

	if err := co.UserService.Delete(request.Context(), id); err != nil {
		response.Error(err)
	}

	response.JSON(http.StatusOK, dto.ResponseSuccess)
}

// Delete 删除单个用户
func (co *Controller) FindByID(response *goyave.Response, request *goyave.Request) {
	id := request.RouteParams["id"]

	user, err := co.UserService.FindByID(request.Context(), id)
	if err != nil {
		response.Error(err)
	}

	response.JSON(http.StatusOK, dto.CommonResponse{
		Message: dto.ResponseSuccessMessage,
		Data:    user,
	})
}

// UpdatePassword 更改密码
func (co *Controller) UpdatePassword(response *goyave.Response, request *goyave.Request) {
	req := typeutil.MustConvert[*dto.UserChangePassword](request.Data)
	if err := co.UserService.UpdatePassword(request.Context(), req.ConfirmPassword, req.NewPassword); err != nil {
		response.Error(err)
	}

	response.JSON(http.StatusOK, dto.ResponseSuccess)
}

// ResetPassword 重置用户密码
func (co *Controller) ResetPassword(response *goyave.Response, request *goyave.Request) {
	id := request.RouteParams["id"]

	pwd, err := co.UserService.ResetPassword(request.Context(), id)
	if err != nil {
		response.Error(err)
	}

	response.JSON(http.StatusOK, dto.CommonResponse{
		Message: dto.ResponseSuccessMessage,
		Data:    dto.UserResetPasswordResponse{Password: pwd},
	})
}
