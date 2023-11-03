package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	UserRpcConf    zrpc.RpcClientConf
	LikeRpcConf    zrpc.RpcServerConf
	CollectRpcConf zrpc.RpcServerConf

	MysqlConf struct {
		Datasource string
	}
}
