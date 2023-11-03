package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	UserRpcConf zrpc.RpcClientConf
	MysqlConf   struct {
		Datasource string
		//Address  string
		//Username string
		//Password string
		//DbName   string
	}
}
