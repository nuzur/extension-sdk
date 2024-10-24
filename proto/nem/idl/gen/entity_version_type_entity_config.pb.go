// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.24.4
// source: entity_version_type_entity_config.proto

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

type EntityVersionTypeEntityConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityUuid       string `protobuf:"bytes,1,opt,name=entity_uuid,json=entityUuid,proto3" json:"entity_uuid,omitempty"`
	VersionFieldUuid string `protobuf:"bytes,2,opt,name=version_field_uuid,json=versionFieldUuid,proto3" json:"version_field_uuid,omitempty"`
	AppendOnly       bool   `protobuf:"varint,3,opt,name=append_only,json=appendOnly,proto3" json:"append_only,omitempty"`
}

func (x *EntityVersionTypeEntityConfig) Reset() {
	*x = EntityVersionTypeEntityConfig{}
	mi := &file_entity_version_type_entity_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EntityVersionTypeEntityConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityVersionTypeEntityConfig) ProtoMessage() {}

func (x *EntityVersionTypeEntityConfig) ProtoReflect() protoreflect.Message {
	mi := &file_entity_version_type_entity_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityVersionTypeEntityConfig.ProtoReflect.Descriptor instead.
func (*EntityVersionTypeEntityConfig) Descriptor() ([]byte, []int) {
	return file_entity_version_type_entity_config_proto_rawDescGZIP(), []int{0}
}

func (x *EntityVersionTypeEntityConfig) GetEntityUuid() string {
	if x != nil {
		return x.EntityUuid
	}
	return ""
}

func (x *EntityVersionTypeEntityConfig) GetVersionFieldUuid() string {
	if x != nil {
		return x.VersionFieldUuid
	}
	return ""
}

func (x *EntityVersionTypeEntityConfig) GetAppendOnly() bool {
	if x != nil {
		return x.AppendOnly
	}
	return false
}

var File_entity_version_type_entity_config_proto protoreflect.FileDescriptor

var file_entity_version_type_entity_config_proto_rawDesc = []byte{
	0x0a, 0x27, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6e, 0x65, 0x6d, 0x22, 0x8f,
	0x01, 0x0a, 0x1d, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x55, 0x75, 0x69,
	0x64, 0x12, 0x2c, 0x0a, 0x12, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x55, 0x75, 0x69, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x4f, 0x6e, 0x6c, 0x79,
	0x42, 0x33, 0x0a, 0x03, 0x6e, 0x65, 0x6d, 0x42, 0x1d, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x01, 0x5a, 0x0b, 0x6e, 0x65, 0x6d, 0x2f, 0x69, 0x64,
	0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entity_version_type_entity_config_proto_rawDescOnce sync.Once
	file_entity_version_type_entity_config_proto_rawDescData = file_entity_version_type_entity_config_proto_rawDesc
)

func file_entity_version_type_entity_config_proto_rawDescGZIP() []byte {
	file_entity_version_type_entity_config_proto_rawDescOnce.Do(func() {
		file_entity_version_type_entity_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_entity_version_type_entity_config_proto_rawDescData)
	})
	return file_entity_version_type_entity_config_proto_rawDescData
}

var file_entity_version_type_entity_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_entity_version_type_entity_config_proto_goTypes = []any{
	(*EntityVersionTypeEntityConfig)(nil), // 0: nem.EntityVersionTypeEntityConfig
}
var file_entity_version_type_entity_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_entity_version_type_entity_config_proto_init() }
func file_entity_version_type_entity_config_proto_init() {
	if File_entity_version_type_entity_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_entity_version_type_entity_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entity_version_type_entity_config_proto_goTypes,
		DependencyIndexes: file_entity_version_type_entity_config_proto_depIdxs,
		MessageInfos:      file_entity_version_type_entity_config_proto_msgTypes,
	}.Build()
	File_entity_version_type_entity_config_proto = out.File
	file_entity_version_type_entity_config_proto_rawDesc = nil
	file_entity_version_type_entity_config_proto_goTypes = nil
	file_entity_version_type_entity_config_proto_depIdxs = nil
}
