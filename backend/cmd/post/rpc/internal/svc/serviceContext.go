package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"video_clip/cmd/post/model"
	"video_clip/cmd/post/rpc/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	SqlConn      sqlx.SqlConn
	PostModel    model.PostModel
	TagModel     model.TagModel
	PostTagModel model.PostTagModel

	// RPC客户端
	// 点赞服务 - 获得是否点赞
	// LikeRpc likerpc.Like
	// 用户服务 - 获得用户信息
	//UserRpc userclient.User
	// 收藏服务 - 获得是否收藏
	//CollectRpc collectclient.Collect
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.MysqlConf.Datasource)
	return &ServiceContext{
		Config:    c,
		SqlConn:   conn,
		PostModel: model.NewPostModel(conn),
		TagModel:  model.NewTagModel(conn),
		//UserRpc:   userclient.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		//LikeRpc: likeclient.NewLike(zrpc.MustNewClient(c.LikeRpcConf)),
	}
}
