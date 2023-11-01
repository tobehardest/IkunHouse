package auth

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"video_clip/cmd/auth/rpc/auth"

	"video_clip/cmd/api/internal/svc"
	"video_clip/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// todo: add your logic here and delete this line

	registerResp, err := l.svcCtx.AuthClient.Register(l.ctx, &auth.RegisterReq{
		Mobile:   req.Mobile,
		Username: req.UserName,
		Password: req.Password,
		Avatar:   req.Avatar,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	_ = copier.Copy(&resp, registerResp)

	return resp, nil
}
