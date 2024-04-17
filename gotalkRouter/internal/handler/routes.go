package handler

import (
	"net/http"

	"gotalkRouter/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/from/:name",
				Handler: GotalkRouterHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/health",
				Handler: HealthCheckHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/users",
				Handler: UsersHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/users",
				Handler: AddUsersHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/users",
				Handler: DeleteUsersHandler(serverCtx),
			},
		},
	)
}
