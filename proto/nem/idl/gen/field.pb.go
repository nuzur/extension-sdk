// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.24.4
// source: field.proto

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

type FieldStatus int32

const (
	FieldStatus_FIELD_STATUS_INVALID  FieldStatus = 0
	FieldStatus_FIELD_STATUS_ACTIVE   FieldStatus = 1
	FieldStatus_FIELD_STATUS_INACTIVE FieldStatus = 2
)

// Enum value maps for FieldStatus.
var (
	FieldStatus_name = map[int32]string{
		0: "FIELD_STATUS_INVALID",
		1: "FIELD_STATUS_ACTIVE",
		2: "FIELD_STATUS_INACTIVE",
	}
	FieldStatus_value = map[string]int32{
		"FIELD_STATUS_INVALID":  0,
		"FIELD_STATUS_ACTIVE":   1,
		"FIELD_STATUS_INACTIVE": 2,
	}
)

func (x FieldStatus) Enum() *FieldStatus {
	p := new(FieldStatus)
	*p = x
	return p
}

func (x FieldStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FieldStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_field_proto_enumTypes[0].Descriptor()
}

func (FieldStatus) Type() protoreflect.EnumType {
	return &file_field_proto_enumTypes[0]
}

func (x FieldStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FieldStatus.Descriptor instead.
func (FieldStatus) EnumDescriptor() ([]byte, []int) {
	return file_field_proto_rawDescGZIP(), []int{0}
}

type FieldType int32

const (
	FieldType_FIELD_TYPE_INVALID   FieldType = 0
	FieldType_FIELD_TYPE_UUID      FieldType = 1
	FieldType_FIELD_TYPE_INTEGER   FieldType = 2
	FieldType_FIELD_TYPE_FLOAT     FieldType = 3
	FieldType_FIELD_TYPE_DECIMAL   FieldType = 4
	FieldType_FIELD_TYPE_BOOLEAN   FieldType = 5
	FieldType_FIELD_TYPE_CHAR      FieldType = 6
	FieldType_FIELD_TYPE_VARCHAR   FieldType = 7
	FieldType_FIELD_TYPE_TEXT      FieldType = 8
	FieldType_FIELD_TYPE_ENCRYPTED FieldType = 9
	FieldType_FIELD_TYPE_EMAIL     FieldType = 10
	FieldType_FIELD_TYPE_PHONE     FieldType = 11
	FieldType_FIELD_TYPE_URL       FieldType = 12
	FieldType_FIELD_TYPE_LOCATION  FieldType = 13
	FieldType_FIELD_TYPE_COLOR     FieldType = 14
	FieldType_FIELD_TYPE_RICHTEXT  FieldType = 15
	FieldType_FIELD_TYPE_CODE      FieldType = 16
	FieldType_FIELD_TYPE_MARKDOWN  FieldType = 17
	FieldType_FIELD_TYPE_FILE      FieldType = 18
	FieldType_FIELD_TYPE_IMAGE     FieldType = 19
	FieldType_FIELD_TYPE_AUDIO     FieldType = 20
	FieldType_FIELD_TYPE_VIDEO     FieldType = 21
	FieldType_FIELD_TYPE_ENUM      FieldType = 22
	FieldType_FIELD_TYPE_JSON      FieldType = 23
	FieldType_FIELD_TYPE_ARRAY     FieldType = 24
	FieldType_FIELD_TYPE_DATE      FieldType = 25
	FieldType_FIELD_TYPE_DATETIME  FieldType = 26
	FieldType_FIELD_TYPE_TIME      FieldType = 27
	FieldType_FIELD_TYPE_SLUG      FieldType = 28
)

// Enum value maps for FieldType.
var (
	FieldType_name = map[int32]string{
		0:  "FIELD_TYPE_INVALID",
		1:  "FIELD_TYPE_UUID",
		2:  "FIELD_TYPE_INTEGER",
		3:  "FIELD_TYPE_FLOAT",
		4:  "FIELD_TYPE_DECIMAL",
		5:  "FIELD_TYPE_BOOLEAN",
		6:  "FIELD_TYPE_CHAR",
		7:  "FIELD_TYPE_VARCHAR",
		8:  "FIELD_TYPE_TEXT",
		9:  "FIELD_TYPE_ENCRYPTED",
		10: "FIELD_TYPE_EMAIL",
		11: "FIELD_TYPE_PHONE",
		12: "FIELD_TYPE_URL",
		13: "FIELD_TYPE_LOCATION",
		14: "FIELD_TYPE_COLOR",
		15: "FIELD_TYPE_RICHTEXT",
		16: "FIELD_TYPE_CODE",
		17: "FIELD_TYPE_MARKDOWN",
		18: "FIELD_TYPE_FILE",
		19: "FIELD_TYPE_IMAGE",
		20: "FIELD_TYPE_AUDIO",
		21: "FIELD_TYPE_VIDEO",
		22: "FIELD_TYPE_ENUM",
		23: "FIELD_TYPE_JSON",
		24: "FIELD_TYPE_ARRAY",
		25: "FIELD_TYPE_DATE",
		26: "FIELD_TYPE_DATETIME",
		27: "FIELD_TYPE_TIME",
		28: "FIELD_TYPE_SLUG",
	}
	FieldType_value = map[string]int32{
		"FIELD_TYPE_INVALID":   0,
		"FIELD_TYPE_UUID":      1,
		"FIELD_TYPE_INTEGER":   2,
		"FIELD_TYPE_FLOAT":     3,
		"FIELD_TYPE_DECIMAL":   4,
		"FIELD_TYPE_BOOLEAN":   5,
		"FIELD_TYPE_CHAR":      6,
		"FIELD_TYPE_VARCHAR":   7,
		"FIELD_TYPE_TEXT":      8,
		"FIELD_TYPE_ENCRYPTED": 9,
		"FIELD_TYPE_EMAIL":     10,
		"FIELD_TYPE_PHONE":     11,
		"FIELD_TYPE_URL":       12,
		"FIELD_TYPE_LOCATION":  13,
		"FIELD_TYPE_COLOR":     14,
		"FIELD_TYPE_RICHTEXT":  15,
		"FIELD_TYPE_CODE":      16,
		"FIELD_TYPE_MARKDOWN":  17,
		"FIELD_TYPE_FILE":      18,
		"FIELD_TYPE_IMAGE":     19,
		"FIELD_TYPE_AUDIO":     20,
		"FIELD_TYPE_VIDEO":     21,
		"FIELD_TYPE_ENUM":      22,
		"FIELD_TYPE_JSON":      23,
		"FIELD_TYPE_ARRAY":     24,
		"FIELD_TYPE_DATE":      25,
		"FIELD_TYPE_DATETIME":  26,
		"FIELD_TYPE_TIME":      27,
		"FIELD_TYPE_SLUG":      28,
	}
)

func (x FieldType) Enum() *FieldType {
	p := new(FieldType)
	*p = x
	return p
}

func (x FieldType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FieldType) Descriptor() protoreflect.EnumDescriptor {
	return file_field_proto_enumTypes[1].Descriptor()
}

func (FieldType) Type() protoreflect.EnumType {
	return &file_field_proto_enumTypes[1]
}

func (x FieldType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FieldType.Descriptor instead.
func (FieldType) EnumDescriptor() ([]byte, []int) {
	return file_field_proto_rawDescGZIP(), []int{1}
}

type Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid             string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Version          int64                  `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	Identifier       string                 `protobuf:"bytes,3,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Description      string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Type             FieldType              `protobuf:"varint,5,opt,name=type,proto3,enum=nem.FieldType" json:"type,omitempty"`
	TypeConfig       *FieldTypeConfig       `protobuf:"bytes,6,opt,name=type_config,json=typeConfig,proto3" json:"type_config,omitempty"`
	Required         bool                   `protobuf:"varint,7,opt,name=required,proto3" json:"required,omitempty"`
	Key              bool                   `protobuf:"varint,8,opt,name=key,proto3" json:"key,omitempty"`
	KeyAutoIncrement bool                   `protobuf:"varint,9,opt,name=key_auto_increment,json=keyAutoIncrement,proto3" json:"key_auto_increment,omitempty"`
	Unique           bool                   `protobuf:"varint,10,opt,name=unique,proto3" json:"unique,omitempty"`
	Deprecated       bool                   `protobuf:"varint,11,opt,name=deprecated,proto3" json:"deprecated,omitempty"`
	Generated        bool                   `protobuf:"varint,12,opt,name=generated,proto3" json:"generated,omitempty"`
	Status           FieldStatus            `protobuf:"varint,13,opt,name=status,proto3,enum=nem.FieldStatus" json:"status,omitempty"`
	CreatedAt        *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt        *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CreatedByUuid    string                 `protobuf:"bytes,16,opt,name=created_by_uuid,json=createdByUuid,proto3" json:"created_by_uuid,omitempty"`
	UpdatedByUuid    string                 `protobuf:"bytes,17,opt,name=updated_by_uuid,json=updatedByUuid,proto3" json:"updated_by_uuid,omitempty"`
}

func (x *Field) Reset() {
	*x = Field{}
	mi := &file_field_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Field) ProtoMessage() {}

func (x *Field) ProtoReflect() protoreflect.Message {
	mi := &file_field_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Field.ProtoReflect.Descriptor instead.
func (*Field) Descriptor() ([]byte, []int) {
	return file_field_proto_rawDescGZIP(), []int{0}
}

func (x *Field) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Field) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Field) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *Field) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Field) GetType() FieldType {
	if x != nil {
		return x.Type
	}
	return FieldType_FIELD_TYPE_INVALID
}

