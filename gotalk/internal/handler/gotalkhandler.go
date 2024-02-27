package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gotalk/gotalk/internal/logic"
	"gotalk/gotalk/internal/svc"
	"gotalk/gotalk/internal/types"
)

func GotalkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGotalkLogic(r.Context(), svcCtx)
		resp, err := l.Gotalk(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
