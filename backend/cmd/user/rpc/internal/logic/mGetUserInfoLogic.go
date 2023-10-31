package logic

import (
	"context"

	"video_clip/cmd/user/rpc/internal/svc"
	"video_clip/cmd/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type MGetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MGetUserInfoLogic {
	return &MGetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量获得用户信息接口
func (l *MGetUserInfoLogic) MGetUserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &user.UserInfoResponse{}, nil
}
