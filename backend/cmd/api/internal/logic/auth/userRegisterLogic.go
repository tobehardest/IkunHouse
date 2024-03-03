package auth

import (
	"IkunHouse/cmd/api/internal/svc"
	"IkunHouse/cmd/api/internal/types"
	"IkunHouse/cmd/auth/rpc/authCenter"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"IkunHouse/pkg/errx"
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

	registerResp, err := l.svcCtx.AuthClient.Register(l.ctx, &authCenter.RegisterReq{
		Mobile:   req.Mobile,
		Username: req.UserName,
		Password: req.Password,
		Avatar:   req.Avatar,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	resp = &types.RegisterResp{}
	err = copier.Copy(&resp, registerResp)
	if err != nil {
		errors.Wrapf(errx.NewErrCode(errx.COPIER_COPY_ERROR), "username:%s", req.UserName)
	}

	return resp, nil
}
