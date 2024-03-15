package logic

import (
	"context"

	"gotalkRouter/internal/svc"
	"gotalkRouter/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GotalkRouterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGotalkRouterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GotalkRouterLogic {
	return &GotalkRouterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GotalkRouterLogic) GotalkRouter(req *types.Request) (resp *types.Response, err error) {
	resp = new(types.Response)
	resp.Message = req.Name
	return
}
