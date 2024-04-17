package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gotalkRouter/internal/svc"
	"gotalkRouter/internal/types"
)

type UsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UsersLogic {
	return &UsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (usersLogic *UsersLogic) AddUser(req *types.RequestWithUser) (resp *types.ResponseWithMessage, err error) {
	resp = new(types.ResponseWithMessage)
	types.AddUser(req.User)
	resp.Message = "user add success"
	return
}

func (usersLogic *UsersLogic) DeleteUser(req *types.RequestWithUser) (resp *types.ResponseWithMessage, err error) {
	resp = new(types.ResponseWithMessage)
	if types.DeleteUser(req.User) {
		resp.Message = "user delete success"
	} else {
		resp.Message = "user not delete"
	}

	return
}

func (usersLogic *UsersLogic) Users() (resp *types.ResponseUsers, err error) {
	resp = new(types.ResponseUsers)
	resp.Users = types.GetUserList().Users
	return
}
