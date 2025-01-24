// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.3
// source: connection.proto

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

type ConnectionDbType int32

const (
	ConnectionDbType_CONNECTION_DB_TYPE_INVALID  ConnectionDbType = 0
	ConnectionDbType_CONNECTION_DB_TYPE_MYSQL    ConnectionDbType = 1
	ConnectionDbType_CONNECTION_DB_TYPE_POSTGRES ConnectionDbType = 2
)

// Enum value maps for ConnectionDbType.
var (
	ConnectionDbType_name = map[int32]string{
		0: "CONNECTION_DB_TYPE_INVALID",
		1: "CONNECTION_DB_TYPE_MYSQL",
		2: "CONNECTION_DB_TYPE_POSTGRES",
	}
	ConnectionDbType_value = map[string]int32{
		"CONNECTION_DB_TYPE_INVALID":  0,
		"CONNECTION_DB_TYPE_MYSQL":    1,
		"CONNECTION_DB_TYPE_POSTGRES": 2,
	}
)

func (x ConnectionDbType) Enum() *ConnectionDbType {
	p := new(ConnectionDbType)
	*p = x
	return p
}

func (x ConnectionDbType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConnectionDbType) Descriptor() protoreflect.EnumDescriptor {
	return file_connection_proto_enumTypes[0].Descriptor()
}

func (ConnectionDbType) Type() protoreflect.EnumType {
	return &file_connection_proto_enumTypes[0]
}

