// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: like.proto

package like

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LikeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizId    string `protobuf:"bytes,1,opt,name=bizId,proto3" json:"bizId,omitempty"`        // 业务id
	ObjId    int64  `protobuf:"varint,2,opt,name=objId,proto3" json:"objId,omitempty"`       // 点赞对象id
	UserId   int64  `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`     // 用户id
	LikeType int32  `protobuf:"varint,4,opt,name=likeType,proto3" json:"likeType,omitempty"` // 类型
}

func (x *LikeReq) Reset() {
	*x = LikeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeReq) ProtoMessage() {}

func (x *LikeReq) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeReq.ProtoReflect.Descriptor instead.
func (*LikeReq) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{0}
}

func (x *LikeReq) GetBizId() string {
	if x != nil {
		return x.BizId
	}
	return ""
}

func (x *LikeReq) GetObjId() int64 {
	if x != nil {
		return x.ObjId
	}
	return 0
}

func (x *LikeReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *LikeReq) GetLikeType() int32 {
	if x != nil {
		return x.LikeType
	}
	return 0
}

type LikeRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizId      string `protobuf:"bytes,1,opt,name=bizId,proto3" json:"bizId,omitempty"`            // 业务id
	ObjId      int64  `protobuf:"varint,2,opt,name=objId,proto3" json:"objId,omitempty"`           // 点赞对象id
	LikeNum    int64  `protobuf:"varint,3,opt,name=likeNum,proto3" json:"likeNum,omitempty"`       // 点赞数
	DislikeNum int64  `protobuf:"varint,4,opt,name=dislikeNum,proto3" json:"dislikeNum,omitempty"` // 点踩数
}

func (x *LikeRes) Reset() {
	*x = LikeRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeRes) ProtoMessage() {}

func (x *LikeRes) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeRes.ProtoReflect.Descriptor instead.
func (*LikeRes) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{1}
}

func (x *LikeRes) GetBizId() string {
	if x != nil {
		return x.BizId
	}
	return ""
}

func (x *LikeRes) GetObjId() int64 {
	if x != nil {
		return x.ObjId
	}
	return 0
}

func (x *LikeRes) GetLikeNum() int64 {
	if x != nil {
		return x.LikeNum
	}
	return 0
}

func (x *LikeRes) GetDislikeNum() int64 {
	if x != nil {
		return x.DislikeNum
	}
	return 0
}

type IsLikeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizId  string `protobuf:"bytes,1,opt,name=bizId,proto3" json:"bizId,omitempty"`    // 业务id
	ObjId  int64  `protobuf:"varint,2,opt,name=objId,proto3" json:"objId,omitempty"`   // 点赞对象id
	UserId int64  `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"` // 用户id
}

func (x *IsLikeReq) Reset() {
	*x = IsLikeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsLikeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsLikeReq) ProtoMessage() {}

func (x *IsLikeReq) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsLikeReq.ProtoReflect.Descriptor instead.
func (*IsLikeReq) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{2}
}

func (x *IsLikeReq) GetBizId() string {
	if x != nil {
		return x.BizId
	}
	return ""
}

func (x *IsLikeReq) GetObjId() int64 {
	if x != nil {
		return x.ObjId
	}
	return 0
}

