package logic

import (
	"context"
	"strconv"
	"time"

	"IkunHouse/cmd/collect/rpc/collect"
	"IkunHouse/cmd/collect/rpc/internal/svc"

	"IkunHouse/cmd/collect/code"
	"IkunHouse/cmd/collect/model"
	"IkunHouse/cmd/collect/rpc/internal/types"
	"IkunHouse/pkg/utils"
	"github.com/zeromicro/go-zero/core/logx"
)

type CollectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCollectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectLogic {
	return &CollectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 收藏
func (l *CollectLogic) Collect(in *collect.CollectReq) (*collect.CollectRes, error) {
	// 1. 校验
	if in.UserId == 0 {
		return nil, code.CollectUserIdEmpty
	}
	if in.PostId == 0 {
		return nil, code.CollectPostIdEmpty
	}

	collectRecord := &model.CollectRecord{
		BizId:  "1",
		ObjId:  in.PostId,
		UserId: in.UserId,
	}

	// 2. 更新
	// 2.1 保存数据库
	_, err := l.svcCtx.CollectModel.Insert(l.ctx, collectRecord)
	if err != nil {
		// todo
		return nil, err
	}

	// 2.2 更新redis
	collectKey := utils.GetEntityCollectKey(in.UserId, in.BizId)
	collectExist, err := l.svcCtx.BizRedis.ExistsCtx(l.ctx, collectKey)
	if err != nil {
		l.Logger.Errorf("[Follow] Redis Exists error: %v", err)
		return nil, err
	}
	if collectExist {
		_, err = l.svcCtx.BizRedis.ZaddCtx(l.ctx, collectKey, time.Now().Unix(), strconv.FormatInt(in.PostId, 10))
		if err != nil {
			l.Logger.Errorf("[Follow] Redis Zadd error: %v", err)
			return nil, err
		}
		_, err = l.svcCtx.BizRedis.ZremrangebyrankCtx(l.ctx, collectKey, 0, -(types.CacheMaxCollectCount + 1))
		if err != nil {
			l.Logger.Errorf("[Follow] Redis Zremrangebyrank error: %v", err)
		}
	}

	return &collect.CollectRes{}, nil
}
