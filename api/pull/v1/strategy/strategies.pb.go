// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: pull/v1/strategy/strategies.proto

package strategy

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateStrategiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateStrategiesRequest) Reset() {
	*x = CreateStrategiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pull_v1_strategy_strategies_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStrategiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStrategiesRequest) ProtoMessage() {}

func (x *CreateStrategiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pull_v1_strategy_strategies_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStrategiesRequest.ProtoReflect.Descriptor instead.
func (*CreateStrategiesRequest) Descriptor() ([]byte, []int) {
	return file_pull_v1_strategy_strategies_proto_rawDescGZIP(), []int{0}
}

type CreateStrategiesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateStrategiesReply) Reset() {
	*x = CreateStrategiesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pull_v1_strategy_strategies_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStrategiesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStrategiesReply) ProtoMessage() {}

func (x *CreateStrategiesReply) ProtoReflect() protoreflect.Message {
	mi := &file_pull_v1_strategy_strategies_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStrategiesReply.ProtoReflect.Descriptor instead.
func (*CreateStrategiesReply) Descriptor() ([]byte, []int) {
	return file_pull_v1_strategy_strategies_proto_rawDescGZIP(), []int{1}
}

type UpdateStrategiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateStrategiesRequest) Reset() {
	*x = UpdateStrategiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pull_v1_strategy_strategies_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStrategiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStrategiesRequest) ProtoMessage() {}

func (x *UpdateStrategiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pull_v1_strategy_strategies_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStrategiesRequest.ProtoReflect.Descriptor instead.
func (*UpdateStrategiesRequest) Descriptor() ([]byte, []int) {
	return file_pull_v1_strategy_strategies_proto_rawDescGZIP(), []int{2}
}

type UpdateStrategiesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateStrategiesReply) Reset() {
	*x = UpdateStrategiesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pull_v1_strategy_strategies_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStrategiesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStrategiesReply) ProtoMessage() {}

func (x *UpdateStrategiesReply) ProtoReflect() protoreflect.Message {
	mi := &file_pull_v1_strategy_strategies_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStrategiesReply.ProtoReflect.Descriptor instead.
func (*UpdateStrategiesReply) Descriptor() ([]byte, []int) {
	return file_pull_v1_strategy_strategies_proto_rawDescGZIP(), []int{3}
}

type DeleteStrategiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteStrategiesRequest) Reset() {
	*x = DeleteStrategiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pull_v1_strategy_strategies_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStrategiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStrategiesRequest) ProtoMessage() {}

func (x *DeleteStrategiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pull_v1_strategy_strategies_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStrategiesRequest.ProtoReflect.Descriptor instead.
func (*DeleteStrategiesRequest) Descriptor() ([]byte, []int) {
	return file_pull_v1_strategy_strategies_proto_rawDescGZIP(), []int{4}
}

type DeleteStrategiesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteStrategiesReply) Reset() {
	*x = DeleteStrategiesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pull_v1_strategy_strategies_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStrategiesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStrategiesReply) ProtoMessage() {}

func (x *DeleteStrategiesReply) ProtoReflect() protoreflect.Message {
	mi := &file_pull_v1_strategy_strategies_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStrategiesReply.ProtoReflect.Descriptor instead.
func (*DeleteStrategiesReply) Descriptor() ([]byte, []int) {
	return file_pull_v1_strategy_strategies_proto_rawDescGZIP(), []int{5}
}

type GetStrategiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetStrategiesRequest) Reset() {
	*x = GetStrategiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pull_v1_strategy_strategies_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStrategiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStrategiesRequest) ProtoMessage() {}

func (x *GetStrategiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pull_v1_strategy_strategies_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStrategiesRequest.ProtoReflect.Descriptor instead.
func (*GetStrategiesRequest) Descriptor() ([]byte, []int) {
	return file_pull_v1_strategy_strategies_proto_rawDescGZIP(), []int{6}
}

type GetStrategiesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetStrategiesReply) Reset() {
	*x = GetStrategiesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pull_v1_strategy_strategies_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStrategiesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStrategiesReply) ProtoMessage() {}

