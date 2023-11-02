// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: collect.proto

package collect

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
	Collect_Collect_FullMethodName     = "/collect.Collect/Collect"
	Collect_CollectList_FullMethodName = "/collect.Collect/CollectList"
)

// CollectClient is the client API for Collect service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CollectClient interface {
	// 收藏
	Collect(ctx context.Context, in *CollectReq, opts ...grpc.CallOption) (*CollectRes, error)
	// 收藏列表
	CollectList(ctx context.Context, in *CollectListReq, opts ...grpc.CallOption) (*CollectListRes, error)
}

type collectClient struct {
	cc grpc.ClientConnInterface
}

func NewCollectClient(cc grpc.ClientConnInterface) CollectClient {
	return &collectClient{cc}
}

func (c *collectClient) Collect(ctx context.Context, in *CollectReq, opts ...grpc.CallOption) (*CollectRes, error) {
	out := new(CollectRes)
	err := c.cc.Invoke(ctx, Collect_Collect_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectClient) CollectList(ctx context.Context, in *CollectListReq, opts ...grpc.CallOption) (*CollectListRes, error) {
	out := new(CollectListRes)
	err := c.cc.Invoke(ctx, Collect_CollectList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CollectServer is the server API for Collect service.
// All implementations must embed UnimplementedCollectServer
// for forward compatibility
type CollectServer interface {
	// 收藏
	Collect(context.Context, *CollectReq) (*CollectRes, error)
	// 收藏列表
	CollectList(context.Context, *CollectListReq) (*CollectListRes, error)
	mustEmbedUnimplementedCollectServer()
}

// UnimplementedCollectServer must be embedded to have forward compatible implementations.
type UnimplementedCollectServer struct {
}

func (UnimplementedCollectServer) Collect(context.Context, *CollectReq) (*CollectRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Collect not implemented")
}
func (UnimplementedCollectServer) CollectList(context.Context, *CollectListReq) (*CollectListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollectList not implemented")
}
func (UnimplementedCollectServer) mustEmbedUnimplementedCollectServer() {}

// UnsafeCollectServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CollectServer will
// result in compilation errors.
type UnsafeCollectServer interface {
	mustEmbedUnimplementedCollectServer()
}

func RegisterCollectServer(s grpc.ServiceRegistrar, srv CollectServer) {
	s.RegisterService(&Collect_ServiceDesc, srv)
}

func _Collect_Collect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectServer).Collect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Collect_Collect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectServer).Collect(ctx, req.(*CollectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Collect_CollectList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectServer).CollectList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Collect_CollectList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectServer).CollectList(ctx, req.(*CollectListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Collect_ServiceDesc is the grpc.ServiceDesc for Collect service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Collect_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "collect.Collect",
	HandlerType: (*CollectServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Collect",
			Handler:    _Collect_Collect_Handler,
		},
		{
			MethodName: "CollectList",
			Handler:    _Collect_CollectList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "collect.proto",
}