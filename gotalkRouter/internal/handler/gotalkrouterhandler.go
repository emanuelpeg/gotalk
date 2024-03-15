package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gotalkRouter/internal/logic"
	"gotalkRouter/internal/svc"
	"gotalkRouter/internal/types"
)

func GotalkRouterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGotalkRouterLogic(r.Context(), svcCtx)
		resp, err := l.GotalkRouter(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
