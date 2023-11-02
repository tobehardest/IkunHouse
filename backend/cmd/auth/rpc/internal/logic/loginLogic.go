package logic

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"video_clip/cmd/auth/code"
	"video_clip/pkg/errx"
	"video_clip/pkg/utils"

	"video_clip/cmd/auth/rpc/auth"
	"video_clip/cmd/auth/rpc/internal/svc"

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
func (l *LoginLogic) Login(in *auth.LoginReq) (*auth.LoginRes, error) {
	// todo: add your logic here and delete this line
	originPassword := in.Password // 记录一下原始密码(用户登录的密码)
	//sqlStr := "select user_id, username, password from user where username = ?"
	//
	//err = db.Get(user, sqlStr, user.UserName)
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil && err != sql.ErrNoRows {
		// 查询数据库出错
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "userName:%s", in.Username)
	}
	if err == sql.ErrNoRows {
		// 用户不存在
		return nil, errors.Wrapf(code.ErrUserNoExistsError, "userName:%s", in.Username)
	}

	// 生成加密密码与查询到的密码比较
	encryptPassword := utils.Encrypt([]byte(originPassword))
	if user.Password != encryptPassword {
		return nil, errors.Wrapf(code.ErrorPasswordWrong, "userName:%s", in.Username)
	}

	out := &auth.LoginRes{
		UserId: 0,
	}
	return out, nil
}
