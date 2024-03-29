package logic

import (
	"IkunHouse/cmd/follow/rpc/follow"
	"IkunHouse/cmd/follow/rpc/internal/svc"
	"context"
	"gorm.io/gorm"

	"IkunHouse/cmd/follow/code"
	"IkunHouse/cmd/follow/model"
	"IkunHouse/cmd/follow/rpc/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type UnFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnFollowLogic {
	return &UnFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 取消关注
func (l *UnFollowLogic) UnFollow(in *follow.UnFollowReq) (*follow.UnFollowRes, error) {
	// 1. 校验
	if len(in.UserId) == 0 {
		return nil, code.FollowUserIdEmpty
	}
	if len(in.FollowedUserId) == 0 {
		return nil, code.FollowedUserIdEmpty
	}

	followData, err := l.svcCtx.FollowModel.FindByUserIDAndFollowedUserID(l.ctx, in.UserId, in.FollowedUserId)
	if err != nil {
		l.Logger.Errorf("[UnFollow] FollowModel.FindByUserIDAndFollowedUserID err: %v req: %v", err, in)
		return nil, err
	}
	if followData == nil {
		return nil, nil
	}
	if followData.FollowStatus == types.FollowStatusUnfollow {
		return nil, nil
	}

	// 2. 取消关系
	// 2.1 事务更新数据库
	err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		err := model.NewFollowModel(tx).UpdateFields(l.ctx, followData.ID, map[string]interface{}{
			"follow_status": types.FollowStatusUnfollow,
		})
		if err != nil {
			return err
		}
		err = model.NewFollowCountModel(tx).DecrFollowCount(l.ctx, in.UserId)
		if err != nil {
			return err
		}
		return model.NewFollowCountModel(tx).DecrFansCount(l.ctx, in.FollowedUserId)
	})
	if err != nil {
		l.Logger.Errorf("[UnFollow] Transaction error: %v", err)
		return nil, err
	}

	// 2.2 更新redis
	// todo: 保持一致性
	_, err = l.svcCtx.BizRedis.ZremCtx(l.ctx, code.UserFollowKey(in.UserId), in.FollowedUserId)
	if err != nil {
		l.Logger.Errorf("[UnFollow] BizRedis.ZremCtx error: %v", err)
		return nil, err
	}
	_, err = l.svcCtx.BizRedis.ZremCtx(l.ctx, code.UserFansKey(in.FollowedUserId), in.UserId)
	if err != nil {
		l.Logger.Errorf("[UnFollow] BizRedis.ZremCtx error: %v", err)
		return nil, err
	}

	return &follow.UnFollowRes{}, nil
}
