// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.2
// source: cmd/video/rpc/video.proto

package video

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Symbols defined in public import of google/protobuf/timestamp.proto.

type Timestamp = timestamppb.Timestamp

// 上传视频方法
type UploadVideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Media       string  `protobuf:"bytes,1,opt,name=media,proto3" json:"media,omitempty"` // 视频路径
	Uid         string  `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Title       string  `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	CoverUrl    string  `protobuf:"bytes,4,opt,name=cover_url,json=coverUrl,proto3" json:"cover_url,omitempty"` // 封面路径
	Longitude   float64 `protobuf:"fixed64,5,opt,name=longitude,proto3" json:"longitude,omitempty"`             // 经度
	Latitude    float64 `protobuf:"fixed64,6,opt,name=latitude,proto3" json:"latitude,omitempty"`               // 纬度
	Address     string  `protobuf:"bytes,7,opt,name=Address,proto3" json:"Address,omitempty"`
	TagId       []int64 `protobuf:"varint,8,rep,packed,name=tag_id,json=tagId,proto3" json:"tag_id,omitempty"` // 标签id
	VideoSha256 string  `protobuf:"bytes,9,opt,name=VideoSha256,proto3" json:"VideoSha256,omitempty"`          // 哈希值，判断视频是否重复
}

func (x *UploadVideoRequest) Reset() {
	*x = UploadVideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_video_rpc_video_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadVideoRequest) ProtoMessage() {}

func (x *UploadVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_video_rpc_video_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadVideoRequest.ProtoReflect.Descriptor instead.
func (*UploadVideoRequest) Descriptor() ([]byte, []int) {
	return file_cmd_video_rpc_video_proto_rawDescGZIP(), []int{0}
}

func (x *UploadVideoRequest) GetMedia() string {
	if x != nil {
		return x.Media
	}
	return ""
}

func (x *UploadVideoRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *UploadVideoRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UploadVideoRequest) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *UploadVideoRequest) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *UploadVideoRequest) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *UploadVideoRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UploadVideoRequest) GetTagId() []int64 {
	if x != nil {
		return x.TagId
	}
	return nil
}

func (x *UploadVideoRequest) GetVideoSha256() string {
	if x != nil {
		return x.VideoSha256
	}
	return ""
}

type UploadVideoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UploadVideoResponse) Reset() {
	*x = UploadVideoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_video_rpc_video_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadVideoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadVideoResponse) ProtoMessage() {}

