// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	payment "zero-shop/payment/api/internal/handler/payment"
	"zero-shop/payment/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckUserState},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/payment",
					Handler: payment.OrderPayHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/payment/v1"),
	)
}
