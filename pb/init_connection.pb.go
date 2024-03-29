// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: init_connection.proto

package pb

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

type Workspaces struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workspaceinfos []*WorkspaceInfo `protobuf:"bytes,1,rep,name=workspaceinfos,proto3" json:"workspaceinfos,omitempty"`
}

func (x *Workspaces) Reset() {
	*x = Workspaces{}
	if protoimpl.UnsafeEnabled {
		mi := &file_init_connection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Workspaces) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Workspaces) ProtoMessage() {}

func (x *Workspaces) ProtoReflect() protoreflect.Message {
	mi := &file_init_connection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Workspaces.ProtoReflect.Descriptor instead.
func (*Workspaces) Descriptor() ([]byte, []int) {
	return file_init_connection_proto_rawDescGZIP(), []int{0}
}

func (x *Workspaces) GetWorkspaceinfos() []*WorkspaceInfo {
	if x != nil {
		return x.Workspaceinfos
	}
	return nil
}

type WorkspaceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkspaceName string `protobuf:"bytes,1,opt,name=workspace_name,json=workspaceName,proto3" json:"workspace_name,omitempty"`
}

func (x *WorkspaceInfo) Reset() {
	*x = WorkspaceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_init_connection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkspaceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkspaceInfo) ProtoMessage() {}

func (x *WorkspaceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_init_connection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkspaceInfo.ProtoReflect.Descriptor instead.
func (*WorkspaceInfo) Descriptor() ([]byte, []int) {
	return file_init_connection_proto_rawDescGZIP(), []int{1}
}

func (x *WorkspaceInfo) GetWorkspaceName() string {
	if x != nil {
		return x.WorkspaceName
	}
	return ""
}

type OTP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username  string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	IpAddress string `protobuf:"bytes,2,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	Otp       int32  `protobuf:"varint,3,opt,name=otp,proto3" json:"otp,omitempty"`
}

func (x *OTP) Reset() {
	*x = OTP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_init_connection_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OTP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OTP) ProtoMessage() {}

func (x *OTP) ProtoReflect() protoreflect.Message {
	mi := &file_init_connection_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OTP.ProtoReflect.Descriptor instead.
func (*OTP) Descriptor() ([]byte, []int) {
	return file_init_connection_proto_rawDescGZIP(), []int{2}
}

func (x *OTP) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *OTP) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

func (x *OTP) GetOtp() int32 {
	if x != nil {
		return x.Otp
	}
	return 0
}

type OTPResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IfOtpCorrect   bool   `protobuf:"varint,1,opt,name=if_otp_correct,json=ifOtpCorrect,proto3" json:"if_otp_correct,omitempty"`
	ConnectionSlug string `protobuf:"bytes,2,opt,name=connection_slug,json=connectionSlug,proto3" json:"connection_slug,omitempty"`
	PublicKey      string `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *OTPResponse) Reset() {
	*x = OTPResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_init_connection_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OTPResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OTPResponse) ProtoMessage() {}

