package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"video_clip/cmd/api/internal/config"
	"video_clip/cmd/auth/rpc/authclient"
	"video_clip/cmd/comment/rpc/commentclient"
	"video_clip/cmd/follow/rpc/followclient"
	"video_clip/cmd/like/rpc/likeclient"
	"video_clip/cmd/msg/rpc/msgclient"
	"video_clip/cmd/user/rpc/userclient"
	"video_clip/cmd/video/rpc/videoclient"
)

type ServiceContext struct {
	Config config.Config

	AuthClient    authclient.Auth
	UserClient    userclient.User
	VideoClient   videoclient.Video
	MsgClient     msgclient.Msg
	LikeClient    likeclient.Like
	FollowClient  followclient.Follow
	CommentClient commentclient.Comment
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		AuthClient:    authclient.NewAuth(zrpc.MustNewClient(c.AuthRpcConf)),
		UserClient:    userclient.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		VideoClient:   videoclient.NewVideo(zrpc.MustNewClient(c.VideoRpcConf)),
		MsgClient:     msgclient.NewMsg(zrpc.MustNewClient(c.MsgRpcConf)),
		LikeClient:    likeclient.NewLike(zrpc.MustNewClient(c.LikeRpcConf)),
		FollowClient:  followclient.NewFollow(zrpc.MustNewClient(c.FollowRpcConf)),
		CommentClient: commentclient.NewComment(zrpc.MustNewClient(c.CommentRpcConf)),
	}
}
