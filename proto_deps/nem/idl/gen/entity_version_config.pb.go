// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.3
// source: entity_version_config.proto

package gen

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

type EntityVersionConfigGenerator int32

const (
	EntityVersionConfigGenerator_ENTITY_VERSION_CONFIG_GENERATOR_INVALID   EntityVersionConfigGenerator = 0
	EntityVersionConfigGenerator_ENTITY_VERSION_CONFIG_GENERATOR_TIMESTAMP EntityVersionConfigGenerator = 1
)

// Enum value maps for EntityVersionConfigGenerator.
var (
	EntityVersionConfigGenerator_name = map[int32]string{
		0: "ENTITY_VERSION_CONFIG_GENERATOR_INVALID",
		1: "ENTITY_VERSION_CONFIG_GENERATOR_TIMESTAMP",
	}
	EntityVersionConfigGenerator_value = map[string]int32{
		"ENTITY_VERSION_CONFIG_GENERATOR_INVALID":   0,
		"ENTITY_VERSION_CONFIG_GENERATOR_TIMESTAMP": 1,
	}
)

func (x EntityVersionConfigGenerator) Enum() *EntityVersionConfigGenerator {
	p := new(EntityVersionConfigGenerator)
	*p = x
	return p
}

func (x EntityVersionConfigGenerator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EntityVersionConfigGenerator) Descriptor() protoreflect.EnumDescriptor {
	return file_entity_version_config_proto_enumTypes[0].Descriptor()
}

func (EntityVersionConfigGenerator) Type() protoreflect.EnumType {
	return &file_entity_version_config_proto_enumTypes[0]
}

