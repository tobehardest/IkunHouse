package code

import "video_clip/pkg/errx"

var (
	FollowUserIdEmpty   = errx.NewErrCodeMsg(400001, "关注用户id为空")
	FollowedUserIdEmpty = errx.NewErrCodeMsg(400002, "被关注用户id为空")
	CannotFollowSelf    = errx.NewErrCodeMsg(400003, "不能关注自己")
	UserIdEmpty         = errx.NewErrCodeMsg(400004, "用户id为空")
)
