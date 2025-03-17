package filling

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
	Paginate(ctx context.Context, request *filter.Request) (*database.PaginatorDTO[*dto.Filling], error)
	Create(ctx context.Context, filling *dto.FillingCreate) error
	Update(ctx context.Context, filling *dto.FillingUpdate) error
	Delete(ctx context.Context, ids []string) error
	Option(ctx context.Context) ([]*dto.Filling, error)
}

type Controller struct {
	goyave.Component
	service Service
}

func (co *Controller) Init(server *goyave.Server) {
	co.service = server.Service(service.Filling).(Service)
	co.Component.Init(server)
}

func (co *Controller) RegisterRoutes(router *goyave.Router) {
	subRouter := router.Subrouter("/fillings")
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
	org := typeutil.MustConvert[*dto.FillingCreate](request.Data)

	err := co.service.Create(request.Context(), org)
	if err != nil {
		response.Error(err)
		return
	}

	response.JSON(http.StatusOK, dto.ResponseSuccess)
}

// Update 更新
func (co *Controller) Update(response *goyave.Response, request *goyave.Request) {
	req := typeutil.MustConvert[*dto.FillingUpdate](request.Data)

	err := co.service.Update(request.Context(), req)
	if err != nil {
		response.Error(err)
		return
	}

	response.JSON(http.StatusOK, dto.ResponseSuccess)
}

// Delete 删除
func (co *Controller) Delete(response *goyave.Response, request *goyave.Request) {
	req := typeutil.MustConvert[*dto.FillingDelete](request.Data)

	err := co.service.Delete(request.Context(), req.IDs)
	if err != nil {
		response.Error(err)
		return
	}

	response.JSON(http.StatusOK, dto.ResponseSuccess)
}

// Option 获取备案枚举
func (co *Controller) Option(response *goyave.Response, request *goyave.Request) {
	fillings, err := co.service.Option(request.Context())
	if err != nil {
		response.Error(err)
		return
	}

	response.JSON(http.StatusOK, dto.CommonResponse{
		Message: dto.ResponseSuccessMessage,
		Data:    fillings,
	})
}
