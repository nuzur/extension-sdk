// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: project_version.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type ProjectVersionReviewStatus int32

const (
	ProjectVersionReviewStatus_PROJECT_VERSION_REVIEW_STATUS_INVALID   ProjectVersionReviewStatus = 0
	ProjectVersionReviewStatus_PROJECT_VERSION_REVIEW_STATUS_DRAFT     ProjectVersionReviewStatus = 1
	ProjectVersionReviewStatus_PROJECT_VERSION_REVIEW_STATUS_IN_REVIEW ProjectVersionReviewStatus = 2
	ProjectVersionReviewStatus_PROJECT_VERSION_REVIEW_STATUS_APPROVED  ProjectVersionReviewStatus = 3
	ProjectVersionReviewStatus_PROJECT_VERSION_REVIEW_STATUS_REJECTED  ProjectVersionReviewStatus = 4
	ProjectVersionReviewStatus_PROJECT_VERSION_REVIEW_STATUS_PUBLISHED ProjectVersionReviewStatus = 5
)

// Enum value maps for ProjectVersionReviewStatus.
var (
	ProjectVersionReviewStatus_name = map[int32]string{
		0: "PROJECT_VERSION_REVIEW_STATUS_INVALID",
		1: "PROJECT_VERSION_REVIEW_STATUS_DRAFT",
		2: "PROJECT_VERSION_REVIEW_STATUS_IN_REVIEW",
		3: "PROJECT_VERSION_REVIEW_STATUS_APPROVED",
		4: "PROJECT_VERSION_REVIEW_STATUS_REJECTED",
		5: "PROJECT_VERSION_REVIEW_STATUS_PUBLISHED",
	}
	ProjectVersionReviewStatus_value = map[string]int32{
		"PROJECT_VERSION_REVIEW_STATUS_INVALID":   0,
		"PROJECT_VERSION_REVIEW_STATUS_DRAFT":     1,
		"PROJECT_VERSION_REVIEW_STATUS_IN_REVIEW": 2,
		"PROJECT_VERSION_REVIEW_STATUS_APPROVED":  3,
		"PROJECT_VERSION_REVIEW_STATUS_REJECTED":  4,
		"PROJECT_VERSION_REVIEW_STATUS_PUBLISHED": 5,
	}
)

func (x ProjectVersionReviewStatus) Enum() *ProjectVersionReviewStatus {
	p := new(ProjectVersionReviewStatus)
	*p = x
	return p
}

func (x ProjectVersionReviewStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProjectVersionReviewStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_project_version_proto_enumTypes[0].Descriptor()
}

func (ProjectVersionReviewStatus) Type() protoreflect.EnumType {
	return &file_project_version_proto_enumTypes[0]
}

