// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: proto/background_service.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *PublicKey) Reset() {
	*x = PublicKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_background_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicKey) ProtoMessage() {}

func (x *PublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_proto_background_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicKey.ProtoReflect.Descriptor instead.
func (*PublicKey) Descriptor() ([]byte, []int) {
	return file_proto_background_service_proto_rawDescGZIP(), []int{0}
}

func (x *PublicKey) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

type InitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkspaceName string `protobuf:"bytes,1,opt,name=workspace_name,json=workspaceName,proto3" json:"workspace_name,omitempty"`
	Username      string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password      string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	PublicKey     []byte `protobuf:"bytes,4,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Port          string `protobuf:"bytes,5,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *InitRequest) Reset() {
	*x = InitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_background_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitRequest) ProtoMessage() {}

func (x *InitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_background_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitRequest.ProtoReflect.Descriptor instead.
func (*InitRequest) Descriptor() ([]byte, []int) {
	return file_proto_background_service_proto_rawDescGZIP(), []int{1}
}

func (x *InitRequest) GetWorkspaceName() string {
	if x != nil {
		return x.WorkspaceName
	}
	return ""
}

func (x *InitRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *InitRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *InitRequest) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *InitRequest) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

type InitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response int32 `protobuf:"varint,1,opt,name=response,proto3" json:"response,omitempty"` // 200 [Valid / ACK / OK] ||| 4000 [InValid / You Fucked Up Somewhere]
	Port     int32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *InitResponse) Reset() {
	*x = InitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_background_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitResponse) ProtoMessage() {}

func (x *InitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_background_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitResponse.ProtoReflect.Descriptor instead.
func (*InitResponse) Descriptor() ([]byte, []int) {
	return file_proto_background_service_proto_rawDescGZIP(), []int{2}
}

func (x *InitResponse) GetResponse() int32 {
	if x != nil {
		return x.Response
	}
	return 0
}

func (x *InitResponse) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type WorkspaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkspaceName string `protobuf:"bytes,1,opt,name=workspace_name,json=workspaceName,proto3" json:"workspace_name,omitempty"`
	Password      string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *WorkspaceRequest) Reset() {
	*x = WorkspaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_background_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkspaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkspaceRequest) ProtoMessage() {}

