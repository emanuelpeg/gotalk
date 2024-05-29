package logic

import (
	"context"
	"gotalk/gotalk"
	"gotalk/internal/svc"
	"sync"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListenerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	msgs sync.Map
}

var listenerLogic *ListenerLogic

func NewListenerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListenerLogic {
	if listenerLogic == nil {
		listenerLogic = &ListenerLogic{
			ctx:    ctx,
			svcCtx: svcCtx,
			Logger: logx.WithContext(ctx),
			//msgs:   make(map[string]*gotalk.ChatMessage),
		}
	}
	return listenerLogic
}

func (l *ListenerLogic) Listener(in *gotalk.SessionId, stream gotalk.ChatService_ListenerServer) error {
	for {
		value, ok := l.msgs.Load(in.SessionId)
		if ok {
			if err := stream.Send(value.(*gotalk.ChatMessage)); err != nil {
				return err
			}
			l.msgs.Delete(in.SessionId)
		}
	}
}
