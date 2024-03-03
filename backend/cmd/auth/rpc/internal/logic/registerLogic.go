package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/openimsdk/openkf/server/pkg/utils"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"video_clip/cmd/auth/code"
	"video_clip/cmd/auth/rpc/authCenter"
	"video_clip/cmd/auth/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"video_clip/cmd/auth/model"
	"video_clip/pkg/errx"
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
func (l *RegisterLogic) Register(in *authCenter.RegisterReq) (*authCenter.RegisterRes, error) {

	// 1、判断用户存不存在
	user, err := l.svcCtx.UserModel.FindOneByUserName(l.ctx, in.Username)
	if err != nil && err != sqlc.ErrNotFound {
		// 数据库查询出错
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "userName:%s", in.Username)
	}
	if user != nil {
		// 用户已存在
		return nil, errors.Wrapf(code.ErrUserExistsed, "userName:%s", in.Username)
	}

	// 用户不存在
	// 2、生成UID
	userId := utils.GenUUIDWithoutHyphen()

	// 构造一个User实例
	user = &model.User{}
	err = copier.Copy(user, in)
	if err != nil {
		errors.Wrapf(errx.NewErrCode(errx.COPIER_COPY_ERROR), "username:%s", in.Username)
	}
	user.UserId = userId
	// 3、保存进数据库
	_, err = l.svcCtx.UserModel.Insert(l.ctx, user)
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "userName:%s", in.Username)
	}

	return &authCenter.RegisterRes{
		UserId: userId,
	}, nil
}
