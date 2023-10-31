// Code generated by goctl. DO NOT EDIT.
// Source: msg.proto

package server

import (
	"context"

	"video_clip/cmd/msg/internal/logic"
	"video_clip/cmd/msg/internal/svc"
	"video_clip/cmd/msg/msg"
)

type MsgServer struct {
	svcCtx *svc.ServiceContext
	msg.UnimplementedMsgServer
}

func NewMsgServer(svcCtx *svc.ServiceContext) *MsgServer {
	return &MsgServer{
		svcCtx: svcCtx,
	}
}

func (s *MsgServer) Ping(ctx context.Context, in *msg.Request) (*msg.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}