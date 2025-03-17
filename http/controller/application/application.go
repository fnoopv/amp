package application

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
	Paginate(ctx context.Context, request *filter.Request) (*database.PaginatorDTO[*dto.Application], error)
	Create(ctx context.Context, app *dto.ApplicationCreate) error
	Update(ctx context.Context, app *dto.ApplicationUpdate) error
	Delete(ctx context.Context, ids []string) error
}

type Controller struct {
	goyave.Component
	AppService Service
}

func (co *Controller) Init(server *goyave.Server) {
	co.AppService = server.Service(service.Application).(Service)
	co.Component.Init(server)
}

func (co *Controller) RegisterRoutes(router *goyave.Router) {
	subRouter := router.Subrouter("/applications")
	subRouter.Get("/", co.Index).ValidateQuery(filter.Validation)
	subRouter.Post("/", co.Create).ValidateBody(CreateRequest)

	subRouter.Post("/delete", co.Delete).ValidateBody(DeleteRequest)
	subRouter.Post("/update", co.Update).ValidateBody(UpdateRequest)
}

func (co *Controller) Index(response *goyave.Response, request *goyave.Request) {
	paginator, err := co.AppService.Paginate(request.Context(), filter.NewRequest(request.Query))
	if response.WriteDBError(err) {
		return
	}

	response.JSON(http.StatusOK, dto.CommonResponse{
		Message: dto.ResponseSuccessMessage,
		Data:    paginator,
	})
}

func (co *Controller) Create(response *goyave.Response, request *goyave.Request) {
	req := typeutil.MustConvert[*dto.ApplicationCreate](request.Data)
	err := co.AppService.Create(request.Context(), req)
	if response.WriteDBError(err) {
		return
	}
	response.JSON(http.StatusOK, dto.ResponseSuccess)
}

func (co *Controller) Update(response *goyave.Response, request *goyave.Request) {
	req := typeutil.MustConvert[*dto.ApplicationUpdate](request.Data)

	if err := co.AppService.Update(request.Context(), req); err != nil {
		response.Error(err)
		return
	}

	response.JSON(http.StatusOK, dto.ResponseSuccess)
}

func (co *Controller) Delete(response *goyave.Response, request *goyave.Request) {
	req := typeutil.MustConvert[*dto.ApplicationDelete](request.Data)
	if err := co.AppService.Delete(request.Context(), req.IDs); err != nil {
		response.Error(err)
		return
	}

	response.JSON(http.StatusOK, dto.ResponseSuccess)
}