func (x *OTPResponse) ProtoReflect() protoreflect.Message {
	mi := &file_init_connection_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OTPResponse.ProtoReflect.Descriptor instead.
func (*OTPResponse) Descriptor() ([]byte, []int) {
	return file_init_connection_proto_rawDescGZIP(), []int{3}
}

func (x *OTPResponse) GetIfOtpCorrect() bool {
	if x != nil {
		return x.IfOtpCorrect
	}
	return false
}

func (x *OTPResponse) GetConnectionSlug() string {
	if x != nil {
		return x.ConnectionSlug
	}
	return ""
}

func (x *OTPResponse) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

type Certificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnectionPassword string `protobuf:"bytes,1,opt,name=connection_password,json=connectionPassword,proto3" json:"connection_password,omitempty"`
	PublicKey          string `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *Certificate) Reset() {
	*x = Certificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_init_connection_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Certificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Certificate) ProtoMessage() {}

func (x *Certificate) ProtoReflect() protoreflect.Message {
	mi := &file_init_connection_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Certificate.ProtoReflect.Descriptor instead.
func (*Certificate) Descriptor() ([]byte, []int) {
	return file_init_connection_proto_rawDescGZIP(), []int{4}
}

func (x *Certificate) GetConnectionPassword() string {
	if x != nil {
		return x.ConnectionPassword
	}
	return ""
}

func (x *Certificate) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

type CertificateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommandConnectionPort int32 `protobuf:"varint,1,opt,name=command_connection_port,json=commandConnectionPort,proto3" json:"command_connection_port,omitempty"`
}

func (x *CertificateResponse) Reset() {
	*x = CertificateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_init_connection_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertificateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertificateResponse) ProtoMessage() {}

func (x *CertificateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_init_connection_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertificateResponse.ProtoReflect.Descriptor instead.
func (*CertificateResponse) Descriptor() ([]byte, []int) {
	return file_init_connection_proto_rawDescGZIP(), []int{5}
}

func (x *CertificateResponse) GetCommandConnectionPort() int32 {
	if x != nil {
		return x.CommandConnectionPort
	}
	return 0
}

var File_init_connection_proto protoreflect.FileDescriptor

var file_init_connection_proto_rawDesc = []byte{
	0x0a, 0x15, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x54, 0x0a, 0x0a, 0x57, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x46, 0x0a, 0x0e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x36,
	0x0a, 0x0d, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x25, 0x0a, 0x0e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x52, 0x0a, 0x03, 0x4f, 0x54, 0x50, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x70, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69,
	0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x74, 0x70, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6f, 0x74, 0x70, 0x22, 0x7b, 0x0a, 0x0b, 0x4f, 0x54,
	0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x66, 0x5f,
	0x6f, 0x74, 0x70, 0x5f, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0c, 0x69, 0x66, 0x4f, 0x74, 0x70, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x12,
	0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x6c,
	0x75, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x6c, 0x75, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0x5d, 0x0a, 0x0b, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0x4d, 0x0a, 0x13, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a,
	0x17, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x6f, 0x72, 0x74, 0x32, 0xfb, 0x01, 0x0a, 0x0e, 0x49, 0x6e, 0x69, 0x74, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x09, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x4f, 0x54, 0x50, 0x12, 0x14, 0x2e, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x54, 0x50, 0x1a, 0x1c, 0x2e, 0x69, 0x6e,
	0x69, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x54,
	0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x14, 0x45, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x73, 0x12, 0x1c, 0x2e, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x1a,
	0x24, 0x2e, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x10, 0x47, 0x45, 0x54, 0x57, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x2e, 0x69, 0x6e, 0x69, 0x74,
	0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x57, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x1a, 0x1b, 0x2e, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x73, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x6d, 0x79, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_init_connection_proto_rawDescOnce sync.Once
	file_init_connection_proto_rawDescData = file_init_connection_proto_rawDesc
)

func file_init_connection_proto_rawDescGZIP() []byte {
	file_init_connection_proto_rawDescOnce.Do(func() {
		file_init_connection_proto_rawDescData = protoimpl.X.CompressGZIP(file_init_connection_proto_rawDescData)
	})
	return file_init_connection_proto_rawDescData
}

var file_init_connection_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_init_connection_proto_goTypes = []interface{}{
	(*Workspaces)(nil),          // 0: init_connection.Workspaces
	(*WorkspaceInfo)(nil),       // 1: init_connection.WorkspaceInfo
	(*OTP)(nil),                 // 2: init_connection.OTP
	(*OTPResponse)(nil),         // 3: init_connection.OTPResponse
	(*Certificate)(nil),         // 4: init_connection.Certificate
	(*CertificateResponse)(nil), // 5: init_connection.CertificateResponse
}
var file_init_connection_proto_depIdxs = []int32{
	1, // 0: init_connection.Workspaces.workspaceinfos:type_name -> init_connection.WorkspaceInfo
	2, // 1: init_connection.InitConnection.VerifyOTP:input_type -> init_connection.OTP
	4, // 2: init_connection.InitConnection.ExchangeCertificates:input_type -> init_connection.Certificate
	0, // 3: init_connection.InitConnection.GETWorkspaceinfo:input_type -> init_connection.Workspaces
	3, // 4: init_connection.InitConnection.VerifyOTP:output_type -> init_connection.OTPResponse
	5, // 5: init_connection.InitConnection.ExchangeCertificates:output_type -> init_connection.CertificateResponse
	0, // 6: init_connection.InitConnection.GETWorkspaceinfo:output_type -> init_connection.Workspaces
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_init_connection_proto_init() }
func file_init_connection_proto_init() {
	if File_init_connection_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_init_connection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Workspaces); i {
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
		file_init_connection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkspaceInfo); i {
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
		file_init_connection_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OTP); i {
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
		file_init_connection_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OTPResponse); i {
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
		file_init_connection_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Certificate); i {
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
		file_init_connection_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertificateResponse); i {
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
			RawDescriptor: file_init_connection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_init_connection_proto_goTypes,
		DependencyIndexes: file_init_connection_proto_depIdxs,
		MessageInfos:      file_init_connection_proto_msgTypes,
	}.Build()
	File_init_connection_proto = out.File
	file_init_connection_proto_rawDesc = nil
	file_init_connection_proto_goTypes = nil
	file_init_connection_proto_depIdxs = nil
}
