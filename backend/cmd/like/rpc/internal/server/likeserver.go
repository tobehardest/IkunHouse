// Code generated by goctl. DO NOT EDIT.
// Source: like.proto

package server

import (
	"context"

	"IkunHouse/cmd/like/rpc/internal/logic"
	"IkunHouse/cmd/like/rpc/internal/svc"
	"IkunHouse/cmd/like/rpc/like"
)

type LikeServer struct {
	svcCtx *svc.ServiceContext
	like.UnimplementedLikeServer
}

func NewLikeServer(svcCtx *svc.ServiceContext) *LikeServer {
	return &LikeServer{
		svcCtx: svcCtx,
	}
}

func (s *LikeServer) Like(ctx context.Context, in *like.LikeReq) (*like.LikeRes, error) {
	l := logic.NewLikeLogic(ctx, s.svcCtx)
	return l.Like(in)
}

func (s *LikeServer) IsLike(ctx context.Context, in *like.IsLikeReq) (*like.IsLikeRes, error) {
	l := logic.NewIsLikeLogic(ctx, s.svcCtx)
	return l.IsLike(in)
}

func (s *LikeServer) LikeCount(ctx context.Context, in *like.LikeCountReq) (*like.LikeCountRes, error) {
	l := logic.NewLikeCountLogic(ctx, s.svcCtx)
	return l.LikeCount(in)
}
