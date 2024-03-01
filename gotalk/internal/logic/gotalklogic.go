package logic

import (
	"context"

	"gotalk/internal/svc"
	"gotalk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GotalkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGotalkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GotalkLogic {
	return &GotalkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GotalkLogic) Gotalk(req *types.Request) (resp *types.Response, err error) {
	resp = new(types.Response)
	resp.Message = req.Name
	return
}
