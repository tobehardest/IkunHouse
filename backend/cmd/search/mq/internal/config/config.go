package config

import (
	"github.com/zeromicro/go-queue/kq"
)

type Config struct {
	UserKqConsumerConf       kq.KqConf
	ShortVideoKqConsumerConf kq.KqConf
	Es                       struct {
		Addresses []string
		Username  string
		Password  string
	}
}
