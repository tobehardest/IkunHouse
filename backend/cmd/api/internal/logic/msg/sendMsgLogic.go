package msg

import (
	"context"

	"IkunHouse/cmd/api/internal/svc"
	"IkunHouse/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMsgLogic {
	return &SendMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendMsgLogic) SendMsg(req *types.SendMsgReq) (resp *types.SendMsgRes, err error) {
	// todo: add your logic here and delete this line

	return
}
