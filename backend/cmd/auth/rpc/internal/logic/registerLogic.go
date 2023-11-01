package logic

import (
	"context"

	"video_clip/cmd/auth/rpc/auth"
	"video_clip/cmd/auth/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 注册接口
func (l *RegisterLogic) Register(in *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	// todo: add your logic here and delete this line

	return &auth.RegisterResponse{}, nil
}
