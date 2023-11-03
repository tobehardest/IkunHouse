package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Es struct {
		Addresses []string
		Username  string
		Password  string
	}
}
