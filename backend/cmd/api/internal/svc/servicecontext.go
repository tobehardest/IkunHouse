package svc

import (
	"IkunHouse/cmd/api/internal/config"
	"IkunHouse/cmd/auth/rpc/authCenterClient"
	"IkunHouse/cmd/collect/rpc/collectClient"
	"IkunHouse/cmd/comment/rpc/commentClient"
	"IkunHouse/cmd/follow/rpc/followClient"
	"IkunHouse/cmd/like/rpc/likeClient"
	"IkunHouse/cmd/msg/rpc/msgclient"
	"IkunHouse/cmd/search/rpc/searchClient"
	"IkunHouse/cmd/user/rpc/userClient"
	"IkunHouse/cmd/video/rpc/videoclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	AuthClient    authCenterClient.AuthCenter
	UserClient    userClient.User
	VideoClient   videoclient.Video
	MsgClient     msgclient.Msg
	LikeClient    likeClient.Like
	FollowClient  followClient.Follow
	CommentClient commentClient.Comment
	CollectClient collectClient.Collect
	SearchClient  searchClient.Search
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		AuthClient: authCenterClient.NewAuthCenter(zrpc.MustNewClient(c.AuthRpcConf)),
		UserClient: userClient.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		//VideoClient:   videoclient.NewVideo(zrpc.MustNewClient(c.VideoRpcConf)),
		//MsgClient:     msgclient.NewMsg(zrpc.MustNewClient(c.MsgRpcConf)),
		//LikeClient:    likeClient.NewLike(zrpc.MustNewClient(c.LikeRpcConf)),
		FollowClient: followClient.NewFollow(zrpc.MustNewClient(c.FollowRpcConf)),
		//CommentClient: commentClient.NewComment(zrpc.MustNewClient(c.CommentRpcConf)),
		//CollectClient: collectClient.NewCollect(zrpc.MustNewClient(c.CollectRpcConf)),
		//SearchClient:  searchClient.NewSearch(zrpc.MustNewClient(c.SearchRpcConf)),
	}
}
