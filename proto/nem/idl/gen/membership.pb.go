// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.24.4
// source: membership.proto

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

type MembershipStatus int32

const (
	MembershipStatus_MEMBERSHIP_STATUS_INVALID  MembershipStatus = 0
	MembershipStatus_MEMBERSHIP_STATUS_ACTIVE   MembershipStatus = 1
	MembershipStatus_MEMBERSHIP_STATUS_DISABLED MembershipStatus = 2
)

// Enum value maps for MembershipStatus.
var (
	MembershipStatus_name = map[int32]string{
		0: "MEMBERSHIP_STATUS_INVALID",
		1: "MEMBERSHIP_STATUS_ACTIVE",
		2: "MEMBERSHIP_STATUS_DISABLED",
	}
	MembershipStatus_value = map[string]int32{
		"MEMBERSHIP_STATUS_INVALID":  0,
		"MEMBERSHIP_STATUS_ACTIVE":   1,
		"MEMBERSHIP_STATUS_DISABLED": 2,
	}
)

func (x MembershipStatus) Enum() *MembershipStatus {
	p := new(MembershipStatus)
	*p = x
	return p
}

func (x MembershipStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MembershipStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_membership_proto_enumTypes[0].Descriptor()
}

func (MembershipStatus) Type() protoreflect.EnumType {
	return &file_membership_proto_enumTypes[0]
}

func (x MembershipStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MembershipStatus.Descriptor instead.
func (MembershipStatus) EnumDescriptor() ([]byte, []int) {
	return file_membership_proto_rawDescGZIP(), []int{0}
}

type MembershipType int32

const (
	MembershipType_MEMBERSHIP_TYPE_INVALID    MembershipType = 0
	MembershipType_MEMBERSHIP_TYPE_FREE       MembershipType = 1
	MembershipType_MEMBERSHIP_TYPE_SMALL      MembershipType = 2
	MembershipType_MEMBERSHIP_TYPE_MEDIUM     MembershipType = 3
	MembershipType_MEMBERSHIP_TYPE_ENTERPRISE MembershipType = 4
)

// Enum value maps for MembershipType.
var (
	MembershipType_name = map[int32]string{
		0: "MEMBERSHIP_TYPE_INVALID",
		1: "MEMBERSHIP_TYPE_FREE",
		2: "MEMBERSHIP_TYPE_SMALL",
		3: "MEMBERSHIP_TYPE_MEDIUM",
		4: "MEMBERSHIP_TYPE_ENTERPRISE",
	}
	MembershipType_value = map[string]int32{
		"MEMBERSHIP_TYPE_INVALID":    0,
		"MEMBERSHIP_TYPE_FREE":       1,
		"MEMBERSHIP_TYPE_SMALL":      2,
		"MEMBERSHIP_TYPE_MEDIUM":     3,
		"MEMBERSHIP_TYPE_ENTERPRISE": 4,
	}
)

func (x MembershipType) Enum() *MembershipType {
	p := new(MembershipType)
	*p = x
	return p
}

func (x MembershipType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MembershipType) Descriptor() protoreflect.EnumDescriptor {
	return file_membership_proto_enumTypes[1].Descriptor()
}

func (MembershipType) Type() protoreflect.EnumType {
	return &file_membership_proto_enumTypes[1]
}

func (x MembershipType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MembershipType.Descriptor instead.
func (MembershipType) EnumDescriptor() ([]byte, []int) {
	return file_membership_proto_rawDescGZIP(), []int{1}
}

type Membership struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid          string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	OwnerUuid     string                 `protobuf:"bytes,2,opt,name=owner_uuid,json=ownerUuid,proto3" json:"owner_uuid,omitempty"`
	Type          MembershipType         `protobuf:"varint,3,opt,name=type,proto3,enum=nem.MembershipType" json:"type,omitempty"`
	Metadata      string                 `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Status        MembershipStatus       `protobuf:"varint,5,opt,name=status,proto3,enum=nem.MembershipStatus" json:"status,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CreatedByUuid string                 `protobuf:"bytes,8,opt,name=created_by_uuid,json=createdByUuid,proto3" json:"created_by_uuid,omitempty"`
	UpdatedByUuid string                 `protobuf:"bytes,9,opt,name=updated_by_uuid,json=updatedByUuid,proto3" json:"updated_by_uuid,omitempty"`
}

func (x *Membership) Reset() {
	*x = Membership{}
	mi := &file_membership_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Membership) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Membership) ProtoMessage() {}

