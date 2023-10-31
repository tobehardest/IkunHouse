package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	// rpc conf
	UserRpcConf  zrpc.RpcClientConf
	VideoRpcConf zrpc.RpcClientConf
	MsgRpcConf   zrpc.RpcClientConf
}