func (x *UploadVideoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_video_rpc_video_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadVideoResponse.ProtoReflect.Descriptor instead.
func (*UploadVideoResponse) Descriptor() ([]byte, []int) {
	return file_cmd_video_rpc_video_proto_rawDescGZIP(), []int{1}
}

func (x *UploadVideoResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 获得视频方法 -
type GetVideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid     string  `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`                                // 可扩展，获得好友相关视频以及其他
	Tag     string  `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`                                // 根据标签查找
	Count   int64   `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`                           // 获得个数
	VideoId []int64 `protobuf:"varint,4,rep,packed,name=video_id,json=videoId,proto3" json:"video_id,omitempty"` // 已获得的video_id，避免推送重复的
}

func (x *GetVideoRequest) Reset() {
	*x = GetVideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_video_rpc_video_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVideoRequest) ProtoMessage() {}

func (x *GetVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_video_rpc_video_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVideoRequest.ProtoReflect.Descriptor instead.
func (*GetVideoRequest) Descriptor() ([]byte, []int) {
	return file_cmd_video_rpc_video_proto_rawDescGZIP(), []int{2}
}

func (x *GetVideoRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *GetVideoRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *GetVideoRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetVideoRequest) GetVideoId() []int64 {
	if x != nil {
		return x.VideoId
	}
	return nil
}

type GetVideoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoList []*VideoModel `protobuf:"bytes,1,rep,name=video_list,json=videoList,proto3" json:"video_list,omitempty"`
}

func (x *GetVideoResponse) Reset() {
	*x = GetVideoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_video_rpc_video_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVideoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVideoResponse) ProtoMessage() {}

func (x *GetVideoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_video_rpc_video_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVideoResponse.ProtoReflect.Descriptor instead.
func (*GetVideoResponse) Descriptor() ([]byte, []int) {
	return file_cmd_video_rpc_video_proto_rawDescGZIP(), []int{3}
}

func (x *GetVideoResponse) GetVideoList() []*VideoModel {
	if x != nil {
		return x.VideoList
	}
	return nil
}

type VideoModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid        string                 `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Title      string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Media      string                 `protobuf:"bytes,4,opt,name=media,proto3" json:"media,omitempty"`
	CoverUrl   string                 `protobuf:"bytes,5,opt,name=cover_url,json=coverUrl,proto3" json:"cover_url,omitempty"`
	ReadCount  string                 `protobuf:"bytes,6,opt,name=read_count,json=readCount,proto3" json:"read_count,omitempty"`
	Address    string                 `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	Longitude  float64                `protobuf:"fixed64,8,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude   float64                `protobuf:"fixed64,9,opt,name=latitude,proto3" json:"latitude,omitempty"`
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
}

func (x *VideoModel) Reset() {
	*x = VideoModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_video_rpc_video_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoModel) ProtoMessage() {}

func (x *VideoModel) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_video_rpc_video_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoModel.ProtoReflect.Descriptor instead.
func (*VideoModel) Descriptor() ([]byte, []int) {
	return file_cmd_video_rpc_video_proto_rawDescGZIP(), []int{4}
}

func (x *VideoModel) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *VideoModel) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *VideoModel) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *VideoModel) GetMedia() string {
	if x != nil {
		return x.Media
	}
	return ""
}

func (x *VideoModel) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *VideoModel) GetReadCount() string {
	if x != nil {
		return x.ReadCount
	}
	return ""
}

func (x *VideoModel) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *VideoModel) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *VideoModel) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *VideoModel) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

// 添加访问量
type AddReadCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid   string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Count int64  `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *AddReadCountRequest) Reset() {
	*x = AddReadCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_video_rpc_video_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddReadCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReadCountRequest) ProtoMessage() {}

func (x *AddReadCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_video_rpc_video_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReadCountRequest.ProtoReflect.Descriptor instead.
func (*AddReadCountRequest) Descriptor() ([]byte, []int) {
	return file_cmd_video_rpc_video_proto_rawDescGZIP(), []int{5}
}

func (x *AddReadCountRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AddReadCountRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *AddReadCountRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type AddReadCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddReadCountResponse) Reset() {
	*x = AddReadCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_video_rpc_video_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddReadCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReadCountResponse) ProtoMessage() {}

func (x *AddReadCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_video_rpc_video_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReadCountResponse.ProtoReflect.Descriptor instead.
func (*AddReadCountResponse) Descriptor() ([]byte, []int) {
	return file_cmd_video_rpc_video_proto_rawDescGZIP(), []int{6}
}

// 获得所有类型
type GetTypeListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTypeListRequest) Reset() {
	*x = GetTypeListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_video_rpc_video_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTypeListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTypeListRequest) ProtoMessage() {}

func (x *GetTypeListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_video_rpc_video_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTypeListRequest.ProtoReflect.Descriptor instead.
func (*GetTypeListRequest) Descriptor() ([]byte, []int) {
	return file_cmd_video_rpc_video_proto_rawDescGZIP(), []int{7}
}

type GetTypeListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeList []*Type `protobuf:"bytes,1,rep,name=type_list,json=typeList,proto3" json:"type_list,omitempty"`
}

func (x *GetTypeListResponse) Reset() {
	*x = GetTypeListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_video_rpc_video_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTypeListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTypeListResponse) ProtoMessage() {}

func (x *GetTypeListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_video_rpc_video_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTypeListResponse.ProtoReflect.Descriptor instead.
func (*GetTypeListResponse) Descriptor() ([]byte, []int) {
	return file_cmd_video_rpc_video_proto_rawDescGZIP(), []int{8}
}

func (x *GetTypeListResponse) GetTypeList() []*Type {
	if x != nil {
		return x.TypeList
	}
	return nil
}

type Type struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Type) Reset() {
	*x = Type{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_video_rpc_video_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Type) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Type) ProtoMessage() {}

func (x *Type) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_video_rpc_video_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Type.ProtoReflect.Descriptor instead.
func (*Type) Descriptor() ([]byte, []int) {
	return file_cmd_video_rpc_video_proto_rawDescGZIP(), []int{9}
}

func (x *Type) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Type) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ExistVideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoHash256 string `protobuf:"bytes,1,opt,name=videoHash256,proto3" json:"videoHash256,omitempty"`
}

