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
	SearchRpcConf  zrpc.RpcClientConf
}

type EtcdConf struct {
	Hosts              []string
	Key                string
	ID                 int64  `json:",optional"`
	User               string `json:",optional"`
	Pass               string `json:",optional"`
	CertFile           string `json:",optional"`
	CertKeyFile        string `json:",optional=CertFile"`
	CACertFile         string `json:",optional=CertFile"`
	InsecureSkipVerify bool   `json:",optional"`
}
