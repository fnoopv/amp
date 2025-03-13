package route

import (
	"github.com/fnoopv/amp/dto"
	"github.com/fnoopv/amp/http/controller/application"
	"github.com/fnoopv/amp/http/controller/attachment"
	"github.com/fnoopv/amp/http/controller/filling"
	"github.com/fnoopv/amp/http/controller/organization"
	"github.com/fnoopv/amp/http/controller/user"
	"github.com/fnoopv/amp/http/middleware"
	"github.com/fnoopv/amp/service"
	"github.com/golang-jwt/jwt"
	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/auth"
	"goyave.dev/goyave/v5/cors"
	"goyave.dev/goyave/v5/middleware/parse"
)

// Routing is an essential part of any Goyave application.
// Defining routes is the action of associating a URI, sometimes having parameters,
// with a handler which will process the request and respond to it.
//
// This file contains your main route registering function that is passed to server.RegisterRoutes().
//
// Learn more here: https://goyave.dev/basics/routing.html
func Register(server *goyave.Server, router *goyave.Router) {
	apiV1 := router.Subrouter("/api/v1")
	apiV1.CORS(cors.Default())
	apiV1.GlobalMiddleware(&middleware.TraceMiddleare{})
	apiV1.GlobalMiddleware(&parse.Middleware{})

	// 登录控制器
	userService := server.Service(service.User).(auth.UserService[dto.UserInternal])
	jwtController := auth.NewJWTController(userService, "Password")
	jwtController.SigningMethod = jwt.SigningMethodHS256
	jwtController.TokenFunc = func(request *goyave.Request, user *dto.UserInternal) (string, error) {
		jwtService := server.Service(auth.JWTServiceName).(*auth.JWTService)
		return jwtService.GenerateTokenWithClaims(jwt.MapClaims{
			"id":              user.ID,
			"sub":             user.Username,
			"is_mfa_verified": user.IsMFAVerified,
		}, jwt.SigningMethodHS256)
	}
	apiV1.Controller(jwtController)

	// 认证中间件
	authenticator := auth.NewJWTAuthenticator(userService)
	authenticator.SigningMethod = jwt.SigningMethodHS256
	authMiddleware := auth.Middleware(authenticator)
	apiV1.GlobalMiddleware(authMiddleware).SetMeta(auth.MetaAuth, true)

	apiV1.Controller(&user.Controller{})
	apiV1.Controller(&organization.Controller{})
	apiV1.Controller(&application.Controller{})
	apiV1.Controller(&attachment.Controller{})
	apiV1.Controller(&filling.Controller{})
}
