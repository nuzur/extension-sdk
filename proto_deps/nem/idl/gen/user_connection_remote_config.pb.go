// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: user_connection_remote_config.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserConnectionRemoteConfig struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	TeamUuid       string                 `protobuf:"bytes,1,opt,name=team_uuid,json=teamUuid,proto3" json:"team_uuid,omitempty"`
	StoreUuid      string                 `protobuf:"bytes,2,opt,name=store_uuid,json=storeUuid,proto3" json:"store_uuid,omitempty"`
	ConnectionUuid string                 `protobuf:"bytes,3,opt,name=connection_uuid,json=connectionUuid,proto3" json:"connection_uuid,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *UserConnectionRemoteConfig) Reset() {
	*x = UserConnectionRemoteConfig{}
	mi := &file_user_connection_remote_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserConnectionRemoteConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserConnectionRemoteConfig) ProtoMessage() {}

func (x *UserConnectionRemoteConfig) ProtoReflect() protoreflect.Message {
	mi := &file_user_connection_remote_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserConnectionRemoteConfig.ProtoReflect.Descriptor instead.
func (*UserConnectionRemoteConfig) Descriptor() ([]byte, []int) {
	return file_user_connection_remote_config_proto_rawDescGZIP(), []int{0}
}

func (x *UserConnectionRemoteConfig) GetTeamUuid() string {
	if x != nil {
		return x.TeamUuid
	}
	return ""
}

func (x *UserConnectionRemoteConfig) GetStoreUuid() string {
	if x != nil {
		return x.StoreUuid
	}
	return ""
}

func (x *UserConnectionRemoteConfig) GetConnectionUuid() string {
	if x != nil {
		return x.ConnectionUuid
	}
	return ""
}

var File_user_connection_remote_config_proto protoreflect.FileDescriptor

var file_user_connection_remote_config_proto_rawDesc = string([]byte{
	0x0a, 0x23, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6e, 0x65, 0x6d, 0x22, 0x81, 0x01, 0x0a, 0x1a, 0x55,
	0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x61,
	0x6d, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65,
	0x61, 0x6d, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x55, 0x75, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x75, 0x69, 0x64, 0x42, 0x30,
	0x0a, 0x03, 0x6e, 0x65, 0x6d, 0x42, 0x1a, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x50, 0x01, 0x5a, 0x0b, 0x6e, 0x65, 0x6d, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_user_connection_remote_config_proto_rawDescOnce sync.Once
	file_user_connection_remote_config_proto_rawDescData []byte
)

func file_user_connection_remote_config_proto_rawDescGZIP() []byte {
	file_user_connection_remote_config_proto_rawDescOnce.Do(func() {
		file_user_connection_remote_config_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_user_connection_remote_config_proto_rawDesc), len(file_user_connection_remote_config_proto_rawDesc)))
	})
	return file_user_connection_remote_config_proto_rawDescData
}

var file_user_connection_remote_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_user_connection_remote_config_proto_goTypes = []any{
	(*UserConnectionRemoteConfig)(nil), // 0: nem.UserConnectionRemoteConfig
}
var file_user_connection_remote_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_connection_remote_config_proto_init() }
func file_user_connection_remote_config_proto_init() {
	if File_user_connection_remote_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_user_connection_remote_config_proto_rawDesc), len(file_user_connection_remote_config_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_connection_remote_config_proto_goTypes,
		DependencyIndexes: file_user_connection_remote_config_proto_depIdxs,
		MessageInfos:      file_user_connection_remote_config_proto_msgTypes,
	}.Build()
	File_user_connection_remote_config_proto = out.File
	file_user_connection_remote_config_proto_goTypes = nil
	file_user_connection_remote_config_proto_depIdxs = nil
}
