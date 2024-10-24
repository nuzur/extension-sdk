// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.24.4
// source: tcp_ip_connection_type_config.proto

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

type TcpIpConnectionTypeConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Port     string `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *TcpIpConnectionTypeConfig) Reset() {
	*x = TcpIpConnectionTypeConfig{}
	mi := &file_tcp_ip_connection_type_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TcpIpConnectionTypeConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TcpIpConnectionTypeConfig) ProtoMessage() {}

func (x *TcpIpConnectionTypeConfig) ProtoReflect() protoreflect.Message {
	mi := &file_tcp_ip_connection_type_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TcpIpConnectionTypeConfig.ProtoReflect.Descriptor instead.
func (*TcpIpConnectionTypeConfig) Descriptor() ([]byte, []int) {
	return file_tcp_ip_connection_type_config_proto_rawDescGZIP(), []int{0}
}

func (x *TcpIpConnectionTypeConfig) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *TcpIpConnectionTypeConfig) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *TcpIpConnectionTypeConfig) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *TcpIpConnectionTypeConfig) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_tcp_ip_connection_type_config_proto protoreflect.FileDescriptor

var file_tcp_ip_connection_type_config_proto_rawDesc = []byte{
	0x0a, 0x23, 0x74, 0x63, 0x70, 0x5f, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6e, 0x65, 0x6d, 0x22, 0x83, 0x01, 0x0a, 0x19, 0x54,
	0x63, 0x70, 0x49, 0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x42, 0x2f, 0x0a, 0x03, 0x6e, 0x65, 0x6d, 0x42, 0x19, 0x54, 0x63, 0x70, 0x49, 0x70, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x50, 0x01, 0x5a, 0x0b, 0x6e, 0x65, 0x6d, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x67, 0x65,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tcp_ip_connection_type_config_proto_rawDescOnce sync.Once
	file_tcp_ip_connection_type_config_proto_rawDescData = file_tcp_ip_connection_type_config_proto_rawDesc
)

func file_tcp_ip_connection_type_config_proto_rawDescGZIP() []byte {
	file_tcp_ip_connection_type_config_proto_rawDescOnce.Do(func() {
		file_tcp_ip_connection_type_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_tcp_ip_connection_type_config_proto_rawDescData)
	})
	return file_tcp_ip_connection_type_config_proto_rawDescData
}

var file_tcp_ip_connection_type_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tcp_ip_connection_type_config_proto_goTypes = []any{
	(*TcpIpConnectionTypeConfig)(nil), // 0: nem.TcpIpConnectionTypeConfig
}
var file_tcp_ip_connection_type_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tcp_ip_connection_type_config_proto_init() }
func file_tcp_ip_connection_type_config_proto_init() {
	if File_tcp_ip_connection_type_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tcp_ip_connection_type_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tcp_ip_connection_type_config_proto_goTypes,
		DependencyIndexes: file_tcp_ip_connection_type_config_proto_depIdxs,
		MessageInfos:      file_tcp_ip_connection_type_config_proto_msgTypes,
	}.Build()
	File_tcp_ip_connection_type_config_proto = out.File
	file_tcp_ip_connection_type_config_proto_rawDesc = nil
	file_tcp_ip_connection_type_config_proto_goTypes = nil
	file_tcp_ip_connection_type_config_proto_depIdxs = nil
}
