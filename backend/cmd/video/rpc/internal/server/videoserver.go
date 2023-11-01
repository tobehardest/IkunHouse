// Code generated by goctl. DO NOT EDIT.
// Source: video.proto

package server

import (
	"context"

	"video_clip/cmd/video/internal/logic"
	"video_clip/cmd/video/internal/svc"
	"video_clip/cmd/video/video"
)

type VideoServer struct {
	svcCtx *svc.ServiceContext
	video.UnimplementedVideoServer
}

func NewVideoServer(svcCtx *svc.ServiceContext) *VideoServer {
	return &VideoServer{
		svcCtx: svcCtx,
	}
}

func (s *VideoServer) UploadVideo(ctx context.Context, in *video.UploadVideoReq) (*video.UploadVideoRes, error) {
	l := logic.NewUploadVideoLogic(ctx, s.svcCtx)
	return l.UploadVideo(in)
}