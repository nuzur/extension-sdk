// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.24.4
// source: change_request_review.proto

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

type ChangeRequestReviewStatus int32

const (
	ChangeRequestReviewStatus_CHANGE_REQUEST_REVIEW_STATUS_INVALID   ChangeRequestReviewStatus = 0
	ChangeRequestReviewStatus_CHANGE_REQUEST_REVIEW_STATUS_DRAFT     ChangeRequestReviewStatus = 1
	ChangeRequestReviewStatus_CHANGE_REQUEST_REVIEW_STATUS_IN_REVIEW ChangeRequestReviewStatus = 2
	ChangeRequestReviewStatus_CHANGE_REQUEST_REVIEW_STATUS_APPROVED  ChangeRequestReviewStatus = 3
	ChangeRequestReviewStatus_CHANGE_REQUEST_REVIEW_STATUS_REJECTED  ChangeRequestReviewStatus = 4
	ChangeRequestReviewStatus_CHANGE_REQUEST_REVIEW_STATUS_PUBLISHED ChangeRequestReviewStatus = 5
)

// Enum value maps for ChangeRequestReviewStatus.
var (
	ChangeRequestReviewStatus_name = map[int32]string{
		0: "CHANGE_REQUEST_REVIEW_STATUS_INVALID",
		1: "CHANGE_REQUEST_REVIEW_STATUS_DRAFT",
		2: "CHANGE_REQUEST_REVIEW_STATUS_IN_REVIEW",
		3: "CHANGE_REQUEST_REVIEW_STATUS_APPROVED",
		4: "CHANGE_REQUEST_REVIEW_STATUS_REJECTED",
		5: "CHANGE_REQUEST_REVIEW_STATUS_PUBLISHED",
	}
	ChangeRequestReviewStatus_value = map[string]int32{
		"CHANGE_REQUEST_REVIEW_STATUS_INVALID":   0,
		"CHANGE_REQUEST_REVIEW_STATUS_DRAFT":     1,
		"CHANGE_REQUEST_REVIEW_STATUS_IN_REVIEW": 2,
		"CHANGE_REQUEST_REVIEW_STATUS_APPROVED":  3,
		"CHANGE_REQUEST_REVIEW_STATUS_REJECTED":  4,
		"CHANGE_REQUEST_REVIEW_STATUS_PUBLISHED": 5,
	}
)

func (x ChangeRequestReviewStatus) Enum() *ChangeRequestReviewStatus {
	p := new(ChangeRequestReviewStatus)
	*p = x
	return p
}

func (x ChangeRequestReviewStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChangeRequestReviewStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_change_request_review_proto_enumTypes[0].Descriptor()
}

func (ChangeRequestReviewStatus) Type() protoreflect.EnumType {
	return &file_change_request_review_proto_enumTypes[0]
}

func (x ChangeRequestReviewStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChangeRequestReviewStatus.Descriptor instead.
func (ChangeRequestReviewStatus) EnumDescriptor() ([]byte, []int) {
	return file_change_request_review_proto_rawDescGZIP(), []int{0}
}

type ChangeRequestReview struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid      string                    `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	UserUuid  string                    `protobuf:"bytes,2,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	Comment   string                    `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	Status    ChangeRequestReviewStatus `protobuf:"varint,4,opt,name=status,proto3,enum=nem.ChangeRequestReviewStatus" json:"status,omitempty"`
	CreatedAt *timestamppb.Timestamp    `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp    `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *ChangeRequestReview) Reset() {
	*x = ChangeRequestReview{}
	mi := &file_change_request_review_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChangeRequestReview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeRequestReview) ProtoMessage() {}

func (x *ChangeRequestReview) ProtoReflect() protoreflect.Message {
	mi := &file_change_request_review_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeRequestReview.ProtoReflect.Descriptor instead.
func (*ChangeRequestReview) Descriptor() ([]byte, []int) {
	return file_change_request_review_proto_rawDescGZIP(), []int{0}
}

func (x *ChangeRequestReview) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *ChangeRequestReview) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *ChangeRequestReview) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *ChangeRequestReview) GetStatus() ChangeRequestReviewStatus {
	if x != nil {
		return x.Status
	}
	return ChangeRequestReviewStatus_CHANGE_REQUEST_REVIEW_STATUS_INVALID
}

func (x *ChangeRequestReview) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ChangeRequestReview) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_change_request_review_proto protoreflect.FileDescriptor

var file_change_request_review_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6e,
	0x65, 0x6d, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x02, 0x0a, 0x13, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x2a, 0x9b, 0x02, 0x0a, 0x19, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x28, 0x0a, 0x24, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x52, 0x45, 0x51,
	0x55, 0x45, 0x53, 0x54, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x26, 0x0a, 0x22,
	0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x52,
	0x45, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x52, 0x41,
	0x46, 0x54, 0x10, 0x01, 0x12, 0x2a, 0x0a, 0x26, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x52,
	0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x10, 0x02,
	0x12, 0x29, 0x0a, 0x25, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45,
	0x53, 0x54, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x03, 0x12, 0x29, 0x0a, 0x25, 0x43,
	0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x52, 0x45,
	0x56, 0x49, 0x45, 0x57, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x4a, 0x45,
	0x43, 0x54, 0x45, 0x44, 0x10, 0x04, 0x12, 0x2a, 0x0a, 0x26, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45,
	0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48, 0x45, 0x44,
	0x10, 0x05, 0x42, 0x29, 0x0a, 0x03, 0x6e, 0x65, 0x6d, 0x42, 0x13, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x50, 0x01,
	0x5a, 0x0b, 0x6e, 0x65, 0x6d, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_change_request_review_proto_rawDescOnce sync.Once
	file_change_request_review_proto_rawDescData = file_change_request_review_proto_rawDesc
)

func file_change_request_review_proto_rawDescGZIP() []byte {
	file_change_request_review_proto_rawDescOnce.Do(func() {
		file_change_request_review_proto_rawDescData = protoimpl.X.CompressGZIP(file_change_request_review_proto_rawDescData)
	})
	return file_change_request_review_proto_rawDescData
}

var file_change_request_review_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_change_request_review_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_change_request_review_proto_goTypes = []any{
	(ChangeRequestReviewStatus)(0), // 0: nem.ChangeRequestReviewStatus
	(*ChangeRequestReview)(nil),    // 1: nem.ChangeRequestReview
	(*timestamppb.Timestamp)(nil),  // 2: google.protobuf.Timestamp
}
var file_change_request_review_proto_depIdxs = []int32{
	0, // 0: nem.ChangeRequestReview.status:type_name -> nem.ChangeRequestReviewStatus
	2, // 1: nem.ChangeRequestReview.created_at:type_name -> google.protobuf.Timestamp
	2, // 2: nem.ChangeRequestReview.updated_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_change_request_review_proto_init() }
func file_change_request_review_proto_init() {
	if File_change_request_review_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_change_request_review_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_change_request_review_proto_goTypes,
		DependencyIndexes: file_change_request_review_proto_depIdxs,
		EnumInfos:         file_change_request_review_proto_enumTypes,
		MessageInfos:      file_change_request_review_proto_msgTypes,
	}.Build()
	File_change_request_review_proto = out.File
	file_change_request_review_proto_rawDesc = nil
	file_change_request_review_proto_goTypes = nil
	file_change_request_review_proto_depIdxs = nil
}