func (x *Field) GetTypeConfig() *FieldTypeConfig {
	if x != nil {
		return x.TypeConfig
	}
	return nil
}

func (x *Field) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

func (x *Field) GetKey() bool {
	if x != nil {
		return x.Key
	}
	return false
}

func (x *Field) GetKeyAutoIncrement() bool {
	if x != nil {
		return x.KeyAutoIncrement
	}
	return false
}

func (x *Field) GetUnique() bool {
	if x != nil {
		return x.Unique
	}
	return false
}

func (x *Field) GetDeprecated() bool {
	if x != nil {
		return x.Deprecated
	}
	return false
}

func (x *Field) GetGenerated() bool {
	if x != nil {
		return x.Generated
	}
	return false
}

func (x *Field) GetStatus() FieldStatus {
	if x != nil {
		return x.Status
	}
	return FieldStatus_FIELD_STATUS_INVALID
}

func (x *Field) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Field) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Field) GetCreatedByUuid() string {
	if x != nil {
		return x.CreatedByUuid
	}
	return ""
}

func (x *Field) GetUpdatedByUuid() string {
	if x != nil {
		return x.UpdatedByUuid
	}
	return ""
}

var File_field_proto protoreflect.FileDescriptor

var file_field_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6e,
	0x65, 0x6d, 0x1a, 0x17, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf4, 0x04, 0x0a,
	0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x35, 0x0a, 0x0b, 0x74, 0x79,
	0x70, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0a, 0x74, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x2c, 0x0a, 0x12, 0x6b, 0x65, 0x79, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x69, 0x6e, 0x63, 0x72,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x6b, 0x65, 0x79,
	0x41, 0x75, 0x74, 0x6f, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x75,
	0x6e, 0x69, 0x71, 0x75, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x72, 0x65,
	0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x28, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62,
	0x79, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x55, 0x75, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x55,
	0x75, 0x69, 0x64, 0x2a, 0x5b, 0x0a, 0x0b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x18, 0x0a, 0x14, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13,
	0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x43, 0x54,
	0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x02,
	0x2a, 0x97, 0x05, 0x0a, 0x09, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16,
	0x0a, 0x12, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x55, 0x49, 0x44, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x46,
	0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x47, 0x45,
	0x52, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x46, 0x49, 0x45,
	0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x43, 0x49, 0x4d, 0x41, 0x4c, 0x10,
	0x04, 0x12, 0x16, 0x0a, 0x12, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x42, 0x4f, 0x4f, 0x4c, 0x45, 0x41, 0x4e, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x49, 0x45,
	0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x52, 0x10, 0x06, 0x12, 0x16,
	0x0a, 0x12, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x56, 0x41, 0x52,
	0x43, 0x48, 0x41, 0x52, 0x10, 0x07, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x45, 0x58, 0x54, 0x10, 0x08, 0x12, 0x18, 0x0a, 0x14, 0x46,
	0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x4e, 0x43, 0x52, 0x59, 0x50,
	0x54, 0x45, 0x44, 0x10, 0x09, 0x12, 0x14, 0x0a, 0x10, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x10, 0x0a, 0x12, 0x14, 0x0a, 0x10, 0x46,
	0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x10,
	0x0b, 0x12, 0x12, 0x0a, 0x0e, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x52, 0x4c, 0x10, 0x0c, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0d, 0x12, 0x14,
	0x0a, 0x10, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4c,
	0x4f, 0x52, 0x10, 0x0e, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x52, 0x49, 0x43, 0x48, 0x54, 0x45, 0x58, 0x54, 0x10, 0x0f, 0x12, 0x13, 0x0a,
	0x0f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45,
	0x10, 0x10, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x4d, 0x41, 0x52, 0x4b, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x11, 0x12, 0x13, 0x0a, 0x0f, 0x46,
	0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x12,
	0x12, 0x14, 0x0a, 0x10, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49,
	0x4d, 0x41, 0x47, 0x45, 0x10, 0x13, 0x12, 0x14, 0x0a, 0x10, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x55, 0x44, 0x49, 0x4f, 0x10, 0x14, 0x12, 0x14, 0x0a, 0x10,
	0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x56, 0x49, 0x44, 0x45, 0x4f,
	0x10, 0x15, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x10, 0x16, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x49, 0x45, 0x4c, 0x44,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4a, 0x53, 0x4f, 0x4e, 0x10, 0x17, 0x12, 0x14, 0x0a, 0x10,
	0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x52, 0x52, 0x41, 0x59,
	0x10, 0x18, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x44, 0x41, 0x54, 0x45, 0x10, 0x19, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x49, 0x45, 0x4c, 0x44,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x1a,
	0x12, 0x13, 0x0a, 0x0f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54,
	0x49, 0x4d, 0x45, 0x10, 0x1b, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x53, 0x4c, 0x55, 0x47, 0x10, 0x1c, 0x42, 0x1b, 0x0a, 0x03, 0x6e, 0x65,
	0x6d, 0x42, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x01, 0x5a, 0x0b, 0x6e, 0x65, 0x6d, 0x2f,
	0x69, 0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_field_proto_rawDescOnce sync.Once
	file_field_proto_rawDescData = file_field_proto_rawDesc
)

