// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.24.4
// source: extension.proto

package gen

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

type ExtensionExtensionType int32

const (
	ExtensionExtensionType_EXTENSION_EXTENSION_TYPE_INVALID      ExtensionExtensionType = 0
	ExtensionExtensionType_EXTENSION_EXTENSION_TYPE_IMPORTER     ExtensionExtensionType = 1
	ExtensionExtensionType_EXTENSION_EXTENSION_TYPE_SYNCHRONIZER ExtensionExtensionType = 2
	ExtensionExtensionType_EXTENSION_EXTENSION_TYPE_GENERATOR    ExtensionExtensionType = 3
)

// Enum value maps for ExtensionExtensionType.
var (
	ExtensionExtensionType_name = map[int32]string{
		0: "EXTENSION_EXTENSION_TYPE_INVALID",
		1: "EXTENSION_EXTENSION_TYPE_IMPORTER",
		2: "EXTENSION_EXTENSION_TYPE_SYNCHRONIZER",
		3: "EXTENSION_EXTENSION_TYPE_GENERATOR",
	}
	ExtensionExtensionType_value = map[string]int32{
		"EXTENSION_EXTENSION_TYPE_INVALID":      0,
		"EXTENSION_EXTENSION_TYPE_IMPORTER":     1,
		"EXTENSION_EXTENSION_TYPE_SYNCHRONIZER": 2,
		"EXTENSION_EXTENSION_TYPE_GENERATOR":    3,
	}
)

func (x ExtensionExtensionType) Enum() *ExtensionExtensionType {
	p := new(ExtensionExtensionType)
	*p = x
	return p
}

func (x ExtensionExtensionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExtensionExtensionType) Descriptor() protoreflect.EnumDescriptor {
	return file_extension_proto_enumTypes[0].Descriptor()
}

func (ExtensionExtensionType) Type() protoreflect.EnumType {
	return &file_extension_proto_enumTypes[0]
}

func (x ExtensionExtensionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExtensionExtensionType.Descriptor instead.
func (ExtensionExtensionType) EnumDescriptor() ([]byte, []int) {
	return file_extension_proto_rawDescGZIP(), []int{0}
}

type ExtensionStatus int32

const (
	ExtensionStatus_EXTENSION_STATUS_INVALID  ExtensionStatus = 0
	ExtensionStatus_EXTENSION_STATUS_ACTIVE   ExtensionStatus = 1
	ExtensionStatus_EXTENSION_STATUS_DISABLED ExtensionStatus = 2
)

// Enum value maps for ExtensionStatus.
var (
	ExtensionStatus_name = map[int32]string{
		0: "EXTENSION_STATUS_INVALID",
		1: "EXTENSION_STATUS_ACTIVE",
		2: "EXTENSION_STATUS_DISABLED",
	}
	ExtensionStatus_value = map[string]int32{
		"EXTENSION_STATUS_INVALID":  0,
		"EXTENSION_STATUS_ACTIVE":   1,
		"EXTENSION_STATUS_DISABLED": 2,
	}
)

func (x ExtensionStatus) Enum() *ExtensionStatus {
	p := new(ExtensionStatus)
	*p = x
	return p
}

func (x ExtensionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExtensionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_extension_proto_enumTypes[1].Descriptor()
}

func (ExtensionStatus) Type() protoreflect.EnumType {
	return &file_extension_proto_enumTypes[1]
}

func (x ExtensionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExtensionStatus.Descriptor instead.
func (ExtensionStatus) EnumDescriptor() ([]byte, []int) {
	return file_extension_proto_rawDescGZIP(), []int{1}
}

type Extension struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid          string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Version       int64                  `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Url           string                 `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	Repository    string                 `protobuf:"bytes,6,opt,name=repository,proto3" json:"repository,omitempty"`
	ExtensionType ExtensionExtensionType `protobuf:"varint,7,opt,name=extension_type,json=extensionType,proto3,enum=nem.ExtensionExtensionType" json:"extension_type,omitempty"`
	Tags          []string               `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
	Public        bool                   `protobuf:"varint,9,opt,name=public,proto3" json:"public,omitempty"`
	Visibility    []*Visibility          `protobuf:"bytes,10,rep,name=visibility,proto3" json:"visibility,omitempty"`
	Status        ExtensionStatus        `protobuf:"varint,11,opt,name=status,proto3,enum=nem.ExtensionStatus" json:"status,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CreatedByUuid string                 `protobuf:"bytes,14,opt,name=created_by_uuid,json=createdByUuid,proto3" json:"created_by_uuid,omitempty"`
	UpdatedByUuid string                 `protobuf:"bytes,15,opt,name=updated_by_uuid,json=updatedByUuid,proto3" json:"updated_by_uuid,omitempty"`
}

func (x *Extension) Reset() {
	*x = Extension{}
	mi := &file_extension_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Extension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Extension) ProtoMessage() {}

func (x *Extension) ProtoReflect() protoreflect.Message {
	mi := &file_extension_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Extension.ProtoReflect.Descriptor instead.
func (*Extension) Descriptor() ([]byte, []int) {
	return file_extension_proto_rawDescGZIP(), []int{0}
}

func (x *Extension) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Extension) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Extension) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Extension) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Extension) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Extension) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *Extension) GetExtensionType() ExtensionExtensionType {
	if x != nil {
		return x.ExtensionType
	}
	return ExtensionExtensionType_EXTENSION_EXTENSION_TYPE_INVALID
}

