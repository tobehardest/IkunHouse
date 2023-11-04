package like

import (
	"context"
	"github.com/jinzhu/copier"

	"video_clip/cmd/api/internal/svc"
	"video_clip/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"video_clip/cmd/like/rpc/like"
)

type LikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeLogic {
	return &LikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LikeLogic) Like(req *types.LikeReq) (resp *types.LikeRes, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.LikeClient.Like(l.ctx, &like.LikeReq{
		BizId:    req.BizId,
		ObjId:    req.TargetId,
		UserId:   req.UserId,
		LikeType: req.LikeType,
	})

	err = copier.Copy(resp, res)
	if err != nil {

	}
	return resp, nil
}