func (x ConnectionDbType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConnectionDbType.Descriptor instead.
func (ConnectionDbType) EnumDescriptor() ([]byte, []int) {
	return file_connection_proto_rawDescGZIP(), []int{0}
}

type ConnectionStatus int32

const (
	ConnectionStatus_CONNECTION_STATUS_INVALID  ConnectionStatus = 0
	ConnectionStatus_CONNECTION_STATUS_ACTIVE   ConnectionStatus = 1
	ConnectionStatus_CONNECTION_STATUS_DISABLED ConnectionStatus = 2
)

// Enum value maps for ConnectionStatus.
var (
	ConnectionStatus_name = map[int32]string{
		0: "CONNECTION_STATUS_INVALID",
		1: "CONNECTION_STATUS_ACTIVE",
		2: "CONNECTION_STATUS_DISABLED",
	}
	ConnectionStatus_value = map[string]int32{
		"CONNECTION_STATUS_INVALID":  0,
		"CONNECTION_STATUS_ACTIVE":   1,
		"CONNECTION_STATUS_DISABLED": 2,
	}
)

func (x ConnectionStatus) Enum() *ConnectionStatus {
	p := new(ConnectionStatus)
	*p = x
	return p
}

func (x ConnectionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConnectionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_connection_proto_enumTypes[1].Descriptor()
}

func (ConnectionStatus) Type() protoreflect.EnumType {
	return &file_connection_proto_enumTypes[1]
}

func (x ConnectionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConnectionStatus.Descriptor instead.
func (ConnectionStatus) EnumDescriptor() ([]byte, []int) {
	return file_connection_proto_rawDescGZIP(), []int{1}
}

type ConnectionType int32

const (
	ConnectionType_CONNECTION_TYPE_INVALID    ConnectionType = 0
	ConnectionType_CONNECTION_TYPE_TCP_IP     ConnectionType = 1
	ConnectionType_CONNECTION_TYPE_TCP_IP_SSH ConnectionType = 2
)

// Enum value maps for ConnectionType.
var (
	ConnectionType_name = map[int32]string{
		0: "CONNECTION_TYPE_INVALID",
		1: "CONNECTION_TYPE_TCP_IP",
		2: "CONNECTION_TYPE_TCP_IP_SSH",
	}
	ConnectionType_value = map[string]int32{
		"CONNECTION_TYPE_INVALID":    0,
		"CONNECTION_TYPE_TCP_IP":     1,
		"CONNECTION_TYPE_TCP_IP_SSH": 2,
	}
)

func (x ConnectionType) Enum() *ConnectionType {
	p := new(ConnectionType)
	*p = x
	return p
}

func (x ConnectionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConnectionType) Descriptor() protoreflect.EnumDescriptor {
	return file_connection_proto_enumTypes[2].Descriptor()
}

func (ConnectionType) Type() protoreflect.EnumType {
	return &file_connection_proto_enumTypes[2]
}

func (x ConnectionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConnectionType.Descriptor instead.
func (ConnectionType) EnumDescriptor() ([]byte, []int) {
	return file_connection_proto_rawDescGZIP(), []int{2}
}

type Connection struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Uuid           string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	StoreUuid      string                 `protobuf:"bytes,2,opt,name=store_uuid,json=storeUuid,proto3" json:"store_uuid,omitempty"`
	EnviormentUuid string                 `protobuf:"bytes,3,opt,name=enviorment_uuid,json=enviormentUuid,proto3" json:"enviorment_uuid,omitempty"`
	Identifier     string                 `protobuf:"bytes,4,opt,name=identifier,proto3" json:"identifier,omitempty"`
	DbType         ConnectionDbType       `protobuf:"varint,5,opt,name=db_type,json=dbType,proto3,enum=nem.ConnectionDbType" json:"db_type,omitempty"`
	DbTypeConfig   *DbTypeConfig          `protobuf:"bytes,6,opt,name=db_type_config,json=dbTypeConfig,proto3" json:"db_type_config,omitempty"`
	DbVersion      string                 `protobuf:"bytes,7,opt,name=db_version,json=dbVersion,proto3" json:"db_version,omitempty"`
	Type           ConnectionType         `protobuf:"varint,8,opt,name=type,proto3,enum=nem.ConnectionType" json:"type,omitempty"`
	TypeConfig     *ConnectionTypeConfig  `protobuf:"bytes,9,opt,name=type_config,json=typeConfig,proto3" json:"type_config,omitempty"`
	Status         ConnectionStatus       `protobuf:"varint,10,opt,name=status,proto3,enum=nem.ConnectionStatus" json:"status,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt      *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CreatedByUuid  string                 `protobuf:"bytes,13,opt,name=created_by_uuid,json=createdByUuid,proto3" json:"created_by_uuid,omitempty"`
	UpdatedByUuid  string                 `protobuf:"bytes,14,opt,name=updated_by_uuid,json=updatedByUuid,proto3" json:"updated_by_uuid,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Connection) Reset() {
	*x = Connection{}
	mi := &file_connection_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Connection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Connection) ProtoMessage() {}

func (x *Connection) ProtoReflect() protoreflect.Message {
	mi := &file_connection_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Connection.ProtoReflect.Descriptor instead.
func (*Connection) Descriptor() ([]byte, []int) {
	return file_connection_proto_rawDescGZIP(), []int{0}
}

func (x *Connection) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Connection) GetStoreUuid() string {
	if x != nil {
		return x.StoreUuid
	}
	return ""
}

func (x *Connection) GetEnviormentUuid() string {
	if x != nil {
		return x.EnviormentUuid
	}
	return ""
}

func (x *Connection) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *Connection) GetDbType() ConnectionDbType {
	if x != nil {
		return x.DbType
	}
	return ConnectionDbType_CONNECTION_DB_TYPE_INVALID
}

func (x *Connection) GetDbTypeConfig() *DbTypeConfig {
	if x != nil {
		return x.DbTypeConfig
	}
	return nil
}

func (x *Connection) GetDbVersion() string {
	if x != nil {
		return x.DbVersion
	}
	return ""
}

func (x *Connection) GetType() ConnectionType {
	if x != nil {
		return x.Type
	}
	return ConnectionType_CONNECTION_TYPE_INVALID
}

func (x *Connection) GetTypeConfig() *ConnectionTypeConfig {
	if x != nil {
		return x.TypeConfig
	}
	return nil
}

func (x *Connection) GetStatus() ConnectionStatus {
	if x != nil {
		return x.Status
	}
	return ConnectionStatus_CONNECTION_STATUS_INVALID
}

func (x *Connection) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Connection) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Connection) GetCreatedByUuid() string {
	if x != nil {
		return x.CreatedByUuid
	}
	return ""
}

func (x *Connection) GetUpdatedByUuid() string {
	if x != nil {
		return x.UpdatedByUuid
	}
	return ""
}

var File_connection_proto protoreflect.FileDescriptor

