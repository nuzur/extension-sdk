// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.24.4
// source: object_store.proto

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

type ObjectStoreStatus int32

const (
	ObjectStoreStatus_OBJECT_STORE_STATUS_INVALID  ObjectStoreStatus = 0
	ObjectStoreStatus_OBJECT_STORE_STATUS_ACTIVE   ObjectStoreStatus = 1
	ObjectStoreStatus_OBJECT_STORE_STATUS_DISABLED ObjectStoreStatus = 2
)

// Enum value maps for ObjectStoreStatus.
var (
	ObjectStoreStatus_name = map[int32]string{
		0: "OBJECT_STORE_STATUS_INVALID",
		1: "OBJECT_STORE_STATUS_ACTIVE",
		2: "OBJECT_STORE_STATUS_DISABLED",
	}
	ObjectStoreStatus_value = map[string]int32{
		"OBJECT_STORE_STATUS_INVALID":  0,
		"OBJECT_STORE_STATUS_ACTIVE":   1,
		"OBJECT_STORE_STATUS_DISABLED": 2,
	}
)

func (x ObjectStoreStatus) Enum() *ObjectStoreStatus {
	p := new(ObjectStoreStatus)
	*p = x
	return p
}

func (x ObjectStoreStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ObjectStoreStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_object_store_proto_enumTypes[0].Descriptor()
}

func (ObjectStoreStatus) Type() protoreflect.EnumType {
	return &file_object_store_proto_enumTypes[0]
}

func (x ObjectStoreStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ObjectStoreStatus.Descriptor instead.
func (ObjectStoreStatus) EnumDescriptor() ([]byte, []int) {
	return file_object_store_proto_rawDescGZIP(), []int{0}
}

type ObjectStoreType int32

const (
	ObjectStoreType_OBJECT_STORE_TYPE_INVALID ObjectStoreType = 0
	ObjectStoreType_OBJECT_STORE_TYPE_S_3     ObjectStoreType = 1
)

// Enum value maps for ObjectStoreType.
var (
	ObjectStoreType_name = map[int32]string{
		0: "OBJECT_STORE_TYPE_INVALID",
		1: "OBJECT_STORE_TYPE_S_3",
	}
	ObjectStoreType_value = map[string]int32{
		"OBJECT_STORE_TYPE_INVALID": 0,
		"OBJECT_STORE_TYPE_S_3":     1,
	}
)

func (x ObjectStoreType) Enum() *ObjectStoreType {
	p := new(ObjectStoreType)
	*p = x
	return p
}

func (x ObjectStoreType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ObjectStoreType) Descriptor() protoreflect.EnumDescriptor {
	return file_object_store_proto_enumTypes[1].Descriptor()
}

func (ObjectStoreType) Type() protoreflect.EnumType {
	return &file_object_store_proto_enumTypes[1]
}

func (x ObjectStoreType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ObjectStoreType.Descriptor instead.
func (ObjectStoreType) EnumDescriptor() ([]byte, []int) {
	return file_object_store_proto_rawDescGZIP(), []int{1}
}

type ObjectStore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid          string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Identifier    string                 `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Type          ObjectStoreType        `protobuf:"varint,3,opt,name=type,proto3,enum=nem.ObjectStoreType" json:"type,omitempty"`
	TypeConfig    *ObjectStoreTypeConfig `protobuf:"bytes,4,opt,name=type_config,json=typeConfig,proto3" json:"type_config,omitempty"`
	Status        ObjectStoreStatus      `protobuf:"varint,5,opt,name=status,proto3,enum=nem.ObjectStoreStatus" json:"status,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CreatedByUuid string                 `protobuf:"bytes,8,opt,name=created_by_uuid,json=createdByUuid,proto3" json:"created_by_uuid,omitempty"`
	UpdatedByUuid string                 `protobuf:"bytes,9,opt,name=updated_by_uuid,json=updatedByUuid,proto3" json:"updated_by_uuid,omitempty"`
}

func (x *ObjectStore) Reset() {
	*x = ObjectStore{}
	mi := &file_object_store_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ObjectStore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectStore) ProtoMessage() {}

func (x *ObjectStore) ProtoReflect() protoreflect.Message {
	mi := &file_object_store_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectStore.ProtoReflect.Descriptor instead.
func (*ObjectStore) Descriptor() ([]byte, []int) {
	return file_object_store_proto_rawDescGZIP(), []int{0}
}

func (x *ObjectStore) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *ObjectStore) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *ObjectStore) GetType() ObjectStoreType {
	if x != nil {
		return x.Type
	}
	return ObjectStoreType_OBJECT_STORE_TYPE_INVALID
}

