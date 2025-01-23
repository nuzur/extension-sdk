// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.28.3
// source: user_connection_local_config.proto

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

type UserConnectionLocalConfigDbType int32

const (
	UserConnectionLocalConfigDbType_USER_CONNECTION_LOCAL_CONFIG_DB_TYPE_INVALID  UserConnectionLocalConfigDbType = 0
	UserConnectionLocalConfigDbType_USER_CONNECTION_LOCAL_CONFIG_DB_TYPE_MYSQL    UserConnectionLocalConfigDbType = 1
	UserConnectionLocalConfigDbType_USER_CONNECTION_LOCAL_CONFIG_DB_TYPE_POSTGRES UserConnectionLocalConfigDbType = 2
)

// Enum value maps for UserConnectionLocalConfigDbType.
var (
	UserConnectionLocalConfigDbType_name = map[int32]string{
		0: "USER_CONNECTION_LOCAL_CONFIG_DB_TYPE_INVALID",
		1: "USER_CONNECTION_LOCAL_CONFIG_DB_TYPE_MYSQL",
		2: "USER_CONNECTION_LOCAL_CONFIG_DB_TYPE_POSTGRES",
	}
	UserConnectionLocalConfigDbType_value = map[string]int32{
		"USER_CONNECTION_LOCAL_CONFIG_DB_TYPE_INVALID":  0,
		"USER_CONNECTION_LOCAL_CONFIG_DB_TYPE_MYSQL":    1,
		"USER_CONNECTION_LOCAL_CONFIG_DB_TYPE_POSTGRES": 2,
	}
)

func (x UserConnectionLocalConfigDbType) Enum() *UserConnectionLocalConfigDbType {
	p := new(UserConnectionLocalConfigDbType)
	*p = x
	return p
}

func (x UserConnectionLocalConfigDbType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserConnectionLocalConfigDbType) Descriptor() protoreflect.EnumDescriptor {
	return file_user_connection_local_config_proto_enumTypes[0].Descriptor()
}

func (UserConnectionLocalConfigDbType) Type() protoreflect.EnumType {
	return &file_user_connection_local_config_proto_enumTypes[0]
}

func (x UserConnectionLocalConfigDbType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserConnectionLocalConfigDbType.Descriptor instead.
func (UserConnectionLocalConfigDbType) EnumDescriptor() ([]byte, []int) {
	return file_user_connection_local_config_proto_rawDescGZIP(), []int{0}
}

type UserConnectionLocalConfig struct {
	state         protoimpl.MessageState          `protogen:"open.v1"`
	IpAddress     string                          `protobuf:"bytes,1,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	DbType        UserConnectionLocalConfigDbType `protobuf:"varint,2,opt,name=db_type,json=dbType,proto3,enum=nem.UserConnectionLocalConfigDbType" json:"db_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserConnectionLocalConfig) Reset() {
	*x = UserConnectionLocalConfig{}
	mi := &file_user_connection_local_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserConnectionLocalConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserConnectionLocalConfig) ProtoMessage() {}

func (x *UserConnectionLocalConfig) ProtoReflect() protoreflect.Message {
	mi := &file_user_connection_local_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserConnectionLocalConfig.ProtoReflect.Descriptor instead.
func (*UserConnectionLocalConfig) Descriptor() ([]byte, []int) {
	return file_user_connection_local_config_proto_rawDescGZIP(), []int{0}
}

func (x *UserConnectionLocalConfig) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

func (x *UserConnectionLocalConfig) GetDbType() UserConnectionLocalConfigDbType {
	if x != nil {
		return x.DbType
	}
	return UserConnectionLocalConfigDbType_USER_CONNECTION_LOCAL_CONFIG_DB_TYPE_INVALID
}

var File_user_connection_local_config_proto protoreflect.FileDescriptor

var file_user_connection_local_config_proto_rawDesc = []byte{
	0x0a, 0x22, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6e, 0x65, 0x6d, 0x22, 0x79, 0x0a, 0x19, 0x55, 0x73, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x63, 0x61, 0x6c,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x70, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x3d, 0x0a, 0x07, 0x64, 0x62, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x63, 0x61, 0x6c,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x62, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x64, 0x62,
	0x54, 0x79, 0x70, 0x65, 0x2a, 0xb6, 0x01, 0x0a, 0x1f, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x44, 0x62, 0x54, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x2c, 0x55, 0x53, 0x45, 0x52,
	0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x4f, 0x43, 0x41,
	0x4c, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x44, 0x42, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x2e, 0x0a, 0x2a, 0x55, 0x53,
	0x45, 0x52, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x4f,
	0x43, 0x41, 0x4c, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x44, 0x42, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4d, 0x59, 0x53, 0x51, 0x4c, 0x10, 0x01, 0x12, 0x31, 0x0a, 0x2d, 0x55, 0x53,
	0x45, 0x52, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x4f,
	0x43, 0x41, 0x4c, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x44, 0x42, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x50, 0x4f, 0x53, 0x54, 0x47, 0x52, 0x45, 0x53, 0x10, 0x02, 0x42, 0x2f, 0x0a,
	0x03, 0x6e, 0x65, 0x6d, 0x42, 0x19, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50,
	0x01, 0x5a, 0x0b, 0x6e, 0x65, 0x6d, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_connection_local_config_proto_rawDescOnce sync.Once
	file_user_connection_local_config_proto_rawDescData = file_user_connection_local_config_proto_rawDesc
)

func file_user_connection_local_config_proto_rawDescGZIP() []byte {
	file_user_connection_local_config_proto_rawDescOnce.Do(func() {
		file_user_connection_local_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_connection_local_config_proto_rawDescData)
	})
	return file_user_connection_local_config_proto_rawDescData
}

var file_user_connection_local_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_connection_local_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_user_connection_local_config_proto_goTypes = []any{
	(UserConnectionLocalConfigDbType)(0), // 0: nem.UserConnectionLocalConfigDbType
	(*UserConnectionLocalConfig)(nil),    // 1: nem.UserConnectionLocalConfig
}
var file_user_connection_local_config_proto_depIdxs = []int32{
	0, // 0: nem.UserConnectionLocalConfig.db_type:type_name -> nem.UserConnectionLocalConfigDbType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_connection_local_config_proto_init() }
func file_user_connection_local_config_proto_init() {
	if File_user_connection_local_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_connection_local_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_connection_local_config_proto_goTypes,
		DependencyIndexes: file_user_connection_local_config_proto_depIdxs,
		EnumInfos:         file_user_connection_local_config_proto_enumTypes,
		MessageInfos:      file_user_connection_local_config_proto_msgTypes,
	}.Build()
	File_user_connection_local_config_proto = out.File
	file_user_connection_local_config_proto_rawDesc = nil
	file_user_connection_local_config_proto_goTypes = nil
	file_user_connection_local_config_proto_depIdxs = nil
}
