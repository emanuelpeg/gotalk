package logic

import (
	"context"
	"gotalk/client"
	"gotalk/internal/types"

	"gotalk/gotalk"
	"gotalk/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	sessions map[string]types.User
}

func NewUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UsersLogic {
	return &UsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UsersLogic) Users(_ *gotalk.RequestUser) (*gotalk.ResponseUsers, error) {
	routerClient := new(client.RouterClient)
	respUser, err := routerClient.Users()
	var users = make([]*gotalk.User, len(respUser.Users))
	for i, user := range respUser.Users {
		users[i] = &gotalk.User{Name: user.Name, Ip: user.Ip}
	}

	return &gotalk.ResponseUsers{Users: users}, err
}
