package attachment

import (
	"context"
	"net/http"
	"os"
	"path"
	"reflect"

	"github.com/dromara/carbon/v2"
	"github.com/fnoopv/amp/dto"
	"github.com/fnoopv/amp/pkg/sha256sum"
	"github.com/fnoopv/amp/pkg/uid"
	"github.com/fnoopv/amp/service"
	"github.com/google/uuid"
	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/config"
	"goyave.dev/goyave/v5/util/fsutil"
	"goyave.dev/goyave/v5/util/fsutil/osfs"
	"goyave.dev/goyave/v5/util/typeutil"
)

// 注册配置项
func init() {
	config.Register("attachment.directory", config.Entry{
		Value:    "attachments",
		Type:     reflect.String,
		IsSlice:  false,
		Required: false,
	},
	)
}

type Service interface {
	Create(ctx context.Context, att *dto.AttachmentCreate) error
	FindByID(ctx context.Context, id string) (*dto.AttachmentInternal, error)
}

type Controller struct {
	goyave.Component
	AttachmentService Service
}

func (co *Controller) Init(server *goyave.Server) {
	co.AttachmentService = server.Service(service.Attachment).(Service)
	co.Component.Init(server)
}

func (co *Controller) RegisterRoutes(router *goyave.Router) {
	subRouter := router.Subrouter("/attachments")
	subRouter.Post("/", co.Upload)
	subRouter.Get("/info/{id}", co.Info)
	subRouter.Get("/download/{id}", co.Download)
}

func (co *Controller) Upload(response *goyave.Response, request *goyave.Request) {
	user := request.User.(*dto.UserInternal)
	data := request.Data.(map[string]any)
	files := data["files"].([]fsutil.File)

	id, err := uid.Generate()
	if err != nil {
		response.Error(err)
		return
	}

	for _, f := range files {
		storagePath, err := f.Save(&osfs.FS{}, co.Config().GetString("attachment.directory"), id)
		if err != nil {
			response.Error(err)
			return
		}

		sum, err := sha256sum.CalcuLateSHA256Sum(path.Join(co.Config().GetString("attachment.directory"), storagePath))
		if err != nil {
			response.Error(err)
			return
		}

		createAtt := &dto.AttachmentCreate{
			ID:          id,
			Name:        f.Header.Filename,
			Size:        f.Header.Size,
			StoragePath: storagePath,
			Mime:        f.MIMEType,
			UploadAt:    carbon.NewDateTime(carbon.Now()),
			UploaderID:  user.ID,
			SHA256Sum:   sum,
		}

		if err := co.AttachmentService.Create(request.Context(), createAtt); err != nil {
			response.Error(err)
			return
		}

		att := typeutil.MustConvert[*dto.Attachment](createAtt)

		response.JSON(http.StatusOK, dto.CommonResponse{
			Message: "success",
			Data:    att,
		})
	}
}

func (co *Controller) Info(response *goyave.Response, request *goyave.Request) {
	fid := request.RouteParams["id"]
	if err := uuid.Validate(fid); err != nil {
		response.JSON(http.StatusBadRequest, dto.CommonResponse{
			Message: "参数错误",
		})

		return
	}

	att, err := co.AttachmentService.FindByID(request.Context(), fid)
	if err != nil {
		response.Error(err)
		return
	}

	response.JSON(http.StatusOK, dto.CommonResponse{
		Message: dto.SuccessMessage,
		Data:    typeutil.MustConvert[*dto.Attachment](att),
	})
}

func (co *Controller) Download(response *goyave.Response, request *goyave.Request) {
	fid := request.RouteParams["id"]
	if err := uuid.Validate(fid); err != nil {
		response.JSON(http.StatusBadRequest, dto.CommonResponse{
			Message: "参数错误",
		})

		return
	}

	att, err := co.AttachmentService.FindByID(request.Context(), fid)
	if err != nil {
		response.Error(err)
		return
	}

	root, err := os.OpenRoot(co.Config().GetString("attachment.directory"))
	if err != nil {
		response.Error(err)
		return
	}
	fs := root.FS().(fsutil.FS)

	response.Download(fs, att.StoragePath, att.Name)
}
