package logic

import (
	"context"

	"video_clip/cmd/auth/rpc/auth"
	"video_clip/cmd/auth/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 登录接口
func (l *LoginLogic) Login(in *auth.LoginRequest) (*auth.LoginResponse, error) {
	// todo: add your logic here and delete this line

	return &auth.LoginResponse{}, nil
}
