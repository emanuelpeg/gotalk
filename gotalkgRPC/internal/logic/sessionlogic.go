package logic

import (
	"errors"
	"github.com/hashicorp/go-uuid"
	"gotalk/client"
	"gotalk/gotalk"
	"gotalk/internal/types"
	"sync"
)

var lock = &sync.Mutex{}

type SessionLogic struct {
	sessions map[string]gotalk.User
}

var sessionLogic *SessionLogic

func GetSessionLogic() *SessionLogic {
	if sessionLogic == nil {
		lock.Lock()
		defer lock.Unlock()
		if sessionLogic == nil {
			sessionLogic = &SessionLogic{
				sessions: make(map[string]gotalk.User),
			}
		}
	}

	return sessionLogic
}

func (sessionLogic *SessionLogic) StartSession(req *gotalk.RequestWithUserName) (resp *gotalk.ResponseWithSessionId, err error) {
	resp = new(gotalk.ResponseWithSessionId)
	user := new(gotalk.User)
	routerClient := new(client.RouterClient)

	err = sessionLogic.validateUserName(req.UserName)
	if err != nil {
		return nil, err
	}

	user.Name = req.UserName
	user.Ip, err = GetIp()
	if err == nil {
		if routerClient.Register(types.User{Name: user.Name, Ip: user.Ip}) {
			resp.Id, err = uuid.GenerateUUID()
			if err == nil {
				sessionLogic.sessions[resp.Id] = *user
			}
		} else {
			err = errors.New(`error in the registration process`)
		}
	}

	return
}

func (sessionLogic *SessionLogic) validateUserName(userName string) error {
	for _, user := range sessionLogic.sessions {
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

func (sessionLogic *SessionLogic) CloseSession(req *gotalk.RequestWithSessionId) (resp *gotalk.ResponseUser, err error) {
	resp = new(gotalk.ResponseUser)
	user, exists := sessionLogic.sessions[req.Id]

	if exists {
		routerClient := new(client.RouterClient)

		if routerClient.UnRegister(types.User{Name: user.Name, Ip: user.Ip}) {
			delete(sessionLogic.sessions, req.Id)
		} else {
			err = errors.New(`error in the deregistration process`)
		}
	}
	resp.Msg = "Bye!"
	return
}