func (x *Extension) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Extension) GetPublic() bool {
	if x != nil {
		return x.Public
	}
	return false
}

func (x *Extension) GetVisibility() []*Visibility {
	if x != nil {
		return x.Visibility
	}
	return nil
}

func (x *Extension) GetStatus() ExtensionStatus {
	if x != nil {
		return x.Status
	}
	return ExtensionStatus_EXTENSION_STATUS_INVALID
}

func (x *Extension) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Extension) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Extension) GetCreatedByUuid() string {
	if x != nil {
		return x.CreatedByUuid
	}
	return ""
}

func (x *Extension) GetUpdatedByUuid() string {
	if x != nil {
		return x.UpdatedByUuid
	}
	return ""
}

var File_extension_proto protoreflect.FileDescriptor

var file_extension_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x6e, 0x65, 0x6d, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x04, 0x0a, 0x09, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1e, 0x0a,
	0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x42, 0x0a,
	0x0e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0d, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x2f, 0x0a,
	0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x52, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x2c,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14,
	0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79,
	0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x55, 0x75, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x55, 0x75,
	0x69, 0x64, 0x2a, 0xb8, 0x01, 0x0a, 0x16, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a,
	0x20, 0x45, 0x58, 0x54, 0x45, 0x4e, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x58, 0x54, 0x45, 0x4e,
	0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x10, 0x00, 0x12, 0x25, 0x0a, 0x21, 0x45, 0x58, 0x54, 0x45, 0x4e, 0x53, 0x49, 0x4f, 0x4e,
	0x5f, 0x45, 0x58, 0x54, 0x45, 0x4e, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x52, 0x10, 0x01, 0x12, 0x29, 0x0a, 0x25, 0x45, 0x58,
	0x54, 0x45, 0x4e, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x58, 0x54, 0x45, 0x4e, 0x53, 0x49, 0x4f,
	0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x48, 0x52, 0x4f, 0x4e, 0x49,
	0x5a, 0x45, 0x52, 0x10, 0x02, 0x12, 0x26, 0x0a, 0x22, 0x45, 0x58, 0x54, 0x45, 0x4e, 0x53, 0x49,
	0x4f, 0x4e, 0x5f, 0x45, 0x58, 0x54, 0x45, 0x4e, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x10, 0x03, 0x2a, 0x6b, 0x0a,
	0x0f, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1c, 0x0a, 0x18, 0x45, 0x58, 0x54, 0x45, 0x4e, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x1b,
	0x0a, 0x17, 0x45, 0x58, 0x54, 0x45, 0x4e, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x45,
	0x58, 0x54, 0x45, 0x4e, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x42, 0x1f, 0x0a, 0x03, 0x6e, 0x65,
	0x6d, 0x42, 0x09, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x01, 0x5a, 0x0b,
	0x6e, 0x65, 0x6d, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_extension_proto_rawDescOnce sync.Once
	file_extension_proto_rawDescData = file_extension_proto_rawDesc
)

func file_extension_proto_rawDescGZIP() []byte {
	file_extension_proto_rawDescOnce.Do(func() {
		file_extension_proto_rawDescData = protoimpl.X.CompressGZIP(file_extension_proto_rawDescData)
	})
	return file_extension_proto_rawDescData
}

var file_extension_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_extension_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_extension_proto_goTypes = []any{
	(ExtensionExtensionType)(0),   // 0: nem.ExtensionExtensionType
	(ExtensionStatus)(0),          // 1: nem.ExtensionStatus
	(*Extension)(nil),             // 2: nem.Extension
	(*Visibility)(nil),            // 3: nem.Visibility
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_extension_proto_depIdxs = []int32{
	0, // 0: nem.Extension.extension_type:type_name -> nem.ExtensionExtensionType
	3, // 1: nem.Extension.visibility:type_name -> nem.Visibility
	1, // 2: nem.Extension.status:type_name -> nem.ExtensionStatus
	4, // 3: nem.Extension.created_at:type_name -> google.protobuf.Timestamp
	4, // 4: nem.Extension.updated_at:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_extension_proto_init() }
func file_extension_proto_init() {
	if File_extension_proto != nil {
		return
	}
	file_visibility_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_extension_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_extension_proto_goTypes,
		DependencyIndexes: file_extension_proto_depIdxs,
		EnumInfos:         file_extension_proto_enumTypes,
		MessageInfos:      file_extension_proto_msgTypes,
	}.Build()
	File_extension_proto = out.File
	file_extension_proto_rawDesc = nil
	file_extension_proto_goTypes = nil
	file_extension_proto_depIdxs = nil
}
