// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.2
// source: cmd/video/rpc/video.proto

package video

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
	Video_UploadVideo_FullMethodName  = "/video.Video/UploadVideo"
	Video_AddReadCount_FullMethodName = "/video.Video/AddReadCount"
	Video_GetVideoList_FullMethodName = "/video.Video/GetVideoList"
	Video_GetTypeList_FullMethodName  = "/video.Video/GetTypeList"
	Video_IfExistVideo_FullMethodName = "/video.Video/IfExistVideo"
)

// VideoClient is the client API for Video service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoClient interface {
	UploadVideo(ctx context.Context, in *UploadVideoRequest, opts ...grpc.CallOption) (*UploadVideoResponse, error)
	AddReadCount(ctx context.Context, in *AddReadCountRequest, opts ...grpc.CallOption) (*AddReadCountResponse, error)
	GetVideoList(ctx context.Context, in *GetVideoRequest, opts ...grpc.CallOption) (*GetVideoResponse, error)
	// rpc AddCollect(AddCollectRequest) returns(AddCollectResponse);
	// rpc GetCollectByVid(GetCollectByVidRequest) returns(GetCollectByVidResponse);
	// rpc GetCollectByUid(GetCollectByUidRequest) returns(GetCollectByUidResponse);
	GetTypeList(ctx context.Context, in *GetTypeListRequest, opts ...grpc.CallOption) (*GetTypeListResponse, error)
	IfExistVideo(ctx context.Context, in *ExistVideoRequest, opts ...grpc.CallOption) (*ExistVideoResponse, error)
}

type videoClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoClient(cc grpc.ClientConnInterface) VideoClient {
	return &videoClient{cc}
}

func (c *videoClient) UploadVideo(ctx context.Context, in *UploadVideoRequest, opts ...grpc.CallOption) (*UploadVideoResponse, error) {
	out := new(UploadVideoResponse)
	err := c.cc.Invoke(ctx, Video_UploadVideo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) AddReadCount(ctx context.Context, in *AddReadCountRequest, opts ...grpc.CallOption) (*AddReadCountResponse, error) {
	out := new(AddReadCountResponse)
	err := c.cc.Invoke(ctx, Video_AddReadCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) GetVideoList(ctx context.Context, in *GetVideoRequest, opts ...grpc.CallOption) (*GetVideoResponse, error) {
	out := new(GetVideoResponse)
	err := c.cc.Invoke(ctx, Video_GetVideoList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) GetTypeList(ctx context.Context, in *GetTypeListRequest, opts ...grpc.CallOption) (*GetTypeListResponse, error) {
	out := new(GetTypeListResponse)
	err := c.cc.Invoke(ctx, Video_GetTypeList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) IfExistVideo(ctx context.Context, in *ExistVideoRequest, opts ...grpc.CallOption) (*ExistVideoResponse, error) {
	out := new(ExistVideoResponse)
	err := c.cc.Invoke(ctx, Video_IfExistVideo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoServer is the server API for Video service.
// All implementations must embed UnimplementedVideoServer
// for forward compatibility
type VideoServer interface {
	UploadVideo(context.Context, *UploadVideoRequest) (*UploadVideoResponse, error)
	AddReadCount(context.Context, *AddReadCountRequest) (*AddReadCountResponse, error)
	GetVideoList(context.Context, *GetVideoRequest) (*GetVideoResponse, error)
	// rpc AddCollect(AddCollectRequest) returns(AddCollectResponse);
	// rpc GetCollectByVid(GetCollectByVidRequest) returns(GetCollectByVidResponse);
	// rpc GetCollectByUid(GetCollectByUidRequest) returns(GetCollectByUidResponse);
	GetTypeList(context.Context, *GetTypeListRequest) (*GetTypeListResponse, error)
	IfExistVideo(context.Context, *ExistVideoRequest) (*ExistVideoResponse, error)
	mustEmbedUnimplementedVideoServer()
}

// UnimplementedVideoServer must be embedded to have forward compatible implementations.
type UnimplementedVideoServer struct {
}

func (UnimplementedVideoServer) UploadVideo(context.Context, *UploadVideoRequest) (*UploadVideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadVideo not implemented")
}
func (UnimplementedVideoServer) AddReadCount(context.Context, *AddReadCountRequest) (*AddReadCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddReadCount not implemented")
}
func (UnimplementedVideoServer) GetVideoList(context.Context, *GetVideoRequest) (*GetVideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideoList not implemented")
}
func (UnimplementedVideoServer) GetTypeList(context.Context, *GetTypeListRequest) (*GetTypeListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTypeList not implemented")
}
func (UnimplementedVideoServer) IfExistVideo(context.Context, *ExistVideoRequest) (*ExistVideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IfExistVideo not implemented")
}
func (UnimplementedVideoServer) mustEmbedUnimplementedVideoServer() {}

// UnsafeVideoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoServer will
// result in compilation errors.
type UnsafeVideoServer interface {
	mustEmbedUnimplementedVideoServer()
}

func RegisterVideoServer(s grpc.ServiceRegistrar, srv VideoServer) {
	s.RegisterService(&Video_ServiceDesc, srv)
}

func _Video_UploadVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).UploadVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Video_UploadVideo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).UploadVideo(ctx, req.(*UploadVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_AddReadCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddReadCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).AddReadCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Video_AddReadCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).AddReadCount(ctx, req.(*AddReadCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_GetVideoList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).GetVideoList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Video_GetVideoList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).GetVideoList(ctx, req.(*GetVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_GetTypeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTypeListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).GetTypeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Video_GetTypeList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).GetTypeList(ctx, req.(*GetTypeListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_IfExistVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).IfExistVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Video_IfExistVideo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).IfExistVideo(ctx, req.(*ExistVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Video_ServiceDesc is the grpc.ServiceDesc for Video service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Video_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "video.Video",
	HandlerType: (*VideoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadVideo",
			Handler:    _Video_UploadVideo_Handler,
		},
		{
			MethodName: "AddReadCount",
			Handler:    _Video_AddReadCount_Handler,
		},
		{
			MethodName: "GetVideoList",
			Handler:    _Video_GetVideoList_Handler,
		},
		{
			MethodName: "GetTypeList",
			Handler:    _Video_GetTypeList_Handler,
		},
		{
			MethodName: "IfExistVideo",
			Handler:    _Video_IfExistVideo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cmd/video/rpc/video.proto",
}
