package logic

import (
	"context"
	"errors"
	"github.com/hashicorp/go-uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"gotalk/client"
	"gotalk/internal/svc"
	"gotalk/internal/types"
	"sync"
)

var lock = &sync.Mutex{}

type UsersLogic struct {
	logx.Logger
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	sessions map[string]types.User
}

var usersLogic *UsersLogic

func GetUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UsersLogic {
	if usersLogic == nil {
		lock.Lock()
		defer lock.Unlock()
		if usersLogic == nil {
			usersLogic = &UsersLogic{
				Logger:   logx.WithContext(ctx),
				ctx:      ctx,
				svcCtx:   svcCtx,
				sessions: make(map[string]types.User),
			}
		}
	}

	return usersLogic
}

func (usersLogic *UsersLogic) StartSession(req *types.RequestWithUserName) (resp *types.ResponseWithSessionId, err error) {
	resp = new(types.ResponseWithSessionId)
	user := new(types.User)
	routerClient := new(client.RouterClient)

	err = usersLogic.validateUserName(req.UserName)
	if err != nil {
		return nil, err
	}

	user.Name = req.UserName
	user.Ip, err = GetIp()
	if err == nil {
		if routerClient.Register(*user) {
			resp.Id, err = uuid.GenerateUUID()
			if err == nil {
				usersLogic.sessions[resp.Id] = *user
			}
		} else {
			err = errors.New(`error in the registration process`)
		}
	}

	return
}

func (usersLogic *UsersLogic) validateUserName(userName string) error {
	for _, user := range usersLogic.sessions {
		if user.Name == userName {
			return errors.New(`the user name is already in use`)
		}
	}
	return nil
}

func GetIp() (ipStr string, err error) {
	/*ifaces, err := net.Interfaces()

	if err != nil {
		for _, i := range ifaces {
			addrs, err := i.Addrs()

			if err != nil {
				for _, addr := range addrs {
					var ip net.IP
					switch v := addr.(type) {
					case *net.IPNet:
						ip = v.IP
					case *net.IPAddr:
						ip = v.IP
					}
					ipStr = ip.String()
				}
			}
		}
	}*/

	return "127.0.0.1", nil
}

func (usersLogic *UsersLogic) CloseSession(req *types.RequestWithSessionId) (resp *types.Response, err error) {
	resp = new(types.Response)
	user, exists := usersLogic.sessions[req.Id]

	if exists {
		routerClient := new(client.RouterClient)

		if routerClient.UnRegister(user) {
			delete(usersLogic.sessions, req.Id)
		} else {
			err = errors.New(`error in the deregistration process`)
		}
	}

	return
}

func (usersLogic *UsersLogic) Users() (resp *types.ResponseUsers, err error) {
	resp = new(types.ResponseUsers)
	routerClient := new(client.RouterClient)
	respUser, err := routerClient.Users()
	resp.Users = respUser.Users
	return
}
