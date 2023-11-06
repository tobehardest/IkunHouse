package logic

import (
	"context"

	"video_clip/cmd/collect/rpc/collect"
	"video_clip/cmd/collect/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"video_clip/cmd/collect/model"
)

type CollectListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCollectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectListLogic {
	return &CollectListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 收藏列表
func (l *CollectListLogic) CollectList(in *collect.CollectListReq) (*collect.CollectListRes, error) {
	// todo: add your logic here and delete this line

	return &collect.CollectListRes{}, nil
}

func (l *CollectListLogic) GetCollectIdsFromCache(ctx context.Context, userId, cursor, pageSize int64) ([]int64, error) {

	return nil, nil
}

func (l *CollectListLogic) AddCacheCollect(ctx context.Context, userId int64, collects []*model.CollectRecord) error {

	return nil
}
