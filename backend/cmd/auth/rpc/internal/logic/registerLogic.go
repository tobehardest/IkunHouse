package logic

import (
	"context"
	"github.com/pkg/errors"
	"strconv"
	"video_clip/cmd/auth/code"
	"video_clip/pkg/errx"
	"video_clip/pkg/uniqueid/snowflake"

	"video_clip/cmd/auth/rpc/auth"
	"video_clip/cmd/auth/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"video_clip/cmd/auth/model"
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
func (l *RegisterLogic) Register(in *auth.RegisterReq) (*auth.RegisterRes, error) {
	// todo: add your logic here and delete this line

	// 1、判断用户存不存在
	_, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		// 数据库查询出错
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "userName:%s", in.Username)
	}

	// 2、生成UID
	userId, err := snowflake.GetID()
	if err != nil {
		return nil, errors.Wrapf(code.ErrorGenIDFailed, "userName:%s", in.Username)
	}

	// 构造一个User实例
	user := &model.User{
		Id:       strconv.FormatUint(userId, 10),
		Username: in.Username,
		Password: in.Password,
	}
	// 3、保存进数据库
	l.svcCtx.UserModel.Insert(l.ctx, user)

	return &auth.RegisterRes{
		UserId: strconv.FormatUint(userId, 10),
	}, nil
}
