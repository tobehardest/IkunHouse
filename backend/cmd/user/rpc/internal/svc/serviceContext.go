package svc

import (
	"github.com/bwmarrin/snowflake"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/gorm"
	"video_clip/cmd/user/rpc/internal/config"
	"video_clip/pkg/const"
	"video_clip/pkg/redisx"
)

type ServiceContext struct {
	Config config.Config
	//VideoRpc videoclient.Video
	Snowflake *snowflake.Node
	Rdb       *redis.ClusterClient
	Msqyl     *gorm.DB
	SqlConn   sqlx.SqlConn
	//UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	snowflakeNode, _ := snowflake.NewNode(_const.UserRpcMachineId)
	rc := make([]string, 0)
	rc = append(rc, c.RedisCluster.Cluster)
	redisDb := redisx.InitRedis(rc)
	mysqlConn := sqlx.NewMysql(c.MysqlCluster.DataSource)
	return &ServiceContext{
		Config:  c,
		SqlConn: mysqlConn,
		//UserModel: model.NewUserModel(mysqlConn,c.CacheRedis),
		//VideoRpc: videoclient.NewVideo(zrpc.MustNewClient(c.ViderRpcConf)),
		Snowflake: snowflakeNode,
		Rdb:       redisDb,
	}
}
