// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"video_clip/service/api-service/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: LoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JWT},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/",
					Handler: getUserInfoHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/v1/user"),
	)
}
