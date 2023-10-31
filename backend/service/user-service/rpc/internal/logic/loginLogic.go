package logic

import (
	"context"

	"video_clip/service/user-service/rpc/internal/svc"
	"video_clip/service/user-service/rpc/types/user"

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
func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	return &user.LoginResponse{}, nil
}
