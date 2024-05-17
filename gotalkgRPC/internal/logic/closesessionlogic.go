package logic

import (
	"context"

	"gotalk/gotalk"
	"gotalk/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CloseSessionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCloseSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CloseSessionLogic {
	return &CloseSessionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CloseSessionLogic) CloseSession(in *gotalk.RequestWithSessionId) (*gotalk.ResponseUser, error) {
	return GetSessionLogic().CloseSession(in)
}
