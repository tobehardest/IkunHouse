// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: follow.proto

package follow

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Follow_Follow_FullMethodName       = "/follow.Follow/Follow"
	Follow_UnFollow_FullMethodName     = "/follow.Follow/UnFollow"
	Follow_FolloweeList_FullMethodName = "/follow.Follow/FolloweeList"
	Follow_FollowerList_FullMethodName = "/follow.Follow/FollowerList"
	Follow_HasFollowed_FullMethodName  = "/follow.Follow/HasFollowed"
)

// FollowClient is the client API for Follow service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FollowClient interface {
	// 关注
	Follow(ctx context.Context, in *FollowReq, opts ...grpc.CallOption) (*FollowRes, error)
	// 取消关注
	UnFollow(ctx context.Context, in *UnFollowReq, opts ...grpc.CallOption) (*UnFollowRes, error)
	// 关注列表
	FolloweeList(ctx context.Context, in *FolloweeListReq, opts ...grpc.CallOption) (*FolloweeListRes, error)
	// 粉丝列表
	FollowerList(ctx context.Context, in *FollowerListReq, opts ...grpc.CallOption) (*FollowerListRes, error)
	// 是否已关注
	HasFollowed(ctx context.Context, in *HasFollowedReq, opts ...grpc.CallOption) (*HasFollowedRes, error)
}

type followClient struct {
	cc grpc.ClientConnInterface
}

func NewFollowClient(cc grpc.ClientConnInterface) FollowClient {
	return &followClient{cc}
}

func (c *followClient) Follow(ctx context.Context, in *FollowReq, opts ...grpc.CallOption) (*FollowRes, error) {
	out := new(FollowRes)
	err := c.cc.Invoke(ctx, Follow_Follow_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followClient) UnFollow(ctx context.Context, in *UnFollowReq, opts ...grpc.CallOption) (*UnFollowRes, error) {
	out := new(UnFollowRes)
	err := c.cc.Invoke(ctx, Follow_UnFollow_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followClient) FolloweeList(ctx context.Context, in *FolloweeListReq, opts ...grpc.CallOption) (*FolloweeListRes, error) {
	out := new(FolloweeListRes)
	err := c.cc.Invoke(ctx, Follow_FolloweeList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followClient) FollowerList(ctx context.Context, in *FollowerListReq, opts ...grpc.CallOption) (*FollowerListRes, error) {
	out := new(FollowerListRes)
	err := c.cc.Invoke(ctx, Follow_FollowerList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followClient) HasFollowed(ctx context.Context, in *HasFollowedReq, opts ...grpc.CallOption) (*HasFollowedRes, error) {
	out := new(HasFollowedRes)
	err := c.cc.Invoke(ctx, Follow_HasFollowed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FollowServer is the server API for Follow service.
// All implementations must embed UnimplementedFollowServer
// for forward compatibility
type FollowServer interface {
	// 关注
	Follow(context.Context, *FollowReq) (*FollowRes, error)
	// 取消关注
	UnFollow(context.Context, *UnFollowReq) (*UnFollowRes, error)
	// 关注列表
	FolloweeList(context.Context, *FolloweeListReq) (*FolloweeListRes, error)
	// 粉丝列表
	FollowerList(context.Context, *FollowerListReq) (*FollowerListRes, error)
	// 是否已关注
	HasFollowed(context.Context, *HasFollowedReq) (*HasFollowedRes, error)
	mustEmbedUnimplementedFollowServer()
}

// UnimplementedFollowServer must be embedded to have forward compatible implementations.
type UnimplementedFollowServer struct {
}

func (UnimplementedFollowServer) Follow(context.Context, *FollowReq) (*FollowRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Follow not implemented")
}
func (UnimplementedFollowServer) UnFollow(context.Context, *UnFollowReq) (*UnFollowRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnFollow not implemented")
}
func (UnimplementedFollowServer) FolloweeList(context.Context, *FolloweeListReq) (*FolloweeListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FolloweeList not implemented")
}
func (UnimplementedFollowServer) FollowerList(context.Context, *FollowerListReq) (*FollowerListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowerList not implemented")
}
func (UnimplementedFollowServer) HasFollowed(context.Context, *HasFollowedReq) (*HasFollowedRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasFollowed not implemented")
}
func (UnimplementedFollowServer) mustEmbedUnimplementedFollowServer() {}

// UnsafeFollowServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FollowServer will
// result in compilation errors.
type UnsafeFollowServer interface {
	mustEmbedUnimplementedFollowServer()
}

func RegisterFollowServer(s grpc.ServiceRegistrar, srv FollowServer) {
	s.RegisterService(&Follow_ServiceDesc, srv)
}

func _Follow_Follow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServer).Follow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Follow_Follow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServer).Follow(ctx, req.(*FollowReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Follow_UnFollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnFollowReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServer).UnFollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Follow_UnFollow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServer).UnFollow(ctx, req.(*UnFollowReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Follow_FolloweeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FolloweeListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServer).FolloweeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Follow_FolloweeList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServer).FolloweeList(ctx, req.(*FolloweeListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Follow_FollowerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowerListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServer).FollowerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Follow_FollowerList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServer).FollowerList(ctx, req.(*FollowerListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Follow_HasFollowed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HasFollowedReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServer).HasFollowed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Follow_HasFollowed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServer).HasFollowed(ctx, req.(*HasFollowedReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Follow_ServiceDesc is the grpc.ServiceDesc for Follow service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Follow_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "follow.Follow",
	HandlerType: (*FollowServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Follow",
			Handler:    _Follow_Follow_Handler,
		},
		{
			MethodName: "UnFollow",
			Handler:    _Follow_UnFollow_Handler,
		},
		{
			MethodName: "FolloweeList",
			Handler:    _Follow_FolloweeList_Handler,
		},
		{
			MethodName: "FollowerList",
			Handler:    _Follow_FollowerList_Handler,
		},
		{
			MethodName: "HasFollowed",
			Handler:    _Follow_HasFollowed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "follow.proto",
}
