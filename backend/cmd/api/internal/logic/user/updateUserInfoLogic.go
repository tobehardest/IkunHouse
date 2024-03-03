package user

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"video_clip/cmd/api/internal/svc"
	"video_clip/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"video_clip/cmd/user/rpc/user"
	"video_clip/pkg/errx"
)

type UpdateUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(req *types.UpdateUserInfoReq) (resp *types.UpdateUserInfoRes, err error) {
	in := &user.UpdateUserInfoReq{}
	err = copier.Copy(in, req)
	if err != nil {
		errors.Wrapf(errx.NewErrCode(errx.COPIER_COPY_ERROR), "userId:%s", req.UserId)
	}
	_, err = l.svcCtx.UserClient.UpdateUserInfo(l.ctx, in)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
