package auth

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"video_clip/cmd/api/internal/svc"
	"video_clip/cmd/api/internal/types"
	"video_clip/cmd/auth/rpc/authCenter"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line
	loginResp, err := l.svcCtx.AuthClient.Login(l.ctx, &authCenter.LoginReq{
		Username: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.LoginResp{}
	err = copier.Copy(&resp, loginResp)
	if err != nil {
		// todo fill err
	}

	return resp, nil
}
