// Code generated by goctl. DO NOT EDIT.
// Source: msg.proto

package msgclient

import (
	msg2 "IkunHouse/cmd/msg/rpc/msg"
	"context"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = msg2.Request
	Response = msg2.Response

	Msg interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultMsg struct {
		cli zrpc.Client
	}
)

func NewMsg(cli zrpc.Client) Msg {
	return &defaultMsg{
		cli: cli,
	}
}

func (m *defaultMsg) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := msg2.NewMsgClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}
