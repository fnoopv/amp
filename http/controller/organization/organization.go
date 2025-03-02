package organization

import (
	"context"
	"net/http"

	"github.com/fnoopv/amp/dto"
	"github.com/fnoopv/amp/service"
	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/database"
	"goyave.dev/goyave/v5/util/typeutil"
)

type Service interface {
	Paginate(ctx context.Context, page, pageSize int) (*database.PaginatorDTO[*dto.Organization], error)
	Create(ctx context.Context, organization *dto.OrganizationCreate) error
	Update(ctx context.Context, id string, organization *dto.OrganizationUpdate) error
	Delete(ctx context.Context, id string) error
	FindByID(ctx context.Context, id string) (*dto.Organization, error)
}

type Controller struct {
	goyave.Component
	organizationService Service
}

func (co *Controller) Init(server *goyave.Server) {
	co.organizationService = server.Service(service.Organization).(Service)
	co.Component.Init(server)
}

func (co *Controller) RegisterRoutes(router *goyave.Router) {
	subRouter := router.Subrouter("/organizations")
	subRouter.Get("/", co.Index).ValidateQuery(IndexRequest)
	subRouter.Post("/", co.Create).ValidateBody(CreateRequest)

	orgRouter := subRouter.Subrouter("/{id}")
	orgRouter.Get("/", co.FindByID)
	orgRouter.Patch("/", co.Update).ValidateBody(UpdateRequest)
	orgRouter.Delete("/", co.Delete)
}

func (co *Controller) Index(response *goyave.Response, request *goyave.Request) {
	query := typeutil.MustConvert[*dto.OrganizationIndex](request.Query)

	paginator, err := co.organizationService.Paginate(request.Context(), query.Page, query.PageSize)
	if response.WriteDBError(err) {
		return
	}

	response.JSON(http.StatusOK, dto.CommonResponse{
		Message: dto.SuccessMessage,
		Data:    paginator,
	})
}

// Create 创建组织
func (co *Controller) Create(response *goyave.Response, request *goyave.Request) {
	org := typeutil.MustConvert[*dto.OrganizationCreate](request.Data)

	err := co.organizationService.Create(request.Context(), org)
	if err != nil {
		response.Error(err)
	}

	response.JSON(http.StatusOK, dto.SuccessResponse)
}

// Update 更新组织信息
func (co *Controller) Update(response *goyave.Response, request *goyave.Request) {
	id := request.RouteParams["id"]
	org := typeutil.MustConvert[*dto.OrganizationUpdate](request.Data)

	err := co.organizationService.Update(request.Context(), id, org)
	if err != nil {
		response.Error(err)
	}

	response.JSON(http.StatusOK, dto.SuccessResponse)
}

// Delete 删除组织信息
func (co *Controller) Delete(response *goyave.Response, request *goyave.Request) {
	id := request.RouteParams["id"]

	err := co.organizationService.Delete(request.Context(), id)
	if err != nil {
		response.Error(err)
	}

	response.JSON(http.StatusOK, dto.SuccessResponse)
}

// FindByID 根据id获取组织信息
func (co *Controller) FindByID(response *goyave.Response, request *goyave.Request) {
	id := request.RouteParams["id"]

	org, err := co.organizationService.FindByID(request.Context(), id)
	if err != nil {
		response.Error(err)
	}

	response.JSON(http.StatusOK, dto.CommonResponse{
		Message: dto.SuccessMessage,
		Data:    org,
	})
}
