package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"video_clip/cmd/api/internal/config"
	"video_clip/cmd/auth/rpc/authCenterClient"
	"video_clip/cmd/collect/rpc/collectClient"
	"video_clip/cmd/comment/rpc/commentClient"
	"video_clip/cmd/follow/rpc/followclient"
	"video_clip/cmd/like/rpc/likeClient"
	"video_clip/cmd/msg/rpc/msgclient"
	"video_clip/cmd/search/rpc/searchClient"
	"video_clip/cmd/user/rpc/userClient"
	"video_clip/cmd/video/rpc/videoclient"
)

type ServiceContext struct {
	Config config.Config

	AuthClient    authCenterClient.AuthCenter
	UserClient    userClient.User
	VideoClient   videoclient.Video
	MsgClient     msgclient.Msg
	LikeClient    likeClient.Like
	FollowClient  followclient.Follow
	CommentClient commentClient.Comment
	CollectClient collectClient.Collect
	SearchClient  searchClient.Search
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		AuthClient: authCenterClient.NewAuthCenter(zrpc.MustNewClient(c.AuthRpcConf)),
		//UserClient:    userClient.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		//VideoClient:   videoclient.NewVideo(zrpc.MustNewClient(c.VideoRpcConf)),
		//MsgClient:     msgclient.NewMsg(zrpc.MustNewClient(c.MsgRpcConf)),
		//LikeClient:    likeClient.NewLike(zrpc.MustNewClient(c.LikeRpcConf)),
		//FollowClient:  followclient.NewFollow(zrpc.MustNewClient(c.FollowRpcConf)),
		//CommentClient: commentClient.NewComment(zrpc.MustNewClient(c.CommentRpcConf)),
		//CollectClient: collectClient.NewCollect(zrpc.MustNewClient(c.CollectRpcConf)),
		//SearchClient:  searchClient.NewSearch(zrpc.MustNewClient(c.SearchRpcConf)),
	}
}
