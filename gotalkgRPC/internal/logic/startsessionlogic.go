package logic

import (
	"context"

	"gotalk/gotalk"
	"gotalk/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type StartSessionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStartSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartSessionLogic {
	return &StartSessionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StartSessionLogic) StartSession(in *gotalk.RequestWithUserName) (*gotalk.ResponseWithSessionId, error) {
	return GetSessionLogic().StartSession(in)
}
