package main

import (
	"database/sql"
	"embed"
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/fnoopv/amp/database/repository"
	"github.com/fnoopv/amp/http/route"
	"github.com/fnoopv/amp/pkg/external/redis"
	"github.com/fnoopv/amp/pkg/migrate"
	"github.com/fnoopv/amp/service/application"
	"github.com/fnoopv/amp/service/attachment"
	"github.com/fnoopv/amp/service/domain"
	"github.com/fnoopv/amp/service/evaluation"
	"github.com/fnoopv/amp/service/filling"
	"github.com/fnoopv/amp/service/network"
	"github.com/fnoopv/amp/service/organization"
	"github.com/fnoopv/amp/service/user"

	"goyave.dev/filter"
	"goyave.dev/goyave/v5"
	_ "goyave.dev/goyave/v5/database/dialect/postgres"
	"goyave.dev/goyave/v5/util/errors"
	"goyave.dev/goyave/v5/util/fsutil"
	"goyave.dev/goyave/v5/util/session"
)

//go:embed resources
var resources embed.FS

//go:embed database/migrations/*.sql
var migrations embed.FS

func main() {
	resources := fsutil.NewEmbed(resources)
	langFS, err := resources.Sub("resources/lang")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.(*errors.Error).String())
		os.Exit(1)
	}

	opts := goyave.Options{
		LangFS: langFS,
	}

	server, err := goyave.New(opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.(*errors.Error).String())
		os.Exit(1)
	}

	server.Logger.Info("Registering hooks")
	server.RegisterSignalHook()

	server.RegisterStartupHook(func(s *goyave.Server) {
		filter.QueryParamPage = "currentPage"
		filter.QueryParamPerPage = "pageSize"
		// 迁移数据库表
		server.Logger.Info("Migrate database tables ...")
		hostPort := net.JoinHostPort(
			s.Config().GetString("database.host"),
			strconv.Itoa(s.Config().GetInt("database.port")),
		)
		dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
			s.Config().GetString("database.username"),
			s.Config().GetString("database.password"),
			hostPort,
			s.Config().GetString("database.name"))
		err := migrate.Migrate(dsn, migrations)
		if err != nil {
			server.Logger.Error(err)
			os.Exit(3)
		}
		server.Logger.Info("Migrate database tables success")

		// 连接redis
		server.Logger.Info("Connect to redis server ...")
		redisAddress := fmt.Sprintf("%s:%d",
			s.Config().GetString("redis.host"),
			s.Config().GetInt("redis.port"))
		if err := redis.Initialize(redisAddress); err != nil {
			server.Logger.Error(err)
			os.Exit(4)
		}
		server.Logger.Info("Connect redis server success")

		server.Logger.Info("Server is listening", "host", s.Host())
	})

	server.RegisterShutdownHook(func(s *goyave.Server) {
		s.Logger.Info("Close redis connection ...")
		redis.Client.Close()
		s.Logger.Info("Close redis connection success")

		s.Logger.Info("Server is shutting down")
	})

	registerServices(server)

	server.Logger.Info("Registering routes")
	server.RegisterRoutes(route.Register)

	if err := server.Start(); err != nil {
		server.Logger.Error(err)
		os.Exit(2)
	}
}

func registerServices(server *goyave.Server) {
	server.Logger.Info("Registering services")

	session := session.GORM(server.DB(), &sql.TxOptions{})
	businessAttachmentRepository := repository.NewBusinessAttachment(server.DB())

	attachmentRepository := repository.NewAttachment(server.DB())
	attachmentService := attachment.NewService(attachmentRepository)
	server.RegisterService(attachmentService)

	userRepository := repository.NewUser(server.DB())
	userService := user.NewService(userRepository)
	server.RegisterService(userService)

	organizationRepository := repository.NewOrganization(server.DB())
	organizationService := organization.NewService(organizationRepository)
	server.RegisterService(organizationService)

	fillingRepository := repository.NewFilling(server.DB())
	fillingService := filling.NewService(
		session,
		fillingRepository,
		businessAttachmentRepository,
		attachmentRepository,
	)
	server.RegisterService(fillingService)

	networkRepository := repository.NewNetwork(server.DB())
	networkService := network.NewService(session, networkRepository, fillingRepository)
	server.RegisterService(networkService)

	domainRepository := repository.NewDomain(server.DB())
	domainService := domain.NewService(session, domainRepository, fillingRepository)
	server.RegisterService(domainService)

	appRepository := repository.NewApplication(server.DB())
	appService := application.NewService(session, appRepository, fillingRepository, networkRepository)
	server.RegisterService(appService)

	evaluationRepository := repository.NewEvaluation(server.DB())
	evaluationService := evaluation.NewService(
		session,
		evaluationRepository,
		businessAttachmentRepository,
		attachmentRepository,
	)
	server.RegisterService(evaluationService)
}
