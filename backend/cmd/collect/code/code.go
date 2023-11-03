package code

import "video_clip/pkg/errx"

var (
	CollectPostIdEmpty = errx.NewErrCodeMsg(500002, "收藏帖子id为空")
	CollectUserIdEmpty = errx.NewErrCodeMsg(500001, "收藏用户Id为空")
)
