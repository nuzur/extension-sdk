// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.3
// source: relationship_node.proto

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

type RelationshipNodeType int32

const (
	RelationshipNodeType_RELATIONSHIP_NODE_TYPE_INVALID RelationshipNodeType = 0
	RelationshipNodeType_RELATIONSHIP_NODE_TYPE_ENTITY  RelationshipNodeType = 1
	RelationshipNodeType_RELATIONSHIP_NODE_TYPE_SERVICE RelationshipNodeType = 2
)

// Enum value maps for RelationshipNodeType.
var (
	RelationshipNodeType_name = map[int32]string{
		0: "RELATIONSHIP_NODE_TYPE_INVALID",
		1: "RELATIONSHIP_NODE_TYPE_ENTITY",
		2: "RELATIONSHIP_NODE_TYPE_SERVICE",
	}
	RelationshipNodeType_value = map[string]int32{
		"RELATIONSHIP_NODE_TYPE_INVALID": 0,
		"RELATIONSHIP_NODE_TYPE_ENTITY":  1,
		"RELATIONSHIP_NODE_TYPE_SERVICE": 2,
	}
)

func (x RelationshipNodeType) Enum() *RelationshipNodeType {
	p := new(RelationshipNodeType)
	*p = x
	return p
}

func (x RelationshipNodeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RelationshipNodeType) Descriptor() protoreflect.EnumDescriptor {
	return file_relationship_node_proto_enumTypes[0].Descriptor()
}

func (RelationshipNodeType) Type() protoreflect.EnumType {
	return &file_relationship_node_proto_enumTypes[0]
}

func (x RelationshipNodeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RelationshipNodeType.Descriptor instead.
func (RelationshipNodeType) EnumDescriptor() ([]byte, []int) {
	return file_relationship_node_proto_rawDescGZIP(), []int{0}
}

type RelationshipNode struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	Uuid          string                      `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Type          RelationshipNodeType        `protobuf:"varint,2,opt,name=type,proto3,enum=nem.RelationshipNodeType" json:"type,omitempty"`
	TypeConfig    *RelationshipNodeTypeConfig `protobuf:"bytes,3,opt,name=type_config,json=typeConfig,proto3" json:"type_config,omitempty"`
	CreatedAt     *timestamppb.Timestamp      `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp      `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CreatedByUuid string                      `protobuf:"bytes,6,opt,name=created_by_uuid,json=createdByUuid,proto3" json:"created_by_uuid,omitempty"`
	UpdatedByUuid string                      `protobuf:"bytes,7,opt,name=updated_by_uuid,json=updatedByUuid,proto3" json:"updated_by_uuid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RelationshipNode) Reset() {
	*x = RelationshipNode{}
	mi := &file_relationship_node_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RelationshipNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationshipNode) ProtoMessage() {}

func (x *RelationshipNode) ProtoReflect() protoreflect.Message {
	mi := &file_relationship_node_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationshipNode.ProtoReflect.Descriptor instead.
func (*RelationshipNode) Descriptor() ([]byte, []int) {
	return file_relationship_node_proto_rawDescGZIP(), []int{0}
}

func (x *RelationshipNode) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *RelationshipNode) GetType() RelationshipNodeType {
	if x != nil {
		return x.Type
	}
	return RelationshipNodeType_RELATIONSHIP_NODE_TYPE_INVALID
}

func (x *RelationshipNode) GetTypeConfig() *RelationshipNodeTypeConfig {
	if x != nil {
		return x.TypeConfig
	}
	return nil
}

func (x *RelationshipNode) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *RelationshipNode) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *RelationshipNode) GetCreatedByUuid() string {
	if x != nil {
		return x.CreatedByUuid
	}
	return ""
}

func (x *RelationshipNode) GetUpdatedByUuid() string {
	if x != nil {
		return x.UpdatedByUuid
	}
	return ""
}

var File_relationship_node_proto protoreflect.FileDescriptor

var file_relationship_node_proto_rawDesc = []byte{
	0x0a, 0x17, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x6e,
	0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6e, 0x65, 0x6d, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x23, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x02, 0x0a, 0x10, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x68, 0x69, 0x70, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x2d, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x6e, 0x65,
	0x6d, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x4e, 0x6f,
	0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x40, 0x0a, 0x0b,
	0x74, 0x79, 0x70, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x68, 0x69, 0x70, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x0a, 0x74, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x62, 0x79, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x55, 0x75, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79,
	0x55, 0x75, 0x69, 0x64, 0x2a, 0x81, 0x01, 0x0a, 0x14, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x68, 0x69, 0x70, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a,
	0x1e, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x48, 0x49, 0x50, 0x5f, 0x4e, 0x4f,
	0x44, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10,
	0x00, 0x12, 0x21, 0x0a, 0x1d, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x48, 0x49,
	0x50, 0x5f, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x4e, 0x54, 0x49,
	0x54, 0x59, 0x10, 0x01, 0x12, 0x22, 0x0a, 0x1e, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x53, 0x48, 0x49, 0x50, 0x5f, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53,
	0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x02, 0x42, 0x26, 0x0a, 0x03, 0x6e, 0x65, 0x6d, 0x42,
	0x10, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x4e, 0x6f, 0x64,
	0x65, 0x50, 0x01, 0x5a, 0x0b, 0x6e, 0x65, 0x6d, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relationship_node_proto_rawDescOnce sync.Once
	file_relationship_node_proto_rawDescData = file_relationship_node_proto_rawDesc
)

func file_relationship_node_proto_rawDescGZIP() []byte {
	file_relationship_node_proto_rawDescOnce.Do(func() {
		file_relationship_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_relationship_node_proto_rawDescData)
	})
	return file_relationship_node_proto_rawDescData
}

var file_relationship_node_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_relationship_node_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_relationship_node_proto_goTypes = []any{
	(RelationshipNodeType)(0),          // 0: nem.RelationshipNodeType
	(*RelationshipNode)(nil),           // 1: nem.RelationshipNode
	(*RelationshipNodeTypeConfig)(nil), // 2: nem.RelationshipNodeTypeConfig
	(*timestamppb.Timestamp)(nil),      // 3: google.protobuf.Timestamp
}
var file_relationship_node_proto_depIdxs = []int32{
	0, // 0: nem.RelationshipNode.type:type_name -> nem.RelationshipNodeType
	2, // 1: nem.RelationshipNode.type_config:type_name -> nem.RelationshipNodeTypeConfig
	3, // 2: nem.RelationshipNode.created_at:type_name -> google.protobuf.Timestamp
	3, // 3: nem.RelationshipNode.updated_at:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_relationship_node_proto_init() }
func file_relationship_node_proto_init() {
	if File_relationship_node_proto != nil {
		return
	}
	file_relationship_node_type_config_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_relationship_node_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relationship_node_proto_goTypes,
		DependencyIndexes: file_relationship_node_proto_depIdxs,
		EnumInfos:         file_relationship_node_proto_enumTypes,
		MessageInfos:      file_relationship_node_proto_msgTypes,
	}.Build()
	File_relationship_node_proto = out.File
	file_relationship_node_proto_rawDesc = nil
	file_relationship_node_proto_goTypes = nil
	file_relationship_node_proto_depIdxs = nil
}