func (x *Membership) ProtoReflect() protoreflect.Message {
	mi := &file_membership_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Membership.ProtoReflect.Descriptor instead.
func (*Membership) Descriptor() ([]byte, []int) {
	return file_membership_proto_rawDescGZIP(), []int{0}
}

func (x *Membership) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Membership) GetOwnerUuid() string {
	if x != nil {
		return x.OwnerUuid
	}
	return ""
}

func (x *Membership) GetType() MembershipType {
	if x != nil {
		return x.Type
	}
	return MembershipType_MEMBERSHIP_TYPE_INVALID
}

func (x *Membership) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *Membership) GetStatus() MembershipStatus {
	if x != nil {
		return x.Status
	}
	return MembershipStatus_MEMBERSHIP_STATUS_INVALID
}

func (x *Membership) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Membership) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Membership) GetCreatedByUuid() string {
	if x != nil {
		return x.CreatedByUuid
	}
	return ""
}

func (x *Membership) GetUpdatedByUuid() string {
	if x != nil {
		return x.UpdatedByUuid
	}
	return ""
}

var File_membership_proto protoreflect.FileDescriptor

var file_membership_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x6e, 0x65, 0x6d, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf9, 0x02, 0x0a, 0x0a, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x15, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x62, 0x79, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x55, 0x75, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79,
	0x55, 0x75, 0x69, 0x64, 0x2a, 0x6f, 0x0a, 0x10, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68,
	0x69, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x19, 0x4d, 0x45, 0x4d, 0x42,
	0x45, 0x52, 0x53, 0x48, 0x49, 0x50, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x4d, 0x45, 0x4d, 0x42, 0x45,
	0x52, 0x53, 0x48, 0x49, 0x50, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x43, 0x54,
	0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x53,
	0x48, 0x49, 0x50, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42,
	0x4c, 0x45, 0x44, 0x10, 0x02, 0x2a, 0x9e, 0x01, 0x0a, 0x0e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x68, 0x69, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x45, 0x4d, 0x42,
	0x45, 0x52, 0x53, 0x48, 0x49, 0x50, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x53,
	0x48, 0x49, 0x50, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x52, 0x45, 0x45, 0x10, 0x01, 0x12,
	0x19, 0x0a, 0x15, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x53, 0x48, 0x49, 0x50, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x53, 0x4d, 0x41, 0x4c, 0x4c, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x45,
	0x4d, 0x42, 0x45, 0x52, 0x53, 0x48, 0x49, 0x50, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x45,
	0x44, 0x49, 0x55, 0x4d, 0x10, 0x03, 0x12, 0x1e, 0x0a, 0x1a, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52,
	0x53, 0x48, 0x49, 0x50, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x4e, 0x54, 0x45, 0x52, 0x50,
	0x52, 0x49, 0x53, 0x45, 0x10, 0x04, 0x42, 0x20, 0x0a, 0x03, 0x6e, 0x65, 0x6d, 0x42, 0x0a, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x50, 0x01, 0x5a, 0x0b, 0x6e, 0x65, 0x6d,
	0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_membership_proto_rawDescOnce sync.Once
	file_membership_proto_rawDescData = file_membership_proto_rawDesc
)

func file_membership_proto_rawDescGZIP() []byte {
	file_membership_proto_rawDescOnce.Do(func() {
		file_membership_proto_rawDescData = protoimpl.X.CompressGZIP(file_membership_proto_rawDescData)
	})
	return file_membership_proto_rawDescData
}

var file_membership_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_membership_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_membership_proto_goTypes = []any{
	(MembershipStatus)(0),         // 0: nem.MembershipStatus
	(MembershipType)(0),           // 1: nem.MembershipType
	(*Membership)(nil),            // 2: nem.Membership
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_membership_proto_depIdxs = []int32{
	1, // 0: nem.Membership.type:type_name -> nem.MembershipType
	0, // 1: nem.Membership.status:type_name -> nem.MembershipStatus
	3, // 2: nem.Membership.created_at:type_name -> google.protobuf.Timestamp
	3, // 3: nem.Membership.updated_at:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_membership_proto_init() }
func file_membership_proto_init() {
	if File_membership_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_membership_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_membership_proto_goTypes,
		DependencyIndexes: file_membership_proto_depIdxs,
		EnumInfos:         file_membership_proto_enumTypes,
		MessageInfos:      file_membership_proto_msgTypes,
	}.Build()
	File_membership_proto = out.File
	file_membership_proto_rawDesc = nil
	file_membership_proto_goTypes = nil
	file_membership_proto_depIdxs = nil
}