func (x *ExistVideoRequest) Reset() {
	*x = ExistVideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_video_rpc_video_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExistVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExistVideoRequest) ProtoMessage() {}

func (x *ExistVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_video_rpc_video_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExistVideoRequest.ProtoReflect.Descriptor instead.
func (*ExistVideoRequest) Descriptor() ([]byte, []int) {
	return file_cmd_video_rpc_video_proto_rawDescGZIP(), []int{10}
}

func (x *ExistVideoRequest) GetVideoHash256() string {
	if x != nil {
		return x.VideoHash256
	}
	return ""
}

type ExistVideoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exist bool   `protobuf:"varint,1,opt,name=exist,proto3" json:"exist,omitempty"` // 是否存在
	Media string `protobuf:"bytes,2,opt,name=media,proto3" json:"media,omitempty"`  // 视频链接(存在时才返回)
}

func (x *ExistVideoResponse) Reset() {
	*x = ExistVideoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_video_rpc_video_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExistVideoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExistVideoResponse) ProtoMessage() {}

func (x *ExistVideoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_video_rpc_video_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExistVideoResponse.ProtoReflect.Descriptor instead.
func (*ExistVideoResponse) Descriptor() ([]byte, []int) {
	return file_cmd_video_rpc_video_proto_rawDescGZIP(), []int{11}
}

func (x *ExistVideoResponse) GetExist() bool {
	if x != nil {
		return x.Exist
	}
	return false
}

func (x *ExistVideoResponse) GetMedia() string {
	if x != nil {
		return x.Media
	}
	return ""
}

var File_cmd_video_rpc_video_proto protoreflect.FileDescriptor

var file_cmd_video_rpc_video_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6d, 0x64, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x61, 0x67,
	0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x03, 0x28, 0x03, 0x52, 0x05, 0x74, 0x61, 0x67, 0x49, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x68, 0x61, 0x32, 0x35, 0x36, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x68, 0x61, 0x32,
	0x35, 0x36, 0x22, 0x25, 0x0a, 0x13, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x66, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49,
	0x64, 0x22, 0x44, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x09, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xa7, 0x02, 0x0a, 0x0a, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72,
	0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f,
	0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c,
	0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x4d, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x52, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x16, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x52, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3f,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x09, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x74, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22,
	0x2a, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x37, 0x0a, 0x11, 0x45,
	0x78, 0x69, 0x73, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x22, 0x0a, 0x0c, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x48, 0x61, 0x73, 0x68, 0x32, 0x35, 0x36,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x48, 0x61, 0x73,
	0x68, 0x32, 0x35, 0x36, 0x22, 0x40, 0x0a, 0x12, 0x45, 0x78, 0x69, 0x73, 0x74, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x78, 0x69, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x32, 0xe2, 0x02, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x12, 0x44, 0x0a, 0x0b, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12,
	0x19, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x52, 0x65, 0x61,
	0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x41,
	0x64, 0x64, 0x52, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65,
	0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3f, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x16, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x44, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x19, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x49, 0x66, 0x45, 0x78, 0x69, 0x73,
	0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x18, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x45,
	0x78, 0x69, 0x73, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e,
	0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x50, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_video_rpc_video_proto_rawDescOnce sync.Once
	file_cmd_video_rpc_video_proto_rawDescData = file_cmd_video_rpc_video_proto_rawDesc
)

