// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.3
// source: object_store_type_config.proto

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

type ObjectStoreTypeConfig struct {
	state         protoimpl.MessageState   `protogen:"open.v1"`
	S3            *ObjectStoreS3TypeConfig `protobuf:"bytes,1,opt,name=s3,proto3" json:"s3,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ObjectStoreTypeConfig) Reset() {
	*x = ObjectStoreTypeConfig{}
	mi := &file_object_store_type_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ObjectStoreTypeConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectStoreTypeConfig) ProtoMessage() {}

func (x *ObjectStoreTypeConfig) ProtoReflect() protoreflect.Message {
	mi := &file_object_store_type_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectStoreTypeConfig.ProtoReflect.Descriptor instead.
func (*ObjectStoreTypeConfig) Descriptor() ([]byte, []int) {
	return file_object_store_type_config_proto_rawDescGZIP(), []int{0}
}

func (x *ObjectStoreTypeConfig) GetS3() *ObjectStoreS3TypeConfig {
	if x != nil {
		return x.S3
	}
	return nil
}

var File_object_store_type_config_proto protoreflect.FileDescriptor

var file_object_store_type_config_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x6e, 0x65, 0x6d, 0x1a, 0x22, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x5f, 0x73, 0x5f, 0x33, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x15, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x2c, 0x0a, 0x02, 0x73, 0x33, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x53, 0x33, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x02, 0x73, 0x33,
	0x42, 0x2b, 0x0a, 0x03, 0x6e, 0x65, 0x6d, 0x42, 0x15, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x01,
	0x5a, 0x0b, 0x6e, 0x65, 0x6d, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_object_store_type_config_proto_rawDescOnce sync.Once
	file_object_store_type_config_proto_rawDescData = file_object_store_type_config_proto_rawDesc
)

func file_object_store_type_config_proto_rawDescGZIP() []byte {
	file_object_store_type_config_proto_rawDescOnce.Do(func() {
		file_object_store_type_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_object_store_type_config_proto_rawDescData)
	})
	return file_object_store_type_config_proto_rawDescData
}

var file_object_store_type_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_object_store_type_config_proto_goTypes = []any{
	(*ObjectStoreTypeConfig)(nil),   // 0: nem.ObjectStoreTypeConfig
	(*ObjectStoreS3TypeConfig)(nil), // 1: nem.ObjectStoreS3TypeConfig
}
var file_object_store_type_config_proto_depIdxs = []int32{
	1, // 0: nem.ObjectStoreTypeConfig.s3:type_name -> nem.ObjectStoreS3TypeConfig
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_object_store_type_config_proto_init() }
func file_object_store_type_config_proto_init() {
	if File_object_store_type_config_proto != nil {
		return
	}
	file_object_store_s_3_type_config_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_object_store_type_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_object_store_type_config_proto_goTypes,
		DependencyIndexes: file_object_store_type_config_proto_depIdxs,
		MessageInfos:      file_object_store_type_config_proto_msgTypes,
	}.Build()
	File_object_store_type_config_proto = out.File
	file_object_store_type_config_proto_rawDesc = nil
	file_object_store_type_config_proto_goTypes = nil
	file_object_store_type_config_proto_depIdxs = nil
}
