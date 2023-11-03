// Code generated by goctl. DO NOT EDIT.
// Source: post.proto

package server

import (
	"context"

	"video_clip/cmd/post/rpc/internal/logic"
	"video_clip/cmd/post/rpc/internal/svc"
	"video_clip/cmd/post/rpc/post"
)

type PostServer struct {
	svcCtx *svc.ServiceContext
	post.UnimplementedPostServer
}

func NewPostServer(svcCtx *svc.ServiceContext) *PostServer {
	return &PostServer{
		svcCtx: svcCtx,
	}
}

// 上传视频
func (s *PostServer) UploadVideo(ctx context.Context, in *post.UploadVideoRequest) (*post.UploadVideoResponse, error) {
	l := logic.NewUploadVideoLogic(ctx, s.svcCtx)
	return l.UploadVideo(in)
}

// 获得视频 - 默认/分类
func (s *PostServer) GetVideoList(ctx context.Context, in *post.GetVideoListRequest) (*post.GetVideoListResponse, error) {
	l := logic.NewGetVideoListLogic(ctx, s.svcCtx)
	return l.GetVideoList(in)
}

// 获得热门视频
func (s *PostServer) GetHotVideoList(ctx context.Context, in *post.GetHotVideoRequest) (*post.GetHotVideoResponse, error) {
	l := logic.NewGetHotVideoListLogic(ctx, s.svcCtx)
	return l.GetHotVideoList(in)
}

// 根据用户id查找视频
func (s *PostServer) GetVideoByUid(ctx context.Context, in *post.GetVideoByUidRequest) (*post.GetVideoByUidResponse, error) {
	l := logic.NewGetVideoByUidLogic(ctx, s.svcCtx)
	return l.GetVideoByUid(in)
}

// 根据搜索查询视频
func (s *PostServer) SearchVideo(ctx context.Context, in *post.SearchVideoRequest) (*post.SearchVideoResponse, error) {
	l := logic.NewSearchVideoLogic(ctx, s.svcCtx)
	return l.SearchVideo(in)
}

// 获得类别
func (s *PostServer) GetTagList(ctx context.Context, in *post.GetTagListRequest) (*post.GetTagListResponse, error) {
	l := logic.NewGetTagListLogic(ctx, s.svcCtx)
	return l.GetTagList(in)
}

// 评论数操作
func (s *PostServer) SetCommentNum(ctx context.Context, in *post.CommentNumRequest) (*post.CommentNumResponse, error) {
	l := logic.NewSetCommentNumLogic(ctx, s.svcCtx)
	return l.SetCommentNum(in)
}

// 点赞数操作
func (s *PostServer) SetLikeNum(ctx context.Context, in *post.LikeNumRequest) (*post.LikeNumResponse, error) {
	l := logic.NewSetLikeNumLogic(ctx, s.svcCtx)
	return l.SetLikeNum(in)
}

// 收藏数操作
func (s *PostServer) CollectNum(ctx context.Context, in *post.CollectNumRequest) (*post.CollectNumResponse, error) {
	l := logic.NewCollectNumLogic(ctx, s.svcCtx)
	return l.CollectNum(in)
}
