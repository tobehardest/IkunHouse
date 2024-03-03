package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"time"
	"video_clip/cmd/follow/code"
	"video_clip/cmd/follow/model"
	"video_clip/cmd/follow/rpc/follow"
	"video_clip/cmd/follow/rpc/internal/svc"
	"video_clip/cmd/follow/rpc/internal/types"
)

type FollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 关注
func (l *FollowLogic) Follow(in *follow.FollowReq) (*follow.FollowRes, error) {
	// 1. 校验
	if len(in.UserId) == 0 {
		return nil, code.FollowUserIdEmpty
	}
	if len(in.FollowedUserId) == 0 {
		return nil, code.FollowedUserIdEmpty
	}
	if in.UserId == in.FollowedUserId {
		return nil, code.CannotFollowSelf
	}
	followData, err := l.svcCtx.FollowModel.FindByUserIDAndFollowedUserID(l.ctx, in.UserId, in.FollowedUserId)
	if err != nil {
		l.Logger.Errorf("[Follow] FollowModel.FindByUserIDAndFollowedUserID err: %v req: %v", err, in)
		return nil, err
	}
	if followData != nil && followData.FollowStatus == types.FollowStatusFollow {
		return &follow.FollowRes{}, nil
	}

	// 2. 更新关注状态和关注数
	// 2.1 更新数据库
	err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		if followData != nil {
			err = model.NewFollowModel(tx).UpdateFields(l.ctx, followData.ID, map[string]interface{}{
				"follow_status": types.FollowStatusFollow,
			})
		} else {
			err = model.NewFollowModel(tx).Insert(l.ctx, &model.Follow{
				UserID:         in.UserId,
				FollowedUserID: in.FollowedUserId,
				FollowStatus:   types.FollowStatusFollow,
				CreateTime:     time.Now(),
				UpdateTime:     time.Now(),
			})
		}

		if err != nil {
			return err
		}
		err = model.NewFollowCountModel(tx).IncrFollowCount(l.ctx, in.UserId)
		if err != nil {
			return err
		}
		return model.NewFollowCountModel(tx).IncrFansCount(l.ctx, in.FollowedUserId)
	})

	if err != nil {
		l.Logger.Errorf("[Follow] Transaction error: %v", err)
		return nil, err
	}

	// 2.2 更新redis
	followExist, err := l.svcCtx.BizRedis.ExistsCtx(l.ctx, code.UserFollowKey(in.UserId))
	if err != nil {
		l.Logger.Errorf("[Follow] Redis Exists error: %v", err)
		return nil, err
	}
	if followExist {
		_, err = l.svcCtx.BizRedis.ZaddCtx(l.ctx, code.UserFollowKey(in.UserId), time.Now().Unix(), in.FollowedUserId)
		if err != nil {
			l.Logger.Errorf("[Follow] Redis Zadd error: %v", err)
			return nil, err
		}
		_, err = l.svcCtx.BizRedis.ZremrangebyrankCtx(l.ctx, code.UserFollowKey(in.UserId), 0, -(types.CacheMaxFollowCount + 1))
		if err != nil {
			l.Logger.Errorf("[Follow] Redis Zremrangebyrank error: %v", err)
		}
	}
	fansExist, err := l.svcCtx.BizRedis.ExistsCtx(l.ctx, code.UserFansKey(in.FollowedUserId))
	if err != nil {
		l.Logger.Errorf("[Follow] Redis Exists error: %v", err)
		return nil, err
	}
	if fansExist {
		_, err = l.svcCtx.BizRedis.ZaddCtx(l.ctx, code.UserFansKey(in.FollowedUserId), time.Now().Unix(), in.UserId)
		if err != nil {
			l.Logger.Errorf("[Follow] Redis Zadd error: %v", err)
			return nil, err
		}
		_, err = l.svcCtx.BizRedis.ZremrangebyrankCtx(l.ctx, code.UserFansKey(in.FollowedUserId), 0, -(types.CacheMaxFansCount + 1))
		if err != nil {
			l.Logger.Errorf("[Follow] Redis Zremrangebyrank error: %v", err)
		}
	}

	return &follow.FollowRes{}, nil
}
