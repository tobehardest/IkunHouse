// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userClient

import (
	"context"

	"IkunHouse/cmd/user/rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetUserInfoDetailReq = user.GetUserInfoDetailReq
	GetUserInfoDetailRes = user.GetUserInfoDetailRes
	UpdateUserInfoReq    = user.UpdateUserInfoReq
	UpdateUserInfoRes    = user.UpdateUserInfoRes
	UserInfo             = user.UserInfo

	User interface {
		// 更新用户信息
		UpdateUserInfo(ctx context.Context, in *UpdateUserInfoReq, opts ...grpc.CallOption) (*UpdateUserInfoRes, error)
		// 获取用户主页信息
		GetUserInfoDetail(ctx context.Context, in *GetUserInfoDetailReq, opts ...grpc.CallOption) (*GetUserInfoDetailRes, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

// 更新用户信息
func (m *defaultUser) UpdateUserInfo(ctx context.Context, in *UpdateUserInfoReq, opts ...grpc.CallOption) (*UpdateUserInfoRes, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UpdateUserInfo(ctx, in, opts...)
}

// 获取用户主页信息
func (m *defaultUser) GetUserInfoDetail(ctx context.Context, in *GetUserInfoDetailReq, opts ...grpc.CallOption) (*GetUserInfoDetailRes, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUserInfoDetail(ctx, in, opts...)
}