func (x *ObjectStore) GetTypeConfig() *ObjectStoreTypeConfig {
	if x != nil {
		return x.TypeConfig
	}
	return nil
}

func (x *ObjectStore) GetStatus() ObjectStoreStatus {
	if x != nil {
		return x.Status
	}
	return ObjectStoreStatus_OBJECT_STORE_STATUS_INVALID
}

func (x *ObjectStore) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ObjectStore) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *ObjectStore) GetCreatedByUuid() string {
	if x != nil {
		return x.CreatedByUuid
	}
	return ""
}

func (x *ObjectStore) GetUpdatedByUuid() string {
	if x != nil {
		return x.UpdatedByUuid
	}
	return ""
}

var File_object_store_proto protoreflect.FileDescriptor

var file_object_store_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6e, 0x65, 0x6d, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x03, 0x0a, 0x0b, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x28,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x6e,
	0x65, 0x6d, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x74, 0x79, 0x70, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x6e, 0x65, 0x6d, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0a, 0x74, 0x79, 0x70, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x55,
	0x75, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62,
	0x79, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x55, 0x75, 0x69, 0x64, 0x2a, 0x76, 0x0a, 0x11, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1f, 0x0a, 0x1b, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x45,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10,
	0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x53, 0x54, 0x4f, 0x52,
	0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10,
	0x01, 0x12, 0x20, 0x0a, 0x1c, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x53, 0x54, 0x4f, 0x52,
	0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45,
	0x44, 0x10, 0x02, 0x2a, 0x4b, 0x0a, 0x0f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54,
	0x5f, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f,
	0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x5f, 0x33, 0x10, 0x01,
	0x42, 0x21, 0x0a, 0x03, 0x6e, 0x65, 0x6d, 0x42, 0x0b, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x50, 0x01, 0x5a, 0x0b, 0x6e, 0x65, 0x6d, 0x2f, 0x69, 0x64, 0x6c, 0x2f,
	0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_object_store_proto_rawDescOnce sync.Once
	file_object_store_proto_rawDescData = file_object_store_proto_rawDesc
)

func file_object_store_proto_rawDescGZIP() []byte {
	file_object_store_proto_rawDescOnce.Do(func() {
		file_object_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_object_store_proto_rawDescData)
	})
	return file_object_store_proto_rawDescData
}

var file_object_store_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_object_store_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_object_store_proto_goTypes = []any{
	(ObjectStoreStatus)(0),        // 0: nem.ObjectStoreStatus
	(ObjectStoreType)(0),          // 1: nem.ObjectStoreType
	(*ObjectStore)(nil),           // 2: nem.ObjectStore
	(*ObjectStoreTypeConfig)(nil), // 3: nem.ObjectStoreTypeConfig
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_object_store_proto_depIdxs = []int32{
	1, // 0: nem.ObjectStore.type:type_name -> nem.ObjectStoreType
	3, // 1: nem.ObjectStore.type_config:type_name -> nem.ObjectStoreTypeConfig
	0, // 2: nem.ObjectStore.status:type_name -> nem.ObjectStoreStatus
	4, // 3: nem.ObjectStore.created_at:type_name -> google.protobuf.Timestamp
	4, // 4: nem.ObjectStore.updated_at:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_object_store_proto_init() }
func file_object_store_proto_init() {
	if File_object_store_proto != nil {
		return
	}
	file_object_store_type_config_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_object_store_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_object_store_proto_goTypes,
		DependencyIndexes: file_object_store_proto_depIdxs,
		EnumInfos:         file_object_store_proto_enumTypes,
		MessageInfos:      file_object_store_proto_msgTypes,
	}.Build()
	File_object_store_proto = out.File
	file_object_store_proto_rawDesc = nil
	file_object_store_proto_goTypes = nil
	file_object_store_proto_depIdxs = nil
}