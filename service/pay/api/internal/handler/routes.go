// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"product/service/pay/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/pay/create",
				Handler: CreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/pay/detail",
				Handler: DetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/pay/callback",
				Handler: CallbackHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