func file_cmd_video_rpc_video_proto_rawDescGZIP() []byte {
	file_cmd_video_rpc_video_proto_rawDescOnce.Do(func() {
		file_cmd_video_rpc_video_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_video_rpc_video_proto_rawDescData)
	})
	return file_cmd_video_rpc_video_proto_rawDescData
}

var file_cmd_video_rpc_video_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_cmd_video_rpc_video_proto_goTypes = []interface{}{
	(*UploadVideoRequest)(nil),    // 0: video.UploadVideoRequest
	(*UploadVideoResponse)(nil),   // 1: video.UploadVideoResponse
	(*GetVideoRequest)(nil),       // 2: video.GetVideoRequest
	(*GetVideoResponse)(nil),      // 3: video.GetVideoResponse
	(*VideoModel)(nil),            // 4: video.VideoModel
	(*AddReadCountRequest)(nil),   // 5: video.AddReadCountRequest
	(*AddReadCountResponse)(nil),  // 6: video.AddReadCountResponse
	(*GetTypeListRequest)(nil),    // 7: video.GetTypeListRequest
	(*GetTypeListResponse)(nil),   // 8: video.GetTypeListResponse
	(*Type)(nil),                  // 9: video.Type
	(*ExistVideoRequest)(nil),     // 10: video.ExistVideoRequest
	(*ExistVideoResponse)(nil),    // 11: video.ExistVideoResponse
	(*timestamppb.Timestamp)(nil), // 12: google.protobuf.Timestamp
}
var file_cmd_video_rpc_video_proto_depIdxs = []int32{
	4,  // 0: video.GetVideoResponse.video_list:type_name -> video.VideoModel
	12, // 1: video.VideoModel.create_time:type_name -> google.protobuf.Timestamp
	9,  // 2: video.GetTypeListResponse.type_list:type_name -> video.Type
	0,  // 3: video.Video.UploadVideo:input_type -> video.UploadVideoRequest
	5,  // 4: video.Video.AddReadCount:input_type -> video.AddReadCountRequest
	2,  // 5: video.Video.GetVideoList:input_type -> video.GetVideoRequest
	7,  // 6: video.Video.GetTypeList:input_type -> video.GetTypeListRequest
	10, // 7: video.Video.IfExistVideo:input_type -> video.ExistVideoRequest
	1,  // 8: video.Video.UploadVideo:output_type -> video.UploadVideoResponse
	6,  // 9: video.Video.AddReadCount:output_type -> video.AddReadCountResponse
	3,  // 10: video.Video.GetVideoList:output_type -> video.GetVideoResponse
	8,  // 11: video.Video.GetTypeList:output_type -> video.GetTypeListResponse
	11, // 12: video.Video.IfExistVideo:output_type -> video.ExistVideoResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_cmd_video_rpc_video_proto_init() }
func file_cmd_video_rpc_video_proto_init() {
	if File_cmd_video_rpc_video_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_video_rpc_video_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadVideoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cmd_video_rpc_video_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadVideoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cmd_video_rpc_video_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVideoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cmd_video_rpc_video_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVideoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cmd_video_rpc_video_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoModel); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cmd_video_rpc_video_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddReadCountRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cmd_video_rpc_video_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddReadCountResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cmd_video_rpc_video_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTypeListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cmd_video_rpc_video_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTypeListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cmd_video_rpc_video_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Type); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cmd_video_rpc_video_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExistVideoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cmd_video_rpc_video_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExistVideoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cmd_video_rpc_video_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cmd_video_rpc_video_proto_goTypes,
		DependencyIndexes: file_cmd_video_rpc_video_proto_depIdxs,
		MessageInfos:      file_cmd_video_rpc_video_proto_msgTypes,
	}.Build()
	File_cmd_video_rpc_video_proto = out.File
	file_cmd_video_rpc_video_proto_rawDesc = nil
	file_cmd_video_rpc_video_proto_goTypes = nil
	file_cmd_video_rpc_video_proto_depIdxs = nil
}
