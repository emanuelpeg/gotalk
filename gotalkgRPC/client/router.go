package client

import (
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"gotalk/internal/types"
)

type RouterClient struct {
}

func (routerClient *RouterClient) Register(user types.User) bool {
	client := resty.New()
	_, err := client.R().
		SetBody(types.RequestWithUser{User: user}).
		Put("http://127.0.0.1:8889/users")
	return err == nil
}

func (routerClient *RouterClient) UnRegister(user types.User) bool {
	client := resty.New()
	_, err := client.R().
		SetBody(types.RequestWithUser{User: user}).
		Delete("http://127.0.0.1:8889/users")
	return err == nil
}

func (routerClient *RouterClient) Users() (response types.ResponseUsers, err error) {
	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		Get("http://127.0.0.1:8889/users")

	if err == nil {
		if resp.IsSuccess() {
			err = json.Unmarshal(resp.Body(), &response)
		} else {
			err = errors.New(string(resp.StatusCode()) + ":" + resp.String())
		}
	}
	return
}