var file_connection_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x6e, 0x65, 0x6d, 0x1a, 0x1c, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x64, 0x62, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xea, 0x04, 0x0a,
	0x0a, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x55, 0x75, 0x69, 0x64, 0x12, 0x27,
	0x0a, 0x0f, 0x65, 0x6e, 0x76, 0x69, 0x6f, 0x72, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x6e, 0x76, 0x69, 0x6f, 0x72, 0x6d,
	0x65, 0x6e, 0x74, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x07, 0x64, 0x62, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x62, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x06, 0x64, 0x62, 0x54, 0x79, 0x70, 0x65, 0x12, 0x37, 0x0a, 0x0e, 0x64, 0x62, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x44, 0x62, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x0c, 0x64, 0x62, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x62, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x62, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x27, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e,
	0x6e, 0x65, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x74, 0x79, 0x70, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x6e, 0x65, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0a, 0x74, 0x79, 0x70, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x55, 0x75, 0x69,
	0x64, 0x12, 0x26, 0x0a, 0x0f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x5f,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x79, 0x55, 0x75, 0x69, 0x64, 0x2a, 0x71, 0x0a, 0x10, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x62, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a,
	0x1a, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x42, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a,
	0x18, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x42, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x4d, 0x59, 0x53, 0x51, 0x4c, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x43,
	0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x42, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x50, 0x4f, 0x53, 0x54, 0x47, 0x52, 0x45, 0x53, 0x10, 0x02, 0x2a, 0x6f, 0x0a, 0x10,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1d, 0x0a, 0x19, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12,
	0x1c, 0x0a, 0x18, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x1e, 0x0a,
	0x1a, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x2a, 0x69, 0x0a,
	0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1b, 0x0a, 0x17, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16,
	0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x54, 0x43, 0x50, 0x5f, 0x49, 0x50, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x4f, 0x4e, 0x4e,
	0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x43, 0x50, 0x5f,
	0x49, 0x50, 0x5f, 0x53, 0x53, 0x48, 0x10, 0x02, 0x42, 0x20, 0x0a, 0x03, 0x6e, 0x65, 0x6d, 0x42,
	0x0a, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x01, 0x5a, 0x0b, 0x6e,
	0x65, 0x6d, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_connection_proto_rawDescOnce sync.Once
	file_connection_proto_rawDescData = file_connection_proto_rawDesc
)

func file_connection_proto_rawDescGZIP() []byte {
	file_connection_proto_rawDescOnce.Do(func() {
		file_connection_proto_rawDescData = protoimpl.X.CompressGZIP(file_connection_proto_rawDescData)
	})
	return file_connection_proto_rawDescData
}

var file_connection_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_connection_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_connection_proto_goTypes = []any{
	(ConnectionDbType)(0),         // 0: nem.ConnectionDbType
	(ConnectionStatus)(0),         // 1: nem.ConnectionStatus
	(ConnectionType)(0),           // 2: nem.ConnectionType
	(*Connection)(nil),            // 3: nem.Connection
	(*DbTypeConfig)(nil),          // 4: nem.DbTypeConfig
	(*ConnectionTypeConfig)(nil),  // 5: nem.ConnectionTypeConfig
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_connection_proto_depIdxs = []int32{
	0, // 0: nem.Connection.db_type:type_name -> nem.ConnectionDbType
	4, // 1: nem.Connection.db_type_config:type_name -> nem.DbTypeConfig
	2, // 2: nem.Connection.type:type_name -> nem.ConnectionType
	5, // 3: nem.Connection.type_config:type_name -> nem.ConnectionTypeConfig
	1, // 4: nem.Connection.status:type_name -> nem.ConnectionStatus
	6, // 5: nem.Connection.created_at:type_name -> google.protobuf.Timestamp
	6, // 6: nem.Connection.updated_at:type_name -> google.protobuf.Timestamp
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_connection_proto_init() }
func file_connection_proto_init() {
	if File_connection_proto != nil {
		return
	}
	file_connection_type_config_proto_init()
	file_db_type_config_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_connection_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_connection_proto_goTypes,
		DependencyIndexes: file_connection_proto_depIdxs,
		EnumInfos:         file_connection_proto_enumTypes,
		MessageInfos:      file_connection_proto_msgTypes,
	}.Build()
	File_connection_proto = out.File
	file_connection_proto_rawDesc = nil
	file_connection_proto_goTypes = nil
	file_connection_proto_depIdxs = nil
}
