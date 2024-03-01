package logic

import (
	"context"

	"gotalkclient/internal/svc"
	"gotalkclient/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GotalkclientLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGotalkclientLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GotalkclientLogic {
	return &GotalkclientLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GotalkclientLogic) Gotalkclient(req *types.Request) (resp *types.Response, err error) {
	resp = new(types.Response)
	resp.Message = "client - " + req.Name
	return
}
