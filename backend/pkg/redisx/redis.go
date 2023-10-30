package redisx

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func InitRedis(redisCluster []string) *redis.ClusterClient {
	// Redis集群连接参数
	clusterOptions := &redis.ClusterOptions{
		Addrs: redisCluster,
	}

	// 创建redis集群客户端
	rdb := redis.NewClusterClient(clusterOptions)

	// 连接redis集群
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic("连接redis失败, error=" + err.Error())
	}
	fmt.Println("redis连接成功")
	return rdb
}