func (x *GetStrategiesReply) ProtoReflect() protoreflect.Message {
	mi := &file_pull_v1_strategy_strategies_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStrategiesReply.ProtoReflect.Descriptor instead.
func (*GetStrategiesReply) Descriptor() ([]byte, []int) {
	return file_pull_v1_strategy_strategies_proto_rawDescGZIP(), []int{7}
}

type ListStrategiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListStrategiesRequest) Reset() {
	*x = ListStrategiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pull_v1_strategy_strategies_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStrategiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStrategiesRequest) ProtoMessage() {}

func (x *ListStrategiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pull_v1_strategy_strategies_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStrategiesRequest.ProtoReflect.Descriptor instead.
func (*ListStrategiesRequest) Descriptor() ([]byte, []int) {
	return file_pull_v1_strategy_strategies_proto_rawDescGZIP(), []int{8}
}

type ListStrategiesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListStrategiesReply) Reset() {
	*x = ListStrategiesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pull_v1_strategy_strategies_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStrategiesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStrategiesReply) ProtoMessage() {}

func (x *ListStrategiesReply) ProtoReflect() protoreflect.Message {
	mi := &file_pull_v1_strategy_strategies_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStrategiesReply.ProtoReflect.Descriptor instead.
func (*ListStrategiesReply) Descriptor() ([]byte, []int) {
	return file_pull_v1_strategy_strategies_proto_rawDescGZIP(), []int{9}
}

var File_pull_v1_strategy_strategies_proto protoreflect.FileDescriptor

var file_pull_v1_strategy_strategies_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x75, 0x6c, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x14, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x6c, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x19, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x17, 0x0a, 0x15, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x19, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x17, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x19, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x17, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x16, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x69, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x17, 0x0a, 0x15, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x69, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xad, 0x04, 0x0a, 0x0a, 0x53,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x12, 0x6e, 0x0a, 0x10, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x12, 0x2d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x6c, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x75, 0x6c, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x69, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x6e, 0x0a, 0x10, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x12, 0x2d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x6c, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x75, 0x6c, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x69, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x6e, 0x0a, 0x10, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x12, 0x2d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x6c, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x75, 0x6c, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x69, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x65, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x12, 0x2a, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x75, 0x6c, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x6c,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x68, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69,
	0x65, 0x73, 0x12, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x6c, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x6c, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x69, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x4a, 0x0a, 0x14, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x75, 0x6c, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x50, 0x01, 0x5a, 0x30, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73,
	0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x75, 0x6c,
	0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x3b, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pull_v1_strategy_strategies_proto_rawDescOnce sync.Once
	file_pull_v1_strategy_strategies_proto_rawDescData = file_pull_v1_strategy_strategies_proto_rawDesc
)

func file_pull_v1_strategy_strategies_proto_rawDescGZIP() []byte {
	file_pull_v1_strategy_strategies_proto_rawDescOnce.Do(func() {
		file_pull_v1_strategy_strategies_proto_rawDescData = protoimpl.X.CompressGZIP(file_pull_v1_strategy_strategies_proto_rawDescData)
	})
	return file_pull_v1_strategy_strategies_proto_rawDescData
}

