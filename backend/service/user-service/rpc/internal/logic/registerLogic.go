package logic

import (
	"context"

	"video_clip/service/user-service/rpc/internal/svc"
	"video_clip/service/user-service/rpc/types/user"

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
func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {

	return &user.RegisterResponse{}, nil
}