func (x *IsLikeReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type IsLikeRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserThumbups map[int64]*UserThumbup `protobuf:"bytes,1,rep,name=userThumbups,proto3" json:"userThumbups,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *IsLikeRes) Reset() {
	*x = IsLikeRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsLikeRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsLikeRes) ProtoMessage() {}

func (x *IsLikeRes) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsLikeRes.ProtoReflect.Descriptor instead.
func (*IsLikeRes) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{3}
}

func (x *IsLikeRes) GetUserThumbups() map[int64]*UserThumbup {
	if x != nil {
		return x.UserThumbups
	}
	return nil
}

type UserThumbup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ThumbupTime int64 `protobuf:"varint,2,opt,name=thumbupTime,proto3" json:"thumbupTime,omitempty"`
	LikeType    int32 `protobuf:"varint,3,opt,name=likeType,proto3" json:"likeType,omitempty"` // 类型
}

func (x *UserThumbup) Reset() {
	*x = UserThumbup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserThumbup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserThumbup) ProtoMessage() {}

func (x *UserThumbup) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserThumbup.ProtoReflect.Descriptor instead.
func (*UserThumbup) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{4}
}

func (x *UserThumbup) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserThumbup) GetThumbupTime() int64 {
	if x != nil {
		return x.ThumbupTime
	}
	return 0
}

func (x *UserThumbup) GetLikeType() int32 {
	if x != nil {
		return x.LikeType
	}
	return 0
}

type LikeCountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizId string `protobuf:"bytes,1,opt,name=bizId,proto3" json:"bizId,omitempty"`  // 业务id
	ObjId int64  `protobuf:"varint,2,opt,name=objId,proto3" json:"objId,omitempty"` // 点赞对象id
}

func (x *LikeCountReq) Reset() {
	*x = LikeCountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeCountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeCountReq) ProtoMessage() {}

func (x *LikeCountReq) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeCountReq.ProtoReflect.Descriptor instead.
func (*LikeCountReq) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{5}
}

func (x *LikeCountReq) GetBizId() string {
	if x != nil {
		return x.BizId
	}
	return ""
}

func (x *LikeCountReq) GetObjId() int64 {
	if x != nil {
		return x.ObjId
	}
	return 0
}

type LikeCountRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LikeNum int64 `protobuf:"varint,1,opt,name=likeNum,proto3" json:"likeNum,omitempty"`
}

func (x *LikeCountRes) Reset() {
	*x = LikeCountRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeCountRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeCountRes) ProtoMessage() {}

func (x *LikeCountRes) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeCountRes.ProtoReflect.Descriptor instead.
func (*LikeCountRes) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{6}
}

func (x *LikeCountRes) GetLikeNum() int64 {
	if x != nil {
		return x.LikeNum
	}
	return 0
}

var File_like_proto protoreflect.FileDescriptor

var file_like_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6c, 0x69,
	0x6b, 0x65, 0x22, 0x69, 0x0a, 0x07, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a,
	0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x69,
	0x7a, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x62, 0x6a, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x6f, 0x62, 0x6a, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x69, 0x6b, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x6c, 0x69, 0x6b, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x6f, 0x0a,
	0x07, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x69, 0x7a, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x6f, 0x62, 0x6a, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6f,
	0x62, 0x6a, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x69, 0x6b, 0x65, 0x4e, 0x75, 0x6d, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6c, 0x69, 0x6b, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1e,
	0x0a, 0x0a, 0x64, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x64, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x4e, 0x75, 0x6d, 0x22, 0x4f,
	0x0a, 0x09, 0x49, 0x73, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x62,
	0x69, 0x7a, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x69, 0x7a, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x62, 0x6a, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x6f, 0x62, 0x6a, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0xa6, 0x01, 0x0a, 0x09, 0x49, 0x73, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x12, 0x45, 0x0a,
	0x0c, 0x75, 0x73, 0x65, 0x72, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x49, 0x73, 0x4c, 0x69, 0x6b,
	0x65, 0x52, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x75, 0x70,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x54, 0x68, 0x75, 0x6d,
	0x62, 0x75, 0x70, 0x73, 0x1a, 0x52, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x54, 0x68, 0x75, 0x6d,
	0x62, 0x75, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6c, 0x69, 0x6b,
	0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x75, 0x70, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x63, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72,
	0x54, 0x68, 0x75, 0x6d, 0x62, 0x75, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x75, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x75, 0x70, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x69, 0x6b, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x6c, 0x69, 0x6b, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x0a,
	0x0c, 0x4c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a,
	0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x69,
	0x7a, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x62, 0x6a, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x6f, 0x62, 0x6a, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x0c, 0x4c, 0x69, 0x6b,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x69, 0x6b,
	0x65, 0x4e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6c, 0x69, 0x6b, 0x65,
	0x4e, 0x75, 0x6d, 0x32, 0x8d, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x24, 0x0a, 0x04,
	0x4c, 0x69, 0x6b, 0x65, 0x12, 0x0d, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x4c, 0x69, 0x6b, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x52,
	0x65, 0x73, 0x12, 0x2a, 0x0a, 0x06, 0x49, 0x73, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x0f, 0x2e, 0x6c,
	0x69, 0x6b, 0x65, 0x2e, 0x49, 0x73, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e,
	0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x49, 0x73, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x12, 0x33,
	0x0a, 0x09, 0x4c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x2e, 0x6c, 0x69,
	0x6b, 0x65, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a,
	0x12, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x6c, 0x69, 0x6b, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_like_proto_rawDescOnce sync.Once
	file_like_proto_rawDescData = file_like_proto_rawDesc
)

func file_like_proto_rawDescGZIP() []byte {
	file_like_proto_rawDescOnce.Do(func() {
		file_like_proto_rawDescData = protoimpl.X.CompressGZIP(file_like_proto_rawDescData)
	})
	return file_like_proto_rawDescData
}

var file_like_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_like_proto_goTypes = []interface{}{
	(*LikeReq)(nil),      // 0: like.LikeReq
	(*LikeRes)(nil),      // 1: like.LikeRes
	(*IsLikeReq)(nil),    // 2: like.IsLikeReq
	(*IsLikeRes)(nil),    // 3: like.IsLikeRes
	(*UserThumbup)(nil),  // 4: like.UserThumbup
	(*LikeCountReq)(nil), // 5: like.LikeCountReq
	(*LikeCountRes)(nil), // 6: like.LikeCountRes
	nil,                  // 7: like.IsLikeRes.UserThumbupsEntry
}
var file_like_proto_depIdxs = []int32{
	7, // 0: like.IsLikeRes.userThumbups:type_name -> like.IsLikeRes.UserThumbupsEntry
	4, // 1: like.IsLikeRes.UserThumbupsEntry.value:type_name -> like.UserThumbup
	0, // 2: like.Like.Like:input_type -> like.LikeReq
	2, // 3: like.Like.IsLike:input_type -> like.IsLikeReq
	5, // 4: like.Like.LikeCount:input_type -> like.LikeCountReq
	1, // 5: like.Like.Like:output_type -> like.LikeRes
	3, // 6: like.Like.IsLike:output_type -> like.IsLikeRes
	6, // 7: like.Like.LikeCount:output_type -> like.LikeCountRes
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_like_proto_init() }
func file_like_proto_init() {
	if File_like_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_like_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikeReq); i {
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
		file_like_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikeRes); i {
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
		file_like_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsLikeReq); i {
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
		file_like_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsLikeRes); i {
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
		file_like_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserThumbup); i {
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
		file_like_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikeCountReq); i {
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
		file_like_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikeCountRes); i {
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
			RawDescriptor: file_like_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_like_proto_goTypes,
		DependencyIndexes: file_like_proto_depIdxs,
		MessageInfos:      file_like_proto_msgTypes,
	}.Build()
	File_like_proto = out.File
	file_like_proto_rawDesc = nil
	file_like_proto_goTypes = nil
	file_like_proto_depIdxs = nil
}
