// Code generated by goctl. DO NOT EDIT.
// Source: auth.proto

package server

import (
	"context"

	"video_clip/cmd/auth/rpc/auth"
	"video_clip/cmd/auth/rpc/internal/logic"
	"video_clip/cmd/auth/rpc/internal/svc"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	auth.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

// 注册接口
func (s *UserServer) Register(ctx context.Context, in *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

// 登录接口
func (s *UserServer) Login(ctx context.Context, in *auth.LoginRequest) (*auth.LoginResponse, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}
