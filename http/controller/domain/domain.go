package domain

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
	Paginate(ctx context.Context, request *filter.Request) (*database.PaginatorDTO[*dto.Domain], error)
	Create(ctx context.Context, record *dto.DomainCreate) error
	Update(ctx context.Context, record *dto.DomainUpdate) error
	Delete(ctx context.Context, ids []string) error
	Option(ctx context.Context) ([]*dto.Domain, error)
}

type Controller struct {
	goyave.Component
	service Service
}

func (co *Controller) Init(server *goyave.Server) {
	co.service = server.Service(service.Domain).(Service)
	co.Component.Init(server)
}

func (co *Controller) RegisterRoutes(router *goyave.Router) {
	subRouter := router.Subrouter("/domains")
	subRouter.Get("/", co.Index).ValidateQuery(filter.Validation)
	subRouter.Post("/", co.Create).ValidateBody(CreateRequest)
	subRouter.Get("/options", co.Option)
	subRouter.Post("/update", co.Update).ValidateBody(UpdateRequest)
	subRouter.Post("/delete", co.Delete).ValidateBody(DeleteRequest)
}

func (co *Controller) Index(response *goyave.Response, request *goyave.Request) {
	paginator, err := co.service.Paginate(request.Context(), filter.NewRequest(request.Query))
	if response.WriteDBError(err) {
		return
	}

	response.JSON(http.StatusOK, dto.CommonResponse{
		Message: dto.ResponseSuccessMessage,
		Data:    paginator,
	})
}

// Create 创建
func (co *Controller) Create(response *goyave.Response, request *goyave.Request) {
	org := typeutil.MustConvert[*dto.DomainCreate](request.Data)

	err := co.service.Create(request.Context(), org)
	if err != nil {
		response.Error(err)
		return
	}

	response.JSON(http.StatusOK, dto.ResponseSuccess)
}

// Update 更新
func (co *Controller) Update(response *goyave.Response, request *goyave.Request) {
	req := typeutil.MustConvert[*dto.DomainUpdate](request.Data)

	err := co.service.Update(request.Context(), req)
	if err != nil {
		response.Error(err)
		return
	}

	response.JSON(http.StatusOK, dto.ResponseSuccess)
}

// Delete 删除
func (co *Controller) Delete(response *goyave.Response, request *goyave.Request) {
	req := typeutil.MustConvert[*dto.DomainDelete](request.Data)

	err := co.service.Delete(request.Context(), req.IDs)
	if err != nil {
		response.Error(err)
		return
	}

	response.JSON(http.StatusOK, dto.ResponseSuccess)
}

// Option 获取枚举
func (co *Controller) Option(response *goyave.Response, request *goyave.Request) {
	records, err := co.service.Option(request.Context())
	if err != nil {
		response.Error(err)
		return
	}

	response.JSON(http.StatusOK, dto.CommonResponse{
		Message: dto.ResponseSuccessMessage,
		Data:    records,
	})
}
