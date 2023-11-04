package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	DB struct {
		DataSource string
	}
	OpenIM OpenIM
}

type OpenIM struct {
	Secret     string
	Ip         string
	ApiPort    int
	PlatformID int
	AdminID    string
}
