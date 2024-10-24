// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.24.4
// source: entity.proto

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

type EntityStatus int32

const (
	EntityStatus_ENTITY_STATUS_INVALID  EntityStatus = 0
	EntityStatus_ENTITY_STATUS_ACTIVE   EntityStatus = 1
	EntityStatus_ENTITY_STATUS_DISABLED EntityStatus = 2
)

// Enum value maps for EntityStatus.
var (
	EntityStatus_name = map[int32]string{
		0: "ENTITY_STATUS_INVALID",
		1: "ENTITY_STATUS_ACTIVE",
		2: "ENTITY_STATUS_DISABLED",
	}
	EntityStatus_value = map[string]int32{
		"ENTITY_STATUS_INVALID":  0,
		"ENTITY_STATUS_ACTIVE":   1,
		"ENTITY_STATUS_DISABLED": 2,
	}
)

func (x EntityStatus) Enum() *EntityStatus {
	p := new(EntityStatus)
	*p = x
	return p
}

func (x EntityStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EntityStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_entity_proto_enumTypes[0].Descriptor()
}

func (EntityStatus) Type() protoreflect.EnumType {
	return &file_entity_proto_enumTypes[0]
}

func (x EntityStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EntityStatus.Descriptor instead.
func (EntityStatus) EnumDescriptor() ([]byte, []int) {
	return file_entity_proto_rawDescGZIP(), []int{0}
}

type EntityType int32

const (
	EntityType_ENTITY_TYPE_INVALID    EntityType = 0
	EntityType_ENTITY_TYPE_STANDALONE EntityType = 1
	EntityType_ENTITY_TYPE_DEPENDENT  EntityType = 2
)

// Enum value maps for EntityType.
var (
	EntityType_name = map[int32]string{
		0: "ENTITY_TYPE_INVALID",
		1: "ENTITY_TYPE_STANDALONE",
		2: "ENTITY_TYPE_DEPENDENT",
	}
	EntityType_value = map[string]int32{
		"ENTITY_TYPE_INVALID":    0,
		"ENTITY_TYPE_STANDALONE": 1,
		"ENTITY_TYPE_DEPENDENT":  2,
	}
)

func (x EntityType) Enum() *EntityType {
	p := new(EntityType)
	*p = x
	return p
}

func (x EntityType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EntityType) Descriptor() protoreflect.EnumDescriptor {
	return file_entity_proto_enumTypes[1].Descriptor()
}

func (EntityType) Type() protoreflect.EnumType {
	return &file_entity_proto_enumTypes[1]
}

func (x EntityType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EntityType.Descriptor instead.
func (EntityType) EnumDescriptor() ([]byte, []int) {
	return file_entity_proto_rawDescGZIP(), []int{1}
}

type Entity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid          string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Version       int64                  `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	Identifier    string                 `protobuf:"bytes,3,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Description   string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Fields        []*Field               `protobuf:"bytes,5,rep,name=fields,proto3" json:"fields,omitempty"`
	Type          EntityType             `protobuf:"varint,6,opt,name=type,proto3,enum=nem.EntityType" json:"type,omitempty"`
	TypeConfig    *EntityTypeConfig      `protobuf:"bytes,7,opt,name=type_config,json=typeConfig,proto3" json:"type_config,omitempty"`
	Render        *ElementRender         `protobuf:"bytes,8,opt,name=render,proto3" json:"render,omitempty"`
	Status        EntityStatus           `protobuf:"varint,9,opt,name=status,proto3,enum=nem.EntityStatus" json:"status,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CreatedByUuid string                 `protobuf:"bytes,12,opt,name=created_by_uuid,json=createdByUuid,proto3" json:"created_by_uuid,omitempty"`
	UpdatedByUuid string                 `protobuf:"bytes,13,opt,name=updated_by_uuid,json=updatedByUuid,proto3" json:"updated_by_uuid,omitempty"`
}

func (x *Entity) Reset() {
	*x = Entity{}
	mi := &file_entity_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Entity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entity) ProtoMessage() {}

func (x *Entity) ProtoReflect() protoreflect.Message {
	mi := &file_entity_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entity.ProtoReflect.Descriptor instead.
func (*Entity) Descriptor() ([]byte, []int) {
	return file_entity_proto_rawDescGZIP(), []int{0}
}

func (x *Entity) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Entity) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Entity) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *Entity) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Entity) GetFields() []*Field {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *Entity) GetType() EntityType {
	if x != nil {
		return x.Type
	}
	return EntityType_ENTITY_TYPE_INVALID
}

func (x *Entity) GetTypeConfig() *EntityTypeConfig {
	if x != nil {
		return x.TypeConfig
	}
	return nil
}

func (x *Entity) GetRender() *ElementRender {
	if x != nil {
		return x.Render
	}
	return nil
}

func (x *Entity) GetStatus() EntityStatus {
	if x != nil {
		return x.Status
	}
	return EntityStatus_ENTITY_STATUS_INVALID
}

func (x *Entity) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Entity) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Entity) GetCreatedByUuid() string {
	if x != nil {
		return x.CreatedByUuid
	}
	return ""
}

func (x *Entity) GetUpdatedByUuid() string {
	if x != nil {
		return x.UpdatedByUuid
	}
	return ""
}

var File_entity_proto protoreflect.FileDescriptor

var file_entity_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x6e, 0x65, 0x6d, 0x1a, 0x14, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x96, 0x04, 0x0a, 0x06, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6e,
	0x65, 0x6d, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x12, 0x23, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f,
	0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x36, 0x0a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6e, 0x65, 0x6d,
	0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x0a, 0x74, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2a, 0x0a,
	0x06, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x6e, 0x65, 0x6d, 0x2e, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x52, 0x06, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x6e, 0x65, 0x6d, 0x2e,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x55, 0x75,
	0x69, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79,
	0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x55, 0x75, 0x69, 0x64, 0x2a, 0x5f, 0x0a, 0x0c, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x15, 0x45, 0x4e,
	0x54, 0x49, 0x54, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12,
	0x1a, 0x0a, 0x16, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x2a, 0x5c, 0x0a, 0x0a, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x4e, 0x54,
	0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x53, 0x54, 0x41, 0x4e, 0x44, 0x41, 0x4c, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x19,
	0x0a, 0x15, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45,
	0x50, 0x45, 0x4e, 0x44, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x42, 0x1c, 0x0a, 0x03, 0x6e, 0x65, 0x6d,
	0x42, 0x06, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x01, 0x5a, 0x0b, 0x6e, 0x65, 0x6d, 0x2f,
	0x69, 0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entity_proto_rawDescOnce sync.Once
	file_entity_proto_rawDescData = file_entity_proto_rawDesc
)

func file_entity_proto_rawDescGZIP() []byte {
	file_entity_proto_rawDescOnce.Do(func() {
		file_entity_proto_rawDescData = protoimpl.X.CompressGZIP(file_entity_proto_rawDescData)
	})
	return file_entity_proto_rawDescData
}

var file_entity_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_entity_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_entity_proto_goTypes = []any{
	(EntityStatus)(0),             // 0: nem.EntityStatus
	(EntityType)(0),               // 1: nem.EntityType
	(*Entity)(nil),                // 2: nem.Entity
	(*Field)(nil),                 // 3: nem.Field
	(*EntityTypeConfig)(nil),      // 4: nem.EntityTypeConfig
	(*ElementRender)(nil),         // 5: nem.ElementRender
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_entity_proto_depIdxs = []int32{
	3, // 0: nem.Entity.fields:type_name -> nem.Field
	1, // 1: nem.Entity.type:type_name -> nem.EntityType
	4, // 2: nem.Entity.type_config:type_name -> nem.EntityTypeConfig
	5, // 3: nem.Entity.render:type_name -> nem.ElementRender
	0, // 4: nem.Entity.status:type_name -> nem.EntityStatus
	6, // 5: nem.Entity.created_at:type_name -> google.protobuf.Timestamp
	6, // 6: nem.Entity.updated_at:type_name -> google.protobuf.Timestamp
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_entity_proto_init() }
func file_entity_proto_init() {
	if File_entity_proto != nil {
		return
	}
	file_element_render_proto_init()
	file_entity_type_config_proto_init()
	file_field_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_entity_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entity_proto_goTypes,
		DependencyIndexes: file_entity_proto_depIdxs,
		EnumInfos:         file_entity_proto_enumTypes,
		MessageInfos:      file_entity_proto_msgTypes,
	}.Build()
	File_entity_proto = out.File
	file_entity_proto_rawDesc = nil
	file_entity_proto_goTypes = nil
	file_entity_proto_depIdxs = nil
}