func file_field_proto_rawDescGZIP() []byte {
	file_field_proto_rawDescOnce.Do(func() {
		file_field_proto_rawDescData = protoimpl.X.CompressGZIP(file_field_proto_rawDescData)
	})
	return file_field_proto_rawDescData
}

var file_field_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_field_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_field_proto_goTypes = []any{
	(FieldStatus)(0),              // 0: nem.FieldStatus
	(FieldType)(0),                // 1: nem.FieldType
	(*Field)(nil),                 // 2: nem.Field
	(*FieldTypeConfig)(nil),       // 3: nem.FieldTypeConfig
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_field_proto_depIdxs = []int32{
	1, // 0: nem.Field.type:type_name -> nem.FieldType
	3, // 1: nem.Field.type_config:type_name -> nem.FieldTypeConfig
	0, // 2: nem.Field.status:type_name -> nem.FieldStatus
	4, // 3: nem.Field.created_at:type_name -> google.protobuf.Timestamp
	4, // 4: nem.Field.updated_at:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_field_proto_init() }
func file_field_proto_init() {
	if File_field_proto != nil {
		return
	}
	file_field_type_config_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_field_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_field_proto_goTypes,
		DependencyIndexes: file_field_proto_depIdxs,
		EnumInfos:         file_field_proto_enumTypes,
		MessageInfos:      file_field_proto_msgTypes,
	}.Build()
	File_field_proto = out.File
	file_field_proto_rawDesc = nil
	file_field_proto_goTypes = nil
	file_field_proto_depIdxs = nil
}
