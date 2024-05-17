// Code generated by goctl. DO NOT EDIT.
// Source: healthCheck.proto

package server

import (
	"context"

	"gotalk/gotalk"
	"gotalk/internal/logic"
	"gotalk/internal/svc"
)

type HealthCheckServer struct {
	svcCtx *svc.ServiceContext
	gotalk.UnimplementedHealthCheckServer
}

func NewHealthCheckServer(svcCtx *svc.ServiceContext) *HealthCheckServer {
	return &HealthCheckServer{
		svcCtx: svcCtx,
	}
}

func (s *HealthCheckServer) Health(ctx context.Context, in *gotalk.RequestHealthCheck) (*gotalk.ResponseHealthCheck, error) {
	l := logic.NewHealthLogic(ctx, s.svcCtx)
	return l.Health(in)
}