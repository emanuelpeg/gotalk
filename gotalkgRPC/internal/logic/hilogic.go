package logic

import (
	"context"
	"gotalk/gotalk"
	"gotalk/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type HiLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HiLogic {
	return &HiLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HiLogic) Hi(_ *gotalk.Request) (*gotalk.Response, error) {
	return &gotalk.Response{Msg: "Hi!"}, nil
}
