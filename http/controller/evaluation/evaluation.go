package evaluation

import (
	"context"
	"net/http"

	"github.com/fnoopv/amp/dto"
	"github.com/fnoopv/amp/service"
	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/util/typeutil"
)

type Service interface {
	FindByFillingID(ctx context.Context, fillingID string) ([]*dto.Evaluation, error)
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
	subRouter.Post("/", co.Create).ValidateBody(CreateRequest)
	subRouter.Get("/", co.Index).ValidateQuery(FindRequest)
	subRouter.Post("/update", co.Update).ValidateBody(UpdateRequest)
	subRouter.Post("/delete", co.Delete).ValidateBody(DeleteRequest)
}

func (co *Controller) Index(response *goyave.Response, request *goyave.Request) {
	req, err := typeutil.Convert[*dto.EvaluationFind](request.Query)
	if err != nil {
		co.Logger().Info("invalid filling_id")
		response.JSON(http.StatusBadRequest, dto.CommonResponse{
			Message: request.Lang.Get("param.invalid"),
		})
		return
	}
	evaluations, err := co.service.FindByFillingID(request.Context(), req.FillingID)
	if response.WriteDBError(err) {
		return
	}

	response.JSON(http.StatusOK, dto.CommonResponse{
		Message: dto.SuccessMessage,
		Data:    evaluations,
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
