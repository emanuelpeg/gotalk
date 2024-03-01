package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gotalkclient/internal/logic"
	"gotalkclient/internal/svc"
	"gotalkclient/internal/types"
)

func GotalkclientHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGotalkclientLogic(r.Context(), svcCtx)
		resp, err := l.Gotalkclient(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