func (x EntityVersionConfigGenerator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EntityVersionConfigGenerator.Descriptor instead.
func (EntityVersionConfigGenerator) EnumDescriptor() ([]byte, []int) {
	return file_entity_version_config_proto_rawDescGZIP(), []int{0}
}

type EntityVersionConfigType int32

const (
	EntityVersionConfigType_ENTITY_VERSION_CONFIG_TYPE_INVALID EntityVersionConfigType = 0
	EntityVersionConfigType_ENTITY_VERSION_CONFIG_TYPE_FIELD   EntityVersionConfigType = 1
	EntityVersionConfigType_ENTITY_VERSION_CONFIG_TYPE_ENTITY  EntityVersionConfigType = 2
)

// Enum value maps for EntityVersionConfigType.
var (
	EntityVersionConfigType_name = map[int32]string{
		0: "ENTITY_VERSION_CONFIG_TYPE_INVALID",
		1: "ENTITY_VERSION_CONFIG_TYPE_FIELD",
		2: "ENTITY_VERSION_CONFIG_TYPE_ENTITY",
	}
	EntityVersionConfigType_value = map[string]int32{
		"ENTITY_VERSION_CONFIG_TYPE_INVALID": 0,
		"ENTITY_VERSION_CONFIG_TYPE_FIELD":   1,
		"ENTITY_VERSION_CONFIG_TYPE_ENTITY":  2,
	}
)

func (x EntityVersionConfigType) Enum() *EntityVersionConfigType {
	p := new(EntityVersionConfigType)
	*p = x
	return p
}

func (x EntityVersionConfigType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EntityVersionConfigType) Descriptor() protoreflect.EnumDescriptor {
	return file_entity_version_config_proto_enumTypes[1].Descriptor()
}

func (EntityVersionConfigType) Type() protoreflect.EnumType {
	return &file_entity_version_config_proto_enumTypes[1]
}

func (x EntityVersionConfigType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EntityVersionConfigType.Descriptor instead.
func (EntityVersionConfigType) EnumDescriptor() ([]byte, []int) {
	return file_entity_version_config_proto_rawDescGZIP(), []int{1}
}

type EntityVersionConfig struct {
	state         protoimpl.MessageState       `protogen:"open.v1"`
	Type          EntityVersionConfigType      `protobuf:"varint,1,opt,name=type,proto3,enum=nem.EntityVersionConfigType" json:"type,omitempty"`
	Generator     EntityVersionConfigGenerator `protobuf:"varint,2,opt,name=generator,proto3,enum=nem.EntityVersionConfigGenerator" json:"generator,omitempty"`
	Config        *EntityVersionTypeConfig     `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EntityVersionConfig) Reset() {
	*x = EntityVersionConfig{}
	mi := &file_entity_version_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EntityVersionConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityVersionConfig) ProtoMessage() {}

func (x *EntityVersionConfig) ProtoReflect() protoreflect.Message {
	mi := &file_entity_version_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityVersionConfig.ProtoReflect.Descriptor instead.
func (*EntityVersionConfig) Descriptor() ([]byte, []int) {
	return file_entity_version_config_proto_rawDescGZIP(), []int{0}
}

func (x *EntityVersionConfig) GetType() EntityVersionConfigType {
	if x != nil {
		return x.Type
	}
	return EntityVersionConfigType_ENTITY_VERSION_CONFIG_TYPE_INVALID
}

func (x *EntityVersionConfig) GetGenerator() EntityVersionConfigGenerator {
	if x != nil {
		return x.Generator
	}
	return EntityVersionConfigGenerator_ENTITY_VERSION_CONFIG_GENERATOR_INVALID
}

func (x *EntityVersionConfig) GetConfig() *EntityVersionTypeConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

var File_entity_version_config_proto protoreflect.FileDescriptor

var file_entity_version_config_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6e,
	0x65, 0x6d, 0x1a, 0x20, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x01, 0x0a, 0x13, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x30, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x6e, 0x65, 0x6d,
	0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3f,
	0x0a, 0x09, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x21, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x52, 0x09, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x34, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2a, 0x7a, 0x0a, 0x1c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x2b, 0x0a, 0x27, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f,
	0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x47,
	0x45, 0x4e, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x10, 0x00, 0x12, 0x2d, 0x0a, 0x29, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x56, 0x45, 0x52,
	0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x47, 0x45, 0x4e, 0x45,
	0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x10,
	0x01, 0x2a, 0x8e, 0x01, 0x0a, 0x17, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a,
	0x22, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f,
	0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x24, 0x0a, 0x20, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f,
	0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x10, 0x01, 0x12, 0x25, 0x0a, 0x21, 0x45,
	0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f,
	0x4e, 0x46, 0x49, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59,
	0x10, 0x02, 0x42, 0x29, 0x0a, 0x03, 0x6e, 0x65, 0x6d, 0x42, 0x13, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x01,
	0x5a, 0x0b, 0x6e, 0x65, 0x6d, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entity_version_config_proto_rawDescOnce sync.Once
	file_entity_version_config_proto_rawDescData = file_entity_version_config_proto_rawDesc
)

func file_entity_version_config_proto_rawDescGZIP() []byte {
	file_entity_version_config_proto_rawDescOnce.Do(func() {
		file_entity_version_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_entity_version_config_proto_rawDescData)
	})
	return file_entity_version_config_proto_rawDescData
}

var file_entity_version_config_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_entity_version_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_entity_version_config_proto_goTypes = []any{
	(EntityVersionConfigGenerator)(0), // 0: nem.EntityVersionConfigGenerator
	(EntityVersionConfigType)(0),      // 1: nem.EntityVersionConfigType
	(*EntityVersionConfig)(nil),       // 2: nem.EntityVersionConfig
	(*EntityVersionTypeConfig)(nil),   // 3: nem.EntityVersionTypeConfig
}
var file_entity_version_config_proto_depIdxs = []int32{
	1, // 0: nem.EntityVersionConfig.type:type_name -> nem.EntityVersionConfigType
	0, // 1: nem.EntityVersionConfig.generator:type_name -> nem.EntityVersionConfigGenerator
	3, // 2: nem.EntityVersionConfig.config:type_name -> nem.EntityVersionTypeConfig
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_entity_version_config_proto_init() }
func file_entity_version_config_proto_init() {
	if File_entity_version_config_proto != nil {
		return
	}
	file_entity_version_type_config_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_entity_version_config_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entity_version_config_proto_goTypes,
		DependencyIndexes: file_entity_version_config_proto_depIdxs,
		EnumInfos:         file_entity_version_config_proto_enumTypes,
		MessageInfos:      file_entity_version_config_proto_msgTypes,
	}.Build()
	File_entity_version_config_proto = out.File
	file_entity_version_config_proto_rawDesc = nil
	file_entity_version_config_proto_goTypes = nil
	file_entity_version_config_proto_depIdxs = nil
}
