package code

import "video_clip/pkg/errx"

var (
	FollowUserIdEmpty   = errx.NewErrCodeMsg(40001, "关注用户id为空")
	FollowedUserIdEmpty = errx.NewErrCodeMsg(40002, "被关注用户id为空")
	CannotFollowSelf    = errx.NewErrCodeMsg(40003, "不能关注自己")
	UserIdEmpty         = errx.NewErrCodeMsg(40004, "用户id为空")
)
