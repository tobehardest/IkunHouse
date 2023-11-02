package code

import "video_clip/pkg/errx"

var (
	ErrUserNoExistsError = errx.NewErrMsg("用户不存在")
	ErrorPasswordWrong   = errx.NewErrMsg("密码错误")
	ErrorGenIDFailed     = errx.NewErrMsg("创建用户ID失败")
)
