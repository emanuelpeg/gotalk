package logic

import (
	"context"

	"gotalk/gotalk"
	"gotalk/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type HealthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHealthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HealthLogic {
	return &HealthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HealthLogic) Health(_ *gotalk.RequestHealthCheck) (*gotalk.ResponseHealthCheck, error) {
	return &gotalk.ResponseHealthCheck{Msg: "ok"}, nil
}
