package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"video_clip/cmd/user/model"
	"video_clip/cmd/user/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	//VideoRpc videoclient.Video
	//Snowflake *snowflake.Node
	//Rdb       *redis.ClusterClient
	//Msqyl     *gorm.DB
	//SqlConn   sqlx.SqlConn
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	//snowflakeNode, _ := snowflake.NewNode(_const.UserRpcMachineId)
	//rc := make([]string, 0)
	//rc = append(rc, c.RedisCluster.Cluster)
	//redisDb := redisx.InitRedis(rc)
	//mysqlConn := sqlx.NewMysql(c.MysqlCluster.DataSource)

	//return &ServiceContext{
	//	Config:  c,
	//	SqlConn: mysqlConn,
	//	//UserModel: model.NewUserModel(mysqlConn,c.CacheRedis),
	//	//VideoRpc: videoclient.NewVideo(zrpc.MustNewClient(c.ViderRpcConf)),
	//	Snowflake: snowflakeNode,
	//	Rdb:       redisDb,
	//}
	conn := sqlx.NewMysql(c.DB.DataSource)
	//rds, err := redis.NewRedis(redis.RedisConf{
	//	Host: c.BizRedis.Host,
	//	Pass: c.BizRedis.Pass,
	//	Type: c.BizRedis.Type,
	//})
	//if err != nil {
	//	panic(err)
	//}

	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn),
	}
}
