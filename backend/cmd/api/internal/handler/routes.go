// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	auth "video_clip/cmd/api/internal/handler/auth"
	comment "video_clip/cmd/api/internal/handler/comment"
	follow "video_clip/cmd/api/internal/handler/follow"
	like "video_clip/cmd/api/internal/handler/like"
	msg "video_clip/cmd/api/internal/handler/msg"
	user "video_clip/cmd/api/internal/handler/user"
	video "video_clip/cmd/api/internal/handler/video"
	"video_clip/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: auth.UserRegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: auth.UserLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/generate_token",
				Handler: auth.GenerateTokenHandler(serverCtx),
			},
		},
		rest.WithPrefix("/auth/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/getUserInfo",
				Handler: user.GetUserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/updateUserInfo",
				Handler: user.UpdateUserInfoHandler(serverCtx),
			},
		},
		rest.WithPrefix("/user/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/uploadVideo",
				Handler: video.UploadVideoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getHotVideoList",
				Handler: video.GetHotVideoListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getVideoListByCatagory",
				Handler: video.GetVideoListByCatagoryHandler(serverCtx),
			},
		},
		rest.WithPrefix("/video/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/like",
				Handler: like.LikeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/cancelLike",
				Handler: like.CancelLikeHandler(serverCtx),
			},
		},
		rest.WithPrefix("/like/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/follow",
				Handler: follow.FollowHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/unFollow",
				Handler: follow.UnFollowHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/follow/list",
				Handler: follow.FollowListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/fans/list",
				Handler: follow.FansListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/follow/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/comment",
				Handler: comment.CommentHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/cancel",
				Handler: comment.CancelCommentHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/list",
				Handler: comment.CommentListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/detail",
				Handler: comment.CommentDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: comment.DeleteCommentHandler(serverCtx),
			},
		},
		rest.WithPrefix("/comment/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/sendMsg",
				Handler: msg.SendMsgHandler(serverCtx),
			},
		},
		rest.WithPrefix("/msg/v1"),
	)
}
