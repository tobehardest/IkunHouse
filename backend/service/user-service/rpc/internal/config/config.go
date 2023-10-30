package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"k8s.io/client-go/tools/cache"
)

type Config struct {
	zrpc.RpcServerConf
	ViderRpcConf zrpc.RpcClientConf

	CacheRedis cache.Config

	MysqlCluster struct {
		DataSource string
	}

	RedisCluster struct {
		Cluster string
	}

	//Token struct{
	//	AccessToken string
	//}

}
