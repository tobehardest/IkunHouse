package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"video_clip/cmd/api/internal/config"
	"video_clip/cmd/msg/rpc/msgclient"
	"video_clip/cmd/user/rpc/userclient"
	"video_clip/cmd/video/rpc/videoclient"
)

type ServiceContext struct {
	Config      config.Config
	UserClient  userclient.User
	VideoClient videoclient.Video
	MsgClient   msgclient.Msg
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		UserClient:  userclient.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		VideoClient: videoclient.NewVideo(zrpc.MustNewClient(c.VideoRpcConf)),
		MsgClient:   msgclient.NewMsg(zrpc.MustNewClient(c.MsgRpcConf)),
	}
}
