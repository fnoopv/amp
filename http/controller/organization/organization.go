package organization

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
	Paginate(ctx context.Context, request *filter.Request) (*database.PaginatorDTO[*dto.Organization], error)
	Create(ctx context.Context, organization *dto.OrganizationCreate) error
	Update(ctx context.Context, organization *dto.OrganizationUpdate) error
	Delete(ctx context.Context, ids []string) error
	FindByID(ctx context.Context, id string) (*dto.Organization, error)
	Option(ctx context.Context) ([]*dto.Organization, error)
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
	subRouter.Get("/", co.Index).ValidateQuery(filter.Validation)
	subRouter.Post("/", co.Create).ValidateBody(CreateRequest)
	subRouter.Get("/options", co.Option)

	subRouter.Get("/info/{id}", co.FindByID)
	subRouter.Post("/update", co.Update).ValidateBody(UpdateRequest)
	subRouter.Post("/delete", co.Delete).ValidateBody(DeleteRequest)
}

func (co *Controller) Index(response *goyave.Response, request *goyave.Request) {
	paginator, err := co.organizationService.Paginate(request.Context(), filter.NewRequest(request.Query))
	if response.WriteDBError(err) {
		return
	}

	response.JSON(http.StatusOK, dto.CommonResponse{
		Message: dto.ResponseSuccessMessage,
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

	response.JSON(http.StatusOK, dto.ResponseSuccess)
}

// Update 更新组织信息
func (co *Controller) Update(response *goyave.Response, request *goyave.Request) {
	req := typeutil.MustConvert[*dto.OrganizationUpdate](request.Data)

	err := co.organizationService.Update(request.Context(), req)
	if err != nil {
		response.Error(err)
	}

	response.JSON(http.StatusOK, dto.ResponseSuccess)
}

// Delete 删除组织信息
func (co *Controller) Delete(response *goyave.Response, request *goyave.Request) {
	req := typeutil.MustConvert[*dto.OrganizationDelete](request.Data)

	err := co.organizationService.Delete(request.Context(), req.IDs)
	if err != nil {
		response.Error(err)
	}

	response.JSON(http.StatusOK, dto.ResponseSuccess)
}

// FindByID 根据id获取组织信息
func (co *Controller) FindByID(response *goyave.Response, request *goyave.Request) {
	id := request.RouteParams["id"]

	org, err := co.organizationService.FindByID(request.Context(), id)
	if err != nil {
		response.Error(err)
	}

	response.JSON(http.StatusOK, dto.CommonResponse{
		Message: dto.ResponseSuccessMessage,
		Data:    org,
	})
}

func (co *Controller) Option(response *goyave.Response, request *goyave.Request) {
	orgs, err := co.organizationService.Option(request.Context())
	if err != nil {
		response.Error(err)
		return
	}

	response.JSON(http.StatusOK, dto.CommonResponse{
		Message: dto.ResponseSuccessMessage,
		Data:    orgs,
	})
}