func (x ProjectVersionReviewStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProjectVersionReviewStatus.Descriptor instead.
func (ProjectVersionReviewStatus) EnumDescriptor() ([]byte, []int) {
	return file_project_version_proto_rawDescGZIP(), []int{0}
}

type ProjectVersionStatus int32

const (
	ProjectVersionStatus_PROJECT_VERSION_STATUS_INVALID  ProjectVersionStatus = 0
	ProjectVersionStatus_PROJECT_VERSION_STATUS_ACTIVE   ProjectVersionStatus = 1
	ProjectVersionStatus_PROJECT_VERSION_STATUS_DISABLED ProjectVersionStatus = 2
)

// Enum value maps for ProjectVersionStatus.
var (
	ProjectVersionStatus_name = map[int32]string{
		0: "PROJECT_VERSION_STATUS_INVALID",
		1: "PROJECT_VERSION_STATUS_ACTIVE",
		2: "PROJECT_VERSION_STATUS_DISABLED",
	}
	ProjectVersionStatus_value = map[string]int32{
		"PROJECT_VERSION_STATUS_INVALID":  0,
		"PROJECT_VERSION_STATUS_ACTIVE":   1,
		"PROJECT_VERSION_STATUS_DISABLED": 2,
	}
)

func (x ProjectVersionStatus) Enum() *ProjectVersionStatus {
	p := new(ProjectVersionStatus)
	*p = x
	return p
}

func (x ProjectVersionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProjectVersionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_project_version_proto_enumTypes[1].Descriptor()
}

func (ProjectVersionStatus) Type() protoreflect.EnumType {
	return &file_project_version_proto_enumTypes[1]
}

func (x ProjectVersionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProjectVersionStatus.Descriptor instead.
func (ProjectVersionStatus) EnumDescriptor() ([]byte, []int) {
	return file_project_version_proto_rawDescGZIP(), []int{1}
}

type ProjectVersion struct {
	state           protoimpl.MessageState      `protogen:"open.v1"`
	Uuid            string                      `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Version         int64                       `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	Identifier      string                      `protobuf:"bytes,3,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Description     string                      `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	ProjectUuid     string                      `protobuf:"bytes,5,opt,name=project_uuid,json=projectUuid,proto3" json:"project_uuid,omitempty"`
	Entities        []*Entity                   `protobuf:"bytes,6,rep,name=entities,proto3" json:"entities,omitempty"`
	Relationships   []*Relationship             `protobuf:"bytes,7,rep,name=relationships,proto3" json:"relationships,omitempty"`
	Enums           []*Enum                     `protobuf:"bytes,8,rep,name=enums,proto3" json:"enums,omitempty"`
	Services        []*Service                  `protobuf:"bytes,9,rep,name=services,proto3" json:"services,omitempty"`
	BaseVersionUuid string                      `protobuf:"bytes,10,opt,name=base_version_uuid,json=baseVersionUuid,proto3" json:"base_version_uuid,omitempty"`
	ReviewStatus    ProjectVersionReviewStatus  `protobuf:"varint,11,opt,name=review_status,json=reviewStatus,proto3,enum=nem.ProjectVersionReviewStatus" json:"review_status,omitempty"`
	Deployments     []*ProjectVersionDeployment `protobuf:"bytes,12,rep,name=deployments,proto3" json:"deployments,omitempty"`
	Status          ProjectVersionStatus        `protobuf:"varint,13,opt,name=status,proto3,enum=nem.ProjectVersionStatus" json:"status,omitempty"`
	CreatedAt       *timestamppb.Timestamp      `protobuf:"bytes,14,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       *timestamppb.Timestamp      `protobuf:"bytes,15,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CreatedByUuid   string                      `protobuf:"bytes,16,opt,name=created_by_uuid,json=createdByUuid,proto3" json:"created_by_uuid,omitempty"`
	UpdatedByUuid   string                      `protobuf:"bytes,17,opt,name=updated_by_uuid,json=updatedByUuid,proto3" json:"updated_by_uuid,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ProjectVersion) Reset() {
	*x = ProjectVersion{}
	mi := &file_project_version_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProjectVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectVersion) ProtoMessage() {}

func (x *ProjectVersion) ProtoReflect() protoreflect.Message {
	mi := &file_project_version_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectVersion.ProtoReflect.Descriptor instead.
func (*ProjectVersion) Descriptor() ([]byte, []int) {
	return file_project_version_proto_rawDescGZIP(), []int{0}
}

func (x *ProjectVersion) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *ProjectVersion) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *ProjectVersion) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *ProjectVersion) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProjectVersion) GetProjectUuid() string {
	if x != nil {
		return x.ProjectUuid
	}
	return ""
}

func (x *ProjectVersion) GetEntities() []*Entity {
	if x != nil {
		return x.Entities
	}
	return nil
}

func (x *ProjectVersion) GetRelationships() []*Relationship {
	if x != nil {
		return x.Relationships
	}
	return nil
}

func (x *ProjectVersion) GetEnums() []*Enum {
	if x != nil {
		return x.Enums
	}
	return nil
}

func (x *ProjectVersion) GetServices() []*Service {
	if x != nil {
		return x.Services
	}
	return nil
}

func (x *ProjectVersion) GetBaseVersionUuid() string {
	if x != nil {
		return x.BaseVersionUuid
	}
	return ""
}

func (x *ProjectVersion) GetReviewStatus() ProjectVersionReviewStatus {
	if x != nil {
		return x.ReviewStatus
	}
	return ProjectVersionReviewStatus_PROJECT_VERSION_REVIEW_STATUS_INVALID
}

func (x *ProjectVersion) GetDeployments() []*ProjectVersionDeployment {
	if x != nil {
		return x.Deployments
	}
	return nil
}

func (x *ProjectVersion) GetStatus() ProjectVersionStatus {
	if x != nil {
		return x.Status
	}
	return ProjectVersionStatus_PROJECT_VERSION_STATUS_INVALID
}

func (x *ProjectVersion) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ProjectVersion) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *ProjectVersion) GetCreatedByUuid() string {
	if x != nil {
		return x.CreatedByUuid
	}
	return ""
}

func (x *ProjectVersion) GetUpdatedByUuid() string {
	if x != nil {
		return x.UpdatedByUuid
	}
	return ""
}

var File_project_version_proto protoreflect.FileDescriptor

var file_project_version_proto_rawDesc = string([]byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6e, 0x65, 0x6d, 0x1a, 0x0c, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x65, 0x6e, 0x75, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x05, 0x0a,
	0x0e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a,
	0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x75,
	0x69, 0x64, 0x12, 0x27, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x0d, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x68, 0x69, 0x70, 0x52, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x68, 0x69, 0x70, 0x73, 0x12, 0x1f, 0x0a, 0x05, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x05,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x12, 0x28, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12,
	0x2a, 0x0a, 0x11, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x62, 0x61, 0x73, 0x65,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x75, 0x69, 0x64, 0x12, 0x44, 0x0a, 0x0d, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x0c, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x3f, 0x0a, 0x0b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x19, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x55,
	0x75, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62,
	0x79, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x55, 0x75, 0x69, 0x64, 0x2a, 0xa2, 0x02, 0x0a, 0x1a,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x29, 0x0a, 0x25, 0x50, 0x52,
	0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45,
	0x56, 0x49, 0x45, 0x57, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x27, 0x0a, 0x23, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54,
	0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x52, 0x41, 0x46, 0x54, 0x10, 0x01, 0x12, 0x2b,
	0x0a, 0x27, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f,
	0x4e, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x49, 0x4e, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x10, 0x02, 0x12, 0x2a, 0x0a, 0x26, 0x50,
	0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x52,
	0x45, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x50, 0x50,
	0x52, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x03, 0x12, 0x2a, 0x0a, 0x26, 0x50, 0x52, 0x4f, 0x4a, 0x45,
	0x43, 0x54, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45,
	0x57, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45,
	0x44, 0x10, 0x04, 0x12, 0x2b, 0x0a, 0x27, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x56,
	0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x05,
	0x2a, 0x82, 0x01, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x0a, 0x1e, 0x50, 0x52, 0x4f,
	0x4a, 0x45, 0x43, 0x54, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x21, 0x0a,
	0x1d, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01,
	0x12, 0x23, 0x0a, 0x1f, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x56, 0x45, 0x52, 0x53,
	0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42,
	0x4c, 0x45, 0x44, 0x10, 0x02, 0x42, 0x24, 0x0a, 0x03, 0x6e, 0x65, 0x6d, 0x42, 0x0e, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x01, 0x5a, 0x0b,
	0x6e, 0x65, 0x6d, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_project_version_proto_rawDescOnce sync.Once
	file_project_version_proto_rawDescData []byte
)

func file_project_version_proto_rawDescGZIP() []byte {
	file_project_version_proto_rawDescOnce.Do(func() {
		file_project_version_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_project_version_proto_rawDesc), len(file_project_version_proto_rawDesc)))
	})
	return file_project_version_proto_rawDescData
}

var file_project_version_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_project_version_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_project_version_proto_goTypes = []any{
	(ProjectVersionReviewStatus)(0),  // 0: nem.ProjectVersionReviewStatus
	(ProjectVersionStatus)(0),        // 1: nem.ProjectVersionStatus
	(*ProjectVersion)(nil),           // 2: nem.ProjectVersion
	(*Entity)(nil),                   // 3: nem.Entity
	(*Relationship)(nil),             // 4: nem.Relationship
	(*Enum)(nil),                     // 5: nem.Enum
	(*Service)(nil),                  // 6: nem.Service
	(*ProjectVersionDeployment)(nil), // 7: nem.ProjectVersionDeployment
	(*timestamppb.Timestamp)(nil),    // 8: google.protobuf.Timestamp
}
var file_project_version_proto_depIdxs = []int32{
	3, // 0: nem.ProjectVersion.entities:type_name -> nem.Entity
	4, // 1: nem.ProjectVersion.relationships:type_name -> nem.Relationship
	5, // 2: nem.ProjectVersion.enums:type_name -> nem.Enum
	6, // 3: nem.ProjectVersion.services:type_name -> nem.Service
	0, // 4: nem.ProjectVersion.review_status:type_name -> nem.ProjectVersionReviewStatus
	7, // 5: nem.ProjectVersion.deployments:type_name -> nem.ProjectVersionDeployment
	1, // 6: nem.ProjectVersion.status:type_name -> nem.ProjectVersionStatus
	8, // 7: nem.ProjectVersion.created_at:type_name -> google.protobuf.Timestamp
	8, // 8: nem.ProjectVersion.updated_at:type_name -> google.protobuf.Timestamp
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_project_version_proto_init() }
func file_project_version_proto_init() {
	if File_project_version_proto != nil {
		return
	}
	file_entity_proto_init()
	file_enum_proto_init()
	file_project_version_deployment_proto_init()
	file_relationship_proto_init()
	file_service_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_project_version_proto_rawDesc), len(file_project_version_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_version_proto_goTypes,
		DependencyIndexes: file_project_version_proto_depIdxs,
		EnumInfos:         file_project_version_proto_enumTypes,
		MessageInfos:      file_project_version_proto_msgTypes,
	}.Build()
	File_project_version_proto = out.File
	file_project_version_proto_goTypes = nil
	file_project_version_proto_depIdxs = nil
}
