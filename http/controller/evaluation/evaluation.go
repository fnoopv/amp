package evaluation

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
	Paginate(ctx context.Context, request *filter.Request) (*database.PaginatorDTO[*dto.Evaluation], error)
	Create(ctx context.Context, evaluation *dto.EvaluationCreate) error
	Update(ctx context.Context, evaluation *dto.EvaluationUpdate) error
	Delete(ctx context.Context, ids []string) error
}

type Controller struct {
	goyave.Component
	service Service
}

func (co *Controller) Init(server *goyave.Server) {
	co.service = server.Service(service.Evaluation).(Service)
	co.Component.Init(server)
}

func (co *Controller) RegisterRoutes(router *goyave.Router) {
	subRouter := router.Subrouter("/evaluations")
	subRouter.Get("/", co.Index).ValidateQuery(filter.Validation)
	subRouter.Post("/", co.Create).ValidateBody(CreateRequest)
	subRouter.Post("/update", co.Update).ValidateBody(UpdateRequest)
	subRouter.Post("/delete", co.Delete).ValidateBody(DeleteRequest)
}

func (co *Controller) Index(response *goyave.Response, request *goyave.Request) {
	paginator, err := co.service.Paginate(request.Context(), filter.NewRequest(request.Query))
	if response.WriteDBError(err) {
		return
	}

	response.JSON(http.StatusOK, dto.CommonResponse{
		Message: dto.SuccessMessage,
		Data:    paginator,
	})
}

// Create 创建
func (co *Controller) Create(response *goyave.Response, request *goyave.Request) {
	org := typeutil.MustConvert[*dto.EvaluationCreate](request.Data)

	err := co.service.Create(request.Context(), org)
	if err != nil {
		response.Error(err)
		return
	}

	response.JSON(http.StatusOK, dto.SuccessResponse)
}

// Update 更新
func (co *Controller) Update(response *goyave.Response, request *goyave.Request) {
	req := typeutil.MustConvert[*dto.EvaluationUpdate](request.Data)

	err := co.service.Update(request.Context(), req)
	if err != nil {
		response.Error(err)
		return
	}

	response.JSON(http.StatusOK, dto.SuccessResponse)
}

// Delete 删除
func (co *Controller) Delete(response *goyave.Response, request *goyave.Request) {
	req := typeutil.MustConvert[*dto.EvaluationDelete](request.Data)

	err := co.service.Delete(request.Context(), req.IDs)
	if err != nil {
		response.Error(err)
		return
	}

	response.JSON(http.StatusOK, dto.SuccessResponse)
}
