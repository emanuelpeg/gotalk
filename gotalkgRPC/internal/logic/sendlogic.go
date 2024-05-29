package logic

import (
	"context"
	"gotalk/gotalk"
	"gotalk/internal/svc"
	"io"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendLogic {
	return &SendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendLogic) Send(stream gotalk.ChatService_SendServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&gotalk.Empty{})
		}
		if err != nil {
			return err
		}
		l.Infof("Received message from %s: %s", msg.User, msg.Message)
		sessionId := getSessionId(msg.User, msg.Ip)

		if sessionId != nil {
			NewListenerLogic(l.ctx, l.svcCtx).msgs.Store(*sessionId, msg)
		}

	}
}

func getSessionId(userName string, ip string) *string {
	for sessionId, user := range GetSessionLogic().sessions {
		if user.Name == userName && user.Ip == ip {
			return &sessionId
		}
	}
	return nil
}
