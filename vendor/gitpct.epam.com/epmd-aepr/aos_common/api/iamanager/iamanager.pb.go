// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: iamanager.proto

package iamanager

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CreateKeysReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	SystemId string `protobuf:"bytes,2,opt,name=system_id,json=systemId,proto3" json:"system_id,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CreateKeysReq) Reset() {
	*x = CreateKeysReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iamanager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateKeysReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateKeysReq) ProtoMessage() {}

func (x *CreateKeysReq) ProtoReflect() protoreflect.Message {
	mi := &file_iamanager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateKeysReq.ProtoReflect.Descriptor instead.
func (*CreateKeysReq) Descriptor() ([]byte, []int) {
	return file_iamanager_proto_rawDescGZIP(), []int{0}
}

func (x *CreateKeysReq) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateKeysReq) GetSystemId() string {
	if x != nil {
		return x.SystemId
	}
	return ""
}

func (x *CreateKeysReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CreateKeysRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Csr   string `protobuf:"bytes,2,opt,name=csr,proto3" json:"csr,omitempty"`
	Error string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CreateKeysRsp) Reset() {
	*x = CreateKeysRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iamanager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateKeysRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateKeysRsp) ProtoMessage() {}

func (x *CreateKeysRsp) ProtoReflect() protoreflect.Message {
	mi := &file_iamanager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateKeysRsp.ProtoReflect.Descriptor instead.
func (*CreateKeysRsp) Descriptor() ([]byte, []int) {
	return file_iamanager_proto_rawDescGZIP(), []int{1}
}

func (x *CreateKeysRsp) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateKeysRsp) GetCsr() string {
	if x != nil {
		return x.Csr
	}
	return ""
}

func (x *CreateKeysRsp) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ApplyCertReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Cert string `protobuf:"bytes,2,opt,name=cert,proto3" json:"cert,omitempty"`
}

func (x *ApplyCertReq) Reset() {
	*x = ApplyCertReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iamanager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyCertReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyCertReq) ProtoMessage() {}

func (x *ApplyCertReq) ProtoReflect() protoreflect.Message {
	mi := &file_iamanager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyCertReq.ProtoReflect.Descriptor instead.
func (*ApplyCertReq) Descriptor() ([]byte, []int) {
	return file_iamanager_proto_rawDescGZIP(), []int{2}
}

func (x *ApplyCertReq) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ApplyCertReq) GetCert() string {
	if x != nil {
		return x.Cert
	}
	return ""
}

type ApplyCertRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	CertUrl string `protobuf:"bytes,2,opt,name=cert_url,json=certUrl,proto3" json:"cert_url,omitempty"`
	Error   string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ApplyCertRsp) Reset() {
	*x = ApplyCertRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iamanager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyCertRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyCertRsp) ProtoMessage() {}

func (x *ApplyCertRsp) ProtoReflect() protoreflect.Message {
	mi := &file_iamanager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyCertRsp.ProtoReflect.Descriptor instead.
func (*ApplyCertRsp) Descriptor() ([]byte, []int) {
	return file_iamanager_proto_rawDescGZIP(), []int{3}
}

func (x *ApplyCertRsp) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ApplyCertRsp) GetCertUrl() string {
	if x != nil {
		return x.CertUrl
	}
	return ""
}

func (x *ApplyCertRsp) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetCertReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Issuer []byte `protobuf:"bytes,2,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Serial string `protobuf:"bytes,3,opt,name=serial,proto3" json:"serial,omitempty"`
}

func (x *GetCertReq) Reset() {
	*x = GetCertReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iamanager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCertReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCertReq) ProtoMessage() {}

func (x *GetCertReq) ProtoReflect() protoreflect.Message {
	mi := &file_iamanager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCertReq.ProtoReflect.Descriptor instead.
func (*GetCertReq) Descriptor() ([]byte, []int) {
	return file_iamanager_proto_rawDescGZIP(), []int{4}
}

func (x *GetCertReq) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GetCertReq) GetIssuer() []byte {
	if x != nil {
		return x.Issuer
	}
	return nil
}

func (x *GetCertReq) GetSerial() string {
	if x != nil {
		return x.Serial
	}
	return ""
}

type GetCertRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	CertUrl string `protobuf:"bytes,2,opt,name=cert_url,json=certUrl,proto3" json:"cert_url,omitempty"`
	KeyUrl  string `protobuf:"bytes,3,opt,name=key_url,json=keyUrl,proto3" json:"key_url,omitempty"`
	Error   string `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GetCertRsp) Reset() {
	*x = GetCertRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iamanager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCertRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCertRsp) ProtoMessage() {}

func (x *GetCertRsp) ProtoReflect() protoreflect.Message {
	mi := &file_iamanager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCertRsp.ProtoReflect.Descriptor instead.
func (*GetCertRsp) Descriptor() ([]byte, []int) {
	return file_iamanager_proto_rawDescGZIP(), []int{5}
}

func (x *GetCertRsp) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GetCertRsp) GetCertUrl() string {
	if x != nil {
		return x.CertUrl
	}
	return ""
}

func (x *GetCertRsp) GetKeyUrl() string {
	if x != nil {
		return x.KeyUrl
	}
	return ""
}

func (x *GetCertRsp) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_iamanager_proto protoreflect.FileDescriptor

var file_iamanager_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x69, 0x61, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x69, 0x61, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x22, 0x5c, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x4b, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x63, 0x73, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x73,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x36, 0x0a, 0x0c, 0x41, 0x70, 0x70, 0x6c, 0x79,
	0x43, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x65, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x65, 0x72, 0x74, 0x22,
	0x53, 0x0a, 0x0c, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x43, 0x65, 0x72, 0x74, 0x52, 0x73, 0x70, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x65, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x22, 0x50, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x65, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x22, 0x6a, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x65, 0x72,
	0x74, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x65, 0x72, 0x74,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x65, 0x72, 0x74,
	0x55, 0x72, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6b, 0x65, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x32, 0xcb, 0x01, 0x0a, 0x09, 0x49, 0x41, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x12, 0x42, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x18,
	0x2e, 0x69, 0x61, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x69, 0x61, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x73, 0x52,
	0x73, 0x70, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x43, 0x65, 0x72,
	0x74, 0x12, 0x17, 0x2e, 0x69, 0x61, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x41, 0x70,
	0x70, 0x6c, 0x79, 0x43, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x69, 0x61, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x43, 0x65, 0x72, 0x74,
	0x52, 0x73, 0x70, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x65, 0x72, 0x74,
	0x12, 0x15, 0x2e, 0x69, 0x61, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x69, 0x61, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x65, 0x72, 0x74, 0x52, 0x73, 0x70, 0x22, 0x00,
	0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x70, 0x63, 0x74, 0x2e, 0x65, 0x70, 0x61, 0x6d, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x70, 0x6d, 0x64, 0x2d, 0x61, 0x65, 0x70, 0x72, 0x2f, 0x61, 0x6f,
	0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x61, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_iamanager_proto_rawDescOnce sync.Once
	file_iamanager_proto_rawDescData = file_iamanager_proto_rawDesc
)

func file_iamanager_proto_rawDescGZIP() []byte {
	file_iamanager_proto_rawDescOnce.Do(func() {
		file_iamanager_proto_rawDescData = protoimpl.X.CompressGZIP(file_iamanager_proto_rawDescData)
	})
	return file_iamanager_proto_rawDescData
}

var file_iamanager_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_iamanager_proto_goTypes = []interface{}{
	(*CreateKeysReq)(nil), // 0: iamanager.CreateKeysReq
	(*CreateKeysRsp)(nil), // 1: iamanager.CreateKeysRsp
	(*ApplyCertReq)(nil),  // 2: iamanager.ApplyCertReq
	(*ApplyCertRsp)(nil),  // 3: iamanager.ApplyCertRsp
	(*GetCertReq)(nil),    // 4: iamanager.GetCertReq
	(*GetCertRsp)(nil),    // 5: iamanager.GetCertRsp
}
var file_iamanager_proto_depIdxs = []int32{
	0, // 0: iamanager.IAManager.CreateKeys:input_type -> iamanager.CreateKeysReq
	2, // 1: iamanager.IAManager.ApplyCert:input_type -> iamanager.ApplyCertReq
	4, // 2: iamanager.IAManager.GetCert:input_type -> iamanager.GetCertReq
	1, // 3: iamanager.IAManager.CreateKeys:output_type -> iamanager.CreateKeysRsp
	3, // 4: iamanager.IAManager.ApplyCert:output_type -> iamanager.ApplyCertRsp
	5, // 5: iamanager.IAManager.GetCert:output_type -> iamanager.GetCertRsp
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_iamanager_proto_init() }
func file_iamanager_proto_init() {
	if File_iamanager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_iamanager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateKeysReq); i {
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
		file_iamanager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateKeysRsp); i {
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
		file_iamanager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyCertReq); i {
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
		file_iamanager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyCertRsp); i {
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
		file_iamanager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCertReq); i {
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
		file_iamanager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCertRsp); i {
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
			RawDescriptor: file_iamanager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_iamanager_proto_goTypes,
		DependencyIndexes: file_iamanager_proto_depIdxs,
		MessageInfos:      file_iamanager_proto_msgTypes,
	}.Build()
	File_iamanager_proto = out.File
	file_iamanager_proto_rawDesc = nil
	file_iamanager_proto_goTypes = nil
	file_iamanager_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// IAManagerClient is the client API for IAManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IAManagerClient interface {
	CreateKeys(ctx context.Context, in *CreateKeysReq, opts ...grpc.CallOption) (*CreateKeysRsp, error)
	ApplyCert(ctx context.Context, in *ApplyCertReq, opts ...grpc.CallOption) (*ApplyCertRsp, error)
	GetCert(ctx context.Context, in *GetCertReq, opts ...grpc.CallOption) (*GetCertRsp, error)
}

type iAManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewIAManagerClient(cc grpc.ClientConnInterface) IAManagerClient {
	return &iAManagerClient{cc}
}

func (c *iAManagerClient) CreateKeys(ctx context.Context, in *CreateKeysReq, opts ...grpc.CallOption) (*CreateKeysRsp, error) {
	out := new(CreateKeysRsp)
	err := c.cc.Invoke(ctx, "/iamanager.IAManager/CreateKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAManagerClient) ApplyCert(ctx context.Context, in *ApplyCertReq, opts ...grpc.CallOption) (*ApplyCertRsp, error) {
	out := new(ApplyCertRsp)
	err := c.cc.Invoke(ctx, "/iamanager.IAManager/ApplyCert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAManagerClient) GetCert(ctx context.Context, in *GetCertReq, opts ...grpc.CallOption) (*GetCertRsp, error) {
	out := new(GetCertRsp)
	err := c.cc.Invoke(ctx, "/iamanager.IAManager/GetCert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IAManagerServer is the server API for IAManager service.
type IAManagerServer interface {
	CreateKeys(context.Context, *CreateKeysReq) (*CreateKeysRsp, error)
	ApplyCert(context.Context, *ApplyCertReq) (*ApplyCertRsp, error)
	GetCert(context.Context, *GetCertReq) (*GetCertRsp, error)
}

// UnimplementedIAManagerServer can be embedded to have forward compatible implementations.
type UnimplementedIAManagerServer struct {
}

func (*UnimplementedIAManagerServer) CreateKeys(context.Context, *CreateKeysReq) (*CreateKeysRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateKeys not implemented")
}
func (*UnimplementedIAManagerServer) ApplyCert(context.Context, *ApplyCertReq) (*ApplyCertRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyCert not implemented")
}
func (*UnimplementedIAManagerServer) GetCert(context.Context, *GetCertReq) (*GetCertRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCert not implemented")
}

func RegisterIAManagerServer(s *grpc.Server, srv IAManagerServer) {
	s.RegisterService(&_IAManager_serviceDesc, srv)
}

func _IAManager_CreateKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateKeysReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAManagerServer).CreateKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iamanager.IAManager/CreateKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAManagerServer).CreateKeys(ctx, req.(*CreateKeysReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAManager_ApplyCert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyCertReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAManagerServer).ApplyCert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iamanager.IAManager/ApplyCert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAManagerServer).ApplyCert(ctx, req.(*ApplyCertReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAManager_GetCert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCertReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAManagerServer).GetCert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iamanager.IAManager/GetCert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAManagerServer).GetCert(ctx, req.(*GetCertReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _IAManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "iamanager.IAManager",
	HandlerType: (*IAManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateKeys",
			Handler:    _IAManager_CreateKeys_Handler,
		},
		{
			MethodName: "ApplyCert",
			Handler:    _IAManager_ApplyCert_Handler,
		},
		{
			MethodName: "GetCert",
			Handler:    _IAManager_GetCert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iamanager.proto",
}