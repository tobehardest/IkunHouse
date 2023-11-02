package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	// rpc conf
	AuthRpcConf    zrpc.RpcClientConf
	UserRpcConf    zrpc.RpcClientConf
	VideoRpcConf   zrpc.RpcClientConf
	MsgRpcConf     zrpc.RpcClientConf
	LikeRpcConf    zrpc.RpcClientConf
	FollowRpcConf  zrpc.RpcClientConf
	CommentRpcConf zrpc.RpcClientConf
	CollectRpcConf zrpc.RpcClientConf
}
