// Code generated by goctl. DO NOT EDIT.
// Source: gotalkgRPC.proto

package hello

import (
	"context"
	"gotalk/gotalk"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = gotalk.Request
	Response = gotalk.Response

	Hello interface {
		Hi(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultHello struct {
		cli zrpc.Client
	}
)

func NewHello(cli zrpc.Client) Hello {
	return &defaultHello{
		cli: cli,
	}
}

func (m *defaultHello) Hi(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := gotalk.NewHelloClient(m.cli.Conn())
	return client.Hi(ctx, in, opts...)
}