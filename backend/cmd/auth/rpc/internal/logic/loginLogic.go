package logic

import (
	"IkunHouse/cmd/auth/code"
	"IkunHouse/cmd/auth/rpc/authCenter"
	"IkunHouse/cmd/auth/rpc/internal/svc"
	"IkunHouse/pkg/errx"
	"IkunHouse/pkg/jwtx"
	"context"
	"database/sql"
	"fmt"
	"github.com/openimsdk/openkf/server/pkg/openim/param/request"
	"github.com/openimsdk/openkf/server/pkg/openim/param/response"
	imAuth "github.com/openimsdk/openkf/server/pkg/openim/sdk/auth"
	"github.com/pkg/errors"
	"net"

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
func (l *LoginLogic) Login(in *authCenter.LoginReq) (*authCenter.LoginRes, error) {
	// todo: add your logic here and delete this line
	//originPassword := in.Password // 记录一下原始密码(用户登录的密码)
	//sqlStr := "select user_id, username, password from user where username = ?"
	//
	//err = db.Get(user, sqlStr, user.UserName)
	user, err := l.svcCtx.UserModel.FindOneByUserName(l.ctx, in.Username)
	if err != nil && err != sql.ErrNoRows {
		// 查询数据库出错
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "userName:%s", in.Username)
	}
	if err == sql.ErrNoRows {
		// 用户不存在
		return nil, errors.Wrapf(code.ErrUserNoExistsError, "userName:%s", in.Username)
	}

	if !(user.UserName == in.Username && user.Password == in.Password) {
		return nil, errors.Wrapf(code.ErrorPasswordWrong, "userName:%s", in.Username)
	}

	// 生成加密密码与查询到的密码比较
	//encryptPassword := utils.Encrypt([]byte(originPassword))
	//if user.Password != encryptPassword {
	//	return nil, errors.Wrapf(code.ErrorPasswordWrong, "userName:%s", in.Username)
	//}
	vcToken, cvExpireTimeSeconds, err := jwtx.CreateToken(user.UserId, l.svcCtx.Config.JwtAuth.AccessSecret, l.svcCtx.Config.JwtAuth.AccessExpire)
	if err != nil {
		return nil, errors.Wrapf(code.ErrorGenIDFailed, "userName:%s", in.Username)
	}

	//imParam := &request.UserTokenParams{
	//	Secret:     l.svcCtx.Config.OpenIM.Secret,
	//	UserID:     user.UserId,
	//	PlatformID: uint(l.svcCtx.Config.OpenIM.PlatformID),
	//}
	//
	//imResp, err := l.GetUserIMToken(imParam)
	//if err != nil {
	//	return nil, err
	//}

	out := &authCenter.LoginRes{
		UserId: user.UserId,
		VCToken: &authCenter.TokenRes{
			Token:             vcToken,
			ExpireTimeSeconds: cvExpireTimeSeconds,
		},
		//IMToken: &auth.TokenRes{
		//	Token:             imResp.Token,
		//	ExpireTimeSeconds: uint64(imResp.ExpireTimeSeconds),
		//},
	}
	return out, nil
}

func (l *LoginLogic) GetUserIMToken(param *request.UserTokenParams) (*response.TokenData, error) {
	// TODO: Add operationID
	// Default not use tls/ssl
	host := fmt.Sprintf("http://%s", net.JoinHostPort(l.svcCtx.Config.OpenIM.Ip,
		fmt.Sprintf("%d", l.svcCtx.Config.OpenIM.ApiPort)))

	resp, err := imAuth.GetUserToken(param, "getUserIMToken", host)
	if err != nil {
		return &response.TokenData{}, err
	}

	if resp.ErrCode != 0 {
		return &response.TokenData{}, errors.New(resp.ErrMsg)
	}

	return &resp.Data, nil
}
