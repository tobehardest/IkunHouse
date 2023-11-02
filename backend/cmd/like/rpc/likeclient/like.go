// Code generated by goctl. DO NOT EDIT.
// Source: like.proto

package likeclient

import (
	"context"

	"video_clip/cmd/like/rpc/like"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	IsLikeReq   = like.IsLikeReq
	IsLikeRes   = like.IsLikeRes
	LikeReq     = like.LikeReq
	LikeRes     = like.LikeRes
	UserThumbup = like.UserThumbup

	Like interface {
		Like(ctx context.Context, in *LikeReq, opts ...grpc.CallOption) (*LikeRes, error)
		IsLike(ctx context.Context, in *IsLikeReq, opts ...grpc.CallOption) (*IsLikeRes, error)
	}

	defaultLike struct {
		cli zrpc.Client
	}
)

func NewLike(cli zrpc.Client) Like {
	return &defaultLike{
		cli: cli,
	}
}

func (m *defaultLike) Like(ctx context.Context, in *LikeReq, opts ...grpc.CallOption) (*LikeRes, error) {
	client := like.NewLikeClient(m.cli.Conn())
	return client.Like(ctx, in, opts...)
}

func (m *defaultLike) IsLike(ctx context.Context, in *IsLikeReq, opts ...grpc.CallOption) (*IsLikeRes, error) {
	client := like.NewLikeClient(m.cli.Conn())
	return client.IsLike(ctx, in, opts...)
}