func (x *WorkspaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_background_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkspaceRequest.ProtoReflect.Descriptor instead.
func (*WorkspaceRequest) Descriptor() ([]byte, []int) {
	return file_proto_background_service_proto_rawDescGZIP(), []int{3}
}

func (x *WorkspaceRequest) GetWorkspaceName() string {
	if x != nil {
		return x.WorkspaceName
	}
	return ""
}

func (x *WorkspaceRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type GetWorkspaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response int32 `protobuf:"varint,1,opt,name=response,proto3" json:"response,omitempty"` // 200 [Valid / ACK / OK] ||| 4000 [InValid / You Fucked Up Somewhere]
	Port     int32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *GetWorkspaceResponse) Reset() {
	*x = GetWorkspaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_background_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWorkspaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkspaceResponse) ProtoMessage() {}

func (x *GetWorkspaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_background_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorkspaceResponse.ProtoReflect.Descriptor instead.
func (*GetWorkspaceResponse) Descriptor() ([]byte, []int) {
	return file_proto_background_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetWorkspaceResponse) GetResponse() int32 {
	if x != nil {
		return x.Response
	}
	return 0
}

func (x *GetWorkspaceResponse) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type NotifyPushRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkspaceName string `protobuf:"bytes,1,opt,name=workspace_name,json=workspaceName,proto3" json:"workspace_name,omitempty"`
}

func (x *NotifyPushRequest) Reset() {
	*x = NotifyPushRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_background_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyPushRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyPushRequest) ProtoMessage() {}

func (x *NotifyPushRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_background_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyPushRequest.ProtoReflect.Descriptor instead.
func (*NotifyPushRequest) Descriptor() ([]byte, []int) {
	return file_proto_background_service_proto_rawDescGZIP(), []int{5}
}

func (x *NotifyPushRequest) GetWorkspaceName() string {
	if x != nil {
		return x.WorkspaceName
	}
	return ""
}

type NotifyPushResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response int32 `protobuf:"varint,1,opt,name=response,proto3" json:"response,omitempty"` // 200 [Valid ??]
}

func (x *NotifyPushResponse) Reset() {
	*x = NotifyPushResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_background_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyPushResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyPushResponse) ProtoMessage() {}

func (x *NotifyPushResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_background_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyPushResponse.ProtoReflect.Descriptor instead.
func (*NotifyPushResponse) Descriptor() ([]byte, []int) {
	return file_proto_background_service_proto_rawDescGZIP(), []int{6}
}

func (x *NotifyPushResponse) GetResponse() int32 {
	if x != nil {
		return x.Response
	}
	return 0
}

type DataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkspaceName string `protobuf:"bytes,1,opt,name=workspace_name,json=workspaceName,proto3" json:"workspace_name,omitempty"`
	ConnectionIp  string `protobuf:"bytes,2,opt,name=connection_ip,json=connectionIp,proto3" json:"connection_ip,omitempty"`
}

func (x *DataRequest) Reset() {
	*x = DataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_background_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataRequest) ProtoMessage() {}

func (x *DataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_background_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataRequest.ProtoReflect.Descriptor instead.
func (*DataRequest) Descriptor() ([]byte, []int) {
	return file_proto_background_service_proto_rawDescGZIP(), []int{7}
}

func (x *DataRequest) GetWorkspaceName() string {
	if x != nil {
		return x.WorkspaceName
	}
	return ""
}

func (x *DataRequest) GetConnectionIp() string {
	if x != nil {
		return x.ConnectionIp
	}
	return ""
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filetype uint32 `protobuf:"varint,1,opt,name=filetype,proto3" json:"filetype,omitempty"` // Data[0] or AES Key[1] or AES IV[2]
	Chunk    []byte `protobuf:"bytes,2,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_background_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_proto_background_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_proto_background_service_proto_rawDescGZIP(), []int{8}
}

func (x *Data) GetFiletype() uint32 {
	if x != nil {
		return x.Filetype
	}
	return 0
}

func (x *Data) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

var File_proto_background_service_proto protoreflect.FileDescriptor

var file_proto_background_service_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x1d, 0x0a, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x22, 0x9f, 0x01, 0x0a, 0x0b, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x25, 0x0a, 0x0e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x22, 0x3e, 0x0a, 0x0c, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x22, 0x55, 0x0a, 0x10, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x46, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x22, 0x3a, 0x0a, 0x11, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x50, 0x75, 0x73, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x30, 0x0a,
	0x12, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x59, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25,
	0x0a, 0x0e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x70, 0x22, 0x38, 0x0a, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x32, 0xf8, 0x02, 0x0a, 0x11, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f,
	0x75, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x12, 0x5f, 0x0a, 0x1a, 0x49, 0x6e, 0x69, 0x74, 0x4e, 0x65, 0x77, 0x57, 0x6f, 0x72, 0x6b,
	0x53, 0x70, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1f, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x5e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x12, 0x24, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x67,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x5b, 0x0a, 0x0a, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x50, 0x75, 0x73, 0x68,
	0x12, 0x25, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x50, 0x75, 0x73, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0x55, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x2e, 0x62, 0x61, 0x63, 0x6b,
	0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x62, 0x61, 0x63,
	0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x30, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_background_service_proto_rawDescOnce sync.Once
	file_proto_background_service_proto_rawDescData = file_proto_background_service_proto_rawDesc
)

func file_proto_background_service_proto_rawDescGZIP() []byte {
	file_proto_background_service_proto_rawDescOnce.Do(func() {
		file_proto_background_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_background_service_proto_rawDescData)
	})
	return file_proto_background_service_proto_rawDescData
}

var file_proto_background_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_background_service_proto_goTypes = []any{
	(*PublicKey)(nil),            // 0: background_service.PublicKey
	(*InitRequest)(nil),          // 1: background_service.InitRequest
	(*InitResponse)(nil),         // 2: background_service.InitResponse
	(*WorkspaceRequest)(nil),     // 3: background_service.WorkspaceRequest
	(*GetWorkspaceResponse)(nil), // 4: background_service.GetWorkspaceResponse
	(*NotifyPushRequest)(nil),    // 5: background_service.NotifyPushRequest
	(*NotifyPushResponse)(nil),   // 6: background_service.NotifyPushResponse
	(*DataRequest)(nil),          // 7: background_service.DataRequest
	(*Data)(nil),                 // 8: background_service.Data
	(*emptypb.Empty)(nil),        // 9: google.protobuf.Empty
}
var file_proto_background_service_proto_depIdxs = []int32{
	9, // 0: background_service.BackgroundService.GetPublicKey:input_type -> google.protobuf.Empty
	1, // 1: background_service.BackgroundService.InitNewWorkSpaceConnection:input_type -> background_service.InitRequest
	3, // 2: background_service.BackgroundService.GetWorkspace:input_type -> background_service.WorkspaceRequest
	5, // 3: background_service.BackgroundService.NotifyPush:input_type -> background_service.NotifyPushRequest
	7, // 4: background_service.DataService.GetData:input_type -> background_service.DataRequest
	0, // 5: background_service.BackgroundService.GetPublicKey:output_type -> background_service.PublicKey
	2, // 6: background_service.BackgroundService.InitNewWorkSpaceConnection:output_type -> background_service.InitResponse
	4, // 7: background_service.BackgroundService.GetWorkspace:output_type -> background_service.GetWorkspaceResponse
	6, // 8: background_service.BackgroundService.NotifyPush:output_type -> background_service.NotifyPushResponse
	8, // 9: background_service.DataService.GetData:output_type -> background_service.Data
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_background_service_proto_init() }
func file_proto_background_service_proto_init() {
	if File_proto_background_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_background_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PublicKey); i {
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
		file_proto_background_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*InitRequest); i {
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
		file_proto_background_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*InitResponse); i {
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
		file_proto_background_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*WorkspaceRequest); i {
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
		file_proto_background_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetWorkspaceResponse); i {
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
		file_proto_background_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*NotifyPushRequest); i {
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
		file_proto_background_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*NotifyPushResponse); i {
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
		file_proto_background_service_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*DataRequest); i {
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
		file_proto_background_service_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*Data); i {
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
			RawDescriptor: file_proto_background_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_background_service_proto_goTypes,
		DependencyIndexes: file_proto_background_service_proto_depIdxs,
		MessageInfos:      file_proto_background_service_proto_msgTypes,
	}.Build()
	File_proto_background_service_proto = out.File
	file_proto_background_service_proto_rawDesc = nil
	file_proto_background_service_proto_goTypes = nil
	file_proto_background_service_proto_depIdxs = nil
}
