package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	// 启动配置
	rest.RestConf

	// RPC配置
	UserRpc zrpc.RpcClientConf

	// JWT 配置
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	//OSS
	OSS struct {
		AccessKey string
		SecretKey string
		Bucket    string // 存储空间名
	}
}