var file_pull_v1_strategy_strategies_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pull_v1_strategy_strategies_proto_goTypes = []interface{}{
	(*CreateStrategiesRequest)(nil), // 0: api.pull.v1.strategy.CreateStrategiesRequest
	(*CreateStrategiesReply)(nil),   // 1: api.pull.v1.strategy.CreateStrategiesReply
	(*UpdateStrategiesRequest)(nil), // 2: api.pull.v1.strategy.UpdateStrategiesRequest
	(*UpdateStrategiesReply)(nil),   // 3: api.pull.v1.strategy.UpdateStrategiesReply
	(*DeleteStrategiesRequest)(nil), // 4: api.pull.v1.strategy.DeleteStrategiesRequest
	(*DeleteStrategiesReply)(nil),   // 5: api.pull.v1.strategy.DeleteStrategiesReply
	(*GetStrategiesRequest)(nil),    // 6: api.pull.v1.strategy.GetStrategiesRequest
	(*GetStrategiesReply)(nil),      // 7: api.pull.v1.strategy.GetStrategiesReply
	(*ListStrategiesRequest)(nil),   // 8: api.pull.v1.strategy.ListStrategiesRequest
	(*ListStrategiesReply)(nil),     // 9: api.pull.v1.strategy.ListStrategiesReply
}
var file_pull_v1_strategy_strategies_proto_depIdxs = []int32{
	0, // 0: api.pull.v1.strategy.Strategies.CreateStrategies:input_type -> api.pull.v1.strategy.CreateStrategiesRequest
	2, // 1: api.pull.v1.strategy.Strategies.UpdateStrategies:input_type -> api.pull.v1.strategy.UpdateStrategiesRequest
	4, // 2: api.pull.v1.strategy.Strategies.DeleteStrategies:input_type -> api.pull.v1.strategy.DeleteStrategiesRequest
	6, // 3: api.pull.v1.strategy.Strategies.GetStrategies:input_type -> api.pull.v1.strategy.GetStrategiesRequest
	8, // 4: api.pull.v1.strategy.Strategies.ListStrategies:input_type -> api.pull.v1.strategy.ListStrategiesRequest
	1, // 5: api.pull.v1.strategy.Strategies.CreateStrategies:output_type -> api.pull.v1.strategy.CreateStrategiesReply
	3, // 6: api.pull.v1.strategy.Strategies.UpdateStrategies:output_type -> api.pull.v1.strategy.UpdateStrategiesReply
	5, // 7: api.pull.v1.strategy.Strategies.DeleteStrategies:output_type -> api.pull.v1.strategy.DeleteStrategiesReply
	7, // 8: api.pull.v1.strategy.Strategies.GetStrategies:output_type -> api.pull.v1.strategy.GetStrategiesReply
	9, // 9: api.pull.v1.strategy.Strategies.ListStrategies:output_type -> api.pull.v1.strategy.ListStrategiesReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pull_v1_strategy_strategies_proto_init() }
func file_pull_v1_strategy_strategies_proto_init() {
	if File_pull_v1_strategy_strategies_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pull_v1_strategy_strategies_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStrategiesRequest); i {
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
		file_pull_v1_strategy_strategies_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStrategiesReply); i {
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
		file_pull_v1_strategy_strategies_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStrategiesRequest); i {
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
		file_pull_v1_strategy_strategies_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStrategiesReply); i {
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
		file_pull_v1_strategy_strategies_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStrategiesRequest); i {
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
		file_pull_v1_strategy_strategies_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStrategiesReply); i {
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
		file_pull_v1_strategy_strategies_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStrategiesRequest); i {
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
		file_pull_v1_strategy_strategies_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStrategiesReply); i {
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
		file_pull_v1_strategy_strategies_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStrategiesRequest); i {
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
		file_pull_v1_strategy_strategies_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStrategiesReply); i {
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
			RawDescriptor: file_pull_v1_strategy_strategies_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pull_v1_strategy_strategies_proto_goTypes,
		DependencyIndexes: file_pull_v1_strategy_strategies_proto_depIdxs,
		MessageInfos:      file_pull_v1_strategy_strategies_proto_msgTypes,
	}.Build()
	File_pull_v1_strategy_strategies_proto = out.File
	file_pull_v1_strategy_strategies_proto_rawDesc = nil
	file_pull_v1_strategy_strategies_proto_goTypes = nil
	file_pull_v1_strategy_strategies_proto_depIdxs = nil
}